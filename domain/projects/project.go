package domain

import (
    . "github.com/atitsbest/webform_descriptor/common/valueobjects"
    "strings"
    "time"
)

type ProjectId string

type Project struct {
    ParentId         ProjectId
    OrderDate        time.Time
    ApprovalDate     time.Time
    BMDOrderNumber   string
    Customer         string
    Risk             string
    AccountingMode   string
    Name             string
    Leader           string
    Techs            []string
    OrderAmount      Money
    OrderAmountDays  WorkAmount
    AlreadyInvoiced  bool
    AchieveAmount    Money
    AchiveAmountDays WorkAmount
}

// Properties mit den Daten aus data befüllen.
func (p *Project) FromCSV(data []string) {
    p.Name = data[0]
    p.Leader = data[1]
    p.Techs = strings.Split(data[2], ",")
    if e := p.OrderAmount.FromString(data[3]); e != nil {
        panic(e)
    }
    if e := p.OrderAmountDays.FromString(data[4]); e != nil {
        panic(e)
    }
}
