package fsm

import (
	"net/http"
	"net/url"
	"time"
)

type ContentHandler func(w http.ResponseWriter, r *http.Request) bool

type Resource interface {
	Authorized(auth string) bool
	Body() []byte
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
	ContentTypesProvided() []string
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
	GetAcceptableContentTypeHandler(r *http.Request) ContentHandler
	Options(http.ResponseWriter)
	DeleteResource() bool
	DeleteCompleted() bool
}
