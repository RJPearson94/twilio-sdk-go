// Package functions contains auto-generated files. DO NOT MODIFY
package functions

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateFunctionInput defines the input fields for creating a new function resource
type CreateFunctionInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// CreateFunctionResponse defines the response fields for the created function
type CreateFunctionResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a new function
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#create-a-function-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFunctionInput) (*CreateFunctionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new function
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#create-a-function-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFunctionInput) (*CreateFunctionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Functions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateFunctionInput{}
	}

	response := &CreateFunctionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
