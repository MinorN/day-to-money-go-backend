package controller

import (
	"crypto/rand"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
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
	len := 4
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("[rand.Read fail]", err)
		c.String(500, "发送失败")
		return
	}
	digits := make([]byte, len)
	for i := range b {
		// 需要 '1234' = ['1','2','3','4']
		// 0 对应的编码是 48
		digits[i] = b[i]%10 + 48
	}
	str := string(digits)
	q := database.NewQuery()
	vc, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: body.Email,
		Code:  str,
	})
	if err != nil {
		// TODO 没有校验
		c.Status(400)
		return
	}
	if err := email.SendValidationCode(vc.Email, vc.Code); err != nil {
		log.Print("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	c.Status(200)
}
