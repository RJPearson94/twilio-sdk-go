// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing webhook resources
// See https://www.twilio.com/docs/verify/api/webhooks for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the webhooks resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the webhooks client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
