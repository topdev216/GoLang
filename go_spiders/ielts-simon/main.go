package main

import (
	"fmt"
	"github.com/hu17889/go_spider/core/spider"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"os"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/request"
	"strings"
	"log"
	"github.com/lunny/html2md"
	"html"
)

/*
 * @author Tom
 * @date 2018-07-16
 * @desc Spider Implementation
 */
type SimonProcesser struct {}

func NewSimonProcesser() *SimonProcesser {
	return &SimonProcesser{}
}

func (this *SimonProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		panic(p.Errormsg())
	}

	dom := p.GetHtmlParser()
	dom.Find("h2.date-header").Each(func(i int, selection *goquery.Selection) {
		p.AddField("created_at", selection.Text())

		node := dom.AfterSelection(selection).Find("div.entry div.entry-inner").First()
		p.AddField("title", node.Find("h3.entry-header a").Text())
		p.AddField("permalink", node.Find("h3.entry-header a").AttrOr("href", ""))

		var categories []string
		node.Find("div.entry-footer p.entry-footer-info span.post-footers a").Each(func(j int, sel *goquery.Selection) {
			if text := sel.Text(); text != "Simon" {
				categories = append(categories, text)
			}
		})
		p.AddField("category", strings.Join(categories, ", "))

		if content, err := node.Find("div.entry-content div.entry-body").Html(); err == nil {
			content = html.UnescapeString(content)
			content = html2md.Convert(content)
			content = strings.TrimSpace(content)
			p.AddField("content", content)
		}
	})
}

func (this *SimonProcesser) Finish() {
	log.Println("Finished!")
}
// END

/*
 * @author Tom
 * @date 2018-07-16
 * @desc File Pipeline
 */
type FilePipeline struct {
	fp *os.File
}

func NewFilePipeline() *FilePipeline {
	return &FilePipeline{}
}

func (this *FilePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	fields := items.GetAll()

	tmp := strings.Split(fields["permalink"], "/")
	fileName := strings.TrimRight(tmp[len(tmp) - 1], ".html") + ".md"

	if fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644); err == nil {
		defer fp.Close()

		fp.WriteString(fmt.Sprintf("## %s\n\n", fields["title"]))
		fp.WriteString(fmt.Sprintf("> Simon Corcora | %s | %s\n> Permalink: %s\n\n", fields["created_at"], fields["category"], fields["permalink"]))
		fp.WriteString(fields["content"] + "\n")
	}
}
// END

/*
 * @author Tom
 * @date 2018-07-16
 * @desc Console Pipeline
 */
type ConsolePipeline struct {}

func NewConsolePipeline() *ConsolePipeline {
	return &ConsolePipeline{}
}

func (this *ConsolePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	log.Printf("%v\n", items.GetAll())
}
// END

// Main Entry Point
func main() {
	defer func() {
		if r := recover(); r != nil {
			println(fmt.Sprintf("Error caught: %v", r))
		}
	}()

	spider.NewSpider(NewSimonProcesser(), "simon").
		SetSleepTime("rand", 200, 1000).
		AddRequest(request.NewRequestWithHeaderFile("http://ielts-simon.com", "html", "header.json")).
		SetThreadnum(16).
		AddPipeline(NewConsolePipeline()).
		AddPipeline(NewFilePipeline()).
		Run()
}
