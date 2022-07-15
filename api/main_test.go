package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type FishObj struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Unit  string `json:"unit"`
}

func TestGetallfish(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/fish", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response []FishObj
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetfish(t *testing.T) {
	expectedbody := gin.H{
		"name":  "鮭魚",
		"price": "500",
		"unit":  "斤",
	}
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/fish/salmon", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response FishObj
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, expectedbody["name"], response.Name)
	assert.Equal(t, http.StatusOK, w.Code)
}
