package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

func has_prefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func format_url(url string) string {
	if has_prefix(url, "//") {
		url = "http:" + url
	} else if !has_prefix(url, "http://") && !has_prefix(url, "https://") {
		url = "http://" + url
	}
	return url
}

func get_access_token() string {
	return ""
}

func do_post(url string, param *fasthttp.Args) string {
	code, resp, err := fasthttp.Post(nil, url, param)
	if code != 200 || err != nil {
		panic(err)
	}
	return string(resp)
}

func json_decode(str string) string {
	return str
}

func shorten(url_long string) string {
	api := "https://api.weibo.com/2/short_url/shorten.json"
	postArgs := fasthttp.AcquireArgs()
	postArgs.Add("access_token", get_access_token())
	postArgs.Add("url_long", url_long)
	return json_decode(do_post(api, postArgs))
}

func expand(url_short string) string {
	api := "https://api.weibo.com/2/short_url/expand.json"
	postArgs := fasthttp.AcquireArgs()
	postArgs.Add("access_token", get_access_token())
	postArgs.Add("url_short", url_short)
	return json_decode(do_post(api, postArgs))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t%s [url]\n", os.Args[0])
		return
	}

	url := format_url(os.Args[1])

	if has_prefix(url, "http://t.cn/") || has_prefix(url, "https://t.cn/") {
		fmt.Println(expand(url))
	} else {
		fmt.Println(shorten(url))
	}

}
