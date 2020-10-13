// Package available_phone_numbers contains auto-generated files. DO NOT MODIFY
package available_phone_numbers

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing available phone number resources
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the available phone numbers resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the available phone numbers client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
