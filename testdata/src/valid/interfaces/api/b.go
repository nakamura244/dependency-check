package api

import (
	"github.com/nakamura244/dependency-check/testdata/src/valid/domain"
)

// Repository is struct
type Repository struct {
}

// InsertProjects is ...
func (repo *Repository) InsertProjects(p *domain.B) (*domain.B, error) {
	return nil, nil
}
