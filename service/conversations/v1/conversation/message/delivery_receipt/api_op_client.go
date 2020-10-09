// Package delivery_receipt contains auto-generated files. DO NOT MODIFY
package delivery_receipt

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific delivery receipt resource
// See https://www.twilio.com/docs/conversations/api/receipt-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	messageSid      string
	sid             string
}

// ClientProperties are the properties required to manage the delivery receipt resources
type ClientProperties struct {
	ConversationSid string
	MessageSid      string
	Sid             string
}

// New creates a new instance of the delivery receipt client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		messageSid:      properties.MessageSid,
		sid:             properties.Sid,
	}
}
