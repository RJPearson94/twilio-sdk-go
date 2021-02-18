// Package plugins contains auto-generated files. DO NOT MODIFY
package plugins

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing plugin resources
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	configurationSid string
}

// ClientProperties are the properties required to manage the plugins resources
type ClientProperties struct {
	ConfigurationSid string
}

// New creates a new instance of the plugins client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		configurationSid: properties.ConfigurationSid,
	}
}
