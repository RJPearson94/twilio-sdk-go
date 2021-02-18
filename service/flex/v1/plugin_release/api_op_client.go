// Package plugin_release contains auto-generated files. DO NOT MODIFY
package plugin_release

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific plugin release resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/release for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the plugin release resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the plugin release client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
