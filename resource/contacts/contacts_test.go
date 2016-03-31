package contacts

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/paybyphone/pingdom-go-sdk/integration"
	"github.com/paybyphone/pingdom-go-sdk/pingdom"
)

const errorResponseText = `
{
  "error": {
    "statuscode": 403,
    "statusdesc": "Forbidden",
    "errormessage": "Something went wrong! This string describes what happened."
  }
}
`

const errorResponse = "Forbidden (403): Something went wrong! This string describes what happened."

func newHTTPTestServer(f func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(f))
	return ts
}

func httpErrorTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, errorResponseText, http.StatusForbidden)
	})
}

func setPingdomenv() {
	os.Setenv("PINGDOM_EMAIL_ADDRESS", "nobody@example.com")
	os.Setenv("PINGDOM_PASSWORD", "changeit")
	os.Setenv("PINGDOM_APP_KEY", "abcdefgh0123456789")
}

func unsetPingdomenv() {
	os.Unsetenv("PINGDOM_EMAIL_ADDRESS")
	os.Unsetenv("PINGDOM_PASSWORD")
	os.Unsetenv("PINGDOM_APP_KEY")
}

func pingdomConfig() pingdom.Config {
	return pingdom.Config{
		EmailAddress: "overridden@example.com",
		Password:     "overridden",
		AppKey:       "overridden1234",
	}
}

func getContactListInputData() GetContactListInput {
	return GetContactListInput{
		Limit:  10,
		Offset: 1,
	}
}

const getContactListInputText = "limit=10&offset=1"

func getContactListOutputData() GetContactListOutput {
	return GetContactListOutput{
		Contacts: []ContactListEntry{
			ContactListEntry{
				ID:                 111250,
				Name:               "John Doe",
				Email:              "john@johnsdomain.com",
				CellPhone:          "46-5555555",
				CountryISO:         "SE",
				DefaultSMSProvider: "clickatell",
				DirectTwitter:      true,
				TwitterUser:        "jdoe",
				IPhoneTokens:       []string{"aabbb", "ccddd"},
				AndroidTokens:      []string{"eefff", "gghhh"},
				Paused:             true,
			},
			ContactListEntry{
				ID:                 111251,
				Name:               "Jane Doe",
				Email:              "jane@janesdomain.com",
				CellPhone:          "1-604-664-1234",
				CountryISO:         "CA",
				DefaultSMSProvider: "clickatell",
				DirectTwitter:      true,
				TwitterUser:        "janedoe",
				IPhoneTokens:       []string{"iijjj", "kklll"},
				AndroidTokens:      []string{"mmnnn", "ooppp"},
				Paused:             true,
			},
		},
	}
}

const getContactListOutputText = `
{
	"contacts": [{
		"id": 111250,
		"name": "John Doe",
		"email": "john@johnsdomain.com",
		"cellphone": "46-5555555",
		"countryiso": "SE",
		"defaultsmsprovider": "clickatell",
		"directtwitter": true,
		"twitteruser": "jdoe",
		"iphonetokens": ["aabbb", "ccddd"],
		"androidtokens": ["eefff", "gghhh"],
		"paused": true
	}, {
		"id": 111251,
		"name": "Jane Doe",
		"email": "jane@janesdomain.com",
		"cellphone": "1-604-664-1234",
		"countryiso": "CA",
		"defaultsmsprovider": "clickatell",
		"directtwitter": true,
		"twitteruser": "janedoe",
		"iphonetokens": ["iijjj", "kklll"],
		"androidtokens": ["mmnnn", "ooppp"],
		"paused": true
	}]
}
`

func httpGetContactListTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, getContactListOutputText, http.StatusOK)
	})
}

func ContactConfigurationData() ContactConfiguration {
	return ContactConfiguration{
		Name:               "John Doe",
		Email:              "john@johnsdomain.com",
		CellPhone:          "5555555",
		CountryCode:        "46",
		CountryISO:         "SE",
		DefaultSMSProvider: "clickatell",
		DirectTwitter:      true,
		TwitterUser:        "jdoe",
	}
}

func createContactInputData() CreateContactInput {
	return CreateContactInput{
		ContactConfiguration: ContactConfigurationData(),
	}
}

const ContactConfigurationText = "cellphone=5555555&countrycode=46&countryiso=SE&defaultsmsprovider=clickatell&directtwitter=true&email=john%40johnsdomain.com&name=John+Doe&twitteruser=jdoe"

func createContactOutputData() CreateContactOutput {
	return CreateContactOutput{
		Contact: createContactEntry{
			ID:   111250,
			Name: "John Doe",
		},
	}
}

const createContactOutputText = `
{
	"contact": {
		"id": 111250,
		"name": "John Doe"
	}
}
`

func httpCreateContactTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, createContactOutputText, http.StatusOK)
	})
}

func modifyContactInputData() ModifyContactInput {
	return ModifyContactInput{
		ContactConfiguration: ContactConfigurationData(),
	}
}

func modifyContactOutputData() ModifyContactOutput {
	return ModifyContactOutput{
		Message: "Modification of contact was successful!",
	}
}

const modifyContactOutputText = `
{
	"message": "Modification of contact was successful!"
}
`

func httpModifyContactTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, modifyContactOutputText, http.StatusOK)
	})
}

func deleteContactInputData() DeleteContactInput {
	return DeleteContactInput{
		ContactID: 134536,
	}
}

func deleteContactOutputData() DeleteContactOutput {
	return DeleteContactOutput{
		Message: "Deletion of notification contact was successful!",
	}
}

