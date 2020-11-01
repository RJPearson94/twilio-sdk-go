// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/configuration/webhook"
)

// Client for managing conversation configuration
type Client struct {
	client *client.Client

	Webhook func() *webhook.Client
}

// New creates a new instance of the configuration client
func New(client *client.Client) *Client {
	return &Client{
		client: client,

		Webhook: func() *webhook.Client { return webhook.New(client) },
	}
}
