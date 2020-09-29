// Package revisions contains auto-generated files. DO NOT MODIFY
package revisions

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing flow revision resources
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-revision for more details
type Client struct {
	client *client.Client

	flowSid string
}

// ClientProperties are the properties required to manage the revisions resources
type ClientProperties struct {
	FlowSid string
}

// New creates a new instance of the revisions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		flowSid: properties.FlowSid,
	}
}
