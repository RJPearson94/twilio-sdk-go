// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateWebhookInput defines the input fields for creating a new webhook resource
type CreateWebhookInput struct {
	ConfigurationFilters     *[]string `form:"Configuration.Filters,omitempty"`
	ConfigurationFlowSid     *string   `form:"Configuration.FlowSid,omitempty"`
	ConfigurationMethod      *string   `form:"Configuration.Method,omitempty"`
	ConfigurationReplayAfter *int      `form:"Configuration.ReplayAfter,omitempty"`
	ConfigurationTriggers    *[]string `form:"Configuration.Triggers,omitempty"`
	ConfigurationURL         *string   `form:"Configuration.Url,omitempty"`
	Target                   string    `validate:"required" form:"Target"`
}

type CreateWebhookResponseConfiguration struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
}

// CreateWebhookResponse defines the response fields for the created webhook
type CreateWebhookResponse struct {
	AccountSid      string                             `json:"account_sid"`
	ChatServiceSid  string                             `json:"chat_service_sid"`
	Configuration   CreateWebhookResponseConfiguration `json:"configuration"`
	ConversationSid string                             `json:"conversation_sid"`
	DateCreated     time.Time                          `json:"date_created"`
	DateUpdated     *time.Time                         `json:"date_updated,omitempty"`
	Sid             string                             `json:"sid"`
	Target          string                             `json:"target"`
	URL             string                             `json:"url"`
}

// Create creates a new webhook
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#create-a-conversationscopedwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new webhook
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#create-a-conversationscopedwebhook-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Conversations/{conversationSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
		},
	}

	if input == nil {
		input = &CreateWebhookInput{}
	}

	response := &CreateWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
