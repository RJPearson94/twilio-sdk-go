// Package versions contains auto-generated files. DO NOT MODIFY
package versions

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing plugin version resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-version for more details
type Client struct {
	client *client.Client

	pluginSid string
}

// ClientProperties are the properties required to manage the versions resources
type ClientProperties struct {
	PluginSid string
}

// New creates a new instance of the versions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		pluginSid: properties.PluginSid,
	}
}
