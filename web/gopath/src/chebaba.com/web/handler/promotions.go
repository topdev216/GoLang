package handler

import (
	"github.com/gin-gonic/gin"
)

func GetPromotionsList(c *gin.Context) {
	c.String(200, "GetPromotionsList")
}

func GetPromotionsById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.String(200, id)
}
