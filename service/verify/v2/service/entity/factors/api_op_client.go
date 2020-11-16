// Package factors contains auto-generated files. DO NOT MODIFY
package factors

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing factor resources
// See https://www.twilio.com/docs/verify/api/factor for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	identity   string
	serviceSid string
}

// ClientProperties are the properties required to manage the factors resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
}

// New creates a new instance of the factors client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,
	}
}
