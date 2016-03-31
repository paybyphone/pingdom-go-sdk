package checks

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

func getCheckListInputData() GetCheckListInput {
	return GetCheckListInput{
		Limit:       10,
		Offset:      0,
		IncludeTags: true,
		Tags:        []string{"apache", "nginx"},
	}
}

const getCheckListInputText = "include_tags=true&limit=10&tags=apache%2Cnginx"

func getCheckListOutputData() GetCheckListOutput {
	return GetCheckListOutput{
		Checks: []CheckListEntry{
			CheckListEntry{
				ID:               85975,
				Name:             "My check 1",
				Type:             "http",
				LastErrorTime:    1297446423,
				LastTestTime:     1300977363,
				LastResponseTime: 355,
				Status:           "up",
				Resolution:       1,
				Hostname:         "example.com",
				Created:          0,
				IPv6:             false,
				Tags: []CheckListEntryTags{
					CheckListEntryTags{
						Name:  "apache",
						Type:  "a",
						Count: 2,
					}},
			},
			CheckListEntry{
				ID:               161748,
				Name:             "My check 2",
				Type:             "ping",
				LastErrorTime:    1299194968,
				LastTestTime:     1300977268,
				LastResponseTime: 1141,
				Status:           "up",
				Resolution:       5,
				Hostname:         "mydomain.com",
				Created:          0,
				IPv6:             false,
				Tags: []CheckListEntryTags{
					CheckListEntryTags{
						Name:  "nginx",
						Type:  "u",
						Count: 1,
					},
				},
			},
			CheckListEntry{
				ID:               208655,
				Name:             "My check 3",
				Type:             "http",
				LastErrorTime:    1300527997,
				LastTestTime:     1300977337,
				LastResponseTime: 800,
				Status:           "down",
				Resolution:       1,
				Hostname:         "example.net",
				Created:          0,
				IPv6:             false,
				Tags: []CheckListEntryTags{
					CheckListEntryTags{
						Name:  "apache",
						Type:  "a",
						Count: 2,
					},
				},
			},
		},
	}
}

const getCheckListOutputText = `
{
	"checks": [{
		"hostname": "example.com",
		"id": 85975,
		"lasterrortime": 1297446423,
		"lastresponsetime": 355,
		"lasttesttime": 1300977363,
		"name": "My check 1",
		"resolution": 1,
		"status": "up",
		"type": "http",
		"tags": [{
			"name": "apache",
			"type": "a",
			"count": 2
		}]
	}, {
		"hostname": "mydomain.com",
		"id": 161748,
		"lasterrortime": 1299194968,
		"lastresponsetime": 1141,
		"lasttesttime": 1300977268,
		"name": "My check 2",
		"resolution": 5,
		"status": "up",
		"type": "ping",
		"tags": [{
			"name": "nginx",
			"type": "u",
			"count": 1
		}]
	}, {
		"hostname": "example.net",
		"id": 208655,
		"lasterrortime": 1300527997,
		"lastresponsetime": 800,
		"lasttesttime": 1300977337,
		"name": "My check 3",
		"resolution": 1,
		"status": "down",
		"type": "http",
		"tags": [{
			"name": "apache",
			"type": "a",
			"count": 2
		}]
	}]
}

`

func httpGetCheckListTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, getCheckListOutputText, http.StatusOK)
	})
}

func getDetailedCheckInputData() GetDetailedCheckInput {
	return GetDetailedCheckInput{
		CheckID: 85975,
	}
}

func getDetailedCheckOutputData() GetDetailedCheckOutput {
	return GetDetailedCheckOutput{
		Check: DetailedCheckEntry{
			ID:         85975,
			Name:       "My check 7",
			Hostname:   "s7.mydomain.com",
			Status:     "up",
			Resolution: 1,
			Type: DetailedCheckEntryTypes{
				HTTP: DetailedCheckEntryHTTP{
					URL:  "/",
					Port: 80,
					RequestHeaders: map[string]string{
						"User-Agent": "Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)",
					},
				},
			},
			ContactIds:               []int{1234, 5678},
			SendToEmail:              false,
			SendToSMS:                false,
			SendToTwitter:            false,
			SendToIphone:             false,
			SendToAndroid:            false,
			SendNotificationWhenDown: 0,
			NotifyAgainEvery:         0,
			NotifyWhenBackUp:         false,
			LastErrorTime:            1293143467,
			LastTestTime:             1294064823,
			LastResponseTime:         0,
			Created:                  1240394682,
			IPv6:                     false,
		},
	}
}

