package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)



func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	//加载静态文件
	r.LoadHTMLFiles("./login.html","./index.html")

	r.GET("/login", func(c *gin.Context) {

		//响应返回 html
		c.HTML(http.StatusOK,"login.html",nil)

	})

	r.POST("/login", func(c *gin.Context) {

		//方式一
		//username := c.PostForm("username")
		//password := c.PostForm("password")

		//方式二(取不到，有默认值)
		username := c.DefaultPostForm("sername","anyone")
		password := c.DefaultPostForm("assword","99")

		//响应返回 html 并传参
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Username":username,
			"Password":password,
		})

	})

	r.Run()


}
