package client

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

const errorResponse = "Forbidden (403): Something went wrong! This string describes what happened."

const okResponseText = `
 {
   "check": {
     "id": 138631,
     "name": "My new HTTP check"
   }
 }
 `

func pingdomConfig() pingdom.Config {
	return pingdom.Config{
		EmailAddress: "nobody@example.com",
		Password:     "changeit",
		AppKey:       "0123456789abcdefgh",
	}
}

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

func queryStringDataTestBasic() queryStringDataTestBasicType {
	return queryStringDataTestBasicType{
		ID:   1234,
		Name: "My new HTTP check",
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

func TestClientNew(t *testing.T) {
	c := New(pingdomConfig())

	if c.Config.Endpoint != "https://api.pingdom.com" {
		t.Fatalf("Expected Endpoint to be https://api.paybyphone.com, got %s", c.Config.Endpoint)
	}
	if c.Config.EmailAddress != "nobody@example.com" {
		t.Fatalf("Expected EmailAddress to be nobody@example.com, got %s", c.Config.EmailAddress)
	}
	if c.Config.Password != "changeit" {
		t.Fatalf("Expected Password to be changeit, got %s", c.Config.Password)
	}
	if c.Config.AppKey != "0123456789abcdefgh" {
		t.Fatalf("Expected AppKey to be 0123456789abcdefgh, got %s", c.Config.AppKey)
	}
}

func TestClientSendRequestSuccess(t *testing.T) {
	ts := httpOKTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := queryStringDataTestBasic()
	out := okResponseType{}
	err := c.SendRequest("GET", "/api/v2.0/test", &in, &out)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := okResponse()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestClientSendRequestError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := queryStringDataTestBasic()
	out := okResponseType{}
	err := c.SendRequest("GET", "/api/v2.0/test", &in, &out)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}
