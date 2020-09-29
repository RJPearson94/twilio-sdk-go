// Package environments contains auto-generated files. DO NOT MODIFY
package environments

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing environment resources
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the environments resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the environments client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
