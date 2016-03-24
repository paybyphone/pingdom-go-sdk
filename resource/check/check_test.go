package check

import (
	"os"
	"testing"

	"github.com/paybyphone/pingdom-go-sdk/pingdom"
)

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
