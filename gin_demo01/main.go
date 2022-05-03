package main

import (
	"github.com/gin-gonic/gin"
	)


func sayhello(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"hello golang!",
	})


}


func main() {

	r := gin.Default()

	//指定用户使用 GET 请求访问 /hello 时，执行 sayhello 这个函数
	r.GET("/hello", sayhello)

	//启动服务
	r.Run()
	
}
