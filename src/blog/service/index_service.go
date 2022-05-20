package serviceindex

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Info 详情
func Info(c *gin.Context) {
	c.HTML(http.StatusOK, "info.html", nil)
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}

func Gbook(c *gin.Context) {
	c.HTML(http.StatusOK, "gbook.html", nil)
}

func List(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", nil)
}

func Share(c *gin.Context) {
	c.HTML(http.StatusOK, "share.html", nil)
}
