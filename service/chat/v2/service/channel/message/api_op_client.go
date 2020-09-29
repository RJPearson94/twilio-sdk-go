// Package message contains auto-generated files. DO NOT MODIFY
package message

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific message resource
// See https://www.twilio.com/docs/chat/rest/message-resource for more details
type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the message resources
type ClientProperties struct {
	ChannelSid string
	ServiceSid string
	Sid        string
}

// New creates a new instance of the message client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
