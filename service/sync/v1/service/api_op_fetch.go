// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchServiceResponse defines the response fields for the retrieved service
type FetchServiceResponse struct {
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

// Fetch retrieves an service resource
// See https://www.twilio.com/docs/sync/api/service#fetch-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an service resource
// See https://www.twilio.com/docs/sync/api/service#fetch-a-service-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
