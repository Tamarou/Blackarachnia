package main

import (
	"net/http"

	"github.com/Tamarou/blackarachnia/fsm"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func NewHandler(res fsm.Resource) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		state := fsm.InitialState()
		for state != nil {
			state = state(res, w, r)
		}
	}
}
