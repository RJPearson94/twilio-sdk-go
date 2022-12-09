// Package address contains auto-generated files. DO NOT MODIFY
package address

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchAutoCreationResponse struct {
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

// FetchAddressResponse defines the response fields for the retrieved address configuration
type FetchAddressResponse struct {
	AccountSid   string                    `json:"account_sid"`
	Address      string                    `json:"address"`
	AutoCreation FetchAutoCreationResponse `json:"auto_creation"`
	DateCreated  time.Time                 `json:"date_created"`
	DateUpdated  *time.Time                `json:"date_updated,omitempty"`
	FriendlyName *string                   `json:"friendly_name,omitempty"`
	Sid          string                    `json:"sid"`
	Type         string                    `json:"type"`
	URL          string                    `json:"url"`
}

// Fetch retrieves a address configuration resource
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#fetch-an-addressconfiguration-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAddressResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a address configuration resource
// See https://www.twilio.com/docs/conversations/api/address-configuration-resource#fetch-an-addressconfiguration-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAddressResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Configuration/Addresses/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchAddressResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
