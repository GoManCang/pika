package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", posting)

	//get传参数
	r.GET("/welcome", func(context *gin.Context) {
		firstName := context.DefaultQuery("firstname", "Guest")
		lastName := context.Query("lastname")
		context.String(http.StatusOK, "hello %s %s", firstName, lastName)
	})

	err := r.Run(":8080")
	if err == nil {
		fmt.Println("启动正常")
	} else {
		fmt.Println("启动异常：", err)
	}

	//	Gin uses encoding/json as default json package but you can change to jsoniter by build from other tags.
	//	go build -tags=jsoniter .
}

func posting(context *gin.Context) {
	fmt.Println("post")
}
