// Package public_key contains auto-generated files. DO NOT MODIFY
package public_key

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific public key resource
// See https://www.twilio.com/docs/iam/credentials/api for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the public key resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the public key client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
