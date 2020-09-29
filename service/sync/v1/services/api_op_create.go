// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateServiceInput defines the input fields for creating a new service resource
type CreateServiceInput struct {
	AclEnabled                    *bool   `form:"AclEnabled,omitempty"`
	FriendlyName                  *string `form:"FriendlyName,omitempty"`
	ReachabilityDebouncingEnabled *bool   `form:"ReachabilityDebouncingEnabled,omitempty"`
	ReachabilityDebouncingWindow  *int    `form:"ReachabilityDebouncingWindow,omitempty"`
	ReachabilityWebhooksEnabled   *bool   `form:"ReachabilityWebhooksEnabled,omitempty"`
	WebhookURL                    *string `form:"WebhookUrl,omitempty"`
	WebhooksFromRestEnabled       *bool   `form:"WebhooksFromRestEnabled,omitempty"`
}

// CreateServiceResponse defines the response fields for the created service
type CreateServiceResponse struct {
	AccountSid                    string     `json:"account_sid"`
	AclEnabled                    bool       `json:"acl_enabled"`
	DateCreated                   time.Time  `json:"date_created"`
	DateUpdated                   *time.Time `json:"date_updated,omitempty"`
	FriendlyName                  *string    `json:"friendly_name,omitempty"`
	ReachabilityDebouncingEnabled bool       `json:"reachability_debouncing_enabled"`
	ReachabilityDebouncingWindow  int        `json:"reachability_debouncing_window"`
	ReachabilityWebhooksEnabled   bool       `json:"reachability_webhooks_enabled"`
	Sid                           string     `json:"sid"`
	URL                           string     `json:"url"`
	UniqueName                    *string    `json:"unique_name,omitempty"`
	WebhookURL                    *string    `json:"webhook_url,omitempty"`
	WebhooksFromRestEnabled       bool       `json:"webhooks_from_rest_enabled"`
}

// Create creates a new service
// See https://www.twilio.com/docs/sync/api/service#create-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new service
// See https://www.twilio.com/docs/sync/api/service#create-a-service-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
