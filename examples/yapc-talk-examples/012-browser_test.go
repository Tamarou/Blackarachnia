package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func TestYAPC_NA_2012_Example012_Resource(t *testing.T) {
	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Accept", "*/*")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example012_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "image/gif", res.Header().Get("Content-Type"))
		assert.NotEmpty(t, res.Body)
	})

	t.Run("GET / (Accept: text/html)", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example012_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "text/html", res.Header().Get("Content-Type"))
		assert.Equal(t, `<html><body><ul><li>1.000000 &mdash; text</li><li>1.000000 &mdash; application</li><li>0.900000 &mdash; application</li><li>0.800000 &mdash; *</li></ul></body></html>`, res.Body.String())
	})
}
