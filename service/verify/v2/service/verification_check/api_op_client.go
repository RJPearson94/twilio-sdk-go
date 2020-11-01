// Package verification_check contains auto-generated files. DO NOT MODIFY
package verification_check

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing verification check resources
// See https://www.twilio.com/docs/verify/api/verification-check for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the verification check resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the verification check client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
