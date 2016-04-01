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

// Package sdk is a partial SDK for the Pingdom API.
//
// This is a work in progress, mainly meant to support efforts to get
// Pingdom into Terraform. Many features may be missing, however the request
// framework and data structuring is set up in a way to make it easily
// exstensible, ensuring support for new features can be added in without fuss.
//
// For further usage instructions see the README either in the source
// or at: https://github.com/paybyphone/pingdom-go-sdk/blob/master/README.md
package sdk
