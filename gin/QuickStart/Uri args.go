package main

import "github.com/gin-gonic/gin"

type PersonOne struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person PersonOne

		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"name": person.Name,
			"uuid": person.ID,
		})
	})

	_ = r.Run(":8088")
	/*
		➜  ~ http localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
		HTTP/1.1 200 OK
		Content-Length: 66
		Content-Type: application/json; charset=utf-8
		Date: Fri, 11 Oct 2019 11:32:00 GMT

		{
		    "name": "thinkerou",
		    "uuid": "987fbc97-4bed-5078-9f07-9141ba07c9f3"
		}

		➜  ~ http localhost:8088/thinkerou/987fbc97-4bed-5078-9f07
		HTTP/1.1 400 Bad Request
		Content-Length: 86
		Content-Type: application/json; charset=utf-8
		Date: Fri, 11 Oct 2019 11:32:17 GMT

		{
			"msg": "Key: 'PersonOne.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"
		}

		➜  ~

	*/

}
