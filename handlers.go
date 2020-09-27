package blackarachnia

import (
	"io"
	"net/http"

	"github.com/Tamarou/blackarachnia/fsm"
	types "github.com/Tamarou/blackarachnia/types"
)

func NewHandler(res types.Resource) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		w := &Response{ResponseWriter: wr}
		fsm.Run(res, w, r)
		io.WriteString(w.ResponseWriter, w.Body())
	}
}
