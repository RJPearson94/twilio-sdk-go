// Package notification contains auto-generated files. DO NOT MODIFY
package notification

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing service notification
// See https://www.twilio.com/docs/conversations/api/service-notification-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the notification resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the notification client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
