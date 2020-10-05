/*
#!perl

use strict;
use warnings;

use Web::Machine;

=pod

=cut

{
    package YAPC::NA::2012::Example020::Resource;
    use strict;
    use warnings;
    use Web::Machine::Util qw[ create_header ];

    use parent 'Web::Machine::Resource';

    sub content_types_provided { [ { 'text/html' => 'to_html' } ] }

    sub to_html { '<html><body><h1>Hello World</h1></body></html>' }

    sub is_authorized {
        my ($self, $auth_header) = @_;
        if ( $auth_header ) {
            return 1 if $auth_header->username eq 'foo' && $auth_header->password eq 'bar';
        }
        return create_header( 'WWWAuthenticate' => [ 'Basic' => ( realm => 'Webmachine' ) ] );
    }

}

Web::Machine->new( resource => 'YAPC::NA::2012::Example020::Resource' )->to_app;
*/
package main

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
)

type YAPC_NA_2012_Example020_Resource struct {
	blackarachnia.Resource
}

func (res YAPC_NA_2012_Example020_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(
		handlerMap.Map("text/html", res.toHTML),
	)
}

func (res YAPC_NA_2012_Example020_Resource) toHTML(w http.ResponseWriter, r *http.Request) error {
	io.WriteString(w, "<html><body><h1>Hello World</h1></body></html>")
	return nil
}

// lifted directly from net/http/request.go
// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

func (res YAPC_NA_2012_Example020_Resource) Authorized(w http.ResponseWriter, header string) bool {
	if user, pass, _ := parseBasicAuth(header); user != "" && pass != "" {
		if user == "foo" && pass == "bar" {
			return true
		}
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="Webmachine"`)
	return false
}
