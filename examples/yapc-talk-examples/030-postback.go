/*
{
    package YAPC::NA::2012::Example030::Resource;
    use strict;
    use warnings;

    use parent 'Web::Machine::Resource';

    our @MESSAGES = ();
    sub save_message { push @MESSAGES => $_[1] }
    sub get_messages { @MESSAGES }

    sub allowed_methods        { [qw[ GET POST ]] }
    sub content_types_provided { [ { 'text/html' => 'to_html' } ] }

    sub to_html {
        my $self = shift;
        '<html><body><form method="POST"><input type="text" name="message" />'
        . '<input type="submit" /></form><hr/><ul>'
        . (join '' => map { '<li>' . $_ . '</li>' } $self->get_messages)
        . '</ul></body></html>'
    }

    sub process_post {
        my $self = shift;
        $self->save_message( $self->request->param('message') );
        $self->response->header('Location' => '/');
        return \301;
    }

}
*/

package main

import (
	"html/template"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
)

var messages = []string{}

type YAPC_NA_2012_Example030_Resource struct {
	blackarachnia.Resource
}

func (res YAPC_NA_2012_Example030_Resource) AllowedMethods() []string {
	return []string{"GET", "POST"}
}

func (res YAPC_NA_2012_Example030_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(handlerMap.Map("text/html", res.toHTML))
}

func (res YAPC_NA_2012_Example030_Resource) toHTML(w http.ResponseWriter, r *http.Request) error {
	const tpl = `<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul>{{- range . -}}<li>{{- . -}}</li>{{- end -}}</ul></body></html>`
	t, e := template.New("webpage").Parse(tpl)
	if e != nil {
		return e
	}

	return t.Execute(w, messages)
}

func (res YAPC_NA_2012_Example030_Resource) ProcessPost(w http.ResponseWriter, r *http.Request) error {
	messages = append(messages, r.FormValue("message"))
	w.Header().Set("Location", "/")
	http.Error(w, "See Other", http.StatusMovedPermanently) // this shoudl be a 303 but the original has 301
	return nil
}
