// Package member contains auto-generated files. DO NOT MODIFY
package member

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific member resource
// See https://www.twilio.com/docs/chat/rest/member-resource for more details
type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the member resources
type ClientProperties struct {
	ChannelSid string
	ServiceSid string
	Sid        string
}

// New creates a new instance of the member client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
