package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)


// Todo model
type Todo struct {

	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`

}


func main() {

	//连接数据库


	r := gin.Default()

	// 告诉 gin 框架模板文件引用的静态文件（css，js）去哪里找
	r.Static("/static", "static")

	// 告诉 gin 框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/tod", func(c *gin.Context) {

		})

		//查询所有的待办事项
		v1Group.GET("/tod", func(c *gin.Context) {

		})

		// 查看某一个待办事项
		v1Group.GET("/tod/:id", func(c *gin.Context) {

		})

		// 修改某一个待办事项
		v1Group.PUT("/tod/:id", func(c *gin.Context) {

		})

		// 删除某一个待办事项
		v1Group.DELETE("/tod/:id", func(c *gin.Context) {

		})

	}

	r.Run()




}


