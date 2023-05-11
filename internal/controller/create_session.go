package controller

import (
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var requestBody struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}
	q := database.NewQuery()
	_, err := q.FindValidationCode(c, queries.FindValidationCodeParams{
		Email: requestBody.Email,
		Code:  requestBody.Code,
	})
	if err != nil {
		c.String(http.StatusBadRequest, "无效验证码")
		return
	}
	jwt := "xxx"
	c.JSON(
		http.StatusOK,
		gin.H{
			jwt: jwt,
		},
	)
}
