package main

/*
       {
           my $res = $cb->(GET "/");
           is($res->code, 401, '... got the expected status');
           is($res->header('Content-Type'), 'text/plain', '... got the expected Content-Type header');
           is($res->header('WWW-Authenticate'), 'Basic realm="Webmachine"', '... got the expected WWW-Authenticate header');
           is(
               $res->content,
               'Unauthorized',
               '... got the expected content'
           );
       }

       {
           my $res = $cb->(GET "/" => ('Authorization' => 'Basic ' . MIME::Base64::encode_base64('foo:bar')));
           is($res->code, 200, '... got the expected status');
           is($res->header('Content-Type'), 'text/html', '... got the expected Content-Type header');
           is($res->header('Content-Length'), 46, '... got the expected Content-Length header');
           is(
               $res->content,
               '<html><body><h1>Hello World</h1></body></html>',
               '... got the expected content'
           );
       }

       {
           my $res = $cb->(GET "/" => ('Authorization' => 'Basic ' . MIME::Base64::encode_base64('foo:baz')));
           is($res->code, 401, '... got the expected status');
           is($res->header('Content-Type'), 'text/plain', '... got the expected Content-Type header');
           is($res->header('WWW-Authenticate'), 'Basic realm="Webmachine"', '... got the expected WWW-Authenticate header');
           is(
               $res->content,
               'Unauthorized',
               '... got the expected content'
           );
       }
   };
*/

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func TestYAPC_NA_2012_Example020_Resource(t *testing.T) {
	t.Run("GET / 401 (no user/pass)", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example020_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, "text/plain; charset=utf-8", res.Header().Get("Content-Type"))
		assert.Equal(t, `Basic realm="Webmachine"`, res.Header().Get("WWW-Authenticate"))
	})
	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("foo", "bar")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example020_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "text/html", res.Header().Get("Content-Type"))
		assert.Equal(t, "<html><body><h1>Hello World</h1></body></html>", res.Body.String())
	})
	t.Run("GET / 401 (bad user/pass)", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("foo", "baz")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example020_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, "text/plain; charset=utf-8", res.Header().Get("Content-Type"))
		assert.Equal(t, `Basic realm="Webmachine"`, res.Header().Get("WWW-Authenticate"))
	})
}
