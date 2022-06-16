// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchWebhookResponse defines the response fields for the retrieved webhook
type FetchWebhookResponse struct {
	AccountSid    string     `json:"account_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	EventTypes    []string   `json:"event_types"`
	FriendlyName  string     `json:"friendly_name"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	Status        string     `json:"status"`
	URL           string     `json:"url"`
	Version       string     `json:"version"`
	WebhookMethod string     `json:"webhook_method"`
	WebhookURL    string     `json:"webhook_url"`
}

// Fetch retrieves a webhook resource
// See https://www.twilio.com/docs/verify/api/webhooks#fetch-a-webhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Fetch() (*FetchWebhookResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a webhook resource
// See https://www.twilio.com/docs/verify/api/webhooks#fetch-a-webhook-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) FetchWithContext(context context.Context) (*FetchWebhookResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Webhooks/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchWebhookResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
