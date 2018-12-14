package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"models"
)

func ManageInit(r render.Render) {
	err := models.InitTables()
	var result string
	if err != nil {
		result = "initialize failed."
	} else {
		result = "initialize successfully."
	}
	r.HTML(200, "index", result)
}
