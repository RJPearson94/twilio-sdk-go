// Package entities contains auto-generated files. DO NOT MODIFY
package entities

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing entity resources
// See https://www.twilio.com/docs/verify/api/entity for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the entities resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the entities client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
