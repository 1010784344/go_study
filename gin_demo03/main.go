package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)


//http://127.0.0.1:8080/web?name=yang&age=18

func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()


	r.GET("/web", func(c *gin.Context) {

		//方式一
		//name := c.Query("name")
		//age := c.Query("age")

		//方式二(取不到，有默认值)
		name := c.DefaultQuery("name","anyone")
		age := c.DefaultQuery("age","99")

		c.JSON(http.StatusOK,gin.H{

			"name":name,
			"age":age,

		})

	})


	r.Run()


}
