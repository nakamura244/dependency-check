package log

import "github.com/nakamura244/dependency-check/testdata/src/valid/domain"

// Repository is struct
type Repository struct{}

// InfoLog is logging  to info
func (repo *Repository) InfoLog(msg string, fields ...domain.B) {
}
