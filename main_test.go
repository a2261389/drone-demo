package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, r)
	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	if string(body) != "OK" {
		log.Fatal("body is not OK")
	}
}
