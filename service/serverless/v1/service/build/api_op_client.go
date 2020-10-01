// Package build contains auto-generated files. DO NOT MODIFY
package build

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/build/status"
)

// Client for managing a specific build resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Status func() *status.Client
}

// ClientProperties are the properties required to manage the build resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the build client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Status: func() *status.Client {
			return status.New(client, status.ClientProperties{
				BuildSid:   properties.Sid,
				ServiceSid: properties.ServiceSid,
			})
		},
	}
}
