// Package entity contains auto-generated files. DO NOT MODIFY
package entity

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific entity resource
// See https://www.twilio.com/docs/verify/api/entity for more details
type Client struct {
	client *client.Client

	identity   string
	serviceSid string
}

// ClientProperties are the properties required to manage the entity resources
type ClientProperties struct {
	Identity   string
	ServiceSid string
}

// New creates a new instance of the entity client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:   properties.Identity,
		serviceSid: properties.ServiceSid,
	}
}
