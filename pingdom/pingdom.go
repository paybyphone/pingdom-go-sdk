package pingdom

// Config - The configuration for the Pingdom API.
type Config struct {
	// The email address for the Pingdom account.
	EmailAddress string

	// The password for the Pingdom account.
	Password string

	// The application key required for API requests.
	AppKey string
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

// NewRequest - Create a new request instance with configuration set.
func NewRequest(c Config) *Request {
	r := &Request{
		config: c,
	}
	return r
}
