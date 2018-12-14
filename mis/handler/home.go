package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	data := make(map[string]string)
	data["hello"] = "world"
	c.HTML(http.StatusOK, "index.html", data)
}

func AddRecordHandler(c *gin.Context) {
	// 添加数据成功
	c.JSON(http.StatusOK, gin.H{"result": 0})
}

func DelRecordHandler(c *gin.Context) {
	pass := c.Params.ByName("pass")
	if pass && pass == "adminn" {
		// 输入密码后，删除成功
		c.JSON(http.StatusOK, gin.H{"status": 0})
	}
}
