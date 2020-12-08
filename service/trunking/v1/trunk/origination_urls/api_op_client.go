// Package origination_urls contains auto-generated files. DO NOT MODIFY
package origination_urls

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing Origination URL resources
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource for more details
type Client struct {
	client *client.Client

	trunkSid string
}

// ClientProperties are the properties required to manage the origination urls resources
type ClientProperties struct {
	TrunkSid string
}

// New creates a new instance of the origination urls client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		trunkSid: properties.TrunkSid,
	}
}
