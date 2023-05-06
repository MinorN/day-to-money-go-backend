package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

// CreateValidationCode
// @Summary      用来发送邮箱验证码
// @Description  接受邮箱地址，发送验证码
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /validation_codes [get]
func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string
	}
	err := c.ShouldBind(&body)
	if err != nil {
		c.String(400, "参数错误")
		return
	}
	log.Println(body.Email)
}
