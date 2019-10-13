package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		cCP := c.Copy()

		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCP.Request.URL.Path)
		}()
	})

	r.GET("long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done!,in path" + c.Request.URL.Path)
	})

	_ = r.Run(":9999")

}
