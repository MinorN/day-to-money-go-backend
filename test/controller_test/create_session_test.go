package controller_test

import (
	"context"
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	r *gin.Engine
	q *queries.Queries
	c context.Context
)

func setupTest(t *testing.T) func(t *testing.T) {
	q = database.NewQuery()
	r = router.New()
	c = context.Background()
	if err := q.DeleteAllUsers(c); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		database.DB.Close()
	}
}

func TestCreateSession(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	email := "1@qq.com"
	code := "1234"
	q := database.NewQuery()
	_, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: email,
		Code:  code,
	})

	if err != nil {
		log.Fatalln(err)
	}
	user, err := q.CreateUser(c, email)
	if err != nil {
		log.Fatalln(err)
	}
	w := httptest.NewRecorder()
	x := gin.H{"email": email, "code": code}
	bytes, _ := json.Marshal(x)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/session",
		strings.NewReader(string(bytes)),
	)
	r.ServeHTTP(w, req)
	var responseBody struct {
		JWT    string `json:"jwt"`
		UserID int32  `json:"userId"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error("jwt is not a string")
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, user.ID, responseBody.UserID)
}
