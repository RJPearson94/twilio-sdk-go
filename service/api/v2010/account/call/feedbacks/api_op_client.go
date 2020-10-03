// Package feedbacks contains auto-generated files. DO NOT MODIFY
package feedbacks

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing feedback resources
// See https://www.twilio.com/docs/voice/api/feedback-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	callSid    string
}

// ClientProperties are the properties required to manage the feedbacks resources
type ClientProperties struct {
	AccountSid string
	CallSid    string
}

// New creates a new instance of the feedbacks client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		callSid:    properties.CallSid,
	}
}
