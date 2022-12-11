// Package addresses contains auto-generated files. DO NOT MODIFY
package addresses

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateAutoCreationInput struct {
	ConversationServiceSid *string   `form:"ConversationServiceSid,omitempty"`
	Enabled                bool      `form:"Enabled"`
	StudioFlowSid          *string   `form:"StudioFlowSid,omitempty"`
	StudioRetryCount       *int      `form:"StudioRetryCount,omitempty"`
	Type                   string    `validate:"required" form:"Type"`
	WebhookFilters         *[]string `form:"WebhookFilters,omitempty"`
	WebhookMethod          *string   `form:"WebhookMethod,omitempty"`
	WebhookUrl             *string   `form:"WebhookUrl,omitempty"`
}

// CreateAddressInput defines the input fields for creating a new address configuration resource
type CreateAddressInput struct {
	Address      string                   `validate:"required" form:"Address"`
	AutoCreation *CreateAutoCreationInput `form:"AutoCreation,omitempty"`
	FriendlyName *string                  `form:"FriendlyName,omitempty"`
	Type         string                   `validate:"required" form:"Type"`
}

type CreateAutoCreationResponse struct {
	BindingName            *string   `json:"binding_name,omitempty"`
	ConversationServiceSid *string   `json:"conversation_service_sid,omitempty"`
	Enabled                bool      `json:"enabled"`
	StudioFlowSid          *string   `json:"studio_flow_sid,omitempty"`
	StudioRetryCount       *int      `json:"studio_retry_count,omitempty"`
	Type                   string    `json:"type"`
	WebhookFilters         *[]string `json:"webhook_filters,omitempty"`
	WebhookMethod          *string   `json:"webhook_method,omitempty"`
	WebhookUrl             *string   `json:"webhook_url,omitempty"`
}

// CreateAddressResponse defines the response fields for the created address configuration
type CreateAddressResponse struct {
	AccountSid   string                     `json:"account_sid"`
	Address      string                     `json:"address"`
	AutoCreation CreateAutoCreationResponse `json:"auto_creation"`
	DateCreated  time.Time                  `json:"date_created"`
	DateUpdated  *time.Time                 `json:"date_updated,omitempty"`
	FriendlyName *string                    `json:"friendly_name,omitempty"`
	Sid          string                     `json:"sid"`
	Type         string                     `json:"type"`
	URL          string                     `json:"url"`
}

// Create creates a new address configuration
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#create-an-addressconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAddressInput) (*CreateAddressResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new address configuration
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#create-an-addressconfiguration-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAddressInput) (*CreateAddressResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Configuration/Addresses",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateAddressInput{}
	}

	response := &CreateAddressResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
