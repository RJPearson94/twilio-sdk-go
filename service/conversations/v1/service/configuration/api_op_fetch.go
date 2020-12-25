// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchConfigurationResponse defines the response fields for the retrieved service configuration
type FetchConfigurationResponse struct {
	ChatServiceSid                    string `json:"chat_service_sid"`
	DefaultChatServiceRoleSid         string `json:"default_chat_service_role_sid"`
	DefaultConversationCreatorRoleSid string `json:"default_conversation_creator_role_sid"`
	DefaultConversationRoleSid        string `json:"default_conversation_role_sid"`
	ReachabilityEnabled               bool   `json:"reachability_enabled"`
	URL                               string `json:"url"`
}

// Fetch retrieves service configuration resource
// See https://www.twilio.com/docs/conversations/api/service-configuration-resource#fetch-a-serviceconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchConfigurationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves service configuration resource
// See https://www.twilio.com/docs/conversations/api/service-configuration-resource#fetch-a-serviceconfiguration-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchConfigurationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Configuration",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	response := &FetchConfigurationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
