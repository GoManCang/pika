package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	// 带目录的
	r.Static("/assets", "./assets")
	//文件系统
	r.StaticFS("/more_static", http.Dir("/Users/li/tool"))
	// 绝对文件路径
	r.StaticFile("/favicon.jpeg", "./resource/favicon.jpeg")
	_ = r.Run(":9999")
	/*
		http://localhost:9999/favicon.jpeg

		http://localhost:9999/more_static/

		http://localhost:9999/assets/index/favicon.jpeg
	*/
}
