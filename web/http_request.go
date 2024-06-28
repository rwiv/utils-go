package web

import (
	"bytes"
	"net/http"
)

type HttpRequest struct {
	Method  string
	Url     string
	Body    []byte
	Headers map[string]string
}

func (b HttpRequest) Call() (*http.Response, error) {
	req, err := http.NewRequest(b.Method, b.Url, bytes.NewBuffer(b.Body))
	if err != nil {
		return nil, err
	}
	for key, value := range b.Headers {
		req.Header.Add(key, value)
	}

	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return c.Do(req)
}
