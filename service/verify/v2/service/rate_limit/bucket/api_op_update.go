// Package bucket contains auto-generated files. DO NOT MODIFY
package bucket

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateBucketInput defines input fields for updating a rate limit bucket resource
type UpdateBucketInput struct {
	Interval *int `form:"Interval,omitempty"`
	Max      *int `form:"Max,omitempty"`
}

// UpdateBucketResponse defines the response fields for the updated rate limit bucket
type UpdateBucketResponse struct {
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

// Update modifies a rate limit bucket resource
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#update-a-bucket for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateBucketInput) (*UpdateBucketResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a rate limit bucket resource
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#update-a-bucket for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateBucketInput) (*UpdateBucketResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/RateLimits/{rateLimitSid}/Buckets/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":   c.serviceSid,
			"rateLimitSid": c.rateLimitSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateBucketInput{}
	}

	response := &UpdateBucketResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
