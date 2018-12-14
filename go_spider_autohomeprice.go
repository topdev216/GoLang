package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/spider"
)

type PageProcesser struct {
}

func Processer() *PageProcesser {
	return &PageProcesser{}
}

func (this *PageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		log.Fatalln(p.Errormsg())
	}

	p.GetHtmlParser().Find("dl.price-dl").Each(func(i int, series *goquery.Selection) {
		// 获取经销商名称/ID
		dealer_name := strings.TrimSpace(series.Find("div.header-banner div.banner-cont div.cont-text div.text-main").Text())
		dealer_id := strings.Trim(series.Find("div.header-nav ul.nav-ul li.nav_0 a").Attr("href"), "/")
		// 获取车系名称
		series_name := strings.TrimSpace(series.Find("dt.fn-clear div.name p a").Text())

		series.Find("dd table tr").Each(func(j int, model *goquery.Selection) {
			// 过滤第一行th
			row := model.Find("td")
			if row.Is("td") {
				// 获取车型名称
				model_name := strings.TrimSpace(row.Eq(0).Text())

				// 获取指导价
				tmp := strings.TrimSpace(row.Eq(1).Text())
				guide_price := strings.TrimSuffix(tmp, "万")

				// 获取经销商报价
				if row.Eq(2).Find("p").HasClass("text-line") { // 有报价
					tmp = strings.TrimSpace(row.Eq(2).Find("div").Find("a[class='font-bold font-arial']").Text())
				} else { // 无报价
					tmp = strings.TrimSpace(row.Eq(2).Find("p a").Text())
				}
				dealer_price := strings.TrimSuffix(tmp, "万")

				log.Printf("%s(%s), %s, %s, %s, %s\n", dealer_name, dealer_id, series_name, model_name, guide_price, dealer_price)
			}
		})
	})
}

func (this *PageProcesser) Finish() {
	log.Println("Finished.")
}

func main() {
	app := spider.NewSpider(Processer(), "autohome_price")

	/**
		 app.AddUrlWithHeaderFile("http://dealer.autohome.com.cn/1/price.html", "html", "go_spider_autohomeprice_header.json")
		 go_spider_autohomeprice_header.json:
		 {
	         "User-Agent":"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.94 Safari/537.36",
	         "Cookie":"CXID=A19CCB818BB019DB23D791548FBBB498; ssuid=7730855583; pgv_pvi=6130636800; SUID=85D614DA2141900A53BCDDA4000941F7; IPLOC=CN4401; SUV=009B17F377812E8153E04BC3A7FDB688; usid=WV16nbdS8H7gIrS9; ABTEST=7|1427790752|v1; weixinIndexVisited=1; ppinf=5|1427792894|1429002494|Y2xpZW50aWQ6NDoyMDE3fGNydDoxMDoxNDI3NzkyODk0fHJlZm5pY2s6NTQ6JUUzJTgwJThDJUU5JTlCJUFBJUU2JThFJUE5JUU1JUFEJUE0JUU1JTlGJThFJUUzJTgwJThEfHRydXN0OjE6MXx1c2VyaWQ6NDQ6QkE0QTc5QjM1MkZBNDQ2ODAyOUFGQzA5QzlCQzIxMTNAcXEuc29odS5jb218dW5pcW5hbWU6NTQ6JUUzJTgwJThDJUU5JTlCJUFBJUU2JThFJUE5JUU1JUFEJUE0JUU1JTlGJThFJUUzJTgwJThEfA; pprdig=tAacHnbM4bvIEaYKg5s1Pe50JcQirXJlTJ-73JPBnSEG7i6hjW6wXjzIECbWkkSR3tJaYrV0sM8GJjy-zKR3fIeExpUzY69ne1ePaZpg9mH5KXEzQcRDwGpLw25kSMQC4gXWyc4Jjzqz2LqU33RGrqOcQfY8bzd9O-sL7K8EPyM; sct=19; ad=Ukllllllll2Fs2L6lllllVU00S9lllllLPFQPylllltlllllxCxlw@@@@@@@@@@@; SNUID=3D1D782AF0F4E2C18F85DF31F19C366B; ppmdig=142848190200000070fe1e18cdee1174da9ce0f4a533ebd5; wapsogou_qq_nickname=; LSTMV=772%2C80; LCLKINT=546532"
	     }
	*/

	app.AddUrl("http://dealer.autohome.com.cn/1/price.html", "html")
	//app.AddPipeline(pipeline.NewPipelineConsole()).
	//SetSleepTime("rand", 500, 1000).
	app.SetThreadnum(3)
	app.Run()
}
