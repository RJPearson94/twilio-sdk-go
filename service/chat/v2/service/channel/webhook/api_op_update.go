// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateChannelWebhookConfigurationInput struct {
	Filters    *[]string `form:"Filters,omitempty"`
	FlowSid    *string   `form:"FlowSid,omitempty"`
	Method     *string   `form:"Method,omitempty"`
	RetryCount *int      `form:"RetryCount,omitempty"`
	Triggers   *[]string `form:"Triggers,omitempty"`
	URL        *string   `form:"Url,omitempty"`
}

// UpdateChannelWebhookInput defines input fields for updating a webhook resource
type UpdateChannelWebhookInput struct {
	Configuration *UpdateChannelWebhookConfigurationInput `form:"Configuration,omitempty"`
}

type UpdateChannelWebhookConfigurationResponse struct {
	Filters    *[]string `json:"filters,omitempty"`
	FlowSid    *string   `json:"flow_sid,omitempty"`
	Method     *string   `json:"method,omitempty"`
	RetryCount *int      `json:"retry_count,omitempty"`
	Triggers   *[]string `json:"triggers,omitempty"`
	URL        *string   `json:"url,omitempty"`
}

// UpdateChannelWebhookResponse defines the response fields for the updated webhook
type UpdateChannelWebhookResponse struct {
	AccountSid    string                                    `json:"account_sid"`
	ChannelSid    string                                    `json:"channel_sid"`
	Configuration UpdateChannelWebhookConfigurationResponse `json:"configuration"`
	DateCreated   time.Time                                 `json:"date_created"`
	DateUpdated   *time.Time                                `json:"date_updated,omitempty"`
	ServiceSid    string                                    `json:"service_sid"`
	Sid           string                                    `json:"sid"`
	Type          string                                    `json:"type"`
	URL           string                                    `json:"url"`
}

// Update modifies a webhook resource
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#update-a-channelwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateChannelWebhookInput) (*UpdateChannelWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a webhook resource
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#update-a-channelwebhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateChannelWebhookInput) (*UpdateChannelWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateChannelWebhookInput{}
	}

	response := &UpdateChannelWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
