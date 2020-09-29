// Package interaction contains auto-generated files. DO NOT MODIFY
package interaction

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific interaction resource
// See https://www.twilio.com/docs/proxy/api/interaction for more details
type Client struct {
	client *client.Client

	serviceSid string
	sessionSid string
	sid        string
}

// ClientProperties are the properties required to manage the interaction resources
type ClientProperties struct {
	ServiceSid string
	SessionSid string
	Sid        string
}

// New creates a new instance of the interaction client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sessionSid: properties.SessionSid,
		sid:        properties.Sid,
	}
}
