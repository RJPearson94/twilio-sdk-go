// Package trunks contains auto-generated files. DO NOT MODIFY
package trunks

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing trunk resources
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the trunks client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
