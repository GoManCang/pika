package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/JsonP", func(c *gin.Context) {
		data := gin.H{
			"ping": "pong",
		}

		c.JSONP(http.StatusOK, data)
	})

	_ = r.Run(":9001")

	/*
		➜  ~ http http://localhost:9001/JsonP\?callback\=xcc
		HTTP/1.1 200 OK
		Content-Length: 20
		Content-Type: application/javascript; charset=utf-8
		Date: Sat, 12 Oct 2019 11:36:36 GMT

		xcc({"ping":"pong"})

		➜  ~
	*/

}
