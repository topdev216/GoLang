package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// signature=d4aa019d38db754148310f223f575296742226c4&echostr=4610954209150367977&timestamp=1475328193&nonce=1137652625
// mine
// signature=d4aa019d38db754148310f223f575296742226c4&echostr=4610954209150367977&timestamp=1475328193&nonce=1137652625
func index(c *routing.Context) error {
	fmt.Fprint(c, "Hello world")
	return nil
}

func hello(c *routing.Context) error {
	fmt.Fprintf(c, "Name: %v", c.Param("world"))
	return nil
}

func login(c *routing.Context) error {
	fmt.Fprintf(c, "login")
	return nil
}

func check_login(c *routing.Context) error {
	fmt.Fprintf(c, "check_login")
	return nil
}

func auth_wechat(c *routing.Context) error {
	fmt.Printf("%s\n", c.QueryArgs())
	param := []string{
		string(c.QueryArgs().Peek("timestamp")),
		string(c.QueryArgs().Peek("nonce")),
		"Nye3FsXBIaSHLZArO3bw4irU8gq69ANnbE4pH8RkFXO",
	}

	sort.Strings(param)
	str := fmt.Sprintf("%s%s%s", param[0], param[1], param[2])
	if hex.EncodeToString([]byte(fmt.Sprintf("%s", sha1.Sum([]byte(str))))) == string(c.QueryArgs().Peek("signature")) {
		fmt.Fprintf(c, "%s", c.QueryArgs().Peek("echostr"))
	} else {
		fmt.Fprint(c, "Access denied.")
	}
	return nil
}

func main() {
	router := routing.New()

	router.Get("/", index)
	router.Get("/hello/<world>", hello)
	router.Get("/wechat/auth", auth_wechat)

	admin := router.Group("/admin")
	admin.Get("/login", login).Post(check_login)

	panic(fasthttp.ListenAndServe(":80", router.HandleRequest))
}
