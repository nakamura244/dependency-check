package iface

import (
	"net/http"
)

// HTTPAPI is interface
type HTTPAPI interface {
	Do(req *http.Request) (*http.Response, error)
}
