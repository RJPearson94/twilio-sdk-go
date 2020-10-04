// Package application contains auto-generated files. DO NOT MODIFY
package application

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific application resource
// See https://www.twilio.com/docs/usage/api/applications for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the application resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the application client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}
