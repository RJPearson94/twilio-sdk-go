// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateServiceInput defines the input fields for creating a new service resource
type CreateServiceInput struct {
	FriendlyName       string `validate:"required" form:"FriendlyName"`
	IncludeCredentials *bool  `form:"IncludeCredentials,omitempty"`
	UiEditable         *bool  `form:"UiEditable,omitempty"`
	UniqueName         string `validate:"required" form:"UniqueName"`
}

// CreateServiceResponse defines the response fields for the created service
type CreateServiceResponse struct {
	AccountSid         string     `json:"account_sid"`
	DateCreated        time.Time  `json:"date_created"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	DomainBase         string     `json:"domain_base"`
	FriendlyName       string     `json:"friendly_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	Sid                string     `json:"sid"`
	URL                string     `json:"url"`
	UiEditable         bool       `json:"ui_editable"`
	UniqueName         string     `json:"unique_name"`
}

// Create creates a new service
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service#create-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new service
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service#create-a-service-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
