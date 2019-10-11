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
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginJson", func(c *gin.Context) {
		var json Login
		//传入的参数需要是application/json/满足校验条件
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//校验密码
		if json.User != "test" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "your are logged in",
		})
	})

	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//校验密码
		if form.User != "test" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "your are logged in",
		})

	})

	_ = r.Run(":9000")

}
