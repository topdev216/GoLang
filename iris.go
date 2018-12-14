package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/kataras/iris"
)

func auth(ctx *iris.Context) {
	timestamp := ctx.URLParam("timestamp")
	nonce := ctx.URLParam("nonce")
	token := "Nye3FsXBIaSHLZArO3bw4irU8gq69ANnbE4pH8RkFXO"

	str := fmt.Sprintf("%s%s%s", timestamp, token, nonce)
	if hex.EncodeToString([]byte(fmt.Sprintf("%s", sha1.Sum([]byte(str))))) == ctx.URLParam("signature") {
		ctx.HTML(iris.StatusOK, ctx.URLParam("echostr"))
	} else {
		ctx.HTML(iris.StatusOK, "Access denied.")
	}
}

func main() {
	// iris.Favicon("./favicon.ico")

	// handlers
	iris.Get("/auth", auth)

	iris.Listen(":8080")
}
