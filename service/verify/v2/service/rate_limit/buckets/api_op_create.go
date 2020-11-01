// Package buckets contains auto-generated files. DO NOT MODIFY
package buckets

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateBucketInput defines the input fields for creating a new rate limit bucket
type CreateBucketInput struct {
	Interval int `validate:"required" form:"Interval"`
	Max      int `validate:"required" form:"Max"`
}

// CreateBucketResponse defines the response fields for the created rate limit bucket
type CreateBucketResponse struct {
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

// Create creates a new rate limit bucket
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#create-a-bucket for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateBucketInput) (*CreateBucketResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new rate limit bucket
// See https://www.twilio.com/docs/verify/api/service-rate-limit-buckets#create-a-bucket for more details
func (c Client) CreateWithContext(context context.Context, input *CreateBucketInput) (*CreateBucketResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/RateLimits/{rateLimitSid}/Buckets",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":   c.serviceSid,
			"rateLimitSid": c.rateLimitSid,
		},
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	response := &CreateBucketResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
