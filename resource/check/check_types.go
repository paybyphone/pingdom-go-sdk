package check

// checkListEntryTags - unexported type for a check list entry's tags.
type checkListEntryTags struct {
	_ struct{}

	// The tag name.
	Name string

	// The tag type - "a" for auto-tagged, "u" for user-tagged.
	Type string

	// The tag count (undocumented in API, unsure of exact meaning).
	Count int
}

// checkListEntry - holds a single check from GetCheckListOutput.
type checkListEntry struct {
	_ struct{}

	// The check identifier.
	ID int

	// The check name.
	Name string

	// The check type.
	Type string

	// Timestamp of last error (if any). Format is UNIX timestamp.
	LastErrorTime int

	// Timestamp of last test (if any). Format is UNIX timestamp
	LatestTime int

	// Response time (in milliseconds) of last test.
	LastResponseTime int

	// The current check status.
	Status string

	// How often the check should be checked, in minutes.
	Resolution int

	// The target host.
	Hostname string

	// The time the check was created (UNIX timestamp).
	Created int

	// The check uses IPv6 instead of IPv4.
	IPv6 bool

	// Any tags for the check.
	Tags []checkListEntryTags
}

// GetCheckListInput - Input to send the GetCheckListInput function.
type GetCheckListInput struct {
	_ struct{}

	// Limits the number of returned probes to the specified quantity.
	// Max value is 25000.
	Limit int

	// Offset for the check listing. Requires Limit.
	Offset int

	// Include tag list for each check.
	IncludeTags bool `json:"include_tags"`

	// A tag list to search on.
	Tags []string
}

// GetCheckListOutput - Output for the GetCheckList function.
type GetCheckListOutput struct {
	_ struct{}

	// The list of matched checks.
	Checks []checkListEntry
}

