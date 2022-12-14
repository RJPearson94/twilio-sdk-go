// Package address contains auto-generated files. DO NOT MODIFY
package address

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateAutoCreationInput struct {
	ConversationServiceSid *string   `form:"ConversationServiceSid,omitempty"`
	Enabled                *bool     `form:"Enabled,omitempty"`
	StudioFlowSid          *string   `form:"StudioFlowSid,omitempty"`
	StudioRetryCount       *int      `form:"StudioRetryCount,omitempty"`
	Type                   *string   `form:"Type,omitempty"`
	WebhookFilters         *[]string `form:"WebhookFilters,omitempty"`
	WebhookMethod          *string   `form:"WebhookMethod,omitempty"`
	WebhookUrl             *string   `form:"WebhookUrl,omitempty"`
}

// UpdateAddressInput defines input fields for updating a address configuration resource
type UpdateAddressInput struct {
	Address      *string                  `form:"Address,omitempty"`
	AutoCreation *UpdateAutoCreationInput `form:"AutoCreation,omitempty"`
	FriendlyName *string                  `form:"FriendlyName,omitempty"`
	Type         *string                  `form:"Type,omitempty"`
}

type UpdateAutoCreationResponse struct {
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

// UpdateAddressResponse defines the response fields for the updated address configuration
type UpdateAddressResponse struct {
	AccountSid   string                     `json:"account_sid"`
	Address      string                     `json:"address"`
	AutoCreation UpdateAutoCreationResponse `json:"auto_creation"`
	DateCreated  time.Time                  `json:"date_created"`
	DateUpdated  *time.Time                 `json:"date_updated,omitempty"`
	FriendlyName *string                    `json:"friendly_name,omitempty"`
	Sid          string                     `json:"sid"`
	Type         string                     `json:"type"`
	URL          string                     `json:"url"`
}

// Update modifies a address configuration resource
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#update-an-addressconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateAddressInput) (*UpdateAddressResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a address configuration resource
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#update-an-addressconfiguration-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateAddressInput) (*UpdateAddressResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Configuration/Addresses/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateAddressInput{}
	}

	response := &UpdateAddressResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
