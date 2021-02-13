// Package ip_access_control_lists contains auto-generated files. DO NOT MODIFY
package ip_access_control_lists

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing IP access control list resources
// See https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource for more details
type Client struct {
	client *client.Client

	trunkSid string
}

// ClientProperties are the properties required to manage the ip access control lists resources
type ClientProperties struct {
	TrunkSid string
}

// New creates a new instance of the ip access control lists client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		trunkSid: properties.TrunkSid,
	}
}
