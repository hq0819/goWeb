package main

import (
	"github.com/gin-gonic/gin"
	"goWebF/src/blog/service"
	"net/http"
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

	r.Run(":8080")

}
