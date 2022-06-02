package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"goWebF/src/blog/service"
	"io"
	"net/http"
	"os"
)

var r *gin.Engine = gin.Default()

func init() {
	/*	dir, err := ioutil.ReadDir("/template")
		fmt.Println(dir,err)*/
	//批量处理所有的页面以及静态资源
	r.LoadHTMLGlob("template/*.html")

	r.StaticFS("/static", http.Dir("template/static"))
}
func main() {

	r.GET("/index", serviceindex.Index)
	r.GET("/about", serviceindex.About)
	r.GET("/list", serviceindex.List)
	r.GET("/gbook", serviceindex.Gbook)
	r.GET("/share", serviceindex.Share)
	r.GET("/info", serviceindex.Info)
	r.GET("/login", func(context *gin.Context) {
		context.HTML(200, "login.html", nil)
	})

	r.POST("/signin", func(context *gin.Context) {
		file, err := context.FormFile("files")
		if err != nil {
			fmt.Println("读取文件失败")
		}
		println(file.Filename)
		open, _ := file.Open()

		create, _ := os.Create("src/files/" + file.Filename)
		writer := bufio.NewWriter(create)
		reader := bufio.NewReader(open)
		io.Copy(writer, reader)

	})

	r.Run(":8080")

}
