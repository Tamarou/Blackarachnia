package fsm

import (
	"net/http"
	"reflect"
	"runtime"

	types "github.com/Tamarou/blackarachnia/types"
)

func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

func Run(res types.Resource, w types.Response, r *http.Request) {
	state := initialState()
	for state != nil {
		//		log.Println(nameOf(state))
		state = state(res, w, r)
	}
}
