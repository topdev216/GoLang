package handler

import (
	"github.com/gin-gonic/gin"
)

func GetPriceList(c *gin.Context) {
	c.String(200, "GetPromotionsList")
}

func GetPriceByDealerSeries(c *gin.Context) {
	dealer := c.Params.ByName("dealer")
	series := c.Params.ByName("series")[1:]
	c.String(200, dealer+", "+series)
}
