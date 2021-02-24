// Package messaging_configuration contains auto-generated files. DO NOT MODIFY
package messaging_configuration

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateMessagingConfigurationInput defines input fields for updating a messaging configuration resource
type UpdateMessagingConfigurationInput struct {
	MessagingServiceSid string `validate:"required" form:"MessagingServiceSid"`
}

// UpdateMessagingConfigurationResponse defines the response fields for the updated messaging configuration
type UpdateMessagingConfigurationResponse struct {
	AccountSid          string     `json:"account_sid"`
	Country             string     `json:"country"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	MessagingServiceSid string     `json:"messaging_service_sid"`
	ServiceSid          string     `json:"service_sid"`
	URL                 string     `json:"url"`
}

// Update modifies a messaging configuration resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateMessagingConfigurationInput) (*UpdateMessagingConfigurationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a messaging configuration resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateMessagingConfigurationInput) (*UpdateMessagingConfigurationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/MessagingConfigurations/{countryCode}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"countryCode": c.countryCode,
		},
	}

	if input == nil {
		input = &UpdateMessagingConfigurationInput{}
	}

	response := &UpdateMessagingConfigurationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
