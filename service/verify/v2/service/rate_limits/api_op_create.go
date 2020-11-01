// Package rate_limits contains auto-generated files. DO NOT MODIFY
package rate_limits

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateRateLimitInput defines the input fields for creating a new rate limit
type CreateRateLimitInput struct {
	Description *string `form:"Description,omitempty"`
	UniqueName  string  `validate:"required" form:"UniqueName"`
}

// CreateRateLimitResponse defines the response fields for the created rate limit
type CreateRateLimitResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Description *string    `json:"description,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  string     `json:"unique_name"`
}

// Create creates a new rate limit
// See https://www.twilio.com/docs/verify/api/service-rate-limits#create-a-rate-limit for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateRateLimitInput) (*CreateRateLimitResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new rate limit
// See https://www.twilio.com/docs/verify/api/service-rate-limits#create-a-rate-limit for more details
func (c Client) CreateWithContext(context context.Context, input *CreateRateLimitInput) (*CreateRateLimitResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/RateLimits",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateRateLimitInput{}
	}

	response := &CreateRateLimitResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
