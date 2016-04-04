// Copyright 2016 PayByPhone Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pingdom contains any top-level configuration structures
// necessary to work with the rest of the SDK and API.
package pingdom

import (
	"os"
	"strings"
)

// The pingdom API endpoint.
const apiAddress = "https://api.pingdom.com"

// Config contains the configuration for connecting to the Pingdom API.
//
//
// Supplying Configuration to Services
//
// All service constructors (ie: checks, contacts, etc) take zero or
// more of these structs as configuration, like so:
//
//   cfg := pingdom.Config{
//     EmailAddress: "jdoe@example.com",
//     Password:     "password",
//     AppKey:       "appkey",
//   }
//   svc := checks.New(cfg)
//
// Note that default options are set for EmailAddress, Password, and AppKey.
// See the DefaultConfigProvider method for more details.
type Config struct {
	// The email address for the Pingdom account.
	EmailAddress string

	// The password for the Pingdom account.
	Password string

	// The application key required for API requests.
	AppKey string

	// The API endpoint. Changing this is only recommended for testing.
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
