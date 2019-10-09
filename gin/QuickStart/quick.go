package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

	//接受表单
	r.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultQuery("nick", "anonymous")

		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	/*
		POST /post/test?id=1234&page=1 HTTP/1.1
		Content-Type: application/x-www-form-urlencoded

		name=manu&message=this_is_great
	*/
	r.POST("/post/test", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Printf("id:%s;page:%s;name:%s;,message:%s", id, page, name, message)

	})

	/*
		POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
		Content-Type: application/x-www-form-urlencoded

		names[first]=thinkerou&names[second]=tianou
	*/

	r.POST("/post/arr/obj", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		fmt.Printf("ids:%v;names:%v", ids, names)
	})

	/*
			文件上传
		curl -X POST http://localhost:8080/upload \
		  -F "file=@/Users/appleboy/test.zip" \
		  -H "Content-Type: multipart/form-data"
	*/

	r.POST("upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s upload!", file.Filename))
	})

	err := r.Run(":8080")
	if err == nil {
		fmt.Println("启动正常")
	} else {
		fmt.Println("启动异常：", err)
	}

	/**
	//多文件上传
	curl -X POST http://localhost:8080/upload \
	  -F "upload[]=@/Users/appleboy/test1.zip" \
	  -F "upload[]=@/Users/appleboy/test2.zip" \
	  -H "Content-Type: multipart/form-data"
	*/

	r.POST("/multi/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploads", len(files)))
	})

	//	Gin uses encoding/json as default json package but you can change to jsoniter by build from other tags.
	//	go build -tags=jsoniter .
}

func posting(context *gin.Context) {
	fmt.Println("post")
}
