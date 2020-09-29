// Package version contains auto-generated files. DO NOT MODIFY
package version

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/version/content"
)

// Client for managing a specific function version resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function-version for more details
type Client struct {
	client *client.Client

	functionSid string
	serviceSid  string
	sid         string

	Content func() *content.Client
}

// ClientProperties are the properties required to manage the version resources
type ClientProperties struct {
	FunctionSid string
	ServiceSid  string
	Sid         string
}

// New creates a new instance of the version client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		functionSid: properties.FunctionSid,
		serviceSid:  properties.ServiceSid,
		sid:         properties.Sid,

		Content: func() *content.Client {
			return content.New(client, content.ClientProperties{
				FunctionSid: properties.FunctionSid,
				ServiceSid:  properties.ServiceSid,
				VersionSid:  properties.Sid,
			})
		},
	}
}
