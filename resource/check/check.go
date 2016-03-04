package check

// GetDetailedCheckInput - Input to send to the detailed check method.
type GetDetailedCheckInput struct {
	_ struct{}

	// The ID of the check that you want to get a description for.
	CheckID int
}

// GetDetailedCheckOutput - Output for the detailed check method.
type GetDetailedCheckOutput struct {
	_ struct{}

	// The check identifier.
	ID int

	// The check name.
	Name string

	// The target host.
	Hostname string

	// The current check status.
	Status string

	// How often the check should be checked, in minutes.
	Resolution int

	// Contains one element representing the type of check and
	// type-specific settings.
	Type interface{}

	// A list of contact IDs that receive alerts.
	ContactIds []int

	// Send alerts as email.
	SendToEmail bool

	// Send alerts as SMS.
	SendToSMS bool

	// Send alerts through Twitter.
	SendToTwitter bool

	// Send alerts to iPhone.
	SendToIphone bool

	// Send alerts to Android.
	SendToAndroid bool

	// The failure count threshold to send notifications on.
	SendNotificationWhenDown int

	// The check frequency to notify on after a service has failed.
	NotifyAgainEvery int

	// Send a notification after a failed check resolves itself.
	NotifyWhenBackUp bool

	// Timestamp of last error (if any). Format is UNIX timestamp.
	LastErrorTime int

	// Timestamp of last test (if any). Format is UNIX timestamp
	LatestTime int

	// Response time (in milliseconds) of last test.
	LastResponseTime int

	// The time the check was created (UNIX timestamp).
	Created int

	// The check uses IPv6 instead of IPv4.
	IPv6 bool
}

// GetDetailedCheckOutputHTTP - Output for the HTTP check type,
// detailed check method.
type GetDetailedCheckOutputHTTP struct {
	_ struct{}

	// Path to the target on the server.
	URL string

	// true if the connection to the server is encrypted.
	Encryption bool

	// Target port to connect to on the server.
	Port int

	// Username for HTTP authentication.
	Username string

	// Password for HTTP authentication.
	Password string

	// A string the target response should contain.
	ShouldContain string

	// A string the target response should not contain.
	ShouldNotContain string

	// Data that should be posted to the web page, for example submission data
	// for a sign-up or login form. The data needs to be formatted in the same
	// way as a web browser would send it to the web server.
	PostData string

	// Custom headers to send with the HTTP request.
	RequestHeaders map[string]string
}

// GetDetailedCheckOutputHTTPCustom - Output for the Custom HTTP check type,
// detailed check method.
type GetDetailedCheckOutputHTTPCustom struct {
	_ struct{}

	// Path to the target XML file on the server.
	URL string

	// true if the connection to the server is encrypted.
	Encryption bool

	// Target port to connect to on the server.
	Port int

	// Username for HTTP authentication.
	Username string

	// Password for HTTP authentication.
	Password string

	// Additional URLs to target.
	AdditionalURLs []string
}

// GetDetailedCheckOutputTCP - Output for the TCP check type,
// detailed check method.
type GetDetailedCheckOutputTCP struct {
	_ struct{}

	// Path to the target XML file on the server.
	Port int

	// String to send to the server.
	StringToSend string

	// String to expect in response.
	StringToExpect string
}

// GetDetailedCheckOutputPing - Output for the Ping check type,
// detailed check method.
type GetDetailedCheckOutputPing struct {
	_ struct{}
}

// GetDetailedCheckOutputDNS - Output for the DNS check type,
// detailed check method.
type GetDetailedCheckOutputDNS struct {
	_ struct{}

	// DNS server to use.
	DNSServer string

	// Expected IP address from the query.
	ExpectedIP string
}

// GetDetailedCheckOutputUDP - Output for the UDP check type,
// detailed check method.
type GetDetailedCheckOutputUDP struct {
	_ struct{}

	// The target port to check.
	Port int

	// String to send.
	StringToSend string

	// String to expect in response.
	StringToExpect string
}

// GetDetailedCheckOutputSMTP - Output for the SMTP check type,
// detailed check method.
type GetDetailedCheckOutputSMTP struct {
	_ struct{}

	// The target port to check.
	Port int

	// Username for SMTP authentication.
	Username string

	// Password for SMTP authentication.
	Password string

	// Enable STARTTLS on the SMTP connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}

// GetDetailedCheckOutputPOP3 - Output for the POP3 check type,
// detailed check method.
type GetDetailedCheckOutputPOP3 struct {
	_ struct{}

	// The target port to check.
	Port int

	// Enable encryption on the SMTP connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}

// GetDetailedCheckOutputIMAP - Output for the IMAP check type,
// detailed check method.
type GetDetailedCheckOutputIMAP struct {
	_ struct{}

	// The target port to check.
	Port int

	// Enable encryption on the SMTP connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}
