// Package deployment contains auto-generated files. DO NOT MODIFY
package deployment

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific deployment resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment for more details
type Client struct {
	client *client.Client

	environmentSid string
	serviceSid     string
	sid            string
}

// ClientProperties are the properties required to manage the deployment resources
type ClientProperties struct {
	EnvironmentSid string
	ServiceSid     string
	Sid            string
}

// New creates a new instance of the deployment client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		environmentSid: properties.EnvironmentSid,
		serviceSid:     properties.ServiceSid,
		sid:            properties.Sid,
	}
}
