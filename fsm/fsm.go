package fsm

import (
	"net/http"
	"strings"
	"time"
)

type State func(res Resource, w http.ResponseWriter, r *http.Request) (next State)

func chooseMediaType(accept string, choices []string) string {
	return ""
}

func chooseLanguage(accept string, choices []string) string {
	return ""
}

func chooseCharset(accept string, choices []string) string {
	return "text/plain"
}

func chooseEncoding(accept string, choices []string) string {
	return "identity"
}

/*
sub _add_caching_headers {
    my ($resource, $response) = @_;
    if ( my $etag = $resource->generate_etag ) {
        $response->header( 'Etag' => _ensure_quoted_header( $etag ) );
    }
    if ( my $expires = $resource->expires ) {
        $response->header( 'Expires' => $expires );
    }
    if ( my $modified = $resource->last_modified ) {
        $response->header( 'Last-Modified' => $modified );
    }
}

sub _handle_304 {
    my ($resource, $response) = @_;
    $response->headers->remove_header('Content-Type');
    $response->headers->remove_header('Content-Encoding');
    $response->headers->remove_header('Content-Language');
    _add_caching_headers($resource, $response);
    return \304;
}
*/
func handle304(res Resource, w http.ResponseWriter) {
	// remove content headers
	// setup caching headers
	http.Error(w, "Not Modified", http.StatusNotModified)
}

func etagInList(etag string, header string) bool {
	return strings.Contains(header, etag)
}

func methodInList(method string, methodList []string) bool {
	for _, known := range methodList {
		if known == method {
			return true
		}
	}
	return false
}

func InitialState() State {
	return b13
}

// service available
func b13(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.ServiceAvailable() {
		next = b12
	} else {
		http.Error(w, "Not Available", http.StatusServiceUnavailable)
	}
	return
}

//  method impelemented
func b12(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	// TODO verify this
	if methodInList(r.Method, res.KnownMethods()) {
		next = b11
	} else {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	}
	return
}

// URL Too Long
func b11(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.URLTooLong(r.URL) {
		http.Error(w, "URI Too Long", http.StatusRequestURITooLong)
	} else {
		next = b10
	}
	return
}

// method allowed
func b10(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	// TODO verify this
	if methodInList(r.Method, res.AllowedMethods()) {
		next = b9
	} else {
		allow := strings.Join(res.AllowedMethods(), ",")
		w.Header().Set("Allow", allow)
		http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
	}
	return
}

// malformed request
func b9(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.MalformedRequest(r) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		next = b8
	}
	return
}

// is authorized
func b8(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.Authorized(r.Header.Get("Authorization")) {
		next = b7
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	return
}

// forbidden
func b7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.Forbidden() {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		next = b6
	}
	return
}

// content headers ok
func b6(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.ValidContentHeaders(r) {
		next = b5
	} else {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	}
	return
}

// known content type
func b5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.KnownContentType(r.Header.Get("Content-Type")) {
		next = b4
	} else {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
	}
	return
}

// request entity too large
func b4(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.ValidEntityLength(r.Header.Get("Content-Length")) {
		next = b3
	} else {
		http.Error(w, "Request Entity Too Large", http.StatusRequestEntityTooLarge)
	}
	return
}

// method_is_options
func b3(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "OPTIONS" {
		res.Options(w)
		w.WriteHeader(http.StatusOK)
	} else {
		next = c3
	}
	return
}

/*
$STATE_DESC{'c3'} = 'accept_header_exists';
sub c3 {
    my ($resource, $request, $response) = @_;
    my $metadata = _metadata($request);
    if ( !$request->header('Accept') ) {
        $metadata->{'Content-Type'} = create_header( MediaType => (
            pair_key( $resource->content_types_provided->[0] )
        ));
        return \&d4
    }
    return \&c4;
}
*/
// accept_header_exists
func c3(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("Accept") == "" {
		// set Content-Type from first ContentTypeProfided
		next = d4
	} else {
		next = c4
	}
	return
}

// acceptable_media_type_available
func c4(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	ct := chooseMediaType(r.Header.Get("Accept"), res.ContentTypesProvided())
	if ct != "" {
		w.Header().Set("Content-Type", ct)
		next = d4
	} else {
		http.Error(w, "Content-Type Not Acceptable", http.StatusNotAcceptable)
	}
	return
}

