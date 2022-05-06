package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)



func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()


	r.GET("/user/:name/:age", func(c *gin.Context) {

		name := c.Param("name")
		age := c.Param("age")

		//响应返回 json
		c.JSON(http.StatusOK,gin.H{

			"name":name,
			"age":age,

		})

	})


	r.Run()


}
