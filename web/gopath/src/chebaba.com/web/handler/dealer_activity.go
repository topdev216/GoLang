package handler

import (
	"github.com/gin-gonic/gin"
)

func GetActivityList(c *gin.Context) {
	c.JSON(200, gin.H{"result": "OK"})
}

func GetActivityById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{"result": id})
}
