package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const apiVersion = "1.0.0"

type responseError struct {
	code    int
	message string
	errors  interface{}
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/users/:id", func(c *gin.Context) {
		if param := c.Param("id"); param != "" {
			content, err := ioutil.ReadFile("./data/data.json")
			if err != nil {
				res := handleFailResponse(c, 500, "can not open file", nil)
				c.JSON(500, gin.H(res))
				return
			}

			users := []map[string]interface{}{}
			json.Unmarshal(content, &users)
			result := map[string]interface{}{}

			// search user
			for _, user := range users {
				id := strconv.Itoa(int(user["id"].(float64)))
				if id == param {
					result = user
				}
			}
			if _, ok := result["id"]; !ok {
				res := handleFailResponse(c, 404, "user not found", nil)
				c.JSON(404, gin.H(res))
				return
			}
			res := handleSuccessResponse(c, result)
			c.JSON(200, gin.H(res))
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello docker"})
	})
	r.Run()
}

func handleSuccessResponse(c *gin.Context, data interface{}) map[string]interface{} {
	response := gin.H{
		"apiVersion": apiVersion,
		"data":       data,
	}

	return response
}

func handleFailResponse(c *gin.Context, httpCode int, message string, errors interface{}) map[string]interface{} {
	response := gin.H{
		"apiVersion": apiVersion,
		"error": gin.H{
			"code":    httpCode,
			"message": message,
			"errors":  errors,
		},
	}
	if errors == nil || errors == "" {
		delete(response["error"].(gin.H), "errors")
	}

	return response
}
