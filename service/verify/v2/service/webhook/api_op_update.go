// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateWebhookInput defines input fields for updating a webhook resource
type UpdateWebhookInput struct {
	EventTypes   *[]string `form:"EventTypes,omitempty"`
	FriendlyName *string   `form:"FriendlyName,omitempty"`
	Status       *string   `form:"Status,omitempty"`
	WebhookURL   *string   `form:"WebhookUrl,omitempty"`
}

// UpdateWebhookResponse defines the response fields for the updated webhook
type UpdateWebhookResponse struct {
	AccountSid    string     `json:"account_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	EventTypes    []string   `json:"event_types"`
	FriendlyName  string     `json:"friendly_name"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	Status        string     `json:"status"`
	URL           string     `json:"url"`
	WebhookMethod string     `json:"webhook_method"`
	WebhookURL    string     `json:"webhook_url"`
}

// Update modifies a webhook resource
// See https://www.twilio.com/docs/verify/api/webhooks#update-a-webhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a webhook resource
// See https://www.twilio.com/docs/verify/api/webhooks#update-a-webhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
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
