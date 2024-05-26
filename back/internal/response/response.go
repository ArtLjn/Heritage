/*
@Time : 2024/5/23 下午8:43
@Author : ljn
@File : response
@Software: GoLand
*/

package response

import "github.com/gin-gonic/gin"

type BuildResponse interface {
	SetCode(code int) BuildResponse
	SetMsg(msg string) BuildResponse
	SetData(data ...interface{}) BuildResponse
	Build(ctx *gin.Context)
}

type ResponseBuild struct {
	code int
	msg  string
	data []interface{}
}

func (r *ResponseBuild) SetCode(code int) BuildResponse {
	r.code = code
	return r
}

func (r *ResponseBuild) SetMsg(msg string) BuildResponse {
	r.msg = msg
	return r
}

func (r *ResponseBuild) SetData(data ...interface{}) BuildResponse {
	r.data = data
	return r
}

func (r *ResponseBuild) Build(ctx *gin.Context) {
	ctx.JSON(r.code, gin.H{"code": r.code, "msg": r.msg, "data": r.data})
}

func (r *ResponseBuild) NewBuildJsonError(ctx *gin.Context) {
	r.SetCode(400).SetMsg("json error").SetData(nil).Build(ctx)
}

func (r *ResponseBuild) NewBuildSuccess(ctx *gin.Context) {
	r.SetCode(200).SetMsg("success").SetData(nil).Build(ctx)
}
