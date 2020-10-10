// Package roles contains auto-generated files. DO NOT MODIFY
package roles

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing role resources
// See https://www.twilio.com/docs/conversations/api/role-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the roles client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
