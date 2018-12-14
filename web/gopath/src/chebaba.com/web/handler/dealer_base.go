package handler

import (
	"github.com/gin-gonic/gin"
)

func GetDealerList(c *gin.Context) {
	c.String(200, "GetDealerList")
}

func GetDealerById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.String(200, id)
}
