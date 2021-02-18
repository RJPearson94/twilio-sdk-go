// Package version contains auto-generated files. DO NOT MODIFY
package version

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	pluginSid string
	sid       string
}

// ClientProperties are the properties required to manage the version resources
type ClientProperties struct {
	PluginSid string
	Sid       string
}

// New creates a new instance of the version client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		pluginSid: properties.PluginSid,
		sid:       properties.Sid,
	}
}
