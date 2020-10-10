// Package user contains auto-generated files. DO NOT MODIFY
package user

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific user resource
// See https://www.twilio.com/docs/conversations/api/user-resource for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the user resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the user client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
