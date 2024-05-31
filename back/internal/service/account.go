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
	"github.com/gin-gonic/gin"
)

type AccountRepo interface {
	InitAccount()
	Login(login model.AccountLogin) (data []interface{}, err error)
	LogOut(key string)
	VerifyToken(key string) error
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
	if token == "" {
		a.r.SetCode(400).SetMsg("token失效").SetData(nil).Build(ctx)
		return
	}
	err := a.repo.VerifyToken(token)
	if err != nil {
		a.r.SetCode(400).SetMsg("token失效").SetData(nil).Build(ctx)
		return
	}
	a.r.SetCode(200).SetMsg("success").SetData(nil).Build(ctx)
	return
}
