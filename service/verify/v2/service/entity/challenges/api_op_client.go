// Package challenges contains auto-generated files. DO NOT MODIFY
package challenges

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing challenge resources
// See https://www.twilio.com/docs/verify/api/challenge for more details
type Client struct {
	client *client.Client

	identity   string
	serviceSid string
}

// ClientProperties are the properties required to manage the challenges resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
}

// New creates a new instance of the challenges client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,
	}
}
