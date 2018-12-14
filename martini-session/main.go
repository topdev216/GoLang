package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"login/models"
	"net/http" // for port
)

type _SIGNINDATA struct {
	USERNAME string `form:"username"`
	PASSWORD string `form:"password"`
}

func main() {

	m := martini.Classic()          // initialize
	m.Use(render.Renderer())        // use templates, need this directory
	m.Use(martini.Static("static")) // serve /static directory

	store := sessions.NewCookieStore([]byte("localhost"))
	m.Use(sessions.Sessions("testsite", store))

	m.Get("/", func(r render.Render) { // serve this url
		r.HTML(200, "home", nil)
	})

	m.Post("/", binding.Form(_SIGNINDATA{}), func(data _SIGNINDATA, session sessions.Session) string {
		v := session.Get(data.USERNAME)
		if v == 1 {
			return "Sorry! You have signed in."
		}

		if data.USERNAME == "admin" && data.PASSWORD == "admin" {
			session.Set(data.USERNAME, 1)
			return "Congratulations! You have successfully signed in."
		}

		return "Sorry! Please check your username or password."
	})

	http.ListenAndServe(":80", m) // use port 80 as default
	m.Run()
}