const getDetailedCheckOutputText = `
{
	"check": {
		"id": 85975,
		"name": "My check 7",
		"resolution": 1,
		"sendtoemail": false,
		"sendtosms": false,
		"sendtotwitter": false,
		"sendtoiphone": false,
		"sendnotificationwhendown": 0,
		"notifyagainevery": 0,
		"notifywhenbackup": false,
		"created": 1240394682,
		"type": {
			"http": {
				"url": "/",
				"port": 80,
				"requestheaders": {
					"User-Agent": "Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)"
				}
			}
		},
		"hostname": "s7.mydomain.com",
		"status": "up",
		"lasterrortime": 1293143467,
		"lasttesttime": 1294064823,
		"contactids": [1234, 5678]
	}
}
`

func httpGetDetailedCheckTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, getDetailedCheckOutputText, http.StatusOK)
	})
}

func CheckConfigurationData() CheckConfiguration {
	return CheckConfiguration{
		Name:                     "My check",
		Host:                     "example.com",
		Paused:                   true,
		Resolution:               1,
		ContactIDs:               []int{1234, 5678},
		SendToEmail:              true,
		SendToSMS:                true,
		SendToTwitter:            true,
		SendToIphone:             true,
		SendToAndroid:            true,
		SendNotificationWhenDown: 2,
		NotifyAgainEvery:         1,
		NotifyWhenBackUp:         true,
		Tags:                     []string{"foo", "bar"},
		IPv6:                     false,
	}
}

func CheckConfigurationHTTPData() CheckConfigurationHTTP {
	return CheckConfigurationHTTP{
		URL:              "/test",
		Encryption:       true,
		Port:             443,
		Auth:             "foo:bar",
		ShouldContain:    "foo",
		ShouldNotContain: "bar",
		PostData:         "baz",
		RequestHeaders:   []string{"X-Header1:foo", "X-Header2:bar", "X-Header3:baz"},
	}
}

func createCheckInputHTTPData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationHTTP: CheckConfigurationHTTPData(),
	}
	c.Type = "http"
	return c
}

const CheckConfigurationHTTPText = "auth=foo%3Abar&contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=443&postdata=baz&requestheader0=X-Header1%3Afoo&requestheader1=X-Header2%3Abar&requestheader2=X-Header3%3Abaz&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&shouldcontain=foo&shouldnotcontain=bar&tags=foo%2Cbar&type=http&url=%2Ftest"

func CheckConfigurationHTTPCustomData() CheckConfigurationHTTPCustom {
	return CheckConfigurationHTTPCustom{
		URL:            "/test",
		Encryption:     true,
		Port:           443,
		Auth:           "foo:bar",
		AdditionalURLs: []string{"www.mysite.com", "www.myothersite.com"},
	}
}

func createCheckInputHTTPCustomData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:           CheckConfigurationData(),
		CheckConfigurationHTTPCustom: CheckConfigurationHTTPCustomData(),
	}
	c.Type = "httpcustom"
	return c
}

const CheckConfigurationHTTPCustomText = "additionalurls=www.mysite.com%3Bwww.myothersite.com&auth=foo%3Abar&contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=443&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&tags=foo%2Cbar&type=httpcustom&url=%2Ftest"

func CheckConfigurationTCPData() CheckConfigurationTCP {
	return CheckConfigurationTCP{
		Port:           22,
		StringToSend:   "foo",
		StringToExpect: "bar",
	}
}

func createCheckInputTCPData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:    CheckConfigurationData(),
		CheckConfigurationTCP: CheckConfigurationTCPData(),
	}
	c.Type = "tcp"
	return c
}

const CheckConfigurationTCPText = "contactids=1234%2C5678&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=22&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=bar&stringtosend=foo&tags=foo%2Cbar&type=tcp"

func CheckConfigurationPingData() CheckConfigurationPing {
	return CheckConfigurationPing{}
}

func createCheckInputPingData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationPing: CheckConfigurationPingData(),
	}
	c.Type = "ping"
	return c
}

const CheckConfigurationPingText = "contactids=1234%2C5678&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&tags=foo%2Cbar&type=ping"

func CheckConfigurationDNSData() CheckConfigurationDNS {
	return CheckConfigurationDNS{
		NameServer: "ns1.example.com",
		ExpectedIP: "127.0.0.1",
	}
}

func createCheckInputDNSData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:    CheckConfigurationData(),
		CheckConfigurationDNS: CheckConfigurationDNSData(),
	}
	c.Type = "dns"
	return c
}

const CheckConfigurationDNSText = "contactids=1234%2C5678&expectedip=127.0.0.1&host=example.com&name=My+check&nameserver=ns1.example.com&notifyagainevery=1&notifywhenbackup=true&paused=true&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&tags=foo%2Cbar&type=dns"

