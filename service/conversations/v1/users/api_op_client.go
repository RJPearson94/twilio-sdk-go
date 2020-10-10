// Package users contains auto-generated files. DO NOT MODIFY
package users

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing user resources
// See https://www.twilio.com/docs/conversations/api/user-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the users client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
