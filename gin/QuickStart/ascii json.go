package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/sj", func(c *gin.Context) {
		data := gin.H{
			"key":   "GO 语言",
			"value": "language",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	_ = r.Run(":9001")
	/*
			➜  ~ http http://localhost:9001/sj
			HTTP/1.1 200 OK
			Content-Length: 44
			Content-Type: application/json
			Date: Sat, 12 Oct 2019 11:39:18 GMT

			{
			    "key": "GO 语言",
			    "value": "language"
			}

			➜  ~

		{"key":"GO \u8bed\u8a00","value":"language"}

	*/

}
