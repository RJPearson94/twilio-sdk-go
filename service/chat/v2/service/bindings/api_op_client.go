// Package bindings contains auto-generated files. DO NOT MODIFY
package bindings

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing binding resources
// See https://www.twilio.com/docs/chat/rest/binding-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the bindings resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the bindings client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
