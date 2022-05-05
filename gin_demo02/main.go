package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)




func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	//指定用户使用 GET 请求访问 /json 时，匿名函数进行响应
	r.GET("/json", func(c *gin.Context) {

		//方法一：使用 map
		data1 := map[string]interface{}{

			"name":"小王子",
			"message":"hello world",
			"age":18,
		}

		c.JSON(http.StatusOK,data1)

	})


	r.GET("/json_map", func(c *gin.Context) {

		//方法一：使用 gin 封装好的 map
		data2 := gin.H{"name":"小王子", "message":"hello world", "age":18}

		c.JSON(http.StatusOK,data2)

	})


	type msg struct {
		Name string
		Message string
		Age int
	}


	r.GET("/struct", func(c *gin.Context) {

		//方法二：使用结构体
		data3 := msg{
			"小王子",
			"hello world",
			18,
		}

		c.JSON(http.StatusOK,data3)

	})



	r.Run()


}
