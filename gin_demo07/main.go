package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)



func main() {
	// 返回默认的路由引擎
	// 可以这样理解，就像 flask 里面的那个 app 对象，然后都往这个 app 上去注册路由
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK,"index.html",nil)


	})



	r.POST("/upload", func(c *gin.Context) {


		//从请求中读取文件
		f,err := c.FormFile("f1")

		if err != nil {

			c.JSON(http.StatusBadRequest,gin.H{

				"error":err.Error(),
			})

		}else {
			//拼接文件保存路径的两种方式
			//dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./",f.Filename)

			//将读取的文件保存在本地
			c.SaveUploadedFile(f,dst)

			c.JSON(http.StatusOK,gin.H{

				"status":"ok",

			})
		}


	})


	r.Run()


}
