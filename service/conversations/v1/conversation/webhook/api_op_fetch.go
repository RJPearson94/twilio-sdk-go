// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchWebhookResponseConfiguration struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
}

// FetchWebhookResponse defines the response fields for the retrieved webhook
type FetchWebhookResponse struct {
	AccountSid      string                            `json:"account_sid"`
	Configuration   FetchWebhookResponseConfiguration `json:"configuration"`
	ConversationSid string                            `json:"conversation_sid"`
	DateCreated     time.Time                         `json:"date_created"`
	DateUpdated     *time.Time                        `json:"date_updated,omitempty"`
	Sid             string                            `json:"sid"`
	Target          string                            `json:"target"`
	URL             string                            `json:"url"`
}

// Fetch retrieves an webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#fetch-a-conversationscopedwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWebhookResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#fetch-a-conversationscopedwebhook-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchWebhookResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/{conversationSid}/Webhooks/{sid}",
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	response := &FetchWebhookResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
