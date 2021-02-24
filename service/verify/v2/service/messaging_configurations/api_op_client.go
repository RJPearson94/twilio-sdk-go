// Package messaging_configurations contains auto-generated files. DO NOT MODIFY
package messaging_configurations

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing messaging configuration resources
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the messaging configurations resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the messaging configurations client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
