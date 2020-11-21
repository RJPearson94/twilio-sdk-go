// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchServiceLimitsResponse struct {
	ChannelMembers int `json:"channel_members"`
	UserChannels   int `json:"user_channels"`
}

type FetchServiceMediaResponse struct {
	CompatibilityMessage string `json:"compatibility_message"`
	SizeLimitMB          int    `json:"size_limit_mb"`
}

type FetchServiceNotificationsActionResponse struct {
	Enabled  bool    `json:"enabled"`
	Sound    *string `json:"sound,omitempty"`
	Template *string `json:"template,omitempty"`
}

type FetchServiceNotificationsNewMessageResponse struct {
	BadgeCountEnabled *bool   `json:"badge_count_enabled,omitempty"`
	Enabled           bool    `json:"enabled"`
	Sound             *string `json:"sound,omitempty"`
	Template          *string `json:"template,omitempty"`
}

type FetchServiceNotificationsResponse struct {
	AddedToChannel     FetchServiceNotificationsActionResponse     `json:"added_to_channel"`
	InvitedToChannel   FetchServiceNotificationsActionResponse     `json:"invited_to_channel"`
	LogEnabled         bool                                        `json:"log_enabled"`
	NewMessage         FetchServiceNotificationsNewMessageResponse `json:"new_message"`
	RemovedFromChannel FetchServiceNotificationsActionResponse     `json:"removed_from_channel"`
}

// FetchServiceResponse defines the response fields for the retrieved service
type FetchServiceResponse struct {
	AccountSid                   string                            `json:"account_sid"`
	ConsumptionReportInterval    int                               `json:"consumption_report_interval"`
	DateCreated                  time.Time                         `json:"date_created"`
	DateUpdated                  *time.Time                        `json:"date_updated,omitempty"`
	DefaultChannelCreatorRoleSid string                            `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                            `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                            `json:"default_service_role_sid"`
	FriendlyName                 string                            `json:"friendly_name"`
	Limits                       FetchServiceLimitsResponse        `json:"limits"`
	Media                        FetchServiceMediaResponse         `json:"media"`
	Notifications                FetchServiceNotificationsResponse `json:"notifications"`
	PostWebhookRetryCount        *int                              `json:"post_webhook_retry_count,omitempty"`
	PostWebhookURL               *string                           `json:"post_webhook_url,omitempty"`
	PreWebhookRetryCount         *int                              `json:"pre_webhook_retry_count,omitempty"`
	PreWebhookURL                *string                           `json:"pre_webhook_url,omitempty"`
	ReachabilityEnabled          bool                              `json:"reachability_enabled"`
	ReadStatusEnabled            bool                              `json:"read_status_enabled"`
	Sid                          string                            `json:"sid"`
	TypingIndicatorTimeout       int                               `json:"typing_indicator_timeout"`
	URL                          string                            `json:"url"`
	WebhookFilters               *[]string                         `json:"webhook_filters,omitempty"`
	WebhookMethod                *string                           `json:"webhook_method,omitempty"`
}

// Fetch retrieves a service resource
// See https://www.twilio.com/docs/chat/rest/service-resource#fetch-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a service resource
// See https://www.twilio.com/docs/chat/rest/service-resource#fetch-a-service-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
