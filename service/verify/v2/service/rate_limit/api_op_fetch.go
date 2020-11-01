// Package rate_limit contains auto-generated files. DO NOT MODIFY
package rate_limit

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchRateLimitResponse defines the response fields for the retrieved rate limit
type FetchRateLimitResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Description *string    `json:"description,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  string     `json:"unique_name"`
}

// Fetch retrieves a rate limit resource
// See https://www.twilio.com/docs/verify/api/service-rate-limits#fetch-a-rate-limit for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRateLimitResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a rate limit resource
// See https://www.twilio.com/docs/verify/api/service-rate-limits#fetch-a-rate-limit for more details
func (c Client) FetchWithContext(context context.Context) (*FetchRateLimitResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/RateLimits/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchRateLimitResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
