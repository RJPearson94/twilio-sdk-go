// Package factor contains auto-generated files. DO NOT MODIFY
package factor

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific factor resource
// See https://www.twilio.com/docs/verify/api/factor for more details
type Client struct {
	client *client.Client

	identity   string
	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the factor resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
	Sid        string
}

// New creates a new instance of the factor client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
