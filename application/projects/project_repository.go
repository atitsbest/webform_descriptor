package application

import (
    "github.com/atitsbest/webform_descriptor/domain/projects"
)

type ProjectRepository interface {
    All() []domain.Project
    Add(domain.Project)
}
