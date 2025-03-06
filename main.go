package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/display/:number", GetHelloWorld())
	r.Run(":8080")
}

func GetHelloWorld() func(*gin.Context) {
	return func(ctx *gin.Context) {
		numberParam := ctx.Param("number")
		if numberParam == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "nil",
			})
		}
		number, err := strconv.Atoi(number)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "format",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "Hello World!!!",
			"Answer": Square(number),
		})
	}
}

func Square(number int) int {
	return number * number
}