// accept_language_header_exists
func d4(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("Accept-Language") == "" {
		next = e5
	} else {
		next = d5
	}
	return
}

// accept_language_choice_available
func d5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	lang := chooseLanguage(r.Header.Get("Accept-Language"), res.LanguagesProvided())
	if lang != "" {
		w.Header().Set("Content-Language", lang)
		next = e5
	} else {
		http.Error(w, "No Acceptable Language", http.StatusNotAcceptable)
	}
	return
}

// accept_charset_exists
func e5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("Accept-Charset") == "" {
		next = e6
	} else {
		next = f6
	}
	return
}

// accept_charset_choice_available
func e6(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	var charset string
	if r.Header.Get("Accept-Charset") == "" {
		charset = res.DefaultCharset()
	} else {
		charset = chooseCharset(r.Header.Get("Accept-Charset"), res.CharsetsProvided())
		if charset == "" {
			http.Error(w, "No Acceptable Charset", http.StatusNotAcceptable)
			return // escape early
		}
	}
	next = f6
	// reset the Content-Type with the charset here
	return
}

// accept_encoding_exists
func f6(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("Accept-Encoding") != "" {
		next = f7
	} else {
		encoding := chooseEncoding("identity;q=1.0,*;q=0.5", res.EncodingsProvided())
		if encoding != "" {
			if encoding != "identity" {
				w.Header().Set("Content-Encoding", encoding)
			}
			next = g7
		} else {
			http.Error(w, "No Acceptable Encoding", http.StatusNotAcceptable)
		}
	}
	return
}

// accept_encoding_choice_available
func f7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	encoding := chooseEncoding(r.Header.Get("Accept-Encoding"), res.EncodingsProvided())
	if encoding != "" {
		if encoding != "identity" {
			w.Header().Set("Content-Encoding", encoding)
		}
		next = g7
	} else {
		http.Error(w, "No Acceptable Encoding", http.StatusNotAcceptable)
	}
	return
}

// resource_exists
func g7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	// set the variances here since we've finished content negotiation
	if res.Exists() {
		next = g8
	} else {
		next = h7
	}
	return
}

// if_match_exists
func g8(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-Match") != "" {
		next = g9
	} else {
		next = h10
	}
	return
}

// if_match_is_wildcard
func g9(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-Match") != "*" {
		next = g11
	} else {
		next = h10
	}
	return
}

// etag_in_if_match_list
func g11(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if etagInList(res.ETAG(), r.Header.Get("If-Match")) {
		next = h10
	} else {
		http.Error(w, "Precondition Failed", http.StatusPreconditionFailed)
	}
	return
}

// if_match_exists_and_if_match_is_wildcard
func h7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-Match") != "*" {
		http.Error(w, "Precondition Failed", http.StatusPreconditionFailed)
	} else {
		next = i7
	}
	return
}

// if_unmodified_since_exists
func h10(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-Unmodified-Since") != "" {
		next = h11
	} else {
		next = i12
	}
	return
}

// if_unmodified_since_is_valid_date
func h11(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	_, err := time.Parse(time.RFC3339, r.Header.Get("If-Unmodified-Since"))
	if err == nil {
		next = h12
	} else {
		next = i12
	}
	return
}

// last_modified_is_greater_than_if_unmodified_since
func h12(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	date, _ := time.Parse(time.RFC3339, r.Header.Get("If-Unmodified-Since"))
	if date.Before(res.LastModified()) {
		next = i12
	} else {
		http.Error(w, "Precondition Failed", http.StatusPreconditionFailed)
	}
	return
}

// moved_permanently
func i4(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.MovedPermanently() {
		w.Header().Set("Location", res.Location())
		http.Error(w, "Moved Permanently", http.StatusMovedPermanently)
	} else {
		next = p3
	}
	return
}

// method_is_put
func i7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "PUT" {
		next = i4
	} else {
		next = k7
	}
	return
}

// if_none_match_exists
func i12(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-None-Match") != "" {
		next = i13
	} else {
		next = l13
	}
	return
}

// if_none_match_is_wildcard
func i13(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-None-Match") == "*" {
		next = j18
	} else {
		next = k13
	}
	return
}

// method_is_get_or_head
func j18(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "GET" || r.Method == "HEAD" {
		handle304(res, w)
	} else {
		http.Error(w, "Precondition Failed", http.StatusPreconditionFailed)
	}
	return
}

