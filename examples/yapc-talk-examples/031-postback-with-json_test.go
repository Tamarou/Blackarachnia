package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func Test031_Resource(t *testing.T) {
	resource := New_YAPC_NA_2012_Example031_Resource()
	handler := blackarachnia.NewHandler(resource)

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
	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "text/html", res.Header().Get("Content-Type"))
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

		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		want := `<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul><li>foo</li></ul></body></html>`
		assert.Equal(t, want, res.Body.String())
	})
	/*
		{
			my $res = $cb->(PUT "/", Content_Type => 'application/json', Content => '"bar"');
			is($res->code, 204, '... got the expected status');
		}
	*/
	t.Run("PUT /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodPut, "http://example.com/", bytes.NewBufferString(`"bar"`))
		r.Header.Set("Content-Type", "application/json")

		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})
	/*
				{
		            my $res = $cb->(GET "/");
		            is($res->code, 200, '... got the expected status');
		            is($res->header('Content-Type'), 'text/html', '... got the expected Content-Type header');
		            is($res->header('Content-Length'), 150, '... got the expected Content-Length header');
		            is(
		                $res->content,
		                '<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul><li>foo</li><li>bar</li></ul></body></html>',
		                '... got the expected content'
		            );
		        }

	*/
	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "http://example.com/", nil)

		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		want := `<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul><li>foo</li><li>bar</li></ul></body></html>`
		assert.Equal(t, want, res.Body.String())
	})
}
