// Package ip_access_control_lists contains auto-generated files. DO NOT MODIFY
package ip_access_control_lists

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing IP access control list resources
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the ip access control lists resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the ip access control lists client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
