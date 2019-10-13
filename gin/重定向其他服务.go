package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("dataFromOtherResponse", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:9999/favicon.jpeg")
		if err != nil || resp.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := resp.Body
		conLen := resp.ContentLength
		contentType := resp.Header.Get("Content-Type")

		extraHeader := map[string]string{
			"Content-Disposition": `attachment;filename=hpp.png`,
		}

		c.DataFromReader(http.StatusOK, conLen, contentType, reader, extraHeader)
	})
	_ = router.Run(":9090")
	/*
		http://localhost:9090/dataFromOtherResponse
	*/
}
