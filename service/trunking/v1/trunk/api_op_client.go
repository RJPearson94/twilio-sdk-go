// Package trunk contains auto-generated files. DO NOT MODIFY
package trunk

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the trunk resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the trunk client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
