// Package domain contains auto-generated files. DO NOT MODIFY
package domain

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific SIP domain resource
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the domain resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the domain client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}
