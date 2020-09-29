// Package environments contains auto-generated files. DO NOT MODIFY
package environments

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateEnvironmentInput defines the input fields for creating a new environment resource
type CreateEnvironmentInput struct {
	DomainSuffix *string `form:"DomainSuffix,omitempty"`
	UniqueName   string  `validate:"required" form:"UniqueName"`
}

// CreateEnvironmentResponse defines the response fields for the created environment
type CreateEnvironmentResponse struct {
	AccountSid   string     `json:"account_sid"`
	BuildSid     *string    `json:"build_sid,omitempty"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	DomainName   string     `json:"domain_name"`
	DomainSuffix *string    `json:"domain_suffix,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Create creates a new environment
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment#create-an-environment-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateEnvironmentInput) (*CreateEnvironmentResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new environment
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment#create-an-environment-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateEnvironmentInput) (*CreateEnvironmentResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Environments",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateEnvironmentInput{}
	}

	response := &CreateEnvironmentResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
