// Package logs contains auto-generated files. DO NOT MODIFY
package logs

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing log resources
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/logs for more details
type Client struct {
	client *client.Client

	environmentSid string
	serviceSid     string
}

// ClientProperties are the properties required to manage the logs resources
type ClientProperties struct {
	EnvironmentSid string
	ServiceSid     string
}

// New creates a new instance of the logs client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		environmentSid: properties.EnvironmentSid,
		serviceSid:     properties.ServiceSid,
	}
}
