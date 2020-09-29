// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing task queue statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics for more details
type Client struct {
	client *client.Client

	taskQueueSid string
	workspaceSid string
}

// ClientProperties are the properties required to manage the statistics resources
type ClientProperties struct {
	TaskQueueSid string
	WorkspaceSid string
}

// New creates a new instance of the statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		taskQueueSid: properties.TaskQueueSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
