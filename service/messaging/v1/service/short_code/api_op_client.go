// Package short_code contains auto-generated files. DO NOT MODIFY
package short_code

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific short code resource
// See https://www.twilio.com/docs/sms/services/api/shortcode-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the shortcode resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the shortcode client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
