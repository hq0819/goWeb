package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

type User struct {
	name  string
	age   int
	email string
}

func main() {
	r := gin.Default()
	pra := func(name string) string {
		return name + ", hello world"
	}

	r.GET("/index", func(context *gin.Context) {
		r.LoadHTMLGlob("src/templates/*")
		r.SetFuncMap(template.FuncMap{"pra": pra})
		context.HTML(200, "main.html", User{"heqin", 19, "23423"})
	})

	r.Run(":8080")
}
