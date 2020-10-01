package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func Test00Basic(t *testing.T) {
	/*
		   {
		       my $res = $cb->(GET "/");
		       is($res->code, 200, '... got the expected status');
		       is($res->header('Content-Type'), 'application/json', '... got the expected Content-Type header');
		       is($res->header('Content-Length'), 25, '... got the expected Content-Length header');
		       is(
					$res->content,
					'{"message":"Hello World"}',
					'... got the expected content'
		       );
		   }
	*/
	t.Run("default", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example000_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	/*
	   {
	       my $res = $cb->(GET "/" => ('Accept' => 'text/html'));
	       is($res->code, 200, '... got the expected status');
	       is($res->header('Content-Type'), 'text/html', '... got the expected Content-Type header');
	       is($res->header('Content-Length'), 46, '... got the expected Content-Length header');
	       is(
	           $res->content,
	           '<html><body><h1>Hello World</h1></body></html>',
	           '... got the expected content'
	       );
	   }
	*/
	t.Run("Accept 'text/html'", func(t *testing.T) {
		// TODO implement content negotiation
	})
}
