package usecase

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"

	"github.com/nakamura244/dependency-check/testdata/src/valid/domain"
)

type s struct {
}

// Validate is validate
func (s *s) Validate(p *domain.B) error {
	return nil
}

// GetLogRowData is ...
func (s *s) GetLogRowData(decoded string) ([]*domain.B, error) {
	data, err := base64.StdEncoding.DecodeString(decoded)
	if err != nil {
		return nil, err
	}
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	buf.ReadFrom(zr)

	return nil, nil
}
