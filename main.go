package main

import (
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})

	m.Post("/new", binding.Form(Pasta{}), func(pasta Pasta, r render.Render) {
		pasta.UID = uniuri.New()
		pasta.CreatedAt = time.Now()

		DB.C("pastas").Insert(pasta)

		r.Redirect(pasta.UID)
	})

	m.Get("/:uid", func(r render.Render, params martini.Params) {
		pasta, err := PastaGet(params["uid"])

		fmt.Println(pasta)
		fmt.Println(err)

		if err != nil {
			r.HTML(http.StatusNotFound, "404", nil)
			return
		}

		r.HTML(http.StatusOK, "view", pasta)
	})

	m.Run()
}
