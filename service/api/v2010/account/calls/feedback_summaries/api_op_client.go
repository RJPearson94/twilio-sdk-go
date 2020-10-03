// Package feedback_summaries contains auto-generated files. DO NOT MODIFY
package feedback_summaries

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing call resources
// See https://www.twilio.com/docs/voice/api/feedbacksummary-resource for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the feedback summaries resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the feedback summaries client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