func CheckConfigurationUDPData() CheckConfigurationUDP {
	return CheckConfigurationUDP{
		Port:           53,
		StringToSend:   "foo",
		StringToExpect: "bar",
	}
}

func createCheckInputUDPData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:    CheckConfigurationData(),
		CheckConfigurationUDP: CheckConfigurationUDPData(),
	}
	c.Type = "udp"
	return c
}

const CheckConfigurationUDPText = "contactids=1234%2C5678&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=53&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=bar&stringtosend=foo&tags=foo%2Cbar&type=udp"

func CheckConfigurationSMTPData() CheckConfigurationSMTP {
	return CheckConfigurationSMTP{
		Port:           587,
		Auth:           "foo:bar",
		Encryption:     true,
		StringToExpect: "foobar",
	}
}

func createCheckInputSMTPData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationSMTP: CheckConfigurationSMTPData(),
	}
	c.Type = "smtp"
	return c
}

const CheckConfigurationSMTPText = "auth=foo%3Abar&contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=587&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=foobar&tags=foo%2Cbar&type=smtp"

func CheckConfigurationPOP3Data() CheckConfigurationPOP3 {
	return CheckConfigurationPOP3{
		Port:           993,
		Encryption:     true,
		StringToExpect: "foobar",
	}
}

func createCheckInputPOP3Data() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationPOP3: CheckConfigurationPOP3Data(),
	}
	c.Type = "pop3"
	return c
}

const CheckConfigurationPOP3Text = "contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=993&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=foobar&tags=foo%2Cbar&type=pop3"

func CheckConfigurationIMAPData() CheckConfigurationIMAP {
	return CheckConfigurationIMAP{
		Port:           995,
		Encryption:     true,
		StringToExpect: "foobar",
	}
}

func createCheckInputIMAPData() CreateCheckInput {
	c := CreateCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationIMAP: CheckConfigurationIMAPData(),
	}
	c.Type = "imap"
	return c
}

const CheckConfigurationIMAPText = "contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=995&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=foobar&tags=foo%2Cbar&type=imap"

func createCheckOutputData() CreateCheckOutput {
	return CreateCheckOutput{
		Check: CreateCheckEntry{
			ID:   138631,
			Name: "My new HTTP check",
		},
	}
}

const createCheckOutputText = `
{
	"check": {
		"id": 138631,
		"name": "My new HTTP check"
	}
}
`

func httpCreateCheckTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, createCheckOutputText, http.StatusOK)
	})
}

func modifyCheckInputHTTPData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationHTTP: CheckConfigurationHTTPData(),
	}
	return c
}

func modifyCheckInputHTTPCustomData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:           CheckConfigurationData(),
		CheckConfigurationHTTPCustom: CheckConfigurationHTTPCustomData(),
	}
	return c
}

func modifyCheckInputTCPData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:    CheckConfigurationData(),
		CheckConfigurationTCP: CheckConfigurationTCPData(),
	}
	return c
}

func modifyCheckInputPingData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationPing: CheckConfigurationPingData(),
	}
	return c
}

func modifyCheckInputDNSData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:    CheckConfigurationData(),
		CheckConfigurationDNS: CheckConfigurationDNSData(),
	}
	return c
}

func modifyCheckInputUDPData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:    CheckConfigurationData(),
		CheckConfigurationUDP: CheckConfigurationUDPData(),
	}
	return c
}

func modifyCheckInputSMTPData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationSMTP: CheckConfigurationSMTPData(),
	}
	return c
}

func modifyCheckInputPOP3Data() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationPOP3: CheckConfigurationPOP3Data(),
	}
	return c
}

func modifyCheckInputIMAPData() ModifyCheckInput {
	c := ModifyCheckInput{
		CheckConfiguration:     CheckConfigurationData(),
		CheckConfigurationIMAP: CheckConfigurationIMAPData(),
	}
	return c
}

func modifyCheckOutputData() ModifyCheckOutput {
	return ModifyCheckOutput{
		Message: "Modification of check was successful!",
	}
}

const modifyCheckOutputText = `
{
	"message": "Modification of check was successful!"
}
`

func httpModifyCheckTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, modifyCheckOutputText, http.StatusOK)
	})
}

func deleteCheckInputData() DeleteCheckInput {
	return DeleteCheckInput{
		CheckID: 134536,
	}
}

func deleteCheckOutputData() DeleteCheckOutput {
	return DeleteCheckOutput{
		Message: "Deletion of check was successful!",
	}
}

