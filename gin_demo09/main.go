package main

import (

	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	//get 基本方法
	r.GET("/index", func(c *gin.Context) {

		c.JSON(http.StatusOK,gin.H{"index":"GET"})

	})

	//post 基本方法
	r.POST("/index", func(c *gin.Context) {

		c.JSON(http.StatusOK,gin.H{"index":"POST"})
	})


	// 请求方法的大集合
	r.Any("/any", func(c *gin.Context) {

		switch c.Request.Method {

		//case "GET":
		case http.MethodGet:
			c.JSON(http.StatusOK,gin.H{"any":"GET"})
		//case "POST":
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{"any":"POST"})

		}

	})


	// 访问了不存在地址的请求
	r.NoRoute(func(c *gin.Context) {

		c.JSON(http.StatusNotFound,gin.H{"meg":"你所访问的页面不存在！"})
	})


	//给一组路由设置公共前缀
	videoGroup := r.Group("/video")
	{

		videoGroup.GET("/index",func(c *gin.Context) {

			c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
		})

		videoGroup.GET("/first",func(c *gin.Context) {

			c.JSON(http.StatusOK,gin.H{"msg":"/video/first"})
		})

		videoGroup.GET("/last",func(c *gin.Context) {

			c.JSON(http.StatusOK,gin.H{"msg":"/video/last"})
		})

	}



	r.Run(":8081")


}
