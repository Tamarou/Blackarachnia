package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func Test030_Resource(t *testing.T) {
	/*
		{
		            my $res = $cb->(GET "/");
		            is($res->code, 200, '... got the expected status');
		            is($res->header('Content-Type'), 'text/html', '... got the expected Content-Type header');
		            is($res->header('Content-Length'), 126, '... got the expected Content-Length header');
		            is(
		                $res->content,
		                '<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul></ul></body></html>',
		                '... got the expected content'
		            );
		        }
	*/
	resource := YAPC_NA_2012_Example030_Resource{}

	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
		handler := blackarachnia.NewHandler(resource)
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		want := `<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul></ul></body></html>`
		assert.Equal(t, want, res.Body.String())
	})
	/*
	   {
	               my $res = $cb->(POST "/", [ message => 'foo' ]);
	               is($res->code, 301, '... got the expected status');
	               is($res->header('Location'), '/', '... got the right Location header');
	           }

	*/

	t.Run("POST", func(t *testing.T) {
		res := httptest.NewRecorder()
		form := url.Values{"message": {"foo"}}
		r, _ := http.NewRequest(http.MethodPost, "http://example.com/", strings.NewReader(form.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		handler := blackarachnia.NewHandler(resource)
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusMovedPermanently, res.Code)
		assert.Equal(t, "/", res.Header().Get("Location"))

	})

	/*
		{
		            my $res = $cb->(GET "/");
		            is($res->code, 200, '... got the expected status');
		            is($res->header('Content-Type'), 'text/html', '... got the expected Content-Type header');
		            is($res->header('Content-Length'), 138, '... got the expected Content-Length header');
		            is(
		                $res->content,
		                '<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul><li>foo</li></ul></body></html>',
		                '... got the expected content'
		            );
		        }
	*/

	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)
		handler := blackarachnia.NewHandler(resource)
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		want := `<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul><li>foo</li></ul></body></html>`
		assert.Equal(t, want, res.Body.String())
	})

}
