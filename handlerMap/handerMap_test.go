package handlerMap_test

import (
	"net/http"
	"testing"

	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/stretchr/testify/assert"
)

func toVoid(w http.ResponseWriter, r *http.Request) error { return nil }

func TestHandlerMap(t *testing.T) {
	hm := handlerMap.NewHandlerMap(
		handlerMap.Map("text/plain", toVoid),
		handlerMap.Map("text/html", toVoid),
	)
	assert.Equal(t, "text/plain", hm.FirstType())
}
