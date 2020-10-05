/*
#!perl

use strict;
use warnings;

use Web::Machine;

=pod

And of course, you don't have to just provide
text based results ...

=cut

{
    package YAPC::NA::2012::Example012::Resource;
    use strict;
    use warnings;
    use JSON::XS ();
    use GD::Simple;

    use parent 'Web::Machine::Resource';

    sub content_types_provided { [
        { 'image/gif' => 'to_gif'  },
        { 'text/html' => 'to_html' },
    ] }

    sub to_html {
        my $self = shift;
        '<html><body><ul>' .
            (join "" => map {
                '<li>' . $_->[0] . ' &mdash; ' . $_->[1]->type . '</li>'
            } $self->request->header('Accept')->iterable)
        . '</ul><br/><img src="/hello_world.gif" border="1"/></body></html>'
    }

    sub to_gif {
        my $self = shift;
        my $img  = GD::Simple->new( 130, 20 );
        $img->fgcolor('red');
        $img->moveTo(15, 15);
        $img->string( $self->request->path_info );
        $img->gif;
    }
}

Web::Machine->new( resource => 'YAPC::NA::2012::Example012::Resource' )->to_app;
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/Tamarou/blackarachnia/handlerMap"
	"github.com/Tamarou/blackarachnia/types"
	"github.com/adjust/goautoneg"
)

type YAPC_NA_2012_Example012_Resource struct {
	blackarachnia.Resource
}

func (res YAPC_NA_2012_Example012_Resource) ContentTypesProvided() types.HandlerMap {
	return handlerMap.NewHandlerMap(
		handlerMap.Map("image/gif", res.toGIF),
		handlerMap.Map("text/html", res.toHTML),
	)
}

func (res YAPC_NA_2012_Example012_Resource) toHTML(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "<html><body><ul>")
	for _, a := range goautoneg.ParseAccept(r.Header.Get("Accept")) {
		fmt.Fprintf(w, "<li>%f &mdash; %s</li>", a.Q, a.Type)
	}
	fmt.Fprint(w, "</ul></body></html>")
	return nil
}

func (res YAPC_NA_2012_Example012_Resource) toGIF(w http.ResponseWriter, r *http.Request) error {
	pal := []color.Color{color.Black, color.Transparent}
	rect := image.Rect(0, 0, 130, 20)
	img := image.NewPaletted(rect, pal)
	img.SetColorIndex(130/2, 20/2, 1)
	anim := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}
	gif.EncodeAll(w, &anim)
	return nil
}
