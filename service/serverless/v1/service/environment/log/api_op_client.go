// Package log contains auto-generated files. DO NOT MODIFY
package log

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific log resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/logs for more details
type Client struct {
	client *client.Client

	environmentSid string
	serviceSid     string
	sid            string
}

// ClientProperties are the properties required to manage the log resources
type ClientProperties struct {
	EnvironmentSid string
	ServiceSid     string
	Sid            string
}

// New creates a new instance of the log client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		environmentSid: properties.EnvironmentSid,
		serviceSid:     properties.ServiceSid,
		sid:            properties.Sid,
	}
}
