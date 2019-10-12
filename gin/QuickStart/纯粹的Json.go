package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>hello ,Go</b>",
		})
	})

	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello ,Go</b>",
		})
	})

	/*
			➜  ~ http http://localhost:9001/json
			HTTP/1.1 200 OK
			Content-Length: 47
			Content-Type: application/json; charset=utf-8
			Date: Sat, 12 Oct 2019 11:45:44 GMT

			{
			    "html": "<b>hello ,Go</b>"
			}
		{"html":"\u003cb\u003ehello ,Go\u003c/b\u003e"}

			➜  ~ http http://localhost:9001/purejson
			HTTP/1.1 200 OK
			Content-Length: 28
			Content-Type: application/json; charset=utf-8
			Date: Sat, 12 Oct 2019 11:45:56 GMT

			{
			    "html": "<b>hello ,Go</b>"
			}
		{"html":"<b>hello ,Go</b>"}
			➜  ~
	*/

	_ = r.Run(":9001")

}
