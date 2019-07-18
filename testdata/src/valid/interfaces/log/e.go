package log

import "github.com/nakamura244/dependency-check/testdata/src/valid/domain"

// Loghandler is interface
type Loghandler interface {
	Info(msg string, fields ...domain.B)
}
