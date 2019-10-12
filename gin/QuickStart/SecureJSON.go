package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		names := []string{"lena", "asta", "foot"}

		c.SecureJSON(http.StatusOK, names)
	})

	_ = r.Run(":9000")

	/*
		➜  ~ http http://localhost:9000/someJson
		HTTP/1.1 200 OK
		Content-Length: 31
		Content-Type: application/json; charset=utf-8
		Date: Sat, 12 Oct 2019 11:30:20 GMT

		while(1);["lena","asta","foot"]

		➜  ~
	*/
}
