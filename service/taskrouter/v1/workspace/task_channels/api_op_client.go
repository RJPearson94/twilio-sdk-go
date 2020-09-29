// Package task_channels contains auto-generated files. DO NOT MODIFY
package task_channels

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing task channel resources
// See https://www.twilio.com/docs/taskrouter/api/task-channel for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the task channels resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the task channels client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
