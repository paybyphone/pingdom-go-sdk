package request

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/paybyphone/pingdom-go-sdk/pingdom"
)

const errorResponseText = `
{
	"error": {
		"statuscode": 403,
		"statusdesc": "Forbidden",
		"errormessage": "Something went wrong! This string describes what happened."
	}
}
`

const errorResponseNonJSONText = "<html><head><title>Service Unavailable</title></head><body><b>Service Unavailable</b></body></html>"

const okResponseText = `
{
	"check": {
		"id": 138631,
		"name": "My new HTTP check"
	}
}
`

type okResponseCheckType struct {
	ID   int
	Name string
}

type okResponseType struct {
	Check okResponseCheckType
}

func okResponse() okResponseType {
	return okResponseType{
		Check: okResponseCheckType{
			ID:   138631,
			Name: "My new HTTP check",
		},
	}
}

const errorResponse = "Forbidden (403): Something went wrong! This string describes what happened."

func errorResponseNonJSON() string {
	return fmt.Sprintf("Non-API error (503 Service Unavailable): %s", errorResponseNonJSONText)
}

type queryStringDataTestBasicType struct {
	ID   int    `url:"id"`
	Name string `url:"name"`
}

type queryStringDataTestNumberedArrayType struct {
	queryStringDataTestBasicType
	RequestHeader []string `url:"requestheader,numbered"`
}

type queryStringDataTestSemicolonArrayType struct {
	queryStringDataTestBasicType
	AdditionalURLs []string `url:"additionalurls,semicolon"`
}

func queryStringDataTestBasic() queryStringDataTestBasicType {
	return queryStringDataTestBasicType{
		ID:   1234,
		Name: "My new HTTP check",
	}
}

func queryStringDataTestNumberedArray() queryStringDataTestNumberedArrayType {
	return queryStringDataTestNumberedArrayType{
		queryStringDataTestBasicType: queryStringDataTestBasic(),
		RequestHeader:                []string{"MyHeaderA:CoolValueA", "MyHeaderB:CoolValueB"},
	}
}

func queryStringDataTestSemicolonArray() queryStringDataTestSemicolonArrayType {
	return queryStringDataTestSemicolonArrayType{
		queryStringDataTestBasicType: queryStringDataTestBasic(),
		AdditionalURLs:               []string{"www.mysite.com", "www.myothersite.com"},
	}
}

func newHTTPTestServer(f func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(f))
	return ts
}

func httpErrorTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, errorResponseText, http.StatusForbidden)
	})
}

func httpNonJSONErrorTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		http.Error(w, errorResponseNonJSONText, http.StatusServiceUnavailable)
	})
}

func httpOKTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, okResponseText, http.StatusOK)
	})
}

func pingdomConfig() pingdom.Config {
	return pingdom.Config{
		EmailAddress: "nobody@example.com",
		Password:     "changeit",
		AppKey:       "0123456789abcdefgh",
	}
}

func testRequest(c pingdom.Config, in interface{}, out interface{}) *Request {
	r := NewRequest(c)
	r.Method = "POST"
	r.URI = "/api/2.0/checks"
	r.Input = in
	r.Output = out
	return r
}

func testRequestGet(c pingdom.Config, in interface{}, out interface{}) *Request {
	r := NewRequest(c)
	r.Method = "GET"
	r.URI = "/api/2.0/checks"
	r.Input = in
	r.Output = out
	return r
}

func testRequestOptions(c pingdom.Config, in interface{}, out interface{}) *Request {
	r := NewRequest(c)
	r.Method = "OPTIONS"
	r.URI = "/api/2.0/checks"
	r.Input = in
	r.Output = out
	return r
}

func TestDataToQueryStringBaisc(t *testing.T) {
	in := queryStringDataTestBasic()
	out := dataToQueryString(in)
	expected := "id=1234&name=My+new+HTTP+check"

	if out != expected {
		t.Fatalf("expected %s, got %s", expected, out)
	}
}

