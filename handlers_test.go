package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Tamarou/blackarachnia/fsm"
)

type TestResource struct {
	disabled              bool
	methods               []string
	allowedMethods        []string
	malformed             bool
	unauthorized          bool
	forbidden             bool
	invalidContentHeaders bool
	unknownContentType    bool
	Resource
}

func (sr TestResource) KnownContentType(c string) bool           { return !sr.unknownContentType }
func (sr TestResource) Body() []byte                             { return []byte("Hello World!\n") }
func (sr TestResource) ServiceAvailable() bool                   { return !sr.disabled }
func (sr TestResource) KnownMethods() []string                   { return sr.methods }
func (sr TestResource) AllowedMethods() []string                 { return sr.allowedMethods }
func (sr TestResource) MalformedRequest(r *http.Request) bool    { return sr.malformed }
func (sr TestResource) Authorized(auth string) bool              { return !sr.unauthorized }
func (sr TestResource) Forbidden() bool                          { return sr.forbidden }
func (sr TestResource) ValidContentHeaders(r *http.Request) bool { return !sr.invalidContentHeaders }
func (sr TestResource) LastModified() time.Time                  { return time.Now() }

func (sr TestResource) GetAcceptableContentTypeHandler(r *http.Request) fsm.ContentHandler {
	return func(w http.ResponseWriter, r *http.Request) bool {
		w.WriteHeader(http.StatusOK)
		w.Write(sr.Body())
		return true
	}
}

func TestHandlers(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/", nil)

	t.Run("Service Unavailble", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{disabled: true}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		got := rr.Code
		want := http.StatusServiceUnavailable
		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}

	})

	t.Run("Service Available", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		got := rr.Code
		want := http.StatusServiceUnavailable
		if got == want {
			t.Errorf("got %d wanted anything but %d", got, want)
		}

	})

	t.Run("Unknown Method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		got := rr.Code
		want := http.StatusNotImplemented
		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	})

	t.Run("Known Method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			methods: []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		got := rr.Code
		want := http.StatusNotImplemented
		if got == want {
			t.Errorf("got %d wanted anything but %d", got, want)
		}
	})

	t.Run("Allowed Method", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			methods:        []string{"GET", "HEAD"},
			allowedMethods: []string{"HEAD"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		got := rr.Code
		want := http.StatusMethodNotAllowed
		if got != want {
			t.Errorf("got %d wanted %d", rr.Code, want)
		}
		if allowed := rr.Header().Get("Allow"); allowed != "HEAD" {
			t.Errorf("got %v wanted %v", allowed, "HEAD")
		}
	})

	t.Run("Malformed Request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			malformed:      true,
			methods:        []string{"GET"},
			allowedMethods: []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		if code := rr.Code; code != http.StatusBadRequest {
			t.Errorf("got %d wanted %d", code, http.StatusBadRequest)
		}
	})

	t.Run("Unauthorized", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			unauthorized:   true,
			methods:        []string{"GET"},
			allowedMethods: []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		if code := rr.Code; code != http.StatusUnauthorized {
			t.Errorf("got %d wanted %d", code, http.StatusUnauthorized)
		}
	})

	t.Run("Forbidden", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			forbidden:      true,
			methods:        []string{"GET"},
			allowedMethods: []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		if code := rr.Code; code != http.StatusForbidden {
			t.Errorf("got %d wanted %d", code, http.StatusForbidden)
		}
	})

	t.Run("Valid Content Headers", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			invalidContentHeaders: true,
			methods:               []string{"GET"},
			allowedMethods:        []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		if code := rr.Code; code != http.StatusNotImplemented {
			t.Errorf("got %d wanted %d", code, http.StatusNotImplemented)
		}
	})

	t.Run("Unknown Content Type", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := TestResource{
			unknownContentType: true,
			methods:            []string{"GET"},
			allowedMethods:     []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		if code := rr.Code; code != http.StatusUnsupportedMediaType {
			t.Errorf("got %d wanted %d", code, http.StatusUnsupportedMediaType)
		}
	})

	t.Run("200 OK", func(t *testing.T) {
		rr := httptest.NewRecorder()

		r := TestResource{
			methods:        []string{"GET"},
			allowedMethods: []string{"GET"},
		}
		handler := http.HandlerFunc(NewHandler(r))
		handler.ServeHTTP(rr, request)

		got := rr.Body.String()
		want := "Hello World!\n"

		if code := rr.Code; code != http.StatusOK {
			t.Errorf("got %v wanted %v", code, http.StatusOK)
		}

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
}
