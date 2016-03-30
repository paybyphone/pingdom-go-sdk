package checks

import (
	"fmt"

	"github.com/paybyphone/pingdom-go-sdk/pingdom"
	"github.com/paybyphone/pingdom-go-sdk/pingdom/client"
)

// Check is the base client for check-related methods.
type Check struct {
	client.Client
}

// New returns a new instance of the Check API.
func New(configs ...pingdom.Config) *Check {
	c := &Check{
		Client: *client.New(configs...),
	}
	return c
}

// CheckListEntryTags - unexported type for a check list entry's tags.
type CheckListEntryTags struct {
	_ struct{}

	// The tag name.
	Name string

	// The tag type - "a" for auto-tagged, "u" for user-tagged.
	Type string

	// The tag count (undocumented in API, unsure of exact meaning).
	Count int
}

// CheckListEntry - holds a single check from GetCheckListOutput.
type CheckListEntry struct {
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
	LastTestTime int

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
	Tags []CheckListEntryTags
}

// GetCheckListInput - Input to send the GetCheckListInput function.
type GetCheckListInput struct {
	_ struct{}

	// Limits the number of returned probes to the specified quantity.
	// Max value is 25000.
	Limit int `url:"limit,omitempty"`

	// Offset for the check listing. Requires Limit.
	Offset int `url:"offset,omitempty"`

	// Include tag list for each check.
	IncludeTags bool `url:"include_tags,omitempty"`

	// A tag list to search on.
	Tags []string `url:"tags,omitempty,comma"`
}

// GetCheckListOutput - Output for the GetCheckList function.
type GetCheckListOutput struct {
	_ struct{}

	// The list of matched checks.
	Checks []CheckListEntry
}

// GetCheckList gets a list of available checks based on a specific set of filters.
func (c *Check) GetCheckList(in GetCheckListInput) (out GetCheckListOutput, err error) {
	err = c.SendRequest("GET", "/api/2.0/checks", &in, &out)
	return
}

