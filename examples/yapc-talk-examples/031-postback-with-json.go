/*
TODO strip away until we're closer to the example
{
    package YAPC::NA::2012::Example031::Resource;
    use strict;
    use warnings;
    use JSON::XS ();

    use base 'YAPC::NA::2012::Example031::Resource';

    sub allowed_methods        { [qw[ GET PUT POST ]] }
    sub content_types_accepted { [ { 'application/json' => 'from_json' } ] }

    sub from_json {
        my $self = shift;
        $self->save_message( JSON::XS->new->allow_nonref->decode( $self->request->content ) );
    }

    sub process_post {
        my $self = shift;
        return \415 unless $self->request->header('Content-Type')->match('application/x-www-form-urlencoded');
        $self->SUPER::process_post;
    }
}

*/
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
)

type YAPC_NA_2012_Example031_Resource struct {
	messages []string
	blackarachnia.Resource
}

func New_YAPC_NA_2012_Example031_Resource() *YAPC_NA_2012_Example031_Resource {
	return &YAPC_NA_2012_Example031_Resource{
		[]string{},
		blackarachnia.Resource{},
	}
}

func (res *YAPC_NA_2012_Example031_Resource) AllowedMethods() []string {
	return []string{"GET", "POST", "PUT"}
}

func (res *YAPC_NA_2012_Example031_Resource) ContentTypesAccepted() types.HandlerMap {
	return handlerMap.NewHandlerMap(handlerMap.Map("application/json", res.fromJSON))
}

func (res *YAPC_NA_2012_Example031_Resource) fromJSON(w http.ResponseWriter, r *http.Request) error {
	var m string
	if e := json.NewDecoder(r.Body).Decode(&m); e != nil {
		log.Fatal(e)
		return e
	}
	res.messages = append(res.messages, m)
	return nil
}

func (res *YAPC_NA_2012_Example031_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(handlerMap.Map("text/html", res.toHTML))
}

func (res *YAPC_NA_2012_Example031_Resource) toHTML(w http.ResponseWriter, r *http.Request) error {
	const tpl = `<html><body><form method="POST"><input type="text" name="message" /><input type="submit" /></form><hr/><ul>{{- range . -}}<li>{{- . -}}</li>{{- end -}}</ul></body></html>`
	t, e := template.New("webpage").Parse(tpl)
	if e != nil {
		return e
	}

	return t.Execute(w, res.messages)
}

func (res *YAPC_NA_2012_Example031_Resource) ProcessPost(w http.ResponseWriter, r *http.Request) error {
	res.messages = append(res.messages, r.FormValue("message"))
	w.Header().Set("Location", "/")
	http.Error(w, "See Other", http.StatusMovedPermanently) // this shoudl be a 303 but the original has 301
	return nil
}
