// Package reservation contains auto-generated files. DO NOT MODIFY
package reservation

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific worker reservation resource
// See https://www.twilio.com/docs/taskrouter/api/worker-reservation for more details
type Client struct {
	client *client.Client

	sid          string
	workerSid    string
	workspaceSid string
}

// ClientProperties are the properties required to manage the reservation resources
type ClientProperties struct {
	Sid          string
	WorkerSid    string
	WorkspaceSid string
}

// New creates a new instance of the reservation client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workerSid:    properties.WorkerSid,
		workspaceSid: properties.WorkspaceSid,
	}
}
