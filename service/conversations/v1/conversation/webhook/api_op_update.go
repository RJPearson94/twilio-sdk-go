// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateWebhookConfigurationInput struct {
	Filters  *[]string `form:"Filters,omitempty"`
	FlowSid  *string   `form:"FlowSid,omitempty"`
	Method   *string   `form:"Method,omitempty"`
	Triggers *[]string `form:"Triggers,omitempty"`
	URL      *string   `form:"Url,omitempty"`
}

// UpdateWebhookInput defines input fields for updating an webhook resource
type UpdateWebhookInput struct {
	Configuration *UpdateWebhookConfigurationInput `form:"Configuration,omitempty"`
}

type UpdateWebhookConfigurationResponse struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
}

// UpdateWebhookResponse defines the response fields for the updated webhook
type UpdateWebhookResponse struct {
	AccountSid      string                             `json:"account_sid"`
	Configuration   UpdateWebhookConfigurationResponse `json:"configuration"`
	ConversationSid string                             `json:"conversation_sid"`
	DateCreated     time.Time                          `json:"date_created"`
	DateUpdated     *time.Time                         `json:"date_updated,omitempty"`
	Sid             string                             `json:"sid"`
	Target          string                             `json:"target"`
	URL             string                             `json:"url"`
}

// Update modifies a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#update-a-conversationscopedwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#update-a-conversationscopedwebhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{conversationSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	if input == nil {
		input = &UpdateWebhookInput{}
	}

	response := &UpdateWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
