// Package aws_credentials contains auto-generated files. DO NOT MODIFY
package aws_credentials

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing public resources
type Client struct {
	client *client.Client
}

// New creates a new instance of the awscredentials client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
