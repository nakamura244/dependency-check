package infrastructures

import (
	"net/http"
	"time"

	"github.com/nakamura244/dependency-check/testdata/src/valid/config"
	"github.com/nakamura244/dependency-check/testdata/src/valid/domain"
	"github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/iface"
)

// APIhandler is struct
type APIhandler struct {
	httpClient iface.HTTPAPI
}

// NewAPIhandler is constructor
func NewAPIhandler(cfg *config.Config) *APIhandler {
	return &APIhandler{
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
		},
	}
}

// AddNewRequest is add http request
func (e *APIhandler) AddNewRequest(p *domain.B) error {
	return nil
}
