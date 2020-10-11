// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateConversationWebhookInput defines input fields for updating an webhook resource
type UpdateConversationWebhookInput struct {
	ConfigurationFilters  *[]string `form:"Configuration.Filters,omitempty"`
	ConfigurationFlowSid  *string   `form:"Configuration.FlowSid,omitempty"`
	ConfigurationMethod   *string   `form:"Configuration.Method,omitempty"`
	ConfigurationTriggers *[]string `form:"Configuration.Triggers,omitempty"`
	ConfigurationURL      *string   `form:"Configuration.Url,omitempty"`
}

type UpdateConversationWebhookResponseConfiguration struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
}

// UpdateConversationWebhookResponse defines the response fields for the updated webhook
type UpdateConversationWebhookResponse struct {
	AccountSid      string                                         `json:"account_sid"`
	ChatServiceSid  string                                         `json:"chat_service_sid"`
	Configuration   UpdateConversationWebhookResponseConfiguration `json:"configuration"`
	ConversationSid string                                         `json:"conversation_sid"`
	DateCreated     time.Time                                      `json:"date_created"`
	DateUpdated     *time.Time                                     `json:"date_updated,omitempty"`
	Sid             string                                         `json:"sid"`
	Target          string                                         `json:"target"`
	URL             string                                         `json:"url"`
}

// Update modifies a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#update-a-conversationscopedwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateConversationWebhookInput) (*UpdateConversationWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#update-a-conversationscopedwebhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateConversationWebhookInput) (*UpdateConversationWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Conversations/{conversationSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	if input == nil {
		input = &UpdateConversationWebhookInput{}
	}

	response := &UpdateConversationWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}