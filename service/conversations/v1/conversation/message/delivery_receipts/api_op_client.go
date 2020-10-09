// Package delivery_receipts contains auto-generated files. DO NOT MODIFY
package delivery_receipts

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing delivery receipt resources
// See https://www.twilio.com/docs/conversations/api/receipt-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	messageSid      string
}

// ClientProperties are the properties required to manage the delivery receipts resources
type ClientProperties struct {
	ConversationSid string
	MessageSid      string
}

// New creates a new instance of the delivery receipts client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		messageSid:      properties.MessageSid,
	}
}
