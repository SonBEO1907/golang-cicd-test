package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/display", GetHelloWorld())
	r.Run(":8080")
}

func GetHelloWorld() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "Hello World!!!",
		})
	}
}
