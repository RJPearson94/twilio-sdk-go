// Package bindings contains auto-generated files. DO NOT MODIFY
package bindings

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateBindingInput defines the input fields for creating a new service binding
type CreateBindingInput struct {
	Address                     string    `validate:"required" form:"Address"`
	BindingType                 string    `validate:"required" form:"BindingType"`
	CredentialSid               *string   `form:"CredentialSid,omitempty"`
	Identity                    string    `validate:"required" form:"Identity"`
	NotificationProtocolVersion *string   `form:"NotificationProtocolVersion,omitempty"`
	Tags                        *[]string `form:"Tag,omitempty"`
}

// CreateBindingResponse defines the response fields for the retrieved service binding
type CreateBindingResponse struct {
	AccountSid                  string     `json:"account_sid"`
	Address                     string     `json:"address"`
	BindingType                 string     `json:"binding_type"`
	CredentialSid               *string    `json:"credential_sid,omitempty"`
	DateCreated                 time.Time  `json:"date_created"`
	DateUpdated                 *time.Time `json:"date_updated,omitempty"`
	Identity                    string     `json:"identity"`
	NotificationProtocolVersion string     `json:"notification_protocol_version"`
	ServiceSid                  string     `json:"service_sid"`
	Sid                         string     `json:"sid"`
	Tags                        []string   `json:"tags"`
	URL                         string     `json:"url"`
}

// Create creates a service binding resource
// See https://www.twilio.com/docs/notify/api/binding-resource#create-a-binding-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateBindingInput) (*CreateBindingResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a service binding resource
// See https://www.twilio.com/docs/notify/api/binding-resource#create-a-binding-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateBindingInput) (*CreateBindingResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Bindings",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateBindingInput{}
	}

	response := &CreateBindingResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
