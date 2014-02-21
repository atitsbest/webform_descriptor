package handlers

import (
    "log"
    "net/http"
    "time"

    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"

    "github.com/atitsbest/webform_descriptor/application/projects"
    "github.com/atitsbest/webform_descriptor/domain/projects"

    . "github.com/atitsbest/webform_descriptor/common/valueobjects"
)

type (
    ProjectPostModel struct {
        AccountingMode  string
        Leader          string
        Name            string
        OrderAmount     float64
        OrderDate       time.Time
        Risk            string
        State           string
        BMDOrderNumber  string
        Customer        string
        OrderAmountDays uint
        Techs           []string
    }
)

type SpecialX interface{}

func ApiGetProjects(r render.Render, rep application.ProjectRepository) {
    p := rep.All()
    r.JSON(200, p)
}

func ApiGetProject(params martini.Params, r render.Render, rep application.ProjectRepository) {
    ps := rep.All()
    r.JSON(200, ps[0])
}

func ApiPostProject(prj ProjectPostModel, rep application.ProjectRepository, l *log.Logger) (int, string) {
    p := prj.toProject()
    l.Printf("%#v", p)
    rep.Add(p)
    return http.StatusCreated, prj.Name
}

func (m *ProjectPostModel) toProject() domain.Project {
    return domain.Project{
        Name: m.Name,
        // ParentId: nil,
        OrderDate: m.OrderDate,
        // ApprovalDate: time.Tick(0),
        BMDOrderNumber:   m.BMDOrderNumber,
        Customer:         m.Customer,
        Risk:             m.Risk,
        AccountingMode:   m.AccountingMode,
        Leader:           m.Leader,
        Techs:            m.Techs,
        OrderAmount:      Money(m.OrderAmount),
        OrderAmountDays:  WorkAmount(m.OrderAmountDays),
        AlreadyInvoiced:  false,
        AchieveAmount:    Money(0),
        AchiveAmountDays: WorkAmount(0),
    }
}
