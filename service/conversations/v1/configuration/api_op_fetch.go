// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchConfigurationResponse defines the response fields for the retrieved configuration
type FetchConfigurationResponse struct {
	AccountSid                 string  `json:"account_sid"`
	DefaultChatServiceSid      string  `json:"default_chat_service_sid"`
	DefaultClosedTimer         *string `json:"default_closed_timer,omitempty"`
	DefaultInactiveTimer       *string `json:"default_inactive_timer,omitempty"`
	DefaultMessagingServiceSid string  `json:"default_messaging_service_sid"`
	URL                        string  `json:"url"`
}

// Fetch retrieves configuration resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchConfigurationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves configuration resource
func (c Client) FetchWithContext(context context.Context) (*FetchConfigurationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Configuration",
	}

	response := &FetchConfigurationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
