// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing participant resources
// See https://www.twilio.com/docs/proxy/api/participant for more details
type Client struct {
	client *client.Client

	serviceSid string
	sessionSid string
}

// ClientProperties are the properties required to manage the participants resources
type ClientProperties struct {
	ServiceSid string
	SessionSid string
}

// New creates a new instance of the participants client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sessionSid: properties.SessionSid,
	}
}
