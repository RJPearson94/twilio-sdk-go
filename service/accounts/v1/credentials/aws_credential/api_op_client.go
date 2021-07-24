// Package aws_credential contains auto-generated files. DO NOT MODIFY
package aws_credential

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific aws credential resource
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the awscredential resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the awscredential client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
