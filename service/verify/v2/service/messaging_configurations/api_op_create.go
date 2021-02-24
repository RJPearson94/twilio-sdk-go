// Package messaging_configurations contains auto-generated files. DO NOT MODIFY
package messaging_configurations

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateMessagingConfigurationInput defines the input fields for creating a new messaging configuration
type CreateMessagingConfigurationInput struct {
	Country             string `validate:"required" form:"Country"`
	MessagingServiceSid string `validate:"required" form:"MessagingServiceSid"`
}

// CreateMessagingConfigurationResponse defines the response fields for the created messaging configuration
type CreateMessagingConfigurationResponse struct {
	AccountSid          string     `json:"account_sid"`
	Country             string     `json:"country"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	MessagingServiceSid string     `json:"messaging_service_sid"`
	ServiceSid          string     `json:"service_sid"`
	URL                 string     `json:"url"`
}

// Create creates a new messaging configuration
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateMessagingConfigurationInput) (*CreateMessagingConfigurationResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new messaging configuration
func (c Client) CreateWithContext(context context.Context, input *CreateMessagingConfigurationInput) (*CreateMessagingConfigurationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/MessagingConfigurations",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateMessagingConfigurationInput{}
	}

	response := &CreateMessagingConfigurationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
