package usecase

import "github.com/nakamura244/dependency-check/testdata/src/valid/domain"

// LogRepository is interface
type LogRepository interface {
	InfoLog(msg string, fields ...domain.B)
}
