// Package accounts contains auto-generated files. DO NOT MODIFY
package accounts

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing account resources
// See https://www.twilio.com/docs/iam/api/account for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the accounts client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
