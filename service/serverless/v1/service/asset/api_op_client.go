// Package asset contains auto-generated files. DO NOT MODIFY
package asset

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/version"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
)

// Client for managing a specific asset resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Version  func(string) *version.Client
	Versions *versions.Client
}

// ClientProperties are the properties required to manage the asset resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the asset client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Version: func(versionSid string) *version.Client {
			return version.New(client, version.ClientProperties{
				AssetSid:   properties.Sid,
				ServiceSid: properties.ServiceSid,
				Sid:        versionSid,
			})
		},
		Versions: versions.New(client, versions.ClientProperties{
			AssetSid:   properties.Sid,
			ServiceSid: properties.ServiceSid,
		}),
	}
}
