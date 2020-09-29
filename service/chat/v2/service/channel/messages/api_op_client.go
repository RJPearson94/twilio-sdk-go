// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing message resources
// See https://www.twilio.com/docs/chat/rest/message-resource for more details
type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
}

// ClientProperties are the properties required to manage the messages resources
type ClientProperties struct {
	ChannelSid string
	ServiceSid string
}

// New creates a new instance of the messages client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
	}
}
