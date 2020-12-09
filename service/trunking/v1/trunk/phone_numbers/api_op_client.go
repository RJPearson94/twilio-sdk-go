// Package phone_numbers contains auto-generated files. DO NOT MODIFY
package phone_numbers

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing phone number resources
// See https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource for more details
type Client struct {
	client *client.Client

	trunkSid string
}

// ClientProperties are the properties required to manage the phone numbers resources
type ClientProperties struct {
	TrunkSid string
}

// New creates a new instance of the phone numbers client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		trunkSid: properties.TrunkSid,
	}
}
