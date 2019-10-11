package main

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {

	r := gin.Default()

	r.POST("/profile", func(c *gin.Context) {
		var form ProfileForm
		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}

		c.String(http.StatusOK, "ok")
	})
	_ = r.Run(":8888")
	//➜  Desktop curl -X POST -v --form name=user --form "avatar=@./WechatIMG294.jpeg" http://localhost:8888/profile
	//Note: Unnecessary use of -X or --request, POST is already inferred.
	//*   Trying ::1...
	//* TCP_NODELAY set
	//* Connected to localhost (::1) port 8888 (#0)
	//> POST /profile HTTP/1.1
	//> Host: localhost:8888
	//> User-Agent: curl/7.54.0
	//> Accept: */*
	//> Content-Length: 409765
	//> Expect: 100-continue
	//> Content-Type: multipart/form-data; boundary=------------------------271f465146f0a47c
	//>
	//< HTTP/1.1 100 Continue
	//< HTTP/1.1 200 OK
	//< Content-Type: text/plain; charset=utf-8
	//< Date: Fri, 11 Oct 2019 11:53:05 GMT
	//< Content-Length: 2
	//<
	//* Connection #0 to host localhost left intact
	//ok%

}
