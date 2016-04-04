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

// Package client contains generic client structs and methods that are
// designed to be used by specific Pingdom services and resources.
package client

import (
	"github.com/imdario/mergo"
	"github.com/paybyphone/pingdom-go-sdk/pingdom"
	"github.com/paybyphone/pingdom-go-sdk/pingdom/request"
)

// Client encompasses a generic client object that is further extended by
// services. Any common configuration and functionality goes here.
type Client struct {
	// The configuration for this specific connection.
	Config pingdom.Config
}

// New handles logic for either setting a conneciton based on supplied
// configuration, or getting the configuration from a specific provider.
func New(configs ...pingdom.Config) *Client {
	c := &Client{
		Config: pingdom.DefaultConfigProvider(),
	}
	for _, v := range configs {
		mergo.MergeWithOverwrite(&c.Config, v)
	}
	return c
}

// SendRequest sends a request to a request.Request object.
// It's expected that references to specific data types are passed - no
// checking is done to make sure that references are passed.
func (c *Client) SendRequest(method, uri string, in, out interface{}) error {
	r := request.NewRequest(c.Config)
	r.Method = method
	r.URI = uri
	r.Input = in
	r.Output = out
	err := r.Send()
	if err != nil {
		return err
	}
	return nil
}
