// Package test_users contains auto-generated files. DO NOT MODIFY
package test_users

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing test users resources
// See https://www.twilio.com/docs/studio/rest-api/v2/test-user for more details
type Client struct {
	client *client.Client

	flowSid string
}

// ClientProperties are the properties required to manage the test users resources
type ClientProperties struct {
	FlowSid string
}

// New creates a new instance of the test users client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		flowSid: properties.FlowSid,
	}
}
