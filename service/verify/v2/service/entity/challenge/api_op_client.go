// Package challenge contains auto-generated files. DO NOT MODIFY
package challenge

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific challenge resource
// See https://www.twilio.com/docs/verify/api/challenge for more details
type Client struct {
	client *client.Client

	identity   string
	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the challenge resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
	Sid        string
}

// New creates a new instance of the challenge client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
