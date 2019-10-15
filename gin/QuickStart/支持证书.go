package main

import (
	"github.com/gin-gonic/gin"
)

/*
https://github.com/gin-gonic/gin#support-lets-encrypt
*/
func main() {

	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 废弃

}
