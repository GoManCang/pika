package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {

	route := gin.Default()

	route.Any("/testing", startPage)

	_ = route.Run(":9999")
	/*
		http://localhost:9999/testing?name=name&address=%E5%8C%97%E4%BA%AC
	*/

}

func startPage(context *gin.Context) {
	var person Person
	if context.ShouldBindQuery(&person) == nil {
		log.Println("=========bind query string====")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	context.JSON(200, gin.H{
		"status": "Success",
	})
}
