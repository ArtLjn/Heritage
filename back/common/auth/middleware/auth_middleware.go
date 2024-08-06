/*
@Time : 2024/7/7 下午8:22
@Author : ljn
@File : auth_middleware
@Software: GoLand
*/

package middlewa

import (
	"back/conf"
	"back/internal/response"
	"back/internal/service"
	"back/util"
	"github.com/gin-gonic/gin"
)

func IsAdmin(accountService *service.AccountService, c *conf.Conf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
}
