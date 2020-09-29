// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateServiceInput defines input fields for updating a service resource
type UpdateServiceInput struct {
	AclEnabled                    *bool   `form:"AclEnabled,omitempty"`
	FriendlyName                  *string `form:"FriendlyName,omitempty"`
	ReachabilityDebouncingEnabled *bool   `form:"ReachabilityDebouncingEnabled,omitempty"`
	ReachabilityDebouncingWindow  *int    `form:"ReachabilityDebouncingWindow,omitempty"`
	ReachabilityWebhooksEnabled   *bool   `form:"ReachabilityWebhooksEnabled,omitempty"`
	WebhookURL                    *string `form:"WebhookUrl,omitempty"`
	WebhooksFromRestEnabled       *bool   `form:"WebhooksFromRestEnabled,omitempty"`
}

// UpdateServiceResponse defines the response fields for the updated service
type UpdateServiceResponse struct {
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

// Update modifies a service resource
// See https://www.twilio.com/docs/sync/api/service#update-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service resource
// See https://www.twilio.com/docs/sync/api/service#update-a-service-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	response := &UpdateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
