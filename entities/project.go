package entities

import (
  "time"
  "strings"
  vo "github.com/atitsbest/webform_descriptor/valueobjects"
)

type ProjectId string

type Project struct {
  ParentId          ProjectId
  OrderDate         time.Time
  ApprovalDate      time.Time
  BMDOrderNumber    string
  Customer          string
  Risk              string
  AccountingMode    string
  Name              string
  Leader            string
  Techs             []string
  OrderAmount       vo.Money
  OrderAmountDays   vo.WorkAmount
  AlreadyInvoiced   bool
  AchieveAmount     vo.Money
  AchiveAmountDays  vo.WorkAmount
}

// Properties mit den Daten aus data befüllen.
func (p *Project) FromCSV(data []string) {
  p.Name = data[0]
  p.Leader = data[1]
  p.Techs = strings.Split(data[2],",")
  if e := p.OrderAmount.FromString(data[3]); e != nil {panic(e)}
  if e := p.OrderAmountDays.FromString(data[4]); e!=nil {panic(e)}
}


