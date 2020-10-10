// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration/notification"
)

// Client for managing service configuration
// See https://www.twilio.com/docs/conversations/api/service-configuration-resource for more details
type Client struct {
	client *client.Client

	serviceSid string

	Notification func() *notification.Client
}

// ClientProperties are the properties required to manage the configuration resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the configuration client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,

		Notification: func() *notification.Client {
			return notification.New(client, notification.ClientProperties{
				ServiceSid: properties.ServiceSid,
			})
		},
	}
}
