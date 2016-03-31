package testacc

import (
	"os"
	"testing"
)

// SkipIfNotAcc is designed to skip an integration test if TESTACC is not set.
func SkipIfNotAcc(t *testing.T) {
	if os.Getenv("TESTACC") == "" {
		t.Skipf("Skipping integration test as TESTACC is not set.")
	}
}

// PanicIfMissingEnv is designed to panic if the following environment variables
// are not set:
//
//  * PINGDOM_EMAIL_ADDRESS
//  * PINGDOM_PASSWORD
//  * PINGDOM_APP_KEY
//
// Acceptance tests cannot continue if these are not set so there is no point
// in continuing.
func PanicIfMissingEnv() {
	if os.Getenv("PINGDOM_EMAIL_ADDRESS") == "" || os.Getenv("PINGDOM_PASSWORD") == "" || os.Getenv("PINGDOM_APP_KEY") == "" {
		panic("Please ensure the environment variables PINGDOM_EMAIL_ADDRESS, PINGDOM_PASSWORD, and PINGDOM_APP_KEY are set for acceptance tests")
	}
}

// VetAccConditions is a meta-function that ensures that an acceptance test
// meets the conditions necessary to continue.
func VetAccConditions(t *testing.T) {
	SkipIfNotAcc(t)
	PanicIfMissingEnv()
}
