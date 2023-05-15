package controller

import (
	"context"
	"mangosteen/config"
	"mangosteen/internal/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {

	r = gin.Default()
	config.LoadViperConfig()
	database.Connect()

	vcc := SessionController{}
	vcc.RegisterRoutes(r.Group("/api"))

	email := "test@test.com"
	c := context.Background()
	q := database.NewQuery()
	count1, _ := q.CountValidationCodes(c, email)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation_codes",
		strings.NewReader(`{"email":"`+email+`"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(c, email)
	assert.Equal(t, count2-1, count1)
}
