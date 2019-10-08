package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err == nil {
		fmt.Println("启动正常")
	} else {
		fmt.Println("启动异常：", err)
	}
}
