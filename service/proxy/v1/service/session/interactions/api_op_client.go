// Package interactions contains auto-generated files. DO NOT MODIFY
package interactions

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing interaction resources
// See https://www.twilio.com/docs/proxy/api/interaction for more details
type Client struct {
	client *client.Client

	serviceSid string
	sessionSid string
}

// ClientProperties are the properties required to manage the interactions resources
type ClientProperties struct {
	ServiceSid string
	SessionSid string
}

// New creates a new instance of the interactions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sessionSid: properties.SessionSid,
	}
}
