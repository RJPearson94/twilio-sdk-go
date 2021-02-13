// Package ip_address contains auto-generated files. DO NOT MODIFY
package ip_address

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific IP Address resource
type Client struct {
	client *client.Client

	accountSid             string
	ipAccessControlListSid string
	sid                    string
}

// ClientProperties are the properties required to manage the ip address resources
type ClientProperties struct {
	AccountSid             string
	IpAccessControlListSid string
	Sid                    string
}

// New creates a new instance of the ip address client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:             properties.AccountSid,
		ipAccessControlListSid: properties.IpAccessControlListSid,
		sid:                    properties.Sid,
	}
}
