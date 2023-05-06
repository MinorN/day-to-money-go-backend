package controller_test

import (
	"context"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	r := router.New()
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