func TestDataToQueryStringNumberedArray(t *testing.T) {
	in := queryStringDataTestNumberedArray()
	out := dataToQueryString(in)
	expected := "id=1234&name=My+new+HTTP+check&requestheader0=MyHeaderA%3ACoolValueA&requestheader1=MyHeaderB%3ACoolValueB"

	if out != expected {
		t.Fatalf("expected %s, got %s", expected, out)
	}
}

func TestDataToQueryStringSemicolonArray(t *testing.T) {
	in := queryStringDataTestSemicolonArray()
	out := dataToQueryString(in)
	expected := "additionalurls=www.mysite.com%3Bwww.myothersite.com&id=1234&name=My+new+HTTP+check"

	if out != expected {
		t.Fatalf("expected %s, got %s", expected, out)
	}
}

func TestRequestSendSuccess(t *testing.T) {
	ts := httpOKTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	in := queryStringDataTestBasic()
	out := okResponseType{}
	r := testRequest(cfg, &in, &out)
	err := r.Send()

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := okResponse()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestRequestSendSuccessGet(t *testing.T) {
	ts := httpOKTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	in := queryStringDataTestBasic()
	out := okResponseType{}
	r := testRequestGet(cfg, &in, &out)
	err := r.Send()

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := okResponse()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestRequestSendError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	in := queryStringDataTestBasic()
	out := okResponseType{}
	r := testRequest(cfg, &in, &out)
	err := r.Send()

	if err == nil {
		t.Fatalf("Expected error, got success")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestRequestSendNonJSONError(t *testing.T) {
	ts := httpNonJSONErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	in := queryStringDataTestBasic()
	out := okResponseType{}
	r := testRequest(cfg, &in, &out)
	err := r.Send()

	if err == nil {
		t.Fatalf("Expected error, got success")
	}

	expected := errorResponseNonJSON()

	// HTTP server gives a bunch of whitespace after for some reason
	if strings.TrimSpace(err.Error()) != expected {
		t.Fatalf("expected %s (%T), got %s (%T)", expected, expected, err.Error(), err.Error())
	}
}

func TestRequestSendProtocolError(t *testing.T) {
	ts := httpOKTestServer()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	in := queryStringDataTestBasic()
	out := okResponseType{}
	r := testRequest(cfg, &in, &out)
	ts.Close()
	err := r.Send()

	if err == nil {
		t.Fatalf("Expected error, got success")
	}

	expected := "^HTTP protocol error"

	if ok, _ := regexp.MatchString(expected, err.Error()); ok == false {
		t.Fatalf("expected error to matchi %s, got %s", expected, err)
	}
}

func TestRequestUnsupportedMethodError(t *testing.T) {
	ts := httpOKTestServer()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	in := queryStringDataTestBasic()
	out := okResponseType{}
	r := testRequestOptions(cfg, &in, &out)
	ts.Close()
	err := r.Send()

	if err == nil {
		t.Fatalf("Expected error, got success")
	}

	expected := "API request method OPTIONS not supported by Pingdom"

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestErrorResponse(t *testing.T) {
	rr := &requestResponse{
		StatusCode: 403,
		Status:     "403 Forbidden",
		Body:       []byte(errorResponseText),
	}
	er := ErrorResponse{}
	rr.ReadResponseJSON(&er)

	if er.Error.StatusCode != 403 {
		t.Fatalf("Expected StatusCode to be 403 (int), got %v", er.Error.StatusCode)
	}
	if er.Error.StatusDesc != "Forbidden" {
		t.Fatalf("Expected StatusCode to be Forbidden (string), got %v", er.Error.StatusDesc)
	}
	if er.Error.ErrorMessage != "Something went wrong! This string describes what happened." {
		t.Fatalf("Expected StatusCode to be \"Something went wrong! This string describes what happened.\", got %v", er.Error.ErrorMessage)
	}
}
