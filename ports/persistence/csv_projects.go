package persistence

import (
    "encoding/csv"
    "github.com/atitsbest/webform_descriptor/domain/projects"
    "os"
)

/*
* CSV-Implementierung des Project-Repositories.
 */
type CsvProjects struct {
    // Die Projekte werden im Speicher gehalten.
    ps []domain.Project
}

func (self *CsvProjects) All() []domain.Project {
    if len(self.ps) == 0 {
        self.ps = readProjects()
    }
    return self.ps
}

func (self *CsvProjects) Add(e domain.Project) {
    self.ps = append(self.ps, e)
}

func readProjects() []domain.Project {
    // Datei öffnen.
    f, e := os.Open("data/projects.csv")
    defer f.Close()
    if e != nil {
        panic(e)
    }

    // CSV-Daten einlesen.
    cr := csv.NewReader(f)
    ls, e := cr.ReadAll()
    if e != nil {
        panic(e)
    }

    // Liste der Projekte anlegen.
    ps := make([]domain.Project, len(ls)-1)

    // CSV parsen (Header auslassen)
    for i, line := range ls[1:] {
        ps[i].FromCSV(line)
    }
    return ps
}
