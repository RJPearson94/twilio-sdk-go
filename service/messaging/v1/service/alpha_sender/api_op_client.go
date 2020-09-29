// Package alpha_sender contains auto-generated files. DO NOT MODIFY
package alpha_sender

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific alpha sender resource
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the alphasender resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the alphasender client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
