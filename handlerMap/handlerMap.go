package handlerMap

import (
	"net/http"

	types "github.com/Tamarou/blackarachnia/types"
)

type mapping struct {
	contentType string
	handler     types.Handler
}

func Map(contentType string, handler func(http.ResponseWriter, *http.Request) error) mapping {
	h := types.NewHandlerFunc(handler)
	return mapping{contentType, h}
}

type HandlerMap struct {
	store map[string]types.Handler
	order []string
}

func NewHandlerMap(mappings ...mapping) HandlerMap {
	store := map[string]types.Handler{}
	order := []string{}
	for _, m := range mappings {
		order = append(order, m.contentType)
		store[m.contentType] = m.handler
	}
	h := HandlerMap{store, order}
	return h
}

func (h HandlerMap) FirstType() string                    { return h.order[0] }
func (h HandlerMap) FirstHandler() types.Handler          { return h.store[h.order[0]] }
func (h HandlerMap) Get(contentType string) types.Handler { return h.store[contentType] }
func (h HandlerMap) Types() []string                      { return h.order }
