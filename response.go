package blackarachnia

import (
	"net/http"
)

type Response struct {
	body    string
	charset string
	http.ResponseWriter
}

func (r *Response) Write(b []byte) (int, error) {
	r.body = r.body + string(b)
	return len(b), nil
}

func (r *Response) SetCharset(c string) {
	r.charset = c
}

func (r *Response) Body() string {
	return r.body
}
