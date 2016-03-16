package pingdom

import (
	"os"
	"testing"
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

func TestPingdomDefaultConfigProviderWithEnv(t *testing.T) {
	setPingdomenv()
	c := DefaultConfigProvider()
	if c.Endpoint != "https://api.pingdom.com" {
		t.Fatalf("Expected Endpoint to be https://api.paybyphone.com, got %s", c.Endpoint)
	}
	if c.EmailAddress != "nobody@example.com" {
		t.Fatalf("Expected EmailAddress to be nobody@example.com, got %s", c.EmailAddress)
	}
	if c.Password != "changeit" {
		t.Fatalf("Expected Password to be changeit, got %s", c.Password)
	}
	if c.AppKey != "abcdefgh0123456789" {
		t.Fatalf("Expected AppKey to be abcdefgh0123456789, got %s", c.AppKey)
	}
}

func TestPingdomDefaultConfigProviderNoEnv(t *testing.T) {
	unsetPingdomenv()
	c := DefaultConfigProvider()
	if c.Endpoint != "https://api.pingdom.com" {
		t.Fatalf("Expected Endpoint to be https://api.paybyphone.com, got %s", c.Endpoint)
	}
	if c.EmailAddress != "" {
		t.Fatalf("Expected EmailAddress to be empty, got %s", c.EmailAddress)
	}
	if c.Password != "" {
		t.Fatalf("Expected Password to be empty, got %s", c.Password)
	}
	if c.AppKey != "" {
		t.Fatalf("Expected AppKey to be empty, got %s", c.AppKey)
	}
}
