// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific phone number resource
// See https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource for more details
type Client struct {
	client *client.Client

	sid      string
	trunkSid string
}

// ClientProperties are the properties required to manage the phone number resources
type ClientProperties struct {
	Sid      string
	TrunkSid string
}

// New creates a new instance of the phone number client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:      properties.Sid,
		trunkSid: properties.TrunkSid,
	}
}
