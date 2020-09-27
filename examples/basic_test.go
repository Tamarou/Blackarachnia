package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

type Env struct {
	HTTP_ACCEPT          string
	HTTP_ACCEPT_ENCODING string
	HTTP_ACCEPT_LANGUAGE string
	HTTP_CACHE_CONTROL   string
	HTTP_HOST            string
	HTTP_USER_AGENT      string
	PATH_INFO            string
	QUERY_STRING         string
	REMOTE_ADDR          string
	REMOTE_HOST          string
	REQUEST_METHOD       string
	REQUEST_URI          string
	SCRIPT_NAME          string
	SERVER_NAME          string
	SERVER_PORT          string
	SERVER_PROTOCOL      string
}

func newRequest(env Env) *http.Request {
	URI, _ := url.Parse(env.REQUEST_URI)
	r := &http.Request{
		Method: env.REQUEST_METHOD,
		URL:    URI,
		Proto:  env.SERVER_PROTOCOL,
		Header: map[string][]string{
			"Accept":          {env.HTTP_ACCEPT},
			"Accept-Encoding": {env.HTTP_ACCEPT_ENCODING},
			"Accept-Language": {env.HTTP_ACCEPT_LANGUAGE},
			"Cache-Control":   {env.HTTP_CACHE_CONTROL},
		},
		Host: env.HTTP_HOST,
	}
	return r
}

func TestBasicResource(t *testing.T) {
	tests := []*http.Request{
		//trace => "b13,b12,b11,b10,b9,b8,b7,b6,b5,b4,b3,c3,d4,e5,f6,g7,g8,h10,i12,l13,m16,n16,o16,o18,o18b",
		newRequest(Env{
			REQUEST_METHOD:  "GET",
			SERVER_PROTOCOL: "HTTP/1.1",
			SERVER_NAME:     "example.com",
			SCRIPT_NAME:     "/foo",
		}),
		//trace => "b13,b12,b11,b10,b9,b8,b7,b6,b5,b4,b3,c3,c4,d4,d5,e5,f6,f7,g7,g8,h10,i12,l13,m16,n16,o16,o18,o18b",
		newRequest(Env{
			SCRIPT_NAME:          "",
			SERVER_NAME:          "127.0.0.1",
			HTTP_ACCEPT_ENCODING: "gzip, deflate",
			PATH_INFO:            "/",
			HTTP_ACCEPT:          "text/html,application/xhtml+xml,application/xml;q=0.9,/*;q=0.8",
			REQUEST_METHOD:       "GET",
			HTTP_USER_AGENT:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/534.53.11 (KHTML, like Gecko) Version/5.1.3 Safari/534.53.10",
			QUERY_STRING:         "",
			SERVER_PORT:          "5000",
			HTTP_CACHE_CONTROL:   "max-age=0",
			HTTP_ACCEPT_LANGUAGE: "en-us",
			REMOTE_ADDR:          "127.0.0.1",
			SERVER_PROTOCOL:      "HTTP/1.1",
			REQUEST_URI:          "/",
			REMOTE_HOST:          "127.0.0.1",
			HTTP_HOST:            "0:5000",
		}),
	}

	for _, test := range tests {
		r := httptest.NewRecorder()
		handler := blackarachnia.NewHandler(MyResource{})
		handler.ServeHTTP(r, test)

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "text/html", r.Header().Get("Content-Type"))
		//assert.Equal(t, "37", r.Header().Get("Content-Length"))
		assert.Equal(t, "<html><body>Hello World</body></html>", r.Body.String())
	}
}