// DetailedCheckEntryHTTP - Output for the HTTP check type,
// detailed check method.
type DetailedCheckEntryHTTP struct {
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

// DetailedCheckEntryHTTPCustom - Output for the Custom HTTP check type,
// detailed check method.
type DetailedCheckEntryHTTPCustom struct {
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

// DetailedCheckEntryTCP - Output for the TCP check type,
// detailed check method.
type DetailedCheckEntryTCP struct {
	_ struct{}

	// Path to the target XML file on the server.
	Port int

	// String to send to the server.
	StringToSend string

	// String to expect in response.
	StringToExpect string
}

// DetailedCheckEntryPing - Output for the Ping check type,
// detailed check method.
type DetailedCheckEntryPing struct {
	_ struct{}
}

// DetailedCheckEntryDNS - Output for the DNS check type,
// detailed check method.
type DetailedCheckEntryDNS struct {
	_ struct{}

	// DNS server to use.
	DNSServer string

	// Expected IP address from the query.
	ExpectedIP string
}

// DetailedCheckEntryUDP - Output for the UDP check type,
// detailed check method.
type DetailedCheckEntryUDP struct {
	_ struct{}

	// The target port to check.
	Port int

	// String to send.
	StringToSend string

	// String to expect in response.
	StringToExpect string
}

// DetailedCheckEntrySMTP - Output for the SMTP check type,
// detailed check method.
type DetailedCheckEntrySMTP struct {
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

// DetailedCheckEntryPOP3 - Output for the POP3 check type,
// detailed check method.
type DetailedCheckEntryPOP3 struct {
	_ struct{}

	// The target port to check.
	Port int

	// Enable encryption on the POP3 connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}

// DetailedCheckEntryIMAP - Output for the IMAP check type,
// detailed check method.
type DetailedCheckEntryIMAP struct {
	_ struct{}

	// The target port to check.
	Port int

	// Enable encryption on the IMAP connection.
	Encryption bool

	// String to expect in response.
	StringToExpect string
}

// DetailedCheckEntryTypes is a collection of various structs containing
// type-specific details.
type DetailedCheckEntryTypes struct {
	_ struct{}

	HTTP       DetailedCheckEntryHTTP
	HTTPCustom DetailedCheckEntryHTTPCustom
	TCP        DetailedCheckEntryTCP
	Ping       DetailedCheckEntryPing
	DNS        DetailedCheckEntryDNS
	UDP        DetailedCheckEntryUDP
	SMTP       DetailedCheckEntrySMTP
	POP3       DetailedCheckEntryPOP3
	IMAP       DetailedCheckEntryIMAP
}

// DetailedCheckEntry - Unexported entry for GetDetailedCheckOutput.
type DetailedCheckEntry struct {
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
	Type DetailedCheckEntryTypes

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
	LastTestTime int

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
	Check DetailedCheckEntry
}

// GetDetailedCheck gets detailed information about a single check.
func (c *Check) GetDetailedCheck(in GetDetailedCheckInput) (out GetDetailedCheckOutput, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/api/2.0/checks/%d", in.CheckID), nil, &out)
	return
}

// CheckConfiguration - Structure for the create and modify
// check functions.
type CheckConfiguration struct {
	_ struct{}

	// The name of the check.
	Name string `url:"name,omitempty"`

	// The target hostname or IP address.
	Host string `url:"host,omitempty"`

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
	Type string `url:"type,omitempty"`

	// Pause the check upon creation.
	Paused bool `url:"paused,omitempty"`

	// The resolution of the check. Can be one of
	// 1, 5, 15, 30, or 60.
	Resolution int `url:"resolution,omitempty"`

	// An array of contact IDs.
	ContactIDs []int `url:"contactids,comma,omitempty"`

	// Send alerts as email.
	SendToEmail bool `url:"sendtoemail,omitempty"`

	// Send alerts as SMS.
	SendToSMS bool `url:"sendtosms,omitempty"`

	// Send alerts through Twitter.
	SendToTwitter bool `url:"sendtotwitter,omitempty"`

	// Send alerts to iPhone.
	SendToIphone bool `url:"sendtoiphone,omitempty"`

	// Send alerts to Android.
	SendToAndroid bool `url:"sendtoandroid,omitempty"`

	// The failure count threshold to send notifications on.
	SendNotificationWhenDown int `url:"sendnotificationwhendown,omitempty"`

	// The check frequency to notify on after a service has failed.
	NotifyAgainEvery int `url:"notifyagainevery,omitempty"`

	// Send a notification after a failed check resolves itself.
	NotifyWhenBackUp bool `url:"notifywhenbackup,omitempty"`

	// Tags for the check.
	Tags []string `url:"tags,omitempty,comma"`

	// Use IPv6 instead of IPv4.
	//
	// If an IP address is provided as a host, this setting will be
	// overridden by the version of the IP address provided.
	IPv6 bool `url:"ipv6,omitempty"`
}

// CheckConfigurationHTTP - Configuration for the HTTP check type.
type CheckConfigurationHTTP struct {
	_ struct{}

	// Path to the target on the server.
	URL string `url:"url,omitempty"`

	// true if the connection to the server is encrypted.
	Encryption bool `url:"encryption,omitempty"`

	// Target port to connect to on the server.
	Port int `url:"port,omitempty"`

	// Username and password for target HTTP authentication.
	// Example: user:password
	Auth string `url:"auth,omitempty"`

	// A string the target response should contain.
	ShouldContain string `url:"shouldcontain,omitempty"`

	// A string the target response should not contain.
	// If ShouldContain is also set, this parameter is not allowed.
	ShouldNotContain string `url:"shouldnotcontain,omitempty"`

	// Data that should be posted to the web page, for example submission data
	// for a sign-up or login form. The data needs to be formatted in the same
	// way as a web browser would send it to the web server.
	PostData string `url:"postdata,omitempty"`

	// Custom headers to send with the HTTP request. Required in name: value
	// pairs.
	RequestHeaders []string `url:"requestheader,numbered,omitempty"`
}

// CheckConfigurationHTTPCustom - Configuration for the Custom HTTP check type.
type CheckConfigurationHTTPCustom struct {
	_ struct{}

	// Path to the target on the server.
	URL string `url:"url,omitempty"`

	// true if the connection to the server is encrypted.
	Encryption bool `url:"encryption,omitempty"`

	// Target port to connect to on the server.
	Port int `url:"port,omitempty"`

	// Username and password for target HTTP authentication.
	// Example: user:password
	Auth string `url:"auth,omitempty"`

	// Additional URLs to target.
	AdditionalURLs []string `url:"additionalurls,semicolon,omitempty"`
}

// CheckConfigurationTCP - Configuration for the TCP check type.
type CheckConfigurationTCP struct {
	_ struct{}

	// Path to the target XML file on the server.
	Port int `url:"port,omitempty"`

	// String to send to the server.
	StringToSend string `url:"stringtosend,omitempty"`

	// String to expect in response.
	StringToExpect string `url:"stringtoexpect,omitempty"`
}

// CheckConfigurationPing - Configuration for the Ping check type.
type CheckConfigurationPing struct {
	_ struct{}
}

// CheckConfigurationDNS - Configuration for the DNS check type.
type CheckConfigurationDNS struct {
	_ struct{}

	// DNS server to use.
	NameServer string `url:"nameserver,omitempty"`

	// Expected IP address from the query.
	ExpectedIP string `url:"expectedip,omitempty"`
}

// CheckConfigurationUDP - Configuration for the UDP check type.
type CheckConfigurationUDP struct {
	_ struct{}

	// The target port to check.
	Port int `url:"port,omitempty"`

	// String to send.
	StringToSend string `url:"stringtosend,omitempty"`

	// String to expect in response.
	StringToExpect string `url:"stringtoexpect,omitempty"`
}

// CheckConfigurationSMTP - Configuration for the SMTP check type.
type CheckConfigurationSMTP struct {
	_ struct{}

	// The target port to check.
	Port int `url:"port,omitempty"`

	// Username and password for target SMTP authentication.
	// Example: user:password
	Auth string `url:"auth,omitempty"`

	// Enable STARTTLS on the SMTP connection.
	Encryption bool `url:"encryption,omitempty"`

	// String to expect in response.
	StringToExpect string `url:"stringtoexpect,omitempty"`
}

// CheckConfigurationPOP3 - Configuration for the POP3 check type.
type CheckConfigurationPOP3 struct {
	_ struct{}

	// The target port to check.
	Port int `url:"port,omitempty"`

	// Enable STARTTLS on the SMTP connection.
	Encryption bool `url:"encryption,omitempty"`

	// String to expect in response.
	StringToExpect string `url:"stringtoexpect,omitempty"`
}

// CheckConfigurationIMAP - Configuration for the IMAP check type.
type CheckConfigurationIMAP struct {
	_ struct{}

	// The target port to check.
	Port int `url:"port,omitempty"`

	// Enable STARTTLS on the SMTP connection.
	Encryption bool `url:"encryption,omitempty"`

	// String to expect in response.
	StringToExpect string `url:"stringtoexpect,omitempty"`
}

// CreateCheckInput - Input for the CreateCheck function.
// Embeds checkConfiguration structs.
type CreateCheckInput struct {
	_ struct{}

	CheckConfiguration
	CheckConfigurationHTTP
	CheckConfigurationHTTPCustom
	CheckConfigurationTCP
	CheckConfigurationPing
	CheckConfigurationDNS
	CheckConfigurationUDP
	CheckConfigurationSMTP
	CheckConfigurationPOP3
	CheckConfigurationIMAP
}

// CreateCheckOutput - Output for the CreateCheck function.
type CreateCheckOutput struct {
	_ struct{}

	// The check data.
	Check CreateCheckEntry
}

// CreateCheckEntry is the actual check data in the output of CreateCheck.
type CreateCheckEntry struct {
	_ struct{}

	// The ID of the check that you want to get a description for.
	ID int

	// The name of the check.
	Name string
}

// CreateCheck creates a Pingdom service check.
func (c *Check) CreateCheck(in CreateCheckInput) (out CreateCheckOutput, err error) {
	err = c.SendRequest("POST", "/api/2.0/checks", &in, &out)
	return
}

// ModifyCheckInput - Input for the ModifyCheck function.
// Embeds checkConfiguration structs.
type ModifyCheckInput struct {
	_ struct{}

	// The ID of the check to modify.
	CheckID int `url:"-"`

	CheckConfiguration
	CheckConfigurationHTTP
	CheckConfigurationHTTPCustom
	CheckConfigurationTCP
	CheckConfigurationPing
	CheckConfigurationDNS
	CheckConfigurationUDP
	CheckConfigurationSMTP
	CheckConfigurationPOP3
	CheckConfigurationIMAP
}

// ModifyCheckOutput - Output for the ModifyCheck function.
type ModifyCheckOutput struct {
	_ struct{}

	// The success message.
	Message string
}

// ModifyCheck modifies an existing check.
//
// The provided settings will overwrite previous values. To clear an existing
// value, provide an empty value. Note that you cannot change the type of a
// check once it's created.
func (c *Check) ModifyCheck(in ModifyCheckInput) (out ModifyCheckOutput, err error) {
	err = c.SendRequest("PUT", fmt.Sprintf("/api/2.0/checks/%d", in.CheckID), &in, &out)
	return
}

// DeleteCheckInput - Input to send to the DeleteCheck method.
type DeleteCheckInput struct {
	_ struct{}

	// The ID of the check that you want to delete.
	CheckID int
}

// DeleteCheckOutput - Output for the DeleteCheck method.
type DeleteCheckOutput struct {
	_ struct{}

	// The success message.
	Message string
}

// DeleteCheck deletes a check from Pingdom.
func (c *Check) DeleteCheck(in DeleteCheckInput) (out DeleteCheckOutput, err error) {
	err = c.SendRequest("DELETE", fmt.Sprintf("/api/2.0/checks/%d", in.CheckID), nil, &out)
	return
}
