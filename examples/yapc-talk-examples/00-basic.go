/*
#!perl

use strict;
use warnings;

use Web::Machine;

=pod

curl -v http://0:5000/

# fails with a 406
curl -v http://0:5000/ -H 'Accept: image/jpeg'

=cut

{
    package YAPC::NA::2012::Example000::Resource;
    use strict;
    use warnings;
    use JSON::XS qw[ encode_json ];

    use parent 'Web::Machine::Resource';

    sub content_types_provided { [{ 'application/json' => 'to_json' }] }

    sub to_json { encode_json( { message => 'Hello World' } ) }
}

Web::Machine->new( resource => 'YAPC::NA::2012::Example000::Resource' )->to_app;
*/

package main

import (
	"encoding/json"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
)

type YAPC_NA_2012_Example000_Resource struct{ blackarachnia.Resource }

func (res YAPC_NA_2012_Example000_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(
		handlerMap.Map("application/json", toJSON),
	)
}

func toJSON(w http.ResponseWriter, r *http.Request) error {
	enc := json.NewEncoder(w)
	enc.Encode(struct {
		Message string `json:"message"`
	}{"Hello World"})
	return nil
}
