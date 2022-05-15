package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)


// Todo model
type Todo struct {

	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`

}

// 定义一个全局的 DB 对象
var DB *gorm.DB


// 定义一个函数连接数据库
func initMySQL()(err error){

	dsn := "root:123456@(127.0.0.1)/db1?charset=utf8"
	// 连接数据库
	DB,err = gorm.Open("mysql", dsn)

	if err!= nil{

		return
	}
	// 测试是否连接成功
	return DB.DB().Ping()

}


func main() {

	//连接数据库
	err := initMySQL()

	if err!=nil{

		panic(err)

	}

	// 模型绑定
	DB.AutoMigrate(&Todo{})

	// 程序退出关闭数据库连接
	defer DB.Close()

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
		v1Group.POST("/todo", func(c *gin.Context) {

			// 从请求中把数据取出来
			var todo Todo
			c.BindJSON(&todo)

			// 存入数据库
			err = DB.Create(&todo).Error
			if err!=nil{

				c.JSON(http.StatusOK,gin.H{"error":err.Error()})

			}else {

				c.JSON(http.StatusOK,todo)

			}


		})

		//查询所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

			// 创建一个切片对象
			var todolist []Todo

			err = DB.Find(&todolist).Error

			if err != nil{

				c.JSON(http.StatusOK,gin.H{"err":err.Error()})
			}else {

				c.JSON(http.StatusOK,todolist)
			}

			fmt.Println("hello world")

		})

		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		// 修改某一个待办事项
		// 前端传过来两个信息：id 和 提交的修改信息（json 格式）
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

			// 获取id
			id,ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK,gin.H{"error":"id 不存在"})
				return

			}

			// 根据 id 获取数据库信息
			var todo Todo
			err = DB.Where("id=?",id).First(&todo).Error
			if err != nil{

				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
				return
			}

			// 根据提交的 json 信息将数据库信息覆盖
			c.BindJSON(&todo)

			// 将数据库信息保存到数据库
			err = DB.Save(&todo).Error
			if err != nil{

				c.JSON(http.StatusOK,gin.H{

					"error":err.Error()})
			}else {

				c.JSON(http.StatusOK,todo)

			}
		})

		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

			// 获取id
			id,ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK,gin.H{"error":"id 不存在"})
				return

			}

			// 从数据库找到对应的数据并删除
			err = DB.Where("id=?",id).Delete(Todo{}).Error
			if err != nil{

				c.JSON(http.StatusOK,gin.H{
					"error":err.Error()})
			}else {

				c.JSON(http.StatusOK,gin.H{
					id:"deleted"})
			}


		})

	}

	r.Run()




}


