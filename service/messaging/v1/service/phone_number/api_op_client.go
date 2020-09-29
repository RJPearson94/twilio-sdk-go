// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific phone number resource
// See https://www.twilio.com/docs/sms/services/api/phonenumber-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the phonenumber resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the phonenumber client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
