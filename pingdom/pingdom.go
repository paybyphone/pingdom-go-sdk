package pingdom

import (
	"os"
	"strings"
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

	// The API endpoint
	Endpoint string

	// The proxy config, if any.
	Proxy string
}

// DefaultConfigProvider supplies a default configuration:
//  * Endpoint defaults to https://api.pingdom.com. Proxy is unset.
//  * EmailAddress defaults to PINGDOM_EMAIL_ADDRESS, if set, otherwise empty
//  * Password defaults to PINGDOM_PASSWORD, if set, otherwise empty
//  * AppKey defaults to PINGDOM_APP_KEY, if set, otherwise empty
//
// This essentially loads an initial config state for any given
// API service.
func DefaultConfigProvider() Config {
	env := os.Environ()
	cfg := Config{
		Endpoint: apiAddress,
	}

	for _, v := range env {
		d := strings.Split(v, "=")
		switch d[0] {
		case "PINGDOM_EMAIL_ADDRESS":
			cfg.EmailAddress = d[1]
		case "PINGDOM_PASSWORD":
			cfg.Password = d[1]
		case "PINGDOM_APP_KEY":
			cfg.AppKey = d[1]
		}
	}
	return cfg
}
