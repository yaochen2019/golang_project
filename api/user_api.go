package api

import (
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}

}

// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/punlic/user/login [post]
func (m UserApi) Login(ctx *gin.Context) {
	OK(ctx, ResponseJson{
		Msg: "Login success",
	})
}
