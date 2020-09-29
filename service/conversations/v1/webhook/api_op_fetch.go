// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchWebhookResponse defines the response fields for the retrieved webhook
type FetchWebhookResponse struct {
	AccountSid     string   `json:"account_sid"`
	Filters        []string `json:"filters"`
	Method         string   `json:"method"`
	PostWebhookURL *string  `json:"post_webhook_url,omitempty"`
	PreWebhookURL  *string  `json:"pre_webhook_url,omitempty"`
	Target         string   `json:"target"`
	URL            string   `json:"url"`
}

// Fetch retrieves a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-webhook-resource#fetch-a-conversationwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWebhookResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-webhook-resource#fetch-a-conversationwebhook-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchWebhookResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/Webhooks",
	}

	response := &FetchWebhookResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
