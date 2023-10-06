package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	URL    string
	Header map[string]string
	Data   map[string]interface{}
}

func NewRequest(url string) *Request {
	return &Request{URL: url}
}

func (r *Request) SetHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}

func (r *Request) SetHeaders(headers map[string]string) {
	r.Header = headers
}

func (r *Request) SetData(data map[string]interface{}) {
	r.Data = data
}

func (r *Request) POST() (*http.Response, error) {
	data, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for k, v := range r.Header {
		req.Header.Set(k, v)
	}

	return new(http.Client).Do(req)
}
