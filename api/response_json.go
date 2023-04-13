package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

// 判断结构体是否为空
func (m ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status) //AbortWithStatus只会返回一个
		return
	}

	ctx.AbortWithStatusJSON(status, resp)
}

func buildStatus(resp ResponseJson, nDefaultStatus int) int {
	if 0 == resp.Status {
		return nDefaultStatus
	}

	return resp.Status
} //
// 请求成功
func OK(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

// 其他错判
func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

// 服务端错误
func ServerFail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}