const deleteContactOutputText = `
{
	"message": "Deletion of notification contact was successful!"
}
`

func httpDeleteContactTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, deleteContactOutputText, http.StatusOK)
	})
}

func TestContactNewWithEnv(t *testing.T) {
	setPingdomenv()
	c := New()
	if c.Config.Endpoint != "https://api.pingdom.com" {
		t.Fatalf("Expected Endpoint to be https://api.paybyphone.com, got %s", c.Config.Endpoint)
	}
	if c.Config.EmailAddress != "nobody@example.com" {
		t.Fatalf("Expected EmailAddress to be nobody@example.com, got %s", c.Config.EmailAddress)
	}
	if c.Config.Password != "changeit" {
		t.Fatalf("Expected Password to be changeit, got %s", c.Config.Password)
	}
	if c.Config.AppKey != "abcdefgh0123456789" {
		t.Fatalf("Expected AppKey to be abcdefgh0123456789, got %s", c.Config.AppKey)
	}
}

func TestContactNewWithOverride(t *testing.T) {
	setPingdomenv()
	c := New(pingdomConfig())
	if c.Config.Endpoint != "https://api.pingdom.com" {
		t.Fatalf("Expected Endpoint to be https://api.paybyphone.com, got %s", c.Config.Endpoint)
	}
	if c.Config.EmailAddress != "overridden@example.com" {
		t.Fatalf("Expected EmailAddress to be overridden@example.com, got %s", c.Config.EmailAddress)
	}
	if c.Config.Password != "overridden" {
		t.Fatalf("Expected Password to be overridden, got %s", c.Config.Password)
	}
	if c.Config.AppKey != "overridden1234" {
		t.Fatalf("Expected AppKey to be overridden1234, got %s", c.Config.AppKey)
	}
}

func TestGetContactListQueryText(t *testing.T) {
	in := getContactListInputData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := getContactListInputText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestGetContactList(t *testing.T) {
	ts := httpGetContactListTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getContactListInputData()
	out, err := c.GetContactList(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := getContactListOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestGetContactListError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getContactListInputData()
	_, err := c.GetContactList(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestContactConfigurationQueryText(t *testing.T) {
	in := createContactInputData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := ContactConfigurationText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCreateContact(t *testing.T) {
	ts := httpCreateContactTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := createContactInputData()
	out, err := c.CreateContact(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := createContactOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestCreateContactError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := createContactInputData()
	_, err := c.CreateContact(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestModifyContact(t *testing.T) {
	ts := httpModifyContactTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := modifyContactInputData()
	out, err := c.ModifyContact(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := modifyContactOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestModifyContactError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := modifyContactInputData()
	_, err := c.ModifyContact(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestDeleteContact(t *testing.T) {
	ts := httpDeleteContactTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := deleteContactInputData()
	out, err := c.DeleteContact(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := deleteContactOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestDeleteContactError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := deleteContactInputData()
	_, err := c.DeleteContact(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

// testAccContactsCRUDCreate runs the Create section of the CRUD test
// (using CreateContact).
func testAccContactsCRUDCreate(t *testing.T, in CreateContactInput) int {
	c := New()
	out, err := c.CreateContact(in)
	if err != nil {
		t.Fatalf("Error creating contact: %v", err)
	}
	if out.Contact.ID == 0 {
		t.Fatalf("Error reading contact ID from output (out.Contact.ID was empty)")
	}
	return out.Contact.ID
}

// testAccContactsCRUDRead runs the Read section of the CRUD test
// (using GetContactList).
func testAccContactsCRUDRead(t *testing.T, id int, name string) {
	c := New()
	out, err := c.GetContactList(GetContactListInput{})
	if err != nil {
		t.Fatalf("Error listing contacts: %v", err)
	}

	var found bool

	for _, v := range out.Contacts {
		if v.ID == id {
			found = true
			if v.Name != name {
				t.Fatalf("Expected Name to be %s, got %v", name, v.Name)
			}
		}
	}
	if found == false {
		t.Fatalf("Could not find created contact in contact list")
	}
}

// testAccContactsCRUDUpdate runs the Update section of the CRUD test
// (using UpdateContact).
//
// Note that this also contacts  by proxy so that we can contact
// that the update took effect.
func testAccContactsCRUDUpdate(t *testing.T, id int, in ModifyContactInput) {
	c := New()
	in.Name = "John Doe (updated)"
	in.ContactID = id
	_, err := c.ModifyContact(in)
	if err != nil {
		t.Fatalf("Error updating contact: %v", err)
	}

	testAccContactsCRUDRead(t, id, "John Doe (updated)")
}

// testAccContactsCRUDDelete runs the Delete section of the CRUD test
// (using DeleteContact).
func testAccContactsCRUDDelete(t *testing.T, id int) {
	c := New()
	params := DeleteContactInput{
		ContactID: id,
	}
	out, err := c.DeleteContact(params)
	if err != nil {
		t.Fatalf("Error deleting contact: %v", err)
	}
	if strings.HasPrefix(out.Message, "Deletion of notification contact was successful!") == false {
		t.Fatalf("Expected out.Message to start with Deletion of notification contact was successful!, got %v", out.Message)
	}
}

// TestAccContactsCRUD runs a full create-read-update-delete test for a Pingdom
// contact.
func TestAccContactsCRUD(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccContactsCRUDCreate(t, createContactInputData())
	testAccContactsCRUDRead(t, id, "John Doe")
	testAccContactsCRUDUpdate(t, id, modifyContactInputData())
	testAccContactsCRUDDelete(t, id)
}
