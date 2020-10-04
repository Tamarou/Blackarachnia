/*
#!perl

use strict;
use warnings;

use Web::Machine;

=pod

This test shows that the order of content_types_provided
is actually important if you do not specify a media-type.

# JSON is the default ...
curl -v http://0:5000/

# you must ask specifically for HTML
curl -v http://0:5000/ -H 'Accept: text/html'

# but open in a browser and you get HTML
open http://0:5000/

=cut

{
    package YAPC::NA::2012::Example001::Resource;
    use strict;
    use warnings;
    use JSON::XS qw[ encode_json ];

    use parent 'Web::Machine::Resource';

    sub content_types_provided { [
        { 'application/json' => 'to_json' },
        { 'text/html'        => 'to_html' },
    ] }

    sub to_json { encode_json( { message => 'Hello World' } ) }
    sub to_html { '<html><body><h1>Hello World</h1></body></html>' }
}

Web::Machine->new( resource => 'YAPC::NA::2012::Example001::Resource' )->to_app;
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

type YAPC_NA_2012_Example001_Resource struct {
	blackarachnia.Resource
}

func (res YAPC_NA_2012_Example001_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(
		handlerMap.Map("application/json", res.toJSON),
		handlerMap.Map("text/html", res.toHTML),
	)
}

func (res YAPC_NA_2012_Example001_Resource) toJSON(w http.ResponseWriter, r *http.Request) error {
	enc := json.NewEncoder(w)
	enc.Encode(struct {
		Message string `json:"message"`
	}{"Hello World"})
	return nil
}

func (res YAPC_NA_2012_Example001_Resource) toHTML(w http.ResponseWriter, r *http.Request) error {
	io.WriteString(w, "<html><body><h1>Hello World</h1></body></html>")
	return nil
}
