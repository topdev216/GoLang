package main

import (
	"github.com/hu17889/go_spider/core/spider"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
)

// -------------------- Bean Definition --------------------
type Bean struct {
}

func BeanProcesser() *Bean {
	return &Bean{}
}

func (this *Bean) Process(p *page.Page) {}

func (this *Bean) Finish() {}

// -------------------- File Pipeline --------------------
type FilePipeline struct {
}

func NewFilePipeline() *FilePipeline {
	return &FilePipeline{}
}

func (this *FilePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {}

// -------------------- Console Pipeline --------------------
type ConsolePipeline struct {
}

func NewConsolePipeline() *ConsolePipeline {
	return &ConsolePipeline{}
}

func (this *ConsolePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {}

func main() {
	app := spider.NewSpider(BeanProcesser(), "test")
	app.AddUrlWithHeaderFile("", "html", "header.json")
	app.SetThreadnum(16)
	app.AddPipeline(NewConsolePipeline())
	app.AddPipeline(NewFilePipeline())
	app.Run()
}
