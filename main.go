package main

import (
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
		Directory: "templates", // Specify what path to load the templates from.
		Layout:    "layout",    // Specify a layout template. Layouts can call {{ yield }} to render the current template.

	}))

	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})

	m.Post("/new", binding.Form(Pasta{}), func(pasta Pasta, r render.Render) {
		pasta.UID = uniuri.New()
		pasta.CreatedAt = time.Now()

		DB.C("pastas").Insert(pasta)

		r.Redirect(pasta.UID)
		//r.HTML(http.StatusOK, "view", PastaAll())
	})

	m.Get("/:uid", func(r render.Render, params martini.Params) {
		r.HTML(http.StatusOK, "view", PastaGet(params["uid"]))
	})

	// m.NotFound(func(r render.Render) {
	// 	r.HTML(http.StatusNotFound, "view", PastaGet())
	// })

	m.Run()
}
