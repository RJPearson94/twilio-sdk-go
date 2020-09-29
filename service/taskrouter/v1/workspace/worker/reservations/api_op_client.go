// Package reservations contains auto-generated files. DO NOT MODIFY
package reservations

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing worker reservation resources
// See https://www.twilio.com/docs/taskrouter/api/worker-reservation for more details
type Client struct {
	client *client.Client

	workerSid    string
	workspaceSid string
}

// ClientProperties are the properties required to manage the reservations resources
type ClientProperties struct {
	WorkerSid    string
	WorkspaceSid string
}

// New creates a new instance of the reservations client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workerSid:    properties.WorkerSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
