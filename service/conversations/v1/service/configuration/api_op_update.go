// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateConfigurationInput defines input fields for updating a service configuration resource
type UpdateConfigurationInput struct {
	DefaultChatServiceRoleSid         *string `form:"DefaultChatServiceRoleSid,omitempty"`
	DefaultConversationCreatorRoleSid *string `form:"DefaultConversationCreatorRoleSid,omitempty"`
	DefaultConversationRoleSid        *string `form:"DefaultConversationRoleSid,omitempty"`
	ReachabilityEnabled               *bool   `form:"ReachabilityEnabled,omitempty"`
}

// UpdateConfigurationResponse defines the response fields for the updated service configuration
type UpdateConfigurationResponse struct {
	ChatServiceSid                    string `json:"chat_service_sid"`
	DefaultChatServiceRoleSid         string `json:"default_chat_service_role_sid"`
	DefaultConversationCreatorRoleSid string `json:"default_conversation_creator_role_sid"`
	DefaultConversationRoleSid        string `json:"default_conversation_role_sid"`
	ReachabilityEnabled               bool   `json:"reachability_enabled"`
	URL                               string `json:"url"`
}

// Update modifies a service configuration resource
// See https://www.twilio.com/docs/conversations/api/service-configuration-resource#update-a-serviceconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateConfigurationInput) (*UpdateConfigurationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service configuration resource
// See https://www.twilio.com/docs/conversations/api/service-configuration-resource#update-a-serviceconfiguration-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateConfigurationInput) (*UpdateConfigurationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Configuration",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
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
