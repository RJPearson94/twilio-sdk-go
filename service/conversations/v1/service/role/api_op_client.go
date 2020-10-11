// Package role contains auto-generated files. DO NOT MODIFY
package role

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific role resource
// See https://www.twilio.com/docs/conversations/api/role-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the role resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the role client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
