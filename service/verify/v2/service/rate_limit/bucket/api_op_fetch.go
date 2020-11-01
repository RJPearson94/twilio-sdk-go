// Package bucket contains auto-generated files. DO NOT MODIFY
package bucket

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchBucketResponse defines the response fields for the retrieved rate limit bucket
type FetchBucketResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Interval     int        `json:"interval"`
	Max          int        `json:"max"`
	RateLimitSid string     `json:"rate_limit_sid"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a rate limit bucket resource
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#fetch-a-bucket for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchBucketResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a rate limit bucket resource
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#fetch-a-bucket for more details
func (c Client) FetchWithContext(context context.Context) (*FetchBucketResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/RateLimits/{rateLimitSid}/Buckets/{sid}",
		PathParams: map[string]string{
			"serviceSid":   c.serviceSid,
			"rateLimitSid": c.rateLimitSid,
			"sid":          c.sid,
		},
	}

	response := &FetchBucketResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
