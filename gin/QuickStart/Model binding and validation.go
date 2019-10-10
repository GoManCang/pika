package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
https://github.com/gin-gonic/gin#model-binding-and-validation


https://github.com/go-playground/validator


*/

type Login struct {
	User     string `from:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginJson", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	})

}
