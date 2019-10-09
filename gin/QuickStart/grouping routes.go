package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login %s", "ok")
		})

	}

	v2 := router.Group("/v2")

	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login v2 %s", "ok")
		})
	}

	_ = router.Run(":8080")

}
