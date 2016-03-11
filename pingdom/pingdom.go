package pingdom

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

	// The API endpoint
	Endpoint string

	// The proxy config, if any.
	Proxy string
}
