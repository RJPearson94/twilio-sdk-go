// Package phone_numbers contains auto-generated files. DO NOT MODIFY
package phone_numbers

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing phone number resources
// See https://www.twilio.com/docs/proxy/api/phone-number for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the phone numbers resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the phone numbers client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
