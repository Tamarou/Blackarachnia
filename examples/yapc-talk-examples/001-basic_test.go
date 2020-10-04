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
    };

done_testing;

*/
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func TestYAPC_NA_2012_Example001_Resource(t *testing.T) {
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
	t.Run("GET / (*/*)", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example001_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, "{\"message\":\"Hello World\"}\n", res.Body.String())
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
	t.Run("GET / (Accept: text/html)", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Accept", "text/html")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example001_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "text/html", res.Header().Get("Content-Type"))
		assert.Equal(t, "<html><body><h1>Hello World</h1></body></html>", res.Body.String())
	})

}
