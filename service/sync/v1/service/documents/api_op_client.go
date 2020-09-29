// Package documents contains auto-generated files. DO NOT MODIFY
package documents

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing document resources
// See https://www.twilio.com/docs/sync/api/document-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the documents resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the documents client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
