// This is an autogenerated file. DO NOT MODIFY
package worker

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific worker resource
// See https://www.twilio.com/docs/taskrouter/api/worker for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string
}

// ClientProperties are the properties required to manage the worker resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the worker client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,
	}
}
