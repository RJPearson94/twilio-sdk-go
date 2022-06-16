// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateWebhookInput defines the input fields for creating a new webhook
type CreateWebhookInput struct {
	EventTypes   []string `validate:"required" form:"EventTypes"`
	FriendlyName string   `validate:"required" form:"FriendlyName"`
	Status       *string  `form:"Status,omitempty"`
	Version      *string  `form:"Version,omitempty"`
	WebhookURL   string   `validate:"required" form:"WebhookUrl"`
}

// CreateWebhookResponse defines the response fields for the created webhook
type CreateWebhookResponse struct {
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

// Create creates a new webhook
// See https://www.twilio.com/docs/verify/api/webhooks#create-a-webhook for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Create(input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new webhook
// See https://www.twilio.com/docs/verify/api/webhooks#create-a-webhook for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) CreateWithContext(context context.Context, input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
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
