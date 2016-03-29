package check

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-querystring/query"
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
		Checks: []checkListEntry{
			checkListEntry{
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
				Tags: []checkListEntryTags{
					checkListEntryTags{
						Name:  "apache",
						Type:  "a",
						Count: 2,
					}},
			},
			checkListEntry{
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
				Tags: []checkListEntryTags{
					checkListEntryTags{
						Name:  "nginx",
						Type:  "u",
						Count: 1,
					},
				},
			},
			checkListEntry{
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
				Tags: []checkListEntryTags{
					checkListEntryTags{
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
		Check: detailedCheckEntry{
			ID:         85975,
			Name:       "My check 7",
			Hostname:   "s7.mydomain.com",
			Status:     "up",
			Resolution: 1,
			Type: detailedCheckEntryTypes{
				HTTP: detailedCheckEntryHTTP{
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

func checkConfigurationData() checkConfiguration {
	return checkConfiguration{
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

func checkConfigurationHTTPData() checkConfigurationHTTP {
	return checkConfigurationHTTP{
		URL:              "example.com",
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
		checkConfiguration:     checkConfigurationData(),
		checkConfigurationHTTP: checkConfigurationHTTPData(),
	}
	c.Type = "http"
	return c
}

const checkConfigurationHTTPText = "auth=foo%3Abar&contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=443&postdata=baz&requestheader0=X-Header1%3Afoo&requestheader1=X-Header2%3Abar&requestheader2=X-Header3%3Abaz&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&shouldcontain=foo&shouldnotcontain=bar&tags=foo%2Cbar&type=http&url=example.com"

func checkConfigurationHTTPCustomData() checkConfigurationHTTPCustom {
	return checkConfigurationHTTPCustom{
		URL:            "example.com",
		Encryption:     true,
		Port:           443,
		Auth:           "foo:bar",
		AdditionalURLs: []string{"mysite.com", "myothersite.com"},
	}
}

func createCheckInputHTTPCustomData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:           checkConfigurationData(),
		checkConfigurationHTTPCustom: checkConfigurationHTTPCustomData(),
	}
	c.Type = "httpcustom"
	return c
}

const checkConfigurationHTTPCustomText = "additionalurls=mysite.com%3Bmyothersite.com&auth=foo%3Abar&contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=443&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&tags=foo%2Cbar&type=httpcustom&url=example.com"

func checkConfigurationTCPData() checkConfigurationTCP {
	return checkConfigurationTCP{
		Port:           22,
		StringToSend:   "foo",
		StringToExpect: "bar",
	}
}

func createCheckInputTCPData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:    checkConfigurationData(),
		checkConfigurationTCP: checkConfigurationTCPData(),
	}
	c.Type = "tcp"
	return c
}

const checkConfigurationTCPText = "contactids=1234%2C5678&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=22&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=bar&stringtosend=foo&tags=foo%2Cbar&type=tcp"

func checkConfigurationPingData() checkConfigurationPing {
	return checkConfigurationPing{}
}

func createCheckInputPingData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:     checkConfigurationData(),
		checkConfigurationPing: checkConfigurationPingData(),
	}
	c.Type = "ping"
	return c
}

const checkConfigurationPingText = "contactids=1234%2C5678&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&tags=foo%2Cbar&type=ping"

func checkConfigurationDNSData() checkConfigurationDNS {
	return checkConfigurationDNS{
		NameServer: "ns1.example.com",
		ExpectedIP: "127.0.0.1",
	}
}

func createCheckInputDNSData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:    checkConfigurationData(),
		checkConfigurationDNS: checkConfigurationDNSData(),
	}
	c.Type = "dns"
	return c
}

const checkConfigurationDNSText = "contactids=1234%2C5678&expectedip=127.0.0.1&host=example.com&name=My+check&nameserver=ns1.example.com&notifyagainevery=1&notifywhenbackup=true&paused=true&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&tags=foo%2Cbar&type=dns"

func checkConfigurationUDPData() checkConfigurationUDP {
	return checkConfigurationUDP{
		Port:           53,
		StringToSend:   "foo",
		StringToExpect: "bar",
	}
}

func createCheckInputUDPData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:    checkConfigurationData(),
		checkConfigurationUDP: checkConfigurationUDPData(),
	}
	c.Type = "udp"
	return c
}

