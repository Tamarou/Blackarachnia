/*
#!perl

use strict;
use warnings;

use Web::Machine;

=pod

And showing preference is just as simple as changing
the order of items in content_types_provided

# now HTML is the default
curl -v http://0:5000/

# and you must ask specifically for JSON
curl -v http://0:5000/ -H 'Accept: application/json'

=cut

{
    package YAPC::NA::2012::Example002::Resource;
    use strict;
    use warnings;
    use JSON::XS qw[ encode_json ];

    use parent 'Web::Machine::Resource';

    sub content_types_provided { [
        { 'text/html'        => 'to_html' },
        { 'application/json' => 'to_json' },
    ] }

    sub to_json { encode_json( { message => 'Hello World' } ) }
    sub to_html { '<html><body><h1>Hello World</h1></body></html>' }
}

Web::Machine->new( resource => 'YAPC::NA::2012::Example002::Resource' )->to_app;
*/
package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
)

type YAPC_NA_2012_Example002_Resource struct {
	blackarachnia.Resource
}

func (res YAPC_NA_2012_Example002_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(
		handlerMap.Map("text/html", res.toHTML),
		handlerMap.Map("application/json", res.toJSON),
	)
}

func (res YAPC_NA_2012_Example002_Resource) toJSON(w http.ResponseWriter, r *http.Request) error {
	enc := json.NewEncoder(w)
	enc.Encode(struct {
		Message string `json:"message"`
	}{"Hello World"})
	return nil
}

func (res YAPC_NA_2012_Example002_Resource) toHTML(w http.ResponseWriter, r *http.Request) error {
	io.WriteString(w, "<html><body><h1>Hello World</h1></body></html>")
	return nil
}
