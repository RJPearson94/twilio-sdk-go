// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing user channel resources
// See https://www.twilio.com/docs/chat/rest/user-channel-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	userSid    string
}

// ClientProperties are the properties required to manage the channels resources
type ClientProperties struct {
	ServiceSid string
	UserSid    string
}

// New creates a new instance of the channels client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		userSid:    properties.UserSid,
	}
}
