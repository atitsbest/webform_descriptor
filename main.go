package main

import (
    "os"
    "log"
    "encoding/csv"
    "net/http"
    "time"
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"
    "github.com/codegangsta/martini-contrib/binding"

    "github.com/atitsbest/webform_descriptor/entities"
    "github.com/atitsbest/webform_descriptor/valueobjects"
)

type (
    Projects interface {
        All() []entities.Project
        Add(entities.Project)
    }

    /*
    * CSV-Implementierung des Project-Repositories.
    */
    CsvProjects struct {
        // Die Projekte werden im Speicher gehalten.
        ps []entities.Project
    }

    ProjectPostModel struct {
        AccountingMode string
        Leader string
        Name string
        OrderAmount float64
        OrderDate time.Time
        Risk string
        State string
        BMDOrderNumber string
        Customer string
        OrderAmountDays uint
        Techs []string
    }
)

func (m *ProjectPostModel) toProject() entities.Project {
    return entities.Project{
        Name: m.Name,
        // ParentId: nil,
        OrderDate: m.OrderDate,
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

func (p *CsvProjects) All() []entities.Project {
    if len(p.ps) == 0 {
        p.ps =  readProjects();
    }
    return p.ps;
}

func (p *CsvProjects) Add(e entities.Project) {
    p.ps = append(p.ps, e);
}

func projectIndex(r render.Render) {
    r.HTML(200, "projects/index", nil)
}

func apiGetProjects(r render.Render) {
    ps := readProjects()
    r.JSON(200, ps)
}

func apiGetProject(params martini.Params, r render.Render) {
    ps := readProjects()
    r.JSON(200, ps[0])
}

func apiPostProject(prj ProjectPostModel, l *log.Logger) (int, string) {
    l.Printf("%#v", prj.toProject())
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
        Delims: render.Delims{"{%", "%}"},
    }))
    m.Get("/", projectIndex)
    m.Get("/projects", projectIndex)
    m.Get("/projects/new", projectEdit)

    m.Get("/api/projects", apiGetProjects)
    m.Get("/api/projects/:id", apiGetProject)
    m.Post("/api/projects", binding.Bind(ProjectPostModel{}), apiPostProject)
    return m
}

func main() {
    m := initMartini()
    m.Run()
}
