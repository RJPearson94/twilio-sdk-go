// Package activity contains auto-generated files. DO NOT MODIFY
package activity

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific activity resource
// See https://www.twilio.com/docs/taskrouter/api/activity for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string
}

// ClientProperties are the properties required to manage the activity resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the activity client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,
	}
}
