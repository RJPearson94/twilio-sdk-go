// Package rate_limit contains auto-generated files. DO NOT MODIFY
package rate_limit

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/bucket"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/buckets"
)

// Client for managing a specific rate limit resource
// See https://www.twilio.com/docs/verify/api/service-rate-limits for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Bucket  func(string) *bucket.Client
	Buckets *buckets.Client
}

// ClientProperties are the properties required to manage the rate limit resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the rate limit client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Bucket: func(bucketSid string) *bucket.Client {
			return bucket.New(client, bucket.ClientProperties{
				RateLimitSid: properties.Sid,
				ServiceSid:   properties.ServiceSid,
				Sid:          bucketSid,
			})
		},
		Buckets: buckets.New(client, buckets.ClientProperties{
			RateLimitSid: properties.Sid,
			ServiceSid:   properties.ServiceSid,
		}),
	}
}
