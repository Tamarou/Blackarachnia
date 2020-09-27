package types

import (
	"net/http"
	"net/url"
	"time"
)

type Response interface {
	Body() string
	http.ResponseWriter
}

type Handler interface {
	ServeHTTP(Response, *http.Request) error
}

type HandlerFunc func(Response, *http.Request) error

func (f HandlerFunc) ServeHTTP(w Response, r *http.Request) error { return f(w, r) }

func NewHandlerFunc(f func(http.ResponseWriter, *http.Request) error) HandlerFunc {
	return HandlerFunc(func(w Response, r *http.Request) error { f(w, r); return nil })
}

type HandlerMap interface {
	Get(string) Handler
	FirstType() string
	Types() []string
}

type EmptyHandlerMap struct{}

func (e EmptyHandlerMap) Get(s string) Handler {
	return HandlerFunc(func(w Response, r *http.Request) error { return nil })
}
func (e EmptyHandlerMap) FirstType() string { return "" }
func (e EmptyHandlerMap) Types() []string   { return []string{} }

type Resource interface {
	Authorized(auth string) bool
	Forbidden() bool
	KnownContentType(contentType string) bool
	MalformedRequest(r *http.Request) bool
	KnownMethods() []string
	AllowedMethods() []string
	MultipleChoices() bool
	ServiceAvailable() bool
	URLTooLong(u *url.URL) bool
	ValidContentHeaders(r *http.Request) bool
	ValidEntityLength(length string) bool
	ContentTypesAccepted() HandlerMap
	ContentTypesProvided() HandlerMap
	LanguagesProvided() []string
	CharsetsProvided() []string
	DefaultCharset() string
	EncodingsProvided() []string
	Exists() bool
	ETAG() string
	LastModified() time.Time
	MovedPermanently() bool
	MovedTemporarily() bool
	Location() string
	PreviouslyExisted() bool
	AllowMissingPost() bool
	PostIsCreate() bool
	CreatePath() string
	ProcessPost() string
	IsConflict() bool
	Options(http.ResponseWriter)
	DeleteResource() bool
	DeleteCompleted() bool
}
