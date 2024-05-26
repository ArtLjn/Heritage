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
	"github.com/gin-gonic/gin"
	"strconv"
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
	QueryHeritageInheritor(page, size int) map[string]interface{}
	QueryHeritageProject(page, size int) map[string]interface{}
	ReceiveHeritageInheritor()
	ReceiveHeritageProject()
}

type HeritageService struct {
	repo HeritageRepo
	r    response.ResponseBuild
}

func NewHeritageService(repo HeritageRepo) *HeritageService {
	return &HeritageService{
		repo: repo,
	}
}

func (s *HeritageService) CreateHeritageInheritor(ctx *gin.Context) {
	heritageInheritor := model.HeritageInheritorModel
	if err := ctx.ShouldBindJSON(&heritageInheritor); err != nil {
		s.r.NewBuildJsonError(ctx)
		return
	} else if err = s.repo.PublishHeritageInheritor(heritageInheritor); err != nil {
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
	} else if err = s.repo.PublishHeritageProject(heritageProject); err != nil {
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

func (s *HeritageService) QueryHeritageInheritor(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	raw, _ := strconv.Atoi(ctx.Query("raw"))
	data := make(map[string]interface{})
	switch raw {
	case model.HeritageTypeInheritor:
		data = s.repo.QueryHeritageInheritor(page, size)
	case model.HeritageTypeProject:
		data = s.repo.QueryHeritageProject(page, size)
	}
	s.r.SetCode(200).SetMsg("success").SetData(data).Build(ctx)
}
