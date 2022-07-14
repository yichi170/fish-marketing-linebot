package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetallfish(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/fish", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response []map[string]interface{}
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

	// responseData, _ := ioutil.ReadAll(w.Body)
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["name"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedbody["name"], value)
	assert.Equal(t, http.StatusOK, w.Code)
}
