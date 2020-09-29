// Package workflows contains auto-generated files. DO NOT MODIFY
package workflows

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing workflow resources
// See https://www.twilio.com/docs/taskrouter/api/workflow for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the workflows resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the workflows client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
