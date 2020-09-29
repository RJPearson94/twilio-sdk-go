// Package version contains auto-generated files. DO NOT MODIFY
package version

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific asset version resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset for more details
type Client struct {
	client *client.Client

	assetSid   string
	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the version resources
type ClientProperties struct {
	AssetSid   string
	ServiceSid string
	Sid        string
}

// New creates a new instance of the version client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assetSid:   properties.AssetSid,
		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
