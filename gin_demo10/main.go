package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


// 响应函数
func indexhandler(c *gin.Context){

	fmt.Println("index start")

	c.JSON(http.StatusOK,gin.H{"message":"ok"})

	fmt.Println("index end")
}

// 中间件1
func m1(c *gin.Context){

	fmt.Println("m1 start")

	start := time.Now()

	c.Next()//调用后续的处理函数

	count := time.Since(start)

	fmt.Printf("%v\n",count)

	fmt.Println("m1 end")
}

// 中间件1
func m2(c *gin.Context){

	fmt.Println("m2 start")

	c.Next()//调用后续的处理函数

	fmt.Println("m2 end")
}


func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	//get 基本方法
	r.GET("/index", m1,indexhandler)


	r.Run(":8081")


}


// 多个中间件的调用
//func main() {
//	// 返回默认的路由引擎
//	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
//	r := gin.Default()
//
//	//get 基本方法
//	r.GET("/index", m1,m2,indexhandler)
//
//
//	r.Run(":8081")
//
//
//}