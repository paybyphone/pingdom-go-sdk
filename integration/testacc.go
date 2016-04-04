// Copyright 2016 PayByPhone Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package testacc contains helper methods for running acceptance tests.
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
