package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/user"
	"net/http"
)

// HandleUserCheckError 处理用户校验错误
func HandleUserCheckError(ctx *gin.Context, err error) (string, bool) {
	// 自定义错误信息
	var validationErrors validator.ValidationErrors
	if ok := errors.As(err, &validationErrors); !ok {
		// Handle non-validation errors here
		msg := err.Error()

		ctx.JSON(http.StatusOK, user.RegisterResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})

		return "", true
	}

	// 将错误信息拼接成字符串
	var errorMessages string
	for _, fieldErr := range validationErrors {
		errorMessages += CustomErrorMessage(fieldErr) + "; "
	}
	return errorMessages, false
}

func CustomErrorMessage(err validator.FieldError) string {
	switch err.Field() {
	case "Username":
		switch err.Tag() {
		case "required":
			return "请提供用户名（邮箱）."
		case "email":
			return "请提供有效的邮箱地址（用户名）."
		}
	case "Password":
		switch err.Tag() {
		case "required":
			return "请提供密码."
		case "min":
			return "密码长度至少6位."
		case "alphanum":
			return "密码只能包含字母和数字."
		}
	}

	return err.Error()
}
