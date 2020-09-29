// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing all worker statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics for more details
type Client struct {
	client *client.Client

	workerSid    string
	workspaceSid string
}

// ClientProperties are the properties required to manage the statistics resources
type ClientProperties struct {
	WorkerSid    string
	WorkspaceSid string
}

// New creates a new instance of the statistics client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workerSid:    properties.WorkerSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
