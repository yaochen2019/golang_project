package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yaochen2019/gogofly/service/dto"
	"github.com/yaochen2019/gogofly/util"
	"reflect"
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
	var iUserLoginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&iUserLoginDTO) //拿到前端的数据
	if errs != nil {
		Fail(ctx, ResponseJson{
			Msg: parseValidateErrors(errs.(validator.ValidationErrors), &iUserLoginDTO).Error(), //解析具体的错误类型
		})
	}
	OK(ctx, ResponseJson{
		Data: iUserLoginDTO,
	})
}

// 根据不同的错误类型tag匹配响应的错误响应信息message
func parseValidateErrors(errs validator.ValidationErrors, target any) error {
	var errResult error

	//通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprint("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprint("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = util.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}
