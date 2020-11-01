// Package buckets contains auto-generated files. DO NOT MODIFY
package buckets

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing rate limit bucket resources
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets for more details
type Client struct {
	client *client.Client

	rateLimitSid string
	serviceSid   string
}

// ClientProperties are the properties required to manage the buckets resources
type ClientProperties struct {
	RateLimitSid string
	ServiceSid   string
}

// New creates a new instance of the buckets client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		rateLimitSid: properties.RateLimitSid,
		serviceSid:   properties.ServiceSid,
	}
}
