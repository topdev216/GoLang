package main

import (
	"github.com/codegangsta/martini"
)

func main() {
     m:=martini.Classic()
     m.Get("/:name", func(params martini.Params) string {
     		     return "Hello " + params["name"]
     })

     m.Run()
}
