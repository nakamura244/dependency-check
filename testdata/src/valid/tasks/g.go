package tasks

import (
	"github.com/nakamura244/dependency-check/testdata/src/valid/usecase"
)

func action(event string, i usecase.LogRepository, d, e, c int) (uint, uint, error) {
	return 0, 0, nil
}
