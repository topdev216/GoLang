package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func HomeHandler(c *gin.Context) {
	data := make(map[string]string)
	data["hello"] = "world"
	c.JSON(http.StatusOK, data)
}

func HtmlHandler(c *gin.Context) {
	data := make(map[string]string)
	data["title"] = "HTML模板"
	data["text"] = "一段测试文本"
	c.HTML(http.StatusOK, "index.tmpl", data)
}

func ParamHandler(c *gin.Context) {
	data := make(map[string]string)
	data["title"] = "URL中的参数"
	data["text"] = c.Params.ByName("value")
	c.HTML(http.StatusOK, "index.tmpl", data)
}

func QueryHandler(c *gin.Context) {
	data := make(map[string]string)
	data["title"] = "URL查询参数"
	data["text"] = c.DefaultQuery("value", "默认值")
	c.HTML(http.StatusOK, "index.tmpl", data)
}

func PostHandler(c *gin.Context) {
	data := make(map[string]string)
	data["title"] = "POST测试"
	data["text"] = c.DefaultPostForm("value", "默认值")
	c.HTML(http.StatusOK, "index.tmpl", data)
}

func Uploadhandler(c *gin.Context) {
	data := make(map[string]string)
	data["title"] = "文件上传"
	file, header, err := c.Request.FormFile("file")
	data["text"] = header.Filename
	out, err := os.Create("/tmp/" + data["text"])
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "index.tmpl", data)
}
