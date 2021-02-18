// Package plugin_releases contains auto-generated files. DO NOT MODIFY
package plugin_releases

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing plugin release resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/release for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client
}

// New creates a new instance of the plugin releases client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
