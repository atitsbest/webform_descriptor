package main

import (
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/binding"
    "github.com/codegangsta/martini-contrib/render"

    "github.com/atitsbest/webform_descriptor/application/projects"
    "github.com/atitsbest/webform_descriptor/ports/persistence"
    "github.com/atitsbest/webform_descriptor/ports/ui/handlers"
)

type SpecialX interface{}

func projectIndex(r render.Render) {
    r.HTML(200, "projects/index", nil)
}

func projectEdit(r render.Render) {
    r.HTML(200, "projects/form", nil)
}

func initMartini() *martini.ClassicMartini {
    m := martini.Classic()

    // DEPENDENCY-INJECTION
    m.MapTo(&persistence.CsvProjects{}, (*application.ProjectRepository)(nil))

    m.Use(render.Renderer(render.Options{
        Extensions: []string{".html"},
        Layout:     "_layout",
        Delims:     render.Delims{"{%", "%}"},
    }))

    m.Get("/", projectIndex)
    m.Get("/projects", projectIndex)
    m.Get("/projects/new", projectEdit)

    m.Get("/api/projects", handlers.ApiGetProjects)
    m.Get("/api/projects/:id", handlers.ApiGetProject)
    m.Post("/api/projects", binding.Bind(handlers.ProjectPostModel{}), handlers.ApiPostProject)
    return m
}

func main() {
    m := initMartini()

    m.Run()
}
