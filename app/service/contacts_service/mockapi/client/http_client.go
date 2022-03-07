package client

import "net/http"

// HTTPClient interface
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}
