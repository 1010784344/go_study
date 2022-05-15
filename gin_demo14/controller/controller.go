package controller


import (
	"fmt"
	"gin_demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",nil)
}


func CreateTodo(c *gin.Context) {

	// 从请求中把数据取出来
	var todo models.Todo
	c.BindJSON(&todo)

	// 存入数据库
	err := models.CreateATodo(&todo)
	if err!=nil{

		c.JSON(http.StatusOK,gin.H{"error":err.Error()})

	}else {

		c.JSON(http.StatusOK,todo)

	}


}



func GetTodoList(c *gin.Context) {

	// 创建一个切片对象

	todolist, err := models.GetAllTodo()

	if err != nil{

		c.JSON(http.StatusOK,gin.H{"err":err.Error()})
	}else {

		c.JSON(http.StatusOK,todolist)
	}

	fmt.Println("hello world")

}


func UpdateATodo(c *gin.Context) {

	// 获取id
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"id 不存在"})
		return

	}

	// 根据 id 获取数据库信息
	todo,err := models.GetATodo(id)
	if err != nil{

		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}

	// 根据提交的 json 信息将数据库信息覆盖
	c.BindJSON(&todo)

	// 将数据库信息保存到数据库
	err = models.UpdateATodo(todo)
	if err != nil{

		c.JSON(http.StatusOK,gin.H{

			"error":err.Error()})
	}else {

		c.JSON(http.StatusOK,todo)

	}
}



func DeleteATodo(c *gin.Context) {

	// 获取id
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"id 不存在"})
		return

	}

	// 从数据库找到对应的数据并删除
	err := models.DeleteATodo(id)
	if err != nil{

		c.JSON(http.StatusOK,gin.H{
			"error":err.Error()})
	}else {

		c.JSON(http.StatusOK,gin.H{
			id:"deleted"})
	}


}






