package main

import (

	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	// /index --> sogo.com
	r.GET("/index", func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently,"http://sogo.com")


	})


	//  /a --> /b
	r.GET("/a", func(c *gin.Context) {

		c.Request.URL.Path = "/b"//把请求的 url 修改
		r.HandleContext(c)//继续后续的处理

	})

	r.GET("/b", func(c *gin.Context) {

		c.JSON(http.StatusOK,gin.H{

				"message":"b",
		})

	})


	r.Run()


}
