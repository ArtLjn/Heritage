/*
@Time : 2024/5/23 下午4:45
@Author : ljn
@File : router
@Software: GoLand
*/

package router

import (
	"back/conf"
	"back/internal/response"
	"back/internal/service"
	"back/util"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, c *conf.Conf, heritageService *service.HeritageService,
	accountService *service.AccountService) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	group := r.Group("/api")
	group.POST("/upload", util.GinUploadImg)
	heritageGroup := group.Group("/heritage")
	{
		heritageGroup.POST("/CreateHeritageInheritor", heritageService.CreateHeritageInheritor)
		heritageGroup.POST("/CreateHeritageProject", heritageService.CreateHeritageProject)
		heritageGroup.GET("/queryHeritageInheritor", heritageService.QueryAllHeritageInheritor)
		heritageGroup.GET("/queryHeritageProject", heritageService.QueryAllHeritageProject)
		heritageGroup.GET("/getCategory", heritageService.GetCateGory)
		heritageGroup.GET("/getLevel", heritageService.GetLevel)
		heritageGroup.GET("/queryHeritageTask", heritageService.QueryHeritageTask)
		heritageGroup.GET("/auditHeritageTask", heritageService.AuditHeritageTask)
		heritageGroup.GET("/queryHeritageByLocate", heritageService.QueryHeritageByLocate)
	}
	availedAdmin := func(ctx *gin.Context) {
		f := response.NewResponseBuild()
		token := ctx.GetHeader("Authorization")
		if err := accountService.IsOk(token); err != nil {
			f.SetCode(400).SetMsg(err.Error()).Build(ctx)
			ctx.Abort()
			return
		} else if util.GetLoginName(token) != c.AdminAccount.Username {
			f.SetCode(400).SetMsg("you are not admin").Build(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
	accGroup := group.Group("/account")
	{
		accGroup.POST("/login", accountService.Login)
		accGroup.GET("/logout", accountService.LogOut)
		accGroup.GET("/verifyToken", accountService.VerifyToken)
		accGroup.Use(availedAdmin).GET("/getAllAccount", accountService.FindAllAccount)
		accGroup.Use(availedAdmin).GET("/exportAccount", accountService.ExportAccount)
	}
}
