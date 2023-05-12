package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
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
	user, err := q.FindByEmail(c, requestBody.Email)
	if err != nil {
		user, err = q.CreateUser(c, requestBody.Email)
		if err != nil {
			log.Println("CreateUser fail", err)
			c.String(http.StatusInternalServerError, "请稍后再试")
		}
	}
	// 生成 jwt
	jwt, err := jwt_helper.GenerateJWT(int(user.ID))
	if err != nil {
		log.Println("GenerateJWT fail", err)
		c.String(http.StatusInternalServerError, "请稍后再试")
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"jwt":    jwt,
			"userID": user.ID,
		},
	)
}
