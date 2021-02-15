// Package ip_access_control_list_mapping contains auto-generated files. DO NOT MODIFY
package ip_access_control_list_mapping

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific IP Access Control List Mapping resource
type Client struct {
	client *client.Client

	accountSid string
	domainSid  string
	sid        string
}

// ClientProperties are the properties required to manage the ip access control list mapping resources
type ClientProperties struct {
	AccountSid string
	DomainSid  string
	Sid        string
}

// New creates a new instance of the ip access control list mapping client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		domainSid:  properties.DomainSid,
		sid:        properties.Sid,
	}
}
