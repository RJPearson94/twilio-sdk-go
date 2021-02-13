// Package ip_access_control_list contains auto-generated files. DO NOT MODIFY
package ip_access_control_list

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific IP access control list resource
// See https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource for more details
type Client struct {
	client *client.Client

	sid      string
	trunkSid string
}

// ClientProperties are the properties required to manage the ip access control list resources
type ClientProperties struct {
	Sid      string
	TrunkSid string
}

// New creates a new instance of the ip access control list client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:      properties.Sid,
		trunkSid: properties.TrunkSid,
	}
}
