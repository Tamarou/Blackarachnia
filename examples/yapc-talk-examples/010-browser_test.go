package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tamarou/blackarachnia"
	"github.com/stretchr/testify/assert"
)

func TestYAPC_NA_2012_Example010_Resource(t *testing.T) {
	t.Run("GET /", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Accept", "*/*")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example010_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, `[{"Type":"*","SubType":"*","Q":1,"Params":{}}]`, res.Body.String())
	})
	t.Run("GET / (Accept: text/html)", func(t *testing.T) {
		res := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		handler := blackarachnia.NewHandler(&YAPC_NA_2012_Example010_Resource{})
		handler.ServeHTTP(res, r)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.JSONEq(t, `[{"Type":"text","SubType":"html","Q":1,"Params":{}},{"Type":"application","SubType":"xhtml+xml","Q":1,"Params":{}},{"Type":"application","SubType":"xml","Q":0.8999999761581421,"Params":{}},{"Type":"*","SubType":"*","Q":0.800000011920929,"Params":{}}]`, res.Body.String())
	})
}
