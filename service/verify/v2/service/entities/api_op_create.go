// Package entities contains auto-generated files. DO NOT MODIFY
package entities

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateEntityInput defines the input fields for creating a new entity
type CreateEntityInput struct {
	Identity string `validate:"required" form:"Identity"`
}

// CreateEntityResponse defines the response fields for the created entity
type CreateEntityResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    string     `json:"identity"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Create creates a new entity
// See https://www.twilio.com/docs/verify/api/entity#create-an-entity-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Create(input *CreateEntityInput) (*CreateEntityResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new entity
// See https://www.twilio.com/docs/verify/api/entity#create-an-entity-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) CreateWithContext(context context.Context, input *CreateEntityInput) (*CreateEntityResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Entities",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateEntityInput{}
	}

	response := &CreateEntityResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
