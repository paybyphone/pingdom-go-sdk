package pingdom

import (
	"os"
	"testing"
)

func SetPingdomenv() {
	os.Setenv("PINGDOM_EMAIL_ADDRESS", "nobody@example.com")
	os.Setenv("PINGDOM_PASSWORD", "changeit")
	os.Setenv("PINGDOM_APP_KEY", "abcdefgh0123456789")
}

func UnsetPingdomenv() {
	os.Unsetenv("PINGDOM_EMAIL_ADDRESS")
	os.Unsetenv("PINGDOM_PASSWORD")
	os.Unsetenv("PINGDOM_APP_KEY")
}

func TestPingdomDefaultConfigProvider(t *testing.T) {
}
