package main

import (
	"io"
	"log"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
)

/*

    package My::Resource;
    use strict;
    use warnings;

    use parent "Web::Machine::Resource";

    sub content_types_provided { [{ "text/html" => "to_html" }] }

    sub to_html { "<html><body>Hello World</body></html>" }
}
*/

type MyResource struct{ blackarachnia.Resource }

func (mr MyResource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(
		handlerMap.Map("text/html", mr.ToHTML),
	)
}

func (mr MyResource) ToHTML(w http.ResponseWriter, r *http.Request) error {
	_, e := io.WriteString(w, "<html><body>Hello World</body></html>")
	return e
}

func main() {
	resource := &MyResource{}
	handler := blackarachnia.NewHandler(resource)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
