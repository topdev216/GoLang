package main

/**
 * @author Tom 2016-10-01
 * @desc practice for wechat platform authentication
 * @license MIT
 */

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/expvarhandler"
)

// could be configured in files
var (
	http = flag.String("http", ":8080", "HTTP ListenAndServe at")
)

// controllers, routers are registered that way... ugly
func registerHandlers(ctx *fasthttp.RequestCtx) {
	fmt.Println(ctx.QueryArgs().QueryString())
	switch string(ctx.Path()) {
	case "/":
		// test hello world
		fmt.Fprintf(ctx, "Hello, %s!", "world")
		break
	case "/stats":
		// statistics
		expvarhandler.ExpvarHandler(ctx)
		break
	case "auth":
		timestamp := ctx.QueryArgs().Peek("timestamp")
		nonce := ctx.QueryArgs().Peek("nonce")
		token := "Nye3FsXBIaSHLZArO3bw4irU8gq69ANnbE4pH8RkFXO"

		str := fmt.Sprintf("%s%s%s", timestamp, token, nonce)
		if hex.EncodeToString([]byte(fmt.Sprintf("%s", sha1.Sum([]byte(str))))) == string(ctx.QueryArgs().Peek("signature")) {
			fmt.Fprintf(ctx, "%s", ctx.QueryArgs().Peek("echostr"))
		} else {
			fmt.Fprint(ctx, "Access denied.")
		}
		break
	default:
		// nothing response
		fmt.Fprint(ctx, "")
		break
	}
}

// so ...difficult to use
func initFramework(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain; charset=utf8")
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// register handlers
	registerHandlers(ctx)

	// set cookies
	var cookie fasthttp.Cookie
	cookie.SetKey("city_code")
	cookie.SetValue("510000")
	ctx.Response.Header.SetCookie(&cookie)
}

// entry point you know
func main() {
	// init
	h := initFramework
	// config
	h = fasthttp.CompressHandler(h)
	// serve
	if err := fasthttp.ListenAndServe(*http, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

	select {}
}
