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
