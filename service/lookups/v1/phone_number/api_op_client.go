// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a lookup resource
// See https://www.twilio.com/docs/lookup/api for more details
type Client struct {
	client *client.Client

	phoneNumber string
}

// ClientProperties are the properties required to manage the phone number resources
type ClientProperties struct {
	PhoneNumber string
}

// New creates a new instance of the phone number client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		phoneNumber: properties.PhoneNumber,
	}
}
