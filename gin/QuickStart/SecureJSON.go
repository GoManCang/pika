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
}
