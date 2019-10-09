package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	//gin.ForceConsoleColor()
	gin.DisableConsoleColor()
	router := gin.New()
	router.Use(gin.Logger())
	//恢复中间件可从任何紧急情况中恢复，如果有，则写入500。
	router.Use(gin.Recovery())

	//自定义日志打印
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	g1 := router.Group("/")
	g1.Use(AuthRquired())
	{
		g1.GET("/login", func(c *gin.Context) {
			//panic("hello")
			c.String(http.StatusOK, "hello %s", "1/0.0")
		})
	}

	_ = router.Run(":8080")

}

func AuthRquired() gin.HandlerFunc {

	//认证逻辑
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

func benchEndpoint(context *gin.Context) {
	fmt.Println("bench endpoint")
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func MyBenchLogger() gin.HandlerFunc {
	fmt.Println("bench logger")
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
