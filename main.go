package main

import (
  "os"
  "log"
  "encoding/csv"
  "net/http"
  // "time"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "github.com/codegangsta/martini-contrib/binding"

  "web/proj/entities"
  "web/proj/valueobjects"
)

type (
  ProjectRepository interface {
    All() []entities.Project
  }

  ProjectPostModel struct {
    AccountingMode string
    BMDOrderNumber string
    Customer string
    Leader string
    Name string
    OrderAmount float64
    OrderAmountDays uint
    OrderDate string
    Risk string
    State string
    Techs string
  }
)

func (m *ProjectPostModel) toProject() entities.Project {
  return entities.Project{
    Name: m.Name,
    // ParentId: nil,
    // OrderDate: time.Parse("", m.OrderDate),
    // ApprovalDate: time.Tick(0),
    BMDOrderNumber: m.BMDOrderNumber,
    Customer: m.Customer,
    Risk: m.Risk,
    AccountingMode: m.AccountingMode,
    Leader: m.Leader,
    Techs: m.Techs,
    OrderAmount: valueobjects.Money(m.OrderAmount),
    OrderAmountDays: valueobjects.WorkAmount(m.OrderAmountDays),
    AlreadyInvoiced: false,
    AchieveAmount: valueobjects.Money(0),
    AchiveAmountDays: valueobjects.WorkAmount(0),
  }
}

func readProjects() []entities.Project {
  // Datei öffnen.
  f,e := os.Open("data/projects.csv")
  defer f.Close()
  if e != nil { panic(e) }


  // CSV-Daten einlesen.
  cr := csv.NewReader(f)
  ls,e := cr.ReadAll(); if e != nil { panic(e) }

  // Liste der Projekte anlegen.
  ps := make([]entities.Project, len(ls)-1)

  // CSV parsen (Header auslassen)
  for i,line := range ls[1:] {
    ps[i].FromCSV(line)
  }
  return ps
}

func index(r render.Render) {
  r.HTML(200, "projects/index", nil)
}

func apiGetProjects(r render.Render) {
  ps := readProjects()
  r.JSON(200, ps)
}

func apiPostProject(prj ProjectPostModel, l *log.Logger) (int, string) {
  l.Printf("%#v", prj)
  return http.StatusCreated, prj.Name
}

func projectEdit(r render.Render) {
  r.HTML(200, "projects/form", nil)
}

func initMartini() *martini.ClassicMartini {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Extensions: []string{".html"},
    Layout: "_layout",
  }))
  m.Get("/", index)
  m.Get("/projects/new", projectEdit)

  m.Get("/api/projects", apiGetProjects)
  m.Post("/api/projects", binding.Bind(ProjectPostModel{}), apiPostProject)
  return m
}

func main() {
  m := initMartini()
  m.Run()
}
