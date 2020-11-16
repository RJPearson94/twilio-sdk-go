// Package task_channel contains auto-generated files. DO NOT MODIFY
package task_channel

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific task channel resource
// See https://www.twilio.com/docs/taskrouter/api/task-channel for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string
}

// ClientProperties are the properties required to manage the task channel resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the task channel client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,
	}
}