const deleteCheckOutputText = `
{
	"message": "Deletion of check was successful!"
}
`

func httpDeleteCheckTestServer() *httptest.Server {
	return newHTTPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, deleteCheckOutputText, http.StatusOK)
	})
}

func TestCheckNewWithEnv(t *testing.T) {
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

func TestCheckNewWithOverride(t *testing.T) {
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

func TestGetCheckListQueryText(t *testing.T) {
	in := getCheckListInputData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := getCheckListInputText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestGetCheckList(t *testing.T) {
	ts := httpGetCheckListTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getCheckListInputData()
	out, err := c.GetCheckList(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := getCheckListOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestGetCheckListError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getCheckListInputData()
	_, err := c.GetCheckList(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestGetDetailedCheck(t *testing.T) {
	ts := httpGetDetailedCheckTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getDetailedCheckInputData()
	out, err := c.GetDetailedCheck(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := getDetailedCheckOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestGetDetailedCheckTypes(t *testing.T) {
	ts := httpGetDetailedCheckTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getDetailedCheckInputData()
	out, err := c.GetDetailedCheck(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := "Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)"

	// Deep check some of the HTTP check details
	if out.Check.Type.HTTP.RequestHeaders["User-Agent"] != expected {
		t.Fatalf("expected %v, got %v", expected, out.Check.Type.HTTP.RequestHeaders["User-Agent"])
	}
}

func TestGetDetailedCheckError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := getDetailedCheckInputData()
	_, err := c.GetDetailedCheck(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestCreateCheck(t *testing.T) {
	ts := httpCreateCheckTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := createCheckInputHTTPData()
	out, err := c.CreateCheck(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := createCheckOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestCreateCheckError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := createCheckInputHTTPData()
	_, err := c.CreateCheck(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestCheckConfigurationHTTPQueryText(t *testing.T) {
	in := createCheckInputHTTPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationHTTPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationHTTPCustomQueryText(t *testing.T) {
	in := createCheckInputHTTPCustomData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationHTTPCustomText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationTCPQueryText(t *testing.T) {
	in := createCheckInputTCPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationTCPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationPingQueryText(t *testing.T) {
	in := createCheckInputPingData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationPingText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationDNSQueryText(t *testing.T) {
	in := createCheckInputDNSData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationDNSText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationUDPQueryText(t *testing.T) {
	in := createCheckInputUDPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationUDPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationSMTPQueryText(t *testing.T) {
	in := createCheckInputSMTPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationSMTPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationPOP3QueryText(t *testing.T) {
	in := createCheckInputPOP3Data()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationPOP3Text

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationIMAPQueryText(t *testing.T) {
	in := createCheckInputIMAPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := CheckConfigurationIMAPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestModifyCheck(t *testing.T) {
	ts := httpModifyCheckTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := modifyCheckInputHTTPData()
	out, err := c.ModifyCheck(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := modifyCheckOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestModifyCheckError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := modifyCheckInputHTTPData()
	_, err := c.ModifyCheck(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

func TestDeleteCheck(t *testing.T) {
	ts := httpDeleteCheckTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := deleteCheckInputData()
	out, err := c.DeleteCheck(in)

	if err != nil {
		t.Fatalf("Unexpected request error: %s", err)
	}

	expected := deleteCheckOutputData()

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestDeleteCheckError(t *testing.T) {
	ts := httpErrorTestServer()
	defer ts.Close()
	cfg := pingdomConfig()
	cfg.Endpoint = ts.URL
	c := New(cfg)
	in := deleteCheckInputData()
	_, err := c.DeleteCheck(in)

	if err == nil {
		t.Fatalf("Expected error, none found")
	}

	expected := errorResponse

	if err.Error() != expected {
		t.Fatalf("expected %s, got %s", expected, err)
	}
}

// testAccChecksCRUDCreate runs the Create section of the CRUD test
// (using CreateCheck).
func testAccChecksCRUDCreate(t *testing.T, in CreateCheckInput) int {
	c := New()
	in.ContactIDs = []int{}
	out, err := c.CreateCheck(in)
	if err != nil {
		t.Fatalf("Error creating check: %v", err)
	}
	if out.Check.ID == 0 {
		t.Fatalf("Error reading check ID from output (out.Check.ID was empty)")
	}
	return out.Check.ID
}

// testAccChecksCRUDReadList runs the Read section of the CRUD test
// (using GetCheckList).
//
// This is the first part of the two-part Read test (testing both
// GetCheckList and GetDetailedCheck).
func testAccChecksCRUDReadList(t *testing.T, id int) {
	c := New()
	out, err := c.GetCheckList(GetCheckListInput{})
	if err != nil {
		t.Fatalf("Error listing checks: %v", err)
	}

	var found bool

	for _, v := range out.Checks {
		if v.ID == id {
			found = true
		}
	}
	if found == false {
		t.Fatalf("Could not find created check in check list")
	}
}

// testAccChecksCRUDReadDetail runs the Read section of the CRUD test
// (using GetDetailedCheck).
//
// This is the second part of the two-part Read test (testing both
// GetCheckList and GetDetailedCheck).
func testAccChecksCRUDReadDetail(t *testing.T, id int) {
	c := New()
	params := GetDetailedCheckInput{
		CheckID: id,
	}

	out, err := c.GetDetailedCheck(params)
	if err != nil {
		t.Fatalf("Error getting check detail: %v", err)
	}

	if out.Check.ID != id {
		t.Fatalf("Expected out.Check.ID to be %d, got %v", id, out.Check.ID)
	}

	if out.Check.Hostname != "example.com" {
		t.Fatalf("Expected out.Check.Name to be example.com, got %v", out.Check.Name)
	}
}

// testAccChecksCRUDUpdate runs the Update section of the CRUD test
// (using UpdateCheck).
//
// Note that this also checks GetDetailedCheck by proxy so that we can check
// that the update took effect.
func testAccChecksCRUDUpdate(t *testing.T, id int, in ModifyCheckInput) {
	c := New()
	in.ContactIDs = []int{}
	in.Name = "My check (updated)"
	in.CheckID = id
	_, err := c.ModifyCheck(in)
	if err != nil {
		t.Fatalf("Error updating check: %v", err)
	}

	params := GetDetailedCheckInput{
		CheckID: id,
	}

	out, err := c.GetDetailedCheck(params)
	if err != nil {
		t.Fatalf("Error getting check detail after update: %v", err)
	}

	if out.Check.Name != "My check (updated)" {
		t.Fatalf("Expected out.Check.Name to be My check (updated), got %v", out.Check.Name)
	}
}

// testAccChecksCRUDDelete runs the Delete section of the CRUD test
// (using DeleteCheck).
func testAccChecksCRUDDelete(t *testing.T, id int) {
	c := New()
	params := DeleteCheckInput{
		CheckID: id,
	}
	out, err := c.DeleteCheck(params)
	if err != nil {
		t.Fatalf("Error deleting check: %v", err)
	}
	if strings.HasPrefix(out.Message, "Deletion of check was successful!") == false {
		t.Fatalf("Expected out.Message to start with Deletion of check was successful!, got %v", out.Message)
	}
}

// TestAccChecksCRUDHTTP runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDHTTP(t *testing.T) {
	testacc.VetAccConditions(t)

	create := createCheckInputHTTPData()
	update := modifyCheckInputHTTPData()
	// We can't have shouldnotcontain with shouldcontain, so just empty it
	create.ShouldContain = ""
	update.ShouldContain = ""

	id := testAccChecksCRUDCreate(t, create)
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, update)
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDHTTPCustom runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDHTTPCustom(t *testing.T) {
	testacc.VetAccConditions(t)

	create := createCheckInputHTTPCustomData()
	update := modifyCheckInputHTTPCustomData()
	// Remove AdditionalURLs from the parameter list. Pingdom documents it as this:
	// additionalurls=www.mysite.com;www.myothersite.com
	// But that parameter value produces an error. Can't find any docs giving a working
	// example.
	create.AdditionalURLs = []string{}
	update.AdditionalURLs = []string{}

	id := testAccChecksCRUDCreate(t, create)
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, update)
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDTCP runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDTCP(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputTCPData())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputTCPData())
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDPing runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDPing(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputPingData())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputPingData())
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDDNS runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDDNS(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputDNSData())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputDNSData())
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDUDP runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDUDP(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputUDPData())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputUDPData())
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDSMTP runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDSMTP(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputSMTPData())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputSMTPData())
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDPOP3 runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDPOP3(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputPOP3Data())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputPOP3Data())
	testAccChecksCRUDDelete(t, id)
}

// TestAccChecksCRUDIMAP runs a full create-read-update-delete test for a Pingdom
// check.
func TestAccChecksCRUDIMAP(t *testing.T) {
	testacc.VetAccConditions(t)

	id := testAccChecksCRUDCreate(t, createCheckInputIMAPData())
	testAccChecksCRUDReadList(t, id)
	testAccChecksCRUDReadDetail(t, id)
	testAccChecksCRUDUpdate(t, id, modifyCheckInputIMAPData())
	testAccChecksCRUDDelete(t, id)
}
