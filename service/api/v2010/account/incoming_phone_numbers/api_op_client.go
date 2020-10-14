// Package incoming_phone_numbers contains auto-generated files. DO NOT MODIFY
package incoming_phone_numbers

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing phone number resources
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the incoming phone numbers resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the incoming phone numbers client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
