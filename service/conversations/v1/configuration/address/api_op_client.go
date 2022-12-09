// Package address contains auto-generated files. DO NOT MODIFY
package address

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific address configuration resource
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the address resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the address client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
