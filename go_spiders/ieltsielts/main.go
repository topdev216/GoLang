package main

import (
	"fmt"
	"github.com/hu17889/go_spider/core/spider"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	//"os"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/request"
	"strings"
	"log"
	"github.com/lunny/html2md"
	"html"
)

// ------------------- Ielts Spider
type IeltsProcesser struct {}

func NewIeltsProcesser() *IeltsProcesser {
	return &IeltsProcesser{}
}

func (this *IeltsProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		panic(p.Errormsg())
	}
/*
<article id="post-3380" class="post-3380 post type-post status-publish format-standard hentry category-ielts-speaking">
    <header class="entry-header">
        <h1 class="entry-title">
            <a href="http://ieltsielts.com/model-essay-this-question-was-seen-in-sri-lanka-in-april-2018/" rel="bookmark">Model essay! This question was seen in Sri Lanka in April 2018…</a
        </h1>
	<div class="entry-meta"></div><!-- .entry-meta -->
    </header><!-- .entry-header -->
    <div class="entry-content"></div><!-- .entry-content -->
    <footer class="entry-meta">
        <span class="entry-date">
            <a href="http://ieltsielts.com/model-essay-this-question-was-seen-in-sri-lanka-in-april-2018/" title="5:56 pm" rel="bookmark">
                <time datetime="2018-07-01T17:56:21+00:00" pubdate="">1 July 2018</time>
            </a>
        </span>
        <span class="comments-link">
            <a href="http://ieltsielts.com/model-essay-this-question-was-seen-in-sri-lanka-in-april-2018/#respond">
                <span class="leave-reply">Leave a reply</span>
            </a>
        </span>
    </footer><!-- #entry-meta -->
</article><!-- #post-3380 -->
*/

	dom := p.GetHtmlParser()
	dom.Find("article.post").Each(func(i int, selection *goquery.Selection) {
		p.AddField("title", selection.Find("header.entry-header h1.entry-title a").Text())
		p.AddField("created_at", selection.Find("footer.entry-meta span.entry-date time").AttrOr("datetime", ""))
		p.AddField("permalink", selection.Find("header.entry-header h1.entry-title a").AttrOr("href", ""))
		p.AddField("category", selection.Find("").Text())
		if content, err := selection.Find("div.entry-content").Html(); err == nil {
			content = html.UnescapeString(content)
			content = html2md.Convert(content)
			content = strings.TrimSpace(content)
			p.AddField("content", content)
		}
	})
}

func (this *IeltsProcesser) Finish() {
	log.Println("Finished")
}

// ------------------- File Pipeline
type FilePipeline struct {}

func NewFilePipeline() *FilePipeline {
	return &FilePipeline{}
}

func (this *FilePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
}

// ------------------- Console Pipeline
type ConsolePipeline struct {}

func NewConsolePipeline() *ConsolePipeline {
	return &ConsolePipeline{}
}

func (this *ConsolePipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	log.Printf("%v\n", items.GetAll())
}

// ------------------- Main Entry
func main() {
	defer func() {
		if r := recover(); r != nil {
			println(fmt.Sprintf("%v", r))
		}
	}()

	spider.NewSpider(NewIeltsProcesser(), "ielts").
		SetSleepTime("rand", 200, 1000).
		AddRequest(request.NewRequestWithHeaderFile("http://ieltsielts.com", "html", "header.json")).
		SetThreadnum(16).
		AddPipeline(NewConsolePipeline()).
		//AddPipeline(NewFilePipeline()).
		Run()
}


