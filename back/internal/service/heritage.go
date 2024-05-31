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
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type HeritageRepo interface {
	CreateHeritageInheritor(inheritor *model.HeritageInheritor) error
	CreateHeritageProject(project *model.HeritageProject) error
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
	s.r = response.NewResponseBuild()
	heritageInheritor := model.HeritageInheritorModel
	if err := ctx.ShouldBindJSON(&heritageInheritor); err != nil {
		s.r.NewBuildJsonError(ctx)
		return
	}
	heritageInheritor.Locate = strings.ReplaceAll(heritageInheritor.Locate, "/", "-")
	if err := s.repo.PublishHeritageInheritor(heritageInheritor); err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.NewBuildSuccess(ctx)
}

func (s *HeritageService) CreateHeritageProject(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
	heritageProject := model.HeritageProjectModel
	if err := ctx.ShouldBindJSON(&heritageProject); err != nil {
		s.r.NewBuildJsonError(ctx)
		return
	} else if err = s.repo.PublishHeritageProject(heritageProject); err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.NewBuildSuccess(ctx)
}

func (s *HeritageService) QueryAllHeritageInheritor(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
	list, err := s.repo.QueryAllHeritageInheritor()
	if err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.SetCode(200).SetData(list).Build(ctx)
}

func (s *HeritageService) QueryAllHeritageProject(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
	list, err := s.repo.QueryAllHeritageProject()
	if err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.SetCode(200).SetData(list).Build(ctx)
}

func (s *HeritageService) GetCateGory(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
	s.r.SetCode(200).SetData(s.repo.GetCateGory()).Build(ctx)
}

func (s *HeritageService) GetLevel(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
	s.r.SetCode(200).SetData(s.repo.GetLevel()).Build(ctx)
}

func (s *HeritageService) QueryHeritageTask(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
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

func (s *HeritageService) AuditHeritageTask(ctx *gin.Context) {
	s.r = response.NewResponseBuild()
	header := ctx.Request.Header.Get("Authorization")
	id, _ := strconv.Atoi(ctx.Query("id"))
	allow := ctx.Query("allow")

	res := strconv.CanBackquote(allow)
	if err := s.ac.VerifyToken(header); err != nil {
		s.r.SetCode(400).SetMsg(err.Error()).Build(ctx)
		return
	}
	city := util.GetLoginName(header)
	m, err := s.repo.QueryHeritageTaskById(id)
	if err != nil && city == "" {
		s.r.SetCode(400).SetMsg(err.Error()).Build(ctx)
		return
	} else if m.Locate != city {
		s.r.SetCode(400).SetMsg("审核失败，没有审核权限").Build(ctx)
		return
	}

	if res {
		if m.Type == 1 {
			var ia model.HeritageInheritor
			if err = json.Unmarshal([]byte(m.Field), &ia); err != nil {
				s.r.NewBuildJsonError(ctx)
				return
			}
			if err = s.repo.CreateHeritageInheritor(&ia); err != nil {
				s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
				return
			}
		} else if m.Type == 2 {
			var pa model.HeritageProject
			if err = json.Unmarshal([]byte(m.Field), &pa); err != nil {
				s.r.NewBuildJsonError(ctx)
				return
			}
			if err = s.repo.CreateHeritageProject(&pa); err != nil {
				s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
				return
			}
		}
	}

	err = s.repo.DeleteHeritageTaskById(m.Id)
	if err != nil {
		s.r.SetCode(500).SetMsg(err.Error()).Build(ctx)
		return
	}
	s.r.NewBuildSuccess(ctx)
}
