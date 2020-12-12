// Package plugin contains auto-generated files. DO NOT MODIFY
package plugin

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific plugin resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the plugin resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the plugin client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
