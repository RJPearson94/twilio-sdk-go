// Package addresses contains auto-generated files. DO NOT MODIFY
package addresses

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing address configuration resources
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the addresses client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
