// Package access_tokens contains auto-generated files. DO NOT MODIFY
package access_tokens

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing access token resources
// See https://www.twilio.com/docs/verify/api/access-token for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the access tokens resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the access tokens client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
