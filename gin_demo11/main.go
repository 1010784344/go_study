package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)


type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}


func main() {

	db,err := gorm.Open("mysql","root:123456@(127.0.0.1:3306)/db1?charset=utf8")

	if err !=nil{

		panic(err)
	}

	defer db.Close()

	//创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//插入数据
	u1 := UserInfo{1,"Rose","male","study"}
	db.Create(&u1)

	// 查询数据：First 查询数据库里面的第一条数据
	var u UserInfo
	db.First(&u)
	fmt.Printf("u: %v\n",u)

	// 修改数据
	db.Model(&u).Update("Hobby","work")

	fmt.Printf("u: %v\n",u)

	// 删除数据
	//db.Delete(&u)



















}


