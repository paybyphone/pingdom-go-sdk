package request

import (
	"net/http"
	"net/http/httptest"
	"reflect"
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
