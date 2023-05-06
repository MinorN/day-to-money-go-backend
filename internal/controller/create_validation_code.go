package controller

import (
	"log"
	"mangosteen/internal/email"

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
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.String(400, "参数错误")
		return
	}
	if err := email.SendValidationCode(body.Email, "123456"); err != nil {
		log.Print("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	c.Status(200)
}
