/*
@Time : 2024/5/23 下午8:09
@Author : ljn
@File : heritage
@Software: GoLand
*/

package service

import (
	"back/internal/model"
	"back/internal/response"
	"back/util"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HeritageRepo interface {
	CreateHeritageInheritor(inheritor *model.HeritageInheritor) ([]interface{}, error)
	CreateHeritageProject(project *model.HeritageProject) ([]interface{}, error)
	InitHeritageInheritor()
	InitHeritageProject()
	QueryAllHeritageInheritor() ([]interface{}, error)
	QueryAllHeritageProject() ([]interface{}, error)
	GetCateGory() map[string]interface{}
	GetLevel() map[string]interface{}
	PublishHeritageInheritor(inheritor *model.HeritageInheritor) error
	PublishHeritageProject(project *model.HeritageProject) error
	QueryHeritageInheritor(page, size int, locate string) map[string]interface{}
	QueryHeritageProject(page, size int, locate string) map[string]interface{}
	ReceiveHeritageInheritor()
	ReceiveHeritageProject()
	QueryHeritageTaskById(id int) (model.Heritage, error)
	DeleteHeritageTaskById(id int) error
	QueryHeritageInheritorByLocate(page, size, raw int, locate string) (map[string]interface{}, error)
	QueryHeritageProjectByLocate(page, size, raw int, locate string) (map[string]interface{}, error)
	UpdateHeritageTask(level uint, id int, number string) error
	UpdateHeritageInheritorLevel(number string, rank uint) error
	UpdateHeritageProjectLevel(number string, rank uint) error
	QueryHeritageInheritorBlockChainDataByNumber(number string) []interface{}
	QueryHeritageProjectBlockChainDataByNumber(number string) []interface{}
}

type HeritageService struct {
	repo HeritageRepo
	ac   AccountRepo
	r    response.ResponseBuild
}

func NewHeritageService(repo HeritageRepo, ac AccountRepo) *HeritageService {
	return &HeritageService{
		repo: repo,
		ac:   ac,
	}
}

func (s *HeritageService) CreateHeritageInheritor(ctx *gin.Context) {
	heritageInheritor := model.HeritageInheritorModel
	if err := ctx.ShouldBindJSON(&heritageInheritor); err != nil {
		s.r.NewBuildJsonError(ctx)
		return
	}
	heritageInheritor.Locate = strings.ReplaceAll(heritageInheritor.Locate, "/", "-")
	if len(strings.Split(heritageInheritor.Locate, "-")) != 2 {
		s.r.SetCode(400).SetMsg("请填写你所在的市级单位").Build(ctx)
		return
	} else if err := s.repo.PublishHeritageInheritor(heritageInheritor); err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.NewBuildSuccess(ctx)
}

func (s *HeritageService) CreateHeritageProject(ctx *gin.Context) {
	heritageProject := model.HeritageProjectModel
	if err := ctx.ShouldBindJSON(&heritageProject); err != nil {
		s.r.NewBuildJsonError(ctx)
		return
	}
	heritageProject.Locate = strings.ReplaceAll(heritageProject.Locate, "/", "-")
	if len(strings.Split(heritageProject.Locate, "-")) != 2 {
		s.r.SetCode(400).SetMsg("请填写你所在的市级单位").Build(ctx)
		return
	} else if err := s.repo.PublishHeritageProject(heritageProject); err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.NewBuildSuccess(ctx)
}

func (s *HeritageService) QueryAllHeritageInheritor(ctx *gin.Context) {
	list, err := s.repo.QueryAllHeritageInheritor()
	if err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.SetCode(200).SetData(list).Build(ctx)
}

func (s *HeritageService) QueryAllHeritageProject(ctx *gin.Context) {
	list, err := s.repo.QueryAllHeritageProject()
	if err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.SetCode(200).SetData(list).Build(ctx)
}

func (s *HeritageService) GetCateGory(ctx *gin.Context) {
	s.r.SetCode(200).SetData(s.repo.GetCateGory()).Build(ctx)
}

func (s *HeritageService) GetLevel(ctx *gin.Context) {
	s.r.SetCode(200).SetData(s.repo.GetLevel()).Build(ctx)
}
func (s *HeritageService) QueryHeritageTask(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	raw, _ := strconv.Atoi(ctx.Query("raw"))
	header := ctx.Request.Header.Get("Authorization")
	data := make(map[string]interface{})
	switch raw {
	case model.HeritageTypeInheritor:
		data = s.repo.QueryHeritageInheritor(page, size, header)
	case model.HeritageTypeProject:
		data = s.repo.QueryHeritageProject(page, size, header)
	}
	s.r.SetCode(200).SetMsg("success").SetData(data).Build(ctx)
}
func (s *HeritageService) QueryHeritageByLocate(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	raw := ctx.Query("raw")
	locate := ctx.Query("locate")

	locates := strings.Split(locate, "-")
	switch raw {
	case "1":
		list, err := s.repo.QueryHeritageInheritorByLocate(page, size, len(locates), locate)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
			return
		}
		s.r.SetCode(200).SetData(list).Build(ctx)
	case "2":
		list, err := s.repo.QueryHeritageProjectByLocate(page, size, len(locates), locate)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
			return
		}
		s.r.SetCode(200).SetData(list).Build(ctx)
	}
}
func (s *HeritageService) AuditHeritageTask(ctx *gin.Context) {
	header := ctx.Request.Header.Get("Authorization")
	id, _ := strconv.Atoi(ctx.Query("id"))
	// 是否审核通过: bool
	allow := ctx.Query("allow")
	res := strconv.CanBackquote(allow)

	// Token Verify
	if err := s.ac.VerifyToken(header); err != nil {
		s.r.SetCode(400).SetMsg(err.Error()).Build(ctx)
		return
	}
	// 获取城市名称
	city := util.GetLoginName(header)
	if city == "" {
		s.r.SetCode(400).SetMsg("Token无效").Build(ctx)
		return
	}

	// 根据任务编号查询要审核的项目
	m, err := s.repo.QueryHeritageTaskById(id)
	if err != nil {
		s.r.SetCode(400).SetMsg(err.Error()).Build(ctx)
		return
	}

	// 是否pass
	if res {
		if m.Type == model.HeritageTypeInheritor {
			if err = s.auditHeritageInheritor(&m, city); err != nil {
				s.r.SetCode(400).SetMsg(err.Error()).Build(ctx)
				return
			}
		} else if m.Type == model.HeritageTypeProject {
			if err = s.authHeritageProject(&m, city); err != nil {
				s.r.SetCode(400).SetMsg(err.Error()).Build(ctx)
				return
			}
		}
	}

	// 如果申请的审核级别和审核通过的级别相同 || 审核未通过 则删除该任务条
	if m.ApplyLevel == m.PassLevel || !res {
		err = s.repo.DeleteHeritageTaskById(m.Id)
	}
	if err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.NewBuildSuccess(ctx)
}

