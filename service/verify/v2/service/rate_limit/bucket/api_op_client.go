// Package bucket contains auto-generated files. DO NOT MODIFY
package bucket

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific rate limit bucket resource
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets for more details
type Client struct {
	client *client.Client

	rateLimitSid string
	serviceSid   string
	sid          string
}

// ClientProperties are the properties required to manage the bucket resources
type ClientProperties struct {
	RateLimitSid string
	ServiceSid   string
	Sid          string
}

// New creates a new instance of the bucket client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		rateLimitSid: properties.RateLimitSid,
		serviceSid:   properties.ServiceSid,
		sid:          properties.Sid,
	}
}
