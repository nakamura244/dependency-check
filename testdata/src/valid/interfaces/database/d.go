package database

import (
	"errors"

	"github.com/nakamura244/dependency-check/testdata/src/valid/domain"
)

// SQLRepository is struct
type SQLRepository struct{}

func (repo *SQLRepository) FindProject(id int) (p *domain.B, err error) {
	err = errors.New("aaa")
	return nil, nil
}
