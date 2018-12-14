package main

import (
	"chebaba.com/web/handler"
	"github.com/gin-gonic/gin"
)

func reg(r *gin.Engine) {
	r.GET("/", handler.Index)

	dealerActivity := r.Group("/activity")
	{
		dealerActivity.GET("/", handler.GetActivityList)
		dealerActivity.GET("/:id", handler.GetActivityById)
	}

	dealerBase := r.Group("/dealer")
	{
		dealerBase.GET("/", handler.GetDealerList)
		dealerBase.GET("/:id", handler.GetDealerById)
	}

	promotions := r.Group("/promotions")
	{
		promotions.GET("/", handler.GetPromotionsList)
		promotions.GET("/:id", handler.GetPromotionsById)
	}

	dealerPrice := r.Group("/price")
	{
		dealerPrice.GET("/", handler.GetPriceList)
		dealerPrice.GET("/:dealer/*series", handler.GetPriceByDealerSeries)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("static", "./static")

	reg(r)

	r.Run(":81")
}
