// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateWebhookInput defines the input fields for updating a webhook
type UpdateWebhookInput struct {
	Events        *string `form:"Events,omitempty"`
	UniqueName    *string `form:"UniqueName,omitempty"`
	WebhookMethod *string `form:"WebhookMethod,omitempty"`
	WebhookURL    *string `form:"WebhookUrl,omitempty"`
}

// UpdateWebhookResponse defines the response fields for the updated webhook
type UpdateWebhookResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Events        string     `json:"events"`
	Sid           string     `json:"sid"`
	URL           string     `json:"url"`
	UniqueName    string     `json:"unique_name"`
	WebhookMethod string     `json:"webhook_method"`
	WebhookURL    string     `json:"webhook_url"`
}

// Update modifies an webhook resource
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#update-a-webhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an webhook resource
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#update-a-webhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
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
