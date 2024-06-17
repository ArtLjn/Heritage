/*
@Time : 2024/5/25 下午12:33
@Author : ljn
@File : account
@Software: GoLand
*/

package service

import (
	"back/internal/model"
	"back/internal/response"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
)

type AccountRepo interface {
	InitAccount(map[string][]string)
	Login(login model.AccountLogin) (data []interface{}, err error)
	LogOut(key string)
	VerifyToken(key string) error
	GetAllAccount() []model.Account
	QueryAccountByCity(city string) model.Account
}

type AccountService struct {
	repo AccountRepo
	r    response.ResponseBuild
}

func NewAccountService(repo AccountRepo) *AccountService {
	return &AccountService{repo: repo}
}

func (a *AccountService) Login(ctx *gin.Context) {
	a.r = response.NewResponseBuild()
	acc := model.AccountLogin{}
	if err := ctx.ShouldBind(&acc); err != nil {
		a.r.NewBuildJsonError(ctx)
		return
	}
	data, err := a.repo.Login(acc)
	if err != nil {
		a.r.SetCode(400).SetMsg("登录失败").SetData(nil).Build(ctx)
		return
	}
	a.r.SetCode(200).SetMsg("success").SetData(data).Build(ctx)
	return
}

func (a *AccountService) LogOut(ctx *gin.Context) {
	a.r = response.NewResponseBuild()
	key := ctx.Query("city")
	a.repo.LogOut(key)
	a.r.SetCode(200).SetMsg("success").SetData(nil).Build(ctx)
	return
}

func (a *AccountService) VerifyToken(ctx *gin.Context) {
	a.r = response.NewResponseBuild()
	token := ctx.GetHeader("Authorization")
	if err := a.IsOk(token); err != nil {
		a.r.SetCode(400).SetMsg(err.Error()).SetData(nil).Build(ctx)
		return
	}
	a.r.SetCode(200).SetMsg("success").SetData(nil).Build(ctx)
	return
}

func (a *AccountService) FindAllAccount(ctx *gin.Context) {
	a.r = response.NewResponseBuild()
	data := a.repo.GetAllAccount()
	a.r.SetCode(200).SetMsg("success").SetData(data).Build(ctx)
	return
}

func (a *AccountService) IsOk(token string) error {
	if token == "" {
		return fmt.Errorf("token为空")
	}
	err := a.repo.VerifyToken(token)
	if err != nil {
		return fmt.Errorf("token失效")
	}
	return nil
}

func (a *AccountService) ExportAccount(ctx *gin.Context) {
	a.r = response.NewResponseBuild()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			a.r.SetCode(400).SetMsg("导出失败").SetData(nil).Build(ctx)
			return
		}
	}()
	// 创建一个新的工作簿
	f := excelize.NewFile()

	// 创建一个新的工作表
	sheetName := "Account"
	index, _ := f.NewSheet(sheetName)
	// 字段名称
	account := &model.Account{}
	fields := account.ToFieldList()
	for col, field := range fields {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		if err := f.SetCellValue(sheetName, cell, field); err != nil {
			panic(err)
		} else if err = f.SetColWidth(sheetName, "B", "B", 45); err != nil {
			panic(err)
		} else if err = f.SetColWidth(sheetName, "D", "D", 35); err != nil {
			panic(err)
		}
	}

	accounts := a.repo.GetAllAccount()
	for row, acc := range accounts {
		var e error
		e = f.SetCellValue(sheetName, fmt.Sprintf("A%d", row+2), acc.Id)
		e = f.SetCellValue(sheetName, fmt.Sprintf("B%d", row+2), acc.Address)
		e = f.SetCellValue(sheetName, fmt.Sprintf("C%d", row+2), acc.Password)
		e = f.SetCellValue(sheetName, fmt.Sprintf("D%d", row+2), acc.City)
		e = f.SetCellValue(sheetName, fmt.Sprintf("E%d", row+2), acc.Level)
		if e != nil {
			panic(e)
		}
	}
	// 设置默认工作表
	f.SetActiveSheet(index)

	// 将工作簿保存到内存缓冲区
	var buffer bytes.Buffer
	if err := f.Write(&buffer); err != nil {
		panic(err)
	}
	// 设置HTTP头信息
	ctx.Header("Content-Disposition", "attachment; filename=Account.xlsx")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", fmt.Sprintf("%d", buffer.Len()))
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
}
