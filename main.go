package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello-world", GetHelloWord())
	r.Run(":8080")
}

func GetHelloWord() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":       "Hello World",
			"user":      "Name",
			"image":     "true",
			"something": "true",
			"line":      "demo",
		})
	}
}
