// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateConfigurationInput defines input fields for updating a configuration resource
type UpdateConfigurationInput struct {
	DefaultChatServiceSid      *string `form:"DefaultChatServiceSid,omitempty"`
	DefaultClosedTimer         *string `form:"DefaultClosedTimer,omitempty"`
	DefaultInactiveTimer       *string `form:"DefaultInactiveTimer,omitempty"`
	DefaultMessagingServiceSid *string `form:"DefaultMessagingServiceSid,omitempty"`
}

// UpdateConfigurationResponse defines the response fields for the updated configuration
type UpdateConfigurationResponse struct {
	AccountSid                 string  `json:"account_sid"`
	DefaultChatServiceSid      string  `json:"default_chat_service_sid"`
	DefaultClosedTimer         *string `json:"default_closed_timer,omitempty"`
	DefaultInactiveTimer       *string `json:"default_inactive_timer,omitempty"`
	DefaultMessagingServiceSid string  `json:"default_messaging_service_sid"`
	URL                        string  `json:"url"`
}

// Update modifies a configuration resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateConfigurationInput) (*UpdateConfigurationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a configuration resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateConfigurationInput) (*UpdateConfigurationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Configuration",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &UpdateConfigurationInput{}
	}

	response := &UpdateConfigurationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
