package contact

import (
	"fmt"

	"github.com/paybyphone/pingdom-go-sdk/pingdom"
	"github.com/paybyphone/pingdom-go-sdk/pingdom/client"
)

// Contact is the base client for contact-related methods.
type Contact struct {
	client.Client
}

// New returns a new instance of the Contact API.
func New(configs ...pingdom.Config) *Contact {
	c := &Contact{
		Client: *client.New(configs...),
	}
	return c
}

// contactListEntry holds a single contact from GetContactListOutput.
type contactListEntry struct {
	_ struct{}

	// The contact identifier.
	ID int

	// The contact name.
	Name string

	// The contact email address.
	Email string

	// The contact cell phone number.
	CellPhone string

	//The cell phone country ISO code.
	CountryISO string

	// The default SMS provider.
	DefaultSMSProvider string

	// Send the alert to the Twitter account on this contact as a direct
	// message.
	DirectTwitter bool

	// The Twitter account to direct message (if DirectTwitter is
	// enabled).
	TwitterUser string

	// The iPhone tokens associated with this contact.
	IPhoneTokens []string

	// The Android tokens associated with this contact.
	AndroidTokens []string

	// Pause this contact.
	Paused bool
}

// GetContactListInput contains the input to send the GetContactListInput
// function.
type GetContactListInput struct {
	_ struct{}

	// Limits the number of returned probes to the specified quantity.
	// Max value is 25000.
	Limit int `url:"limit,omitempty"`

	// Offset for the contact listing. Requires Limit.
	Offset int `url:"offset,omitempty"`
}

// GetContactListOutput contains the output for the GetContactList function.
type GetContactListOutput struct {
	_ struct{}

	// The list of matched contacts.
	Contacts []contactListEntry
}

// GetContactList gets a list of available contacts based on a specific set of
// filters.
func (c *Contact) GetContactList(in GetContactListInput) (out GetContactListOutput, err error) {
	err = c.SendRequest("GET", "/api/2.0/contacts", &in, &out)
	return
}

// contactConfiguration is the structure for the create and modify
// contact functions.
type contactConfiguration struct {
	_ struct{}

	// The contact name.
	Name string `url:"name,omitempty"`

	// The contact email address.
	Email string `url:"email,omitempty"`

	// The contact cell phone number, without the country code. In some
	// countries, you will need to exclude leading zeroes. Requires
	// CountryCode and CountryISO.
	CellPhone string `url:"cellphone,omitempty"`

	// The contact cell phone's country code. Requires CountryCode and
	// CountryISO.
	CountryCode string `url:"countrycode,omitempty"`

	//The cell phone country ISO code. For example: US (USA), GB (Britan),
	// or SE (Sweden). Requires CountryCode and CountryISO.
	CountryISO string `url:"countryiso,omitempty"`

	// The default SMS provider. One of: clickatell, bulksms, esendex,
	// or cellsynt.
	DefaultSMSProvider string `url:"defaultsmsprovider,omitempty"`

	// Send the alert to the Twitter account on this contact as a direct
	// message.
	DirectTwitter bool `url:"directtwitter,omitempty"`

	// The Twitter account to direct message (if DirectTwitter is
	// enabled).
	TwitterUser string `url:"twitteruser,omitempty"`
}

// CreateContactInput contains the input for the CreateContact function.
type CreateContactInput struct {
	_ struct{}

	contactConfiguration
}

// CreateContactOutput contains the output for the CreateContact function.
type CreateContactOutput struct {
	_ struct{}

	// The contact data.
	Contact createContactEntry
}

// createContactEntry is the actual contact data in the output of CreateContact.
type createContactEntry struct {
	_ struct{}

	// The ID of the contact that was created.
	ID int

	// The name of the contact that was created.
	Name string
}

// CreateContact creates a contact for use with other Pingdom resources, such as checks.
func (c *Contact) CreateContact(in CreateContactInput) (out CreateContactOutput, err error) {
	err = c.SendRequest("POST", "/api/2.0/contacts", &in, &out)
	return
}

// ModifyContactInput contains the input for the ModifyContact function.
type ModifyContactInput struct {
	_ struct{}

	// The ID of the contact to modify.
	ContactID int `url:"-"`

	// The replacement contact configuration.
	contactConfiguration
}

// ModifyContactOutput contains the output for the ModifyContact function.
type ModifyContactOutput struct {
	_ struct{}

	// The success message.
	Message string
}

// ModifyContact modifies an existing contact.
func (c *Contact) ModifyContact(in ModifyContactInput) (out ModifyContactOutput, err error) {
	err = c.SendRequest("PUT", fmt.Sprintf("/api/2.0/contacts/%d", in.ContactID), &in, &out)
	return
}

// DeleteContactInput contains the input for the DeleteContact method.
type DeleteContactInput struct {
	_ struct{}

	// The ID of the contact that you want to delete.
	ContactID int
}

// DeleteContactOutput contains the output for the DeleteContact method.
type DeleteContactOutput struct {
	_ struct{}

	// The success message.
	Message string
}

// DeleteContact deletes an existing contact from Pingdom.
func (c *Contact) DeleteContact(in DeleteContactInput) (out DeleteContactOutput, err error) {
	err = c.SendRequest("DELETE", fmt.Sprintf("/api/2.0/contacts/%d", in.ContactID), nil, &out)
	return
}
