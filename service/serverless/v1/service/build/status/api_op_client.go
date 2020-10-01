// Package status contains auto-generated files. DO NOT MODIFY
package status

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific build status resource
type Client struct {
	client *client.Client

	buildSid   string
	serviceSid string
}

// ClientProperties are the properties required to manage the status resources
type ClientProperties struct {
	BuildSid   string
	ServiceSid string
}

// New creates a new instance of the status client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		buildSid:   properties.BuildSid,
		serviceSid: properties.ServiceSid,
	}
}
