package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestClientNewStatic(t *testing.T) {
}

func TestClientNewEnv(t *testing.T) {
}
