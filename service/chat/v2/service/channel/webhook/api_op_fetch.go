// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchChannelWebhookConfigurationResponse struct {
	Filters    *[]string `json:"filters,omitempty"`
	FlowSid    *string   `json:"flow_sid,omitempty"`
	Method     *string   `json:"method,omitempty"`
	RetryCount *int      `json:"retry_count,omitempty"`
	Triggers   *[]string `json:"triggers,omitempty"`
	URL        *string   `json:"url,omitempty"`
}

// FetchChannelWebhookResponse defines the response fields for the retrieved webhook
type FetchChannelWebhookResponse struct {
	AccountSid    string                                   `json:"account_sid"`
	ChannelSid    string                                   `json:"channel_sid"`
	Configuration FetchChannelWebhookConfigurationResponse `json:"configuration"`
	DateCreated   time.Time                                `json:"date_created"`
	DateUpdated   *time.Time                               `json:"date_updated,omitempty"`
	ServiceSid    string                                   `json:"service_sid"`
	Sid           string                                   `json:"sid"`
	Type          string                                   `json:"type"`
	URL           string                                   `json:"url"`
}

// Fetch retrieves a webhook resource
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#fetch-a-channelwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelWebhookResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a webhook resource
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#fetch-a-channelwebhook-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChannelWebhookResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Webhooks/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	response := &FetchChannelWebhookResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
