package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-rxy/app/http/controllers/api/v1"
	"go-rxy/app/models/user"
	"go-rxy/app/requests"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析 JSON 对象
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
