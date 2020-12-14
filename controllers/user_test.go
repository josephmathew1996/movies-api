package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"movies-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestAuthHandler(t *testing.T) {
	expectedData := models.User{Email: "josephmathew1401@gmail.com", Password: "123qweasd"}
	b, _ := json.Marshal(expectedData)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/v1/auth", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")
	user.Auth(w, r)
	resp, _ := ioutil.ReadAll(w.Body)
	actual := struct {
		Status     string           `json:"status"`
		StatusCode int              `json:"statusCode"`
		Message    string           `json:"message"`
		Data       models.AuthToken `json:"data"`
	}{}
	json.Unmarshal(resp, &actual)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, 1, actual.Data.UserDetail.ID)
	assert.Equal(t, "Joseph Mathew", actual.Data.UserDetail.Name)
	assert.Equal(t, expectedData.Email, actual.Data.UserDetail.Email)
	assert.Equal(t, "admin", actual.Data.UserDetail.Role)
	assert.Equal(t, "Bearer", actual.Data.TokenType)
}
