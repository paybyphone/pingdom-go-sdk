package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/paybyphone/pingdom-go-sdk/pingdom"
)

// errorResponseErrorType is the actual error object within
// an Error Response.
type errorResponseErrorType struct {
	// The status code.
	StatusCode int

	// The status description.
	StatusDesc string

	// The error message.
	ErrorMessage string
}

// ErrorResponse - structure of a Pingdom JSON error response
type ErrorResponse struct {
	Error errorResponseErrorType
}

// Request - the API request.
type Request struct {
	// The API configuration (user/pass/etc).
	Config pingdom.Config

	// The request method.
	Method string

	// The request URI.
	URI string

	// The request data.
	Input interface{}

	// The output of the request.
	Output interface{}
}

// requestResponse - Unexported struct that encompasses status codes
// and request body in a fashion that can be read after the request
// is closed.
type requestResponse struct {
	// Status code.
	StatusCode int

	// Status code with short-form message.
	Status string

	// Response body.
	Body []byte
}

// BodyString - convert requestResponse.Body to string.
func (r *requestResponse) BodyString() string {
	buf := bytes.NewBuffer(r.Body)
	return buf.String()
}

// readResponseJSON - Read the response body as JSON into variable
// pointed to by v.
func (r *requestResponse) ReadResponseJSON(v interface{}) error {
	err := json.Unmarshal(r.Body, v)
	if err != nil {
		return fmt.Errorf("JSON parsing error: %s", err)
	}
	return nil
}

// newRequestResponse - Create a new requestResponse instance off a HTTP
// response. Warning: This also closes the Body.
func newRequestResponse(r *http.Response) *requestResponse {
	rr := &requestResponse{
		StatusCode: r.StatusCode,
		Status:     r.Status,
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	rr.Body = body
	return rr
}

// Take an interface{} and convert to a x-www-urlencoded query string,
// suitable for use within Pingdom GET/POST/PUT/DELETE requests.
// Returns runtime panic if for some reason d is not a struct.
func dataToQueryString(d interface{}) string {
	v, err := query.Values(d)
	if err != nil {
		panic(err)
	}
	return v.Encode()
}

// Send - Sends a request to the API endpoint, and parse response.
func (r *Request) Send() error {
	qs := dataToQueryString(r.Input)
	var req *http.Request
	var err error
	client := &http.Client{}

	switch r.Method {
	case "GET":
		req, err = http.NewRequest(r.Method, fmt.Sprintf("%s%s?%s", r.Config.Endpoint, r.URI, qs), nil)
	case "POST", "PUT", "DELETE":
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%s", qs))
		req, err = http.NewRequest(r.Method, fmt.Sprintf("%s%s", r.Config.Endpoint, r.URI), &buf)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	default:
		return fmt.Errorf("API request method %s not supported by Pingdom", r.Method)
	}

	if err != nil {
		panic(err)
	}

	req.Header.Add("App-Key", r.Config.AppKey)
	req.SetBasicAuth(r.Config.EmailAddress, r.Config.Password)

	re, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("HTTP protocol error: %s", err)
	}

	resp := newRequestResponse(re)

	// As of right now, every single Pingdom API request returns a 200 error
	// code on success. Anything else for now is an error, and needs to be
	// handled as such.
	if resp.StatusCode != 200 {
		return handleError(resp)
	}

	// Unmarshal response into Output. The service is responsible for
	// this being functional past JSON parsing.
	err = resp.ReadResponseJSON(r.Output)
	if err != nil {
		return err
	}

	return nil
}

//Error. handleError - handles a Pingdom API error response.
func handleError(r *requestResponse) error {
	er := ErrorResponse{}
	err := r.ReadResponseJSON(&er)
	if err != nil {
		// more than likely not JSON, just pull together the body and return it as
		// the error messagea
		return fmt.Errorf("Non-API error (%s): %s", r.Status, r.BodyString())
	}

	// Return a properly formatted error from the appropraite fields.
	return fmt.Errorf("%s (%d): %s", er.Error.StatusDesc, r.StatusCode, er.Error.ErrorMessage)
}

// NewRequest - Create a new request instance with configuration set.
func NewRequest(c pingdom.Config) *Request {
	r := &Request{
		Config: c,
	}
	return r
}
