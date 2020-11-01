// Package rate_limit contains auto-generated files. DO NOT MODIFY
package rate_limit

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateRateLimitInput defines input fields for updating a rate limit resource
type UpdateRateLimitInput struct {
	Description *string `form:"Description,omitempty"`
}

// UpdateRateLimitResponse defines the response fields for the updated rate limit
type UpdateRateLimitResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Description *string    `json:"description,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  string     `json:"unique_name"`
}

// Update modifies a rate limit resource
// See https://www.twilio.com/docs/verify/api/service-rate-limits#update-a-rate-limit for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateRateLimitInput) (*UpdateRateLimitResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a rate limit resource
// See https://www.twilio.com/docs/verify/api/service-rate-limits#update-a-rate-limit for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateRateLimitInput) (*UpdateRateLimitResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/RateLimits/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateRateLimitInput{}
	}

	response := &UpdateRateLimitResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
