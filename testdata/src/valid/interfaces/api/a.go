package api

import (
	"github.com/nakamura244/dependency-check/testdata/src/valid/domain"
)

// APIhandler is interface
type APIhandler interface {
	AddNewRequest(p *domain.B) error
}
