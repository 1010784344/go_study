package main


import (
	"gin_demo/dao"
	"gin_demo/models"
	"gin_demo/routers"


)


func main() {

	//连接数据库
	err := dao.InitMySQL()

	if err!=nil{

		panic(err)

	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 程序退出关闭数据库连接
	defer dao.Close()

	// 路由相关的信息
	r:=routers.SetupRouter()

	r.Run()


}