func (s *HeritageService) auditHeritageInheritor(m *model.Heritage, city string) error {
	var (
		level  uint
		number string
	)
	if le := model.HeritageByLevel(city); le != 0 {
		level = le
	} else {
		level = s.ac.QueryAccountByCity(city).Level
	}
	var ia model.HeritageInheritor
	if err := json.Unmarshal([]byte(m.Field), &ia); err != nil {
		return err
	}
	ia.Level = uint8(level)
	if len(m.Number) != 0 {
		number = m.Number
		chainRecord := s.repo.QueryHeritageInheritorBlockChainDataByNumber(number)
		if len(chainRecord) == 0 {
			return fmt.Errorf("记录不存在于区块链上，请检查number是否正确")
		} else if model.HeritageByLevel(chainRecord[6].(string)) == level {
			// 判断该记录当前等级是否等于当前城市等级，如果相同则表明已经审核过该记录不可以再次审核
			return fmt.Errorf("该记录已经审核不可以再次审核")
		}
		err := s.repo.UpdateHeritageInheritorLevel(number, level)
		if err != nil {
			return err
		}
	} else {
		copyData, err := s.repo.CreateHeritageInheritor(&ia)
		if err != nil {
			return err
		}
		number = copyData[0].(string)
	}

	if err := s.repo.UpdateHeritageTask(level, m.Id, number); err != nil {
		return err
	}
	m.PassLevel = uint8(level)
	return nil
}
func (s *HeritageService) authHeritageProject(m *model.Heritage, city string) error {
	var (
		level  uint
		number string
	)
	if le := model.HeritageByLevel(city); le != 0 {
		level = le
	} else {
		level = s.ac.QueryAccountByCity(city).Level
	}
	var pa model.HeritageProject
	if err := json.Unmarshal([]byte(m.Field), &pa); err != nil {
		return err
	}
	pa.Level = uint8(level)
	// 如果非遗任务中number已经存在，则更新其审核级别
	if len(m.Number) != 0 {
		number = m.Number
		chainRecord := s.repo.QueryHeritageProjectBlockChainDataByNumber(number)
		if len(chainRecord) == 0 {
			return fmt.Errorf("记录不存在于区块链上，请检查number是否正确")
		} else if model.HeritageByLevel(chainRecord[3].(string)) == level {
			// 判断该记录当前等级是否等于当前城市等级，如果相同则表明已经审核过该记录不可以再次审核
			return fmt.Errorf("该记录已经审核不可以再次审核")
		}
		err := s.repo.UpdateHeritageProjectLevel(number, level)
		if err != nil {
			return err
		}
	} else {
		// copyData :上链成功之后返回的数据
		copyData, err := s.repo.CreateHeritageProject(&pa)
		if err != nil {
			return err
		}
		number = copyData[0].(string)
	}

	if err := s.repo.UpdateHeritageTask(level, m.Id, number); err != nil {
		return err
	}
	m.PassLevel = uint8(level)
	return nil
}

func (s *HeritageService) validPermission(level uint8, applyCity, authCity string) error {
	applyCityList := strings.Split(applyCity, "-")
	authCityList := strings.Split(authCity, "-")
	if level == model.National && authCity != "国家" || level == model.Human && authCity != "人类非物质遗产" ||
		level == model.Provincial && applyCityList[0] != authCityList[0] || level == model.City && applyCity != authCity {
		return fmt.Errorf("权限不允许")
	}
	return nil
}
