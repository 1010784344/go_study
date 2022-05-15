package dao

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

// 定义一个全局的 DB 对象
var DB *gorm.DB


// 定义一个函数连接数据库
func InitMySQL()(err error){

	dsn := "root:123456@(127.0.0.1)/db1?charset=utf8"
	// 连接数据库
	DB,err = gorm.Open("mysql", dsn)

	if err!= nil{

		return
	}
	// 测试是否连接成功
	return DB.DB().Ping()

}

func Close()  {
	// 程序退出关闭数据库连接
	DB.Close()
}