// moved_permanently
func k5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.MovedPermanently() {
		w.Header().Set("Location", res.Location())
		http.Error(w, "Moved Permanently", http.StatusMovedPermanently)
	} else {
		next = l5
	}
	return
}

// previously_existed
func k7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.PreviouslyExisted() {
		next = k5
	} else {
		next = l7
	}
	return
}

// etag_in_if_none_match
func k13(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if etagInList(res.ETAG(), r.Header.Get("If-None-Match")) {
		next = j18
	} else {
		next = l13
	}
	return
}

// moved_temporarily
func l5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.MovedTemporarily() {
		w.Header().Set("Location", res.Location())
		http.Error(w, "Moved Temporarily", http.StatusTemporaryRedirect)
	} else {
		next = m5
	}
	return
}

// method_is_post
func l7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "POST" {
		next = m7
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
	return
}

// if_modified_since_exists
func l13(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Header.Get("If-Modified-Since") != "" {
		next = l14
	} else {
		next = m16
	}
	return
}

// if_modified_since_is_valid_date
func l14(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if _, e := http.ParseTime(r.Header.Get("If-Modified-Since")); e != nil {
		return m16
	}
	return l15
}

// if_modified_since_greater_than_now
func l15(res Resource, w http.ResponseWriter, r *http.Request) (next State) {

	date, _ := http.ParseTime(r.Header.Get("If-Modified-Since"))
	if date.After(time.Now()) {
		return m16
	}
	return l17
}

// last_modified_is_greater_than_if_modified_since
func l17(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	date, _ := http.ParseTime(r.Header.Get("If-Modified-Since"))
	if res.LastModified().After(date) {
		return m16
	}
	handle304(res, w)
	return
}

// method_is_post
func m5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "POST" {
		next = n5
	} else {
		http.Error(w, "Gone", http.StatusGone)
	}
	return
}

// allow_post_to_missing_resource
func m7(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.AllowMissingPost() {
		next = n11
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
	return
}

// method_is_delete
func m16(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "DELETE" {
		next = m20
	} else {
		next = n16
	}
	return
}

// delete_enacted_immediately
func m20(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.DeleteResource() {
		next = m20b
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return
}

// did_delete_Complete
func m20b(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.DeleteCompleted() {
		next = o20
	} else {
		http.Error(w, "Accepted", http.StatusAccepted)
	}
	return o18
}

// allow_post_to_missing_resource
func n5(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.AllowMissingPost() {
		next = n11
	} else {
		http.Error(w, "Gone", http.StatusGone)
	}
	return
}

// redirect
func n11(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.PostIsCreate() {
		if URL := res.CreatePath(); URL != "" {
			w.Header().Set("Location", URL)
		}
	} else {
		if body := res.ProcessPost(); body == "" {
			panic("Process Post Invalid")
		}
	}
	if w.Header().Get("Location") != "" {
		http.Error(w, "See Other", http.StatusSeeOther)
		return
	} else {
		panic("Bad Redirect")
	}

	next = p11
	return
}

// method is post
func n16(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "POST" {
		next = n11
	} else {
		next = o16
	}
	return
}

// in conflict
func o14(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.IsConflict() {
		http.Error(w, "Conflict", http.StatusConflict)
		return
	}

	if handler := res.GetAcceptableContentTypeHandler(r); handler != nil {
		_ = handler(w, r)
		return
	}
	return p11
}

// method is put
func o16(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if r.Method == "PUT" {
		next = o14
	} else {
		next = o18
	}
	return
}

// multiple_representations
func o18(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.MultipleChoices() {
		w.WriteHeader(http.StatusMultipleChoices)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	if handler := res.GetAcceptableContentTypeHandler(r); handler != nil {
		_ = handler(w, r)
	}
	return
}

// response body includes entity
func o20(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.Body() != nil {
		next = o18
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
	return
}

// in conflict
func p3(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if res.IsConflict() {
		http.Error(w, "CONFLICT", http.StatusConflict)
		return
	}
	if handler := res.GetAcceptableContentTypeHandler(r); handler != nil {
		if !handler(w, r) {
			next = p11
		}
	} else {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}
	return
}

// new resource
func p11(res Resource, w http.ResponseWriter, r *http.Request) (next State) {
	if w.Header().Get("Location") != "" {
		http.Error(w, "CREATED", http.StatusCreated)
	} else {
		next = o20
	}
	return
}
