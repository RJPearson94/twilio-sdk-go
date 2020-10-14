// Package incoming_phone_number contains auto-generated files. DO NOT MODIFY
package incoming_phone_number

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific phone number resource
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the incoming phone number resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the incoming phone number client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}
