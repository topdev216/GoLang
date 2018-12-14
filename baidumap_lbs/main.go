package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func base64_decode(str string) string {
	dec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str
	}
	return string(dec)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/", func(c *gin.Context) {
		uri := "http://api.map.baidu.com/geodata/v3/poi/create"
		params := url.Values{}
		params.Add("title", c.PostForm("title"))
		params.Add("address", c.PostForm("address"))
		params.Add("latitude", c.PostForm("latitude"))
		params.Add("longitude", c.PostForm("longitude"))
		params.Add("coord_type", "1")
		params.Add("geotable_id", "146844")
		params.Add("ak", "800402c858110e5f2e1c96c8ca810ca9")
		params.Add("hagent", c.PostForm("hagent"))
		params.Add("hvisits", c.PostForm("hvisits"))
		params.Add("hremark", c.PostForm("hremark"))

		// resp, err := http.PostForm(uri, url.Values(params))
		// resp, err := http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
		resp, err := http.Post(uri, "Content-type: application/json", strings.NewReader(params.Encode()))
		// resp, err := http.NewRequest("POST", uri, strings.NewReader(params.Encode())) // 无法获取参数
		// log.Println(strings.NewReader(params.Encode()))
		if err != nil {
			c.String(500, base64_decode(err.Error()))
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.String(500, base64_decode(err.Error()))
			return
		}

		c.String(200, base64_decode(string(body)))
	})

	r.Run(":80")
	log.Println("service's started successfully.")
}
