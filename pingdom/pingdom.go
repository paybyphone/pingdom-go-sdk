package pingdom

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// The pingdom API endpoint.
const apiAddress = "https://api.pingdom.com"

// Config - The configuration for the Pingdom API.
type Config struct {
	// The email address for the Pingdom account.
	EmailAddress string

	// The password for the Pingdom account.
	Password string

	// The application key required for API requests.
	AppKey string
}

// ErrorResponse - structure of a Pingdom JSON error response
type ErrorResponse {
	// The status code.
	StatusCode int

	// The status description.
	StatusDesc string

	// The error message.
	ErrorMessage string
}

// Request - the API request.
type Request struct {
	// The API configuration (user/pass/etc).
	config Config

	// The request method.
	Method string

	// The request URI.
	URI string

	// The request data.
	Input interface{}

	// The request query string data
	formData string

	// The JSON response for non-error responses.
	responseData string

	// The output of the request.
	Output interface{}
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
		req, err = http.NewRequest(r.Method, fmt.Sprintf("%s%s?%s", apiAddress, r.URI, qs), nil)
	case "POST", "PUT", "DELETE":
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%s", qs))
		req, err = http.NewRequest(r.Method, fmt.Sprintf("%s%s", apiAddress, r.URI), &buf)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	default:
		panic(fmt.Errorf("API request method %s not supported by Pingdom", r.Method))
	}

	if err != nil {
		panic(err)
	}

	req.Header.Add("App-Key", r.config.AppKey)
	req.SetBasicAuth(r.config.EmailAddress, r.config.Password)

	resp, err = client.Do(req)

	if err != nil {
		return fmt.Errorf("HTTP protocol error: %s", err)
	}

	// As of right now, every single Pingdom API request returns a 200 error
	// code on success. Anything else for now is an error, and needs to be
	// handled as such.
	if resp.StatusCode != 200 {
		return handleError(resp)
	}

	return nil
}

// handleError - handles a Pingdom API error response.
func handleError(r *http.Response) error {
	// error for redirect errors get a simple message
	if r.StatusCode >= 300 && r.StatusCode < 400 {
		return fmt.Errorf("%s", r.Status)
	}
	// Check to see if have a JSON response body with the correct
	// pingdom error format.
	var buf []byte
	_, err := r.Body.Read(buf)
	if err != nil {
		panic(err)
	}

	er := ErrorResponse{}
	err = json.Unmarshal(buf, &er)
	if err != nil {
		buf := 
		// more than likely not JSON, just pull together the body and return it as the error message
		return fmt.Errorf("%s %s", r.Status, 

	}

}

// NewRequest - Create a new request instance with configuration set.
func NewRequest(c Config) *Request {
	r := &Request{
		config: c,
	}
	return r
}
