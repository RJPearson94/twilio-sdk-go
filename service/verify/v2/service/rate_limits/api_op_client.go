// Package rate_limits contains auto-generated files. DO NOT MODIFY
package rate_limits

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing rate limit resources
// See https://www.twilio.com/docs/verify/api/service-rate-limits for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the rate limits resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the rate limits client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
