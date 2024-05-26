/*
@Time : 2024/5/23 下午4:45
@Author : ljn
@File : router
@Software: GoLand
*/

package router

import (
	"back/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, heritageService *service.HeritageService,
	accountService *service.AccountService) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	group := r.Group("/api")
	heritageGroup := group.Group("/heritage")
	{
		heritageGroup.POST("/CreateHeritageInheritor", heritageService.CreateHeritageInheritor)
		heritageGroup.POST("/CreateHeritageProject", heritageService.CreateHeritageProject)
		heritageGroup.GET("/queryHeritageInheritor", heritageService.QueryAllHeritageInheritor)
		heritageGroup.GET("/queryHeritageProject", heritageService.QueryAllHeritageProject)
		heritageGroup.GET("/getCategory", heritageService.GetCateGory)
		heritageGroup.GET("/getLevel", heritageService.GetLevel)
		heritageGroup.GET("/queryHeritageTask", heritageService.QueryHeritageTask)
	}
	accGroup := group.Group("/account")
	{
		accGroup.POST("/login", accountService.Login)
		accGroup.GET("/logout", accountService.LogOut)
		accGroup.GET("/verifyToken", accountService.VerifyToken)
	}
}
