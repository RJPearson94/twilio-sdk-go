// Package context contains auto-generated files. DO NOT MODIFY
package context

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing execution context resources
// See https://www.twilio.com/docs/studio/rest-api/v2/execution-context for more details
type Client struct {
	client *client.Client

	executionSid string
	flowSid      string
}

// ClientProperties are the properties required to manage the context resources
type ClientProperties struct {
	ExecutionSid string
	FlowSid      string
}

// New creates a new instance of the context client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		executionSid: properties.ExecutionSid,
		flowSid:      properties.FlowSid,
	}
}
