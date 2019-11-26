package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/Tamarou/blackarachnia/fsm"
)

type Resource struct{}

func (r Resource) Exists() bool                               { return true }
func (r Resource) ServiceAvailable() bool                     { return true }
func (r Resource) Authorized(auth string) bool                { return false }
func (r Resource) Forbidden() bool                            { return false }
func (r Resource) AllowMissingPost() bool                     { return false }
func (r Resource) MalformedRequest(req *http.Request) bool    { return false }
func (r Resource) URLTooLong(u *url.URL) bool                 { return false }
func (r Resource) KnownContentType(c string) bool             { return true }
func (r Resource) ValidContentHeaders(req *http.Request) bool { return true }
func (r Resource) ValidEntityLength(length string) bool       { return true }
func (r Resource) Options(w http.ResponseWriter)              {}

func (r Resource) KnownMethods() []string {
	return []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"TRACE",
		"CONNECT",
		"OPTION",
	}
}

func (r Resource) AllowedMethods() []string {
	return []string{
		"GET",
		"HEAD",
	}
}

func (r Resource) DeleteResource() bool           { return false }
func (r Resource) DeleteCompleted() bool          { return false }
func (r Resource) PostIsCreate() bool             { return false }
func (r Resource) CreatePath() string             { return "" }
func (r Resource) BaseURI() string                { return "" } // TODO see where this is used
func (r Resource) ProcessPost() string            { return "" }
func (r Resource) ContentTypesProvided() []string { return []string{} }
func (r Resource) ContentTypesAccepted() []string { return []string{} } // TODO see where this is used
func (r Resource) CharsetsProvided() []string     { return []string{} }
func (r Resource) DefaultCharset() string         { return "" }
func (r Resource) LanguagesProvided() []string    { return []string{} }
func (r Resource) EncodingsProvided() []string    { return []string{"identity"} }
func (r Resource) Variances() []string            { return []string{} } // TODO see where this is used
func (r Resource) IsConflict() bool               { return false }
func (r Resource) MultipleChoices() bool          { return false }
func (r Resource) PreviouslyExisted() bool        { return false }
func (r Resource) MovedPermanently() bool         { return false }
func (r Resource) MovedTemporarily() bool         { return false }
func (r Resource) LastModified() time.Time        { return time.Time{} }
func (r Resource) Expires() time.Time             { return time.Time{} }
func (r Resource) ETAG() string                   { return "" }
func (r Resource) FinishRequest()                 {} // TODO see where this is uesed

func (r Resource) Location() string { return "" }

func (r Resource) GetAcceptableContentTypeHandler(req *http.Request) fsm.ContentHandler {
	return func(w http.ResponseWriter, r *http.Request) bool {
		return true
	}
}