const checkConfigurationUDPText = "contactids=1234%2C5678&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=53&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=bar&stringtosend=foo&tags=foo%2Cbar&type=udp"

func checkConfigurationSMTPData() checkConfigurationSMTP {
	return checkConfigurationSMTP{
		Port:           587,
		Auth:           "foo:bar",
		Encryption:     true,
		StringToExpect: "foobar",
	}
}

func createCheckInputSMTPData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:     checkConfigurationData(),
		checkConfigurationSMTP: checkConfigurationSMTPData(),
	}
	c.Type = "smtp"
	return c
}

const checkConfigurationSMTPText = "auth=foo%3Abar&contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=587&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=foobar&tags=foo%2Cbar&type=smtp"

func checkConfigurationPOP3Data() checkConfigurationPOP3 {
	return checkConfigurationPOP3{
		Port:           993,
		Encryption:     true,
		StringToExpect: "foobar",
	}
}

func createCheckInputPOP3Data() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:     checkConfigurationData(),
		checkConfigurationPOP3: checkConfigurationPOP3Data(),
	}
	c.Type = "pop3"
	return c
}

const checkConfigurationPOP3Text = "contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=993&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=foobar&tags=foo%2Cbar&type=pop3"

func checkConfigurationIMAPData() checkConfigurationIMAP {
	return checkConfigurationIMAP{
		Port:           995,
		Encryption:     true,
		StringToExpect: "foobar",
	}
}

func createCheckInputIMAPData() CreateCheckInput {
	c := CreateCheckInput{
		checkConfiguration:     checkConfigurationData(),
		checkConfigurationIMAP: checkConfigurationIMAPData(),
	}
	c.Type = "imap"
	return c
}

const checkConfigurationIMAPText = "contactids=1234%2C5678&encryption=true&host=example.com&name=My+check&notifyagainevery=1&notifywhenbackup=true&paused=true&port=995&resolution=1&sendnotificationwhendown=2&sendtoandroid=true&sendtoemail=true&sendtoiphone=true&sendtosms=true&sendtotwitter=true&stringtoexpect=foobar&tags=foo%2Cbar&type=imap"

func createCheckOutputData() CreateCheckOutput {
	return CreateCheckOutput{
		Check: createCheckEntry{
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
		checkConfiguration:     checkConfigurationData(),
		checkConfigurationHTTP: checkConfigurationHTTPData(),
	}
	c.Type = "http"
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
	expected := checkConfigurationHTTPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationHTTPCustomQueryText(t *testing.T) {
	in := createCheckInputHTTPCustomData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationHTTPCustomText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationTCPQueryText(t *testing.T) {
	in := createCheckInputTCPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationTCPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationPingQueryText(t *testing.T) {
	in := createCheckInputPingData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationPingText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationDNSQueryText(t *testing.T) {
	in := createCheckInputDNSData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationDNSText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationUDPQueryText(t *testing.T) {
	in := createCheckInputUDPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationUDPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationSMTPQueryText(t *testing.T) {
	in := createCheckInputSMTPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationSMTPText

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationPOP3QueryText(t *testing.T) {
	in := createCheckInputPOP3Data()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationPOP3Text

	if out != expected {
		t.Fatalf("Expected %s, got %s", expected, out)
	}
}

func TestCheckConfigurationIMAPQueryText(t *testing.T) {
	in := createCheckInputIMAPData()
	v, _ := query.Values(in)
	out := v.Encode()
	expected := checkConfigurationIMAPText

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
