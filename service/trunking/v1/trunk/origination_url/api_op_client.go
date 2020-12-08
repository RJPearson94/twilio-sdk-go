// Package origination_url contains auto-generated files. DO NOT MODIFY
package origination_url

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource for more details
type Client struct {
	client *client.Client

	sid      string
	trunkSid string
}

// ClientProperties are the properties required to manage the origination url resources
type ClientProperties struct {
	Sid      string
	TrunkSid string
}

// New creates a new instance of the origination url client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:      properties.Sid,
		trunkSid: properties.TrunkSid,
	}
}
