// Package ip_addresses contains auto-generated files. DO NOT MODIFY
package ip_addresses

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing IP address resources
type Client struct {
	client *client.Client

	accountSid             string
	ipAccessControlListSid string
}

// ClientProperties are the properties required to manage the ip addresses resources
type ClientProperties struct {
	AccountSid             string
	IpAccessControlListSid string
}

// New creates a new instance of the ip addresses client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:             properties.AccountSid,
		ipAccessControlListSid: properties.IpAccessControlListSid,
	}
}