// GetDetailedCheckOutputHTTP - Output for the HTTP check type,
// detailed check method.
type detailedCheckEntryHTTP struct {
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
type detailedCheckEntryHTTPCustom struct {
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
type detailedCheckEntryTCP struct {
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
type detailedCheckEntryPing struct {
	_ struct{}
}

// GetDetailedCheckOutputDNS - Output for the DNS check type,
// detailed check method.
type detailedCheckEntryDNS struct {
	_ struct{}

	// DNS server to use.
	DNSServer string

	// Expected IP address from the query.
	ExpectedIP string
}

// GetDetailedCheckOutputUDP - Output for the UDP check type,
// detailed check method.
type detailedCheckEntryUDP struct {
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
type detailedCheckEntrySMTP struct {
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
type detailedCheckEntryPOP3 struct {
	_ struct{}

	// The target port to check.
	Port int

	// Enable encryption on the POP3 connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}

// GetDetailedCheckOutputIMAP - Output for the IMAP check type,
// detailed check method.
type detailedCheckEntryIMAP struct {
	_ struct{}

	// The target port to check.
	Port int

	// Enable encryption on the IMAP connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}

// detailedCheckEntry - Unexported entry for GetDetailedCheckOutput.
type detailedCheckEntry struct {
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
	Type struct {
		_ struct{}

		detailedCheckEntryHTTP
		detailedCheckEntryHTTPCustom
		detailedCheckEntryTCP
		detailedCheckEntryPing
		detailedCheckEntryDNS
		detailedCheckEntryUDP
		detailedCheckEntrySMTP
		detailedCheckEntryPOP3
		detailedCheckEntryIMAP
	}

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

// GetDetailedCheckInput - Input to send to the detailed check method.
type GetDetailedCheckInput struct {
	_ struct{}

	// The ID of the check that you want to get a description for.
	CheckID int
}

// GetDetailedCheckOutput - Output for the detailed check method.
type GetDetailedCheckOutput struct {
	_ struct{}

	// The detailed check entry.
	Check detailedCheckEntry
}

// checkConfiguration - Structure for the create and modify
// check functions.
type checkConfiguration struct {
	_ struct{}

	// The name of the check.
	Name string `url:",omitempty"`

	// The target hostname or IP address.
	Host string `url:",omitempty"`

	// The type of check. One of:
	//  * http (HTTP check)
	//  * httpcustom (Custom HTTP check)
	//  * tcp (TCP check)
	//  * ping (ping check)
	//  * dns (DNS check)
	//  * udp (UDP check)
	//  * smtp (SMTP check)
	//  * pop3 (POP3 check)
	//  * imap (IMAP check)
	Type string `url:",omitempty"`

	// Pause the check upon creation.
	Paused string `url:",omitempty"`

	// The resolution of the check. Can be one of
	// 1, 5, 15, 30, or 60.
	Resolution int `url:",omitempty"`

	// An array of contact IDs.
	ContactIDs []int `url:",omitempty"`

	// Send alerts as email.
	SendToEmail bool `url:",omitempty"`

	// Send alerts as SMS.
	SendToSMS bool `url:",omitempty"`

	// Send alerts through Twitter.
	SendToTwitter bool `url:",omitempty"`

	// Send alerts to iPhone.
	SendToIphone bool `url:",omitempty"`

	// Send alerts to Android.
	SendToAndroid bool `url:",omitempty"`

	// The failure count threshold to send notifications on.
	SendNotificationWhenDown int `url:",omitempty"`

	// The check frequency to notify on after a service has failed.
	NotifyAgainEvery int `url:",omitempty"`

	// Send a notification after a failed check resolves itself.
	NotifyWhenBackUp bool `url:",omitempty"`

	// Tags for the check.
	Tags []string `url:",omitempty,comma"`

	// Use IPv6 instead of IPv4.
	//
	// If an IP address is provided as a host, this setting will be
	// overridden by the version of the IP address provided.
	IPv6 bool `url:",omitempty"`
}

// checkConfigurationHTTP - Configuration for the HTTP check type.
type checkConfigurationHTTP struct {
	_ struct{}

	// Path to the target on the server.
	URL string `url:",omitempty"`

	// true if the connection to the server is encrypted.
	Encryption bool `url:",omitempty"`

	// Target port to connect to on the server.
	Port int `url:",omitempty"`

	// Username and password for target HTTP authentication.
	// Example: user:password
	Auth string `url:",omitempty"`

	// A string the target response should contain.
	ShouldContain string `url:",omitempty"`

	// A string the target response should not contain.
	// If ShouldContain is also set, this parameter is not allowed.
	ShouldNotContain string `url:",omitempty"`

	// Data that should be posted to the web page, for example submission data
	// for a sign-up or login form. The data needs to be formatted in the same
	// way as a web browser would send it to the web server.
	PostData string `url:",omitempty"`

	// Custom headers to send with the HTTP request. Required in name: value
	// pairs.
	RequestHeaders []string `url:"requestheader,numbered,omitempty"`
}

// checkConfigurationHTTPCustom - Configuration for the Custom HTTP check type.
type checkConfigurationHTTPCustom struct {
	_ struct{}

	// Path to the target XML file on the server.
	URL string `url:",omitempty"`

	// true if the connection to the server is encrypted.
	Encryption bool `url:",omitempty"`

	// Target port to connect to on the server.
	Port int `url:",omitempty"`

	// Username and password for target HTTP authentication.
	// Example: user:password
	Auth string `url:",omitempty"`

	// Additional URLs to target.
	AdditionalURLs []string `url:",semicolon,omitempty"`
}

// checkConfigurationTCP - Configuration for the TCP check type.
type checkConfigurationTCP struct {
	_ struct{}

	// Path to the target XML file on the server.
	Port int `url:",omitempty"`

	// String to send to the server.
	StringToSend string `url:",omitempty"`

	// String to expect in response.
	StringToExpect string `url:",omitempty"`
}

// checkConfigurationPing - Configuration for the Ping check type.
type checkConfigurationPing struct {
	_ struct{}
}

// checkConfigurationDNS - Configuration for the DNS check type.
type checkConfigurationDNS struct {
	_ struct{}

	// DNS server to use.
	DNSServer string `url:",omitempty"`

	// Expected IP address from the query.
	ExpectedIP string `url:",omitempty"`
}

// checkConfigurationUDP - Configuration for the UDP check type.
type checkConfigurationUDP struct {
	_ struct{}

	// The target port to check.
	Port int `url:",omitempty"`

	// String to send.
	StringToSend string `url:",omitempty"`

	// String to expect in response.
	StringToExpect string `url:",omitempty"`
}

// checkConfigurationSMTP - Configuration for the SMTP check type.
type checkConfigurationSMTP struct {
	_ struct{}

	// The target port to check.
	Port int `url:",omitempty"`

	// Username and password for target SMTP authentication.
	// Example: user:password
	Auth string `url:",omitempty"`

	// Enable STARTTLS on the SMTP connection.
	Encryption bool `url:",omitempty"`

	// String to expect in response.
	StringToExpect string `url:",omitempty"`
}

// checkConfigurationPOP3 - Configuration for the POP3 check type.
type checkConfigurationPOP3 struct {
	_ struct{}

	// The target port to check.
	Port int `url:",omitempty"`

	// Enable encryption on the POP3 connection.
	Encryption bool `url:",omitempty"`

	// String to expect in response.
	StringToExpect string `url:",omitempty"`
}

// checkConfigurationIMAP - Configuration for the IMAP check type.
type checkConfigurationIMAP struct {
	_ struct{}

	// The target port to check.
	Port int `url:",omitempty"`

	// Enable encryption on the IMAP connection.
	Encryption bool `url:",omitempty"`

	// String to expect in response.
	StringToExpect string `url:",omitempty"`
}

// CreateCheckInput - Input for the CreateCheck function.
// Embeds checkConfiguration structs.
type CreateCheckInput struct {
	_ struct{}

	checkConfiguration
	checkConfigurationHTTP
	checkConfigurationHTTPCustom
	checkConfigurationTCP
	checkConfigurationPing
	checkConfigurationDNS
	checkConfigurationUDP
	checkConfigurationSMTP
	checkConfigurationPOP3
	checkConfigurationIMAP
}

// CreateCheckOutput - Output for the CreateCheck function.
type CreateCheckOutput struct {
	_ struct{}

	// The ID of the check that you want to get a description for.
	CheckID int

	// The name of the check.
	CheckName string
}

// ModifyCheckInput - Input for the CreateCheck function.
// Embeds checkConfiguration structs.
type ModifyCheckInput struct {
	_ struct{}

	checkConfiguration
	checkConfigurationHTTP
	checkConfigurationHTTPCustom
	checkConfigurationTCP
	checkConfigurationPing
	checkConfigurationDNS
	checkConfigurationUDP
	checkConfigurationSMTP
	checkConfigurationPOP3
	checkConfigurationIMAP
}

// ModifyCheckOutput - Output for the ModifyCheck function.
type ModifyCheckOutput struct {
	_ struct{}

	// The success message.
	Message string
}
