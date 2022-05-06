package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	)

type UserInfo struct {

	Username string `form:"username"`
	Password string `form:"password"`

}


func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	//127.0.0.1:8080/user?username=yang&password=18
	r.GET("/user", func(c *gin.Context) {

		var u UserInfo//声明一个 UserInfo类型的变量 u

		err := c.ShouldBind(&u)

		if err != nil {

			c.JSON(http.StatusOK,gin.H{

				"error":err.Error(),
			})

		}else {
			//打印结构体的值
			fmt.Printf("%#v\n",u)

			c.JSON(http.StatusOK,gin.H{

				"status":"ok",


			})
		}



	})

	//接收表单数据，可以使用 postman 来进行测试，body -- form-data
	r.POST("/user", func(c *gin.Context) {

		var u UserInfo

		err := c.ShouldBind(&u)

		if err != nil {

			c.JSON(http.StatusOK,gin.H{

				"error":err.Error(),
			})

		}else {
			//打印结构体的值
			fmt.Printf("%#v\n",u)

			c.JSON(http.StatusOK,gin.H{

				"status":"ok",


			})
		}



	})

	//接收json数据，可以使用 postman 来进行测试，body -- raw -- json
	r.POST("/json", func(c *gin.Context) {

		var u UserInfo

		err := c.ShouldBind(&u)

		if err != nil {

			c.JSON(http.StatusOK,gin.H{

				"error":err.Error(),
			})

		}else {
			//打印结构体的值
			fmt.Printf("%#v\n",u)

			c.JSON(http.StatusOK,gin.H{

				"status":"ok",


			})
		}



	})


	r.Run()


}
