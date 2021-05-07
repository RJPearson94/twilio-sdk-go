// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateServiceLimitsInput struct {
	ChannelMembers *int `form:"ChannelMembers,omitempty"`
	UserChannels   *int `form:"UserChannels,omitempty"`
}

type UpdateServiceMediaInput struct {
	CompatibilityMessage *string `form:"CompatibilityMessage,omitempty"`
}

type UpdateServiceNotificationsActionInput struct {
	Enabled  *bool   `form:"Enabled,omitempty"`
	Sound    *string `form:"Sound,omitempty"`
	Template *string `form:"Template,omitempty"`
}

type UpdateServiceNotificationsInput struct {
	AddedToChannel     *UpdateServiceNotificationsActionInput     `form:"AddedToChannel,omitempty"`
	InvitedToChannel   *UpdateServiceNotificationsActionInput     `form:"InvitedToChannel,omitempty"`
	LogEnabled         *bool                                      `form:"LogEnabled,omitempty"`
	NewMessage         *UpdateServiceNotificationsNewMessageInput `form:"NewMessage,omitempty"`
	RemovedFromChannel *UpdateServiceNotificationsActionInput     `form:"RemovedFromChannel,omitempty"`
}

type UpdateServiceNotificationsNewMessageInput struct {
	BadgeCountEnabled *bool   `form:"BadgeCountEnabled,omitempty"`
	Enabled           *bool   `form:"Enabled,omitempty"`
	Sound             *string `form:"Sound,omitempty"`
	Template          *string `form:"Template,omitempty"`
}

// UpdateServiceInput defines the input fields for updating a service resource
type UpdateServiceInput struct {
	ConsumptionReportInterval    *int                             `form:"ConsumptionReportInterval,omitempty"`
	DefaultChannelCreatorRoleSid *string                          `form:"DefaultChannelCreatorRoleSid,omitempty"`
	DefaultChannelRoleSid        *string                          `form:"DefaultChannelRoleSid,omitempty"`
	DefaultServiceRoleSid        *string                          `form:"DefaultServiceRoleSid,omitempty"`
	FriendlyName                 *string                          `form:"FriendlyName,omitempty"`
	Limits                       *UpdateServiceLimitsInput        `form:"Limits,omitempty"`
	Media                        *UpdateServiceMediaInput         `form:"Media,omitempty"`
	Notifications                *UpdateServiceNotificationsInput `form:"Notifications,omitempty"`
	PostWebhookRetryCount        *int                             `form:"PostWebhookRetryCount,omitempty"`
	PostWebhookURL               *string                          `form:"PostWebhookUrl,omitempty"`
	PreWebhookRetryCount         *int                             `form:"PreWebhookRetryCount,omitempty"`
	PreWebhookURL                *string                          `form:"PreWebhookUrl,omitempty"`
	ReachabilityEnabled          *bool                            `form:"ReachabilityEnabled,omitempty"`
	ReadStatusEnabled            *bool                            `form:"ReadStatusEnabled,omitempty"`
	TypingIndicatorTimeout       *int                             `form:"TypingIndicatorTimeout,omitempty"`
	WebhookFilters               *[]string                        `form:"WebhookFilters,omitempty"`
	WebhookMethod                *string                          `form:"WebhookMethod,omitempty"`
}

type UpdateServiceLimitsResponse struct {
	ChannelMembers int `json:"channel_members"`
	UserChannels   int `json:"user_channels"`
}

type UpdateServiceMediaResponse struct {
	CompatibilityMessage string `json:"compatibility_message"`
	SizeLimitMB          int    `json:"size_limit_mb"`
}

type UpdateServiceNotificationsActionResponse struct {
	Enabled  bool    `json:"enabled"`
	Sound    *string `json:"sound,omitempty"`
	Template *string `json:"template,omitempty"`
}

type UpdateServiceNotificationsNewMessageResponse struct {
	BadgeCountEnabled *bool   `json:"badge_count_enabled,omitempty"`
	Enabled           bool    `json:"enabled"`
	Sound             *string `json:"sound,omitempty"`
	Template          *string `json:"template,omitempty"`
}

type UpdateServiceNotificationsResponse struct {
	AddedToChannel     UpdateServiceNotificationsActionResponse     `json:"added_to_channel"`
	InvitedToChannel   UpdateServiceNotificationsActionResponse     `json:"invited_to_channel"`
	LogEnabled         bool                                         `json:"log_enabled"`
	NewMessage         UpdateServiceNotificationsNewMessageResponse `json:"new_message"`
	RemovedFromChannel UpdateServiceNotificationsActionResponse     `json:"removed_from_channel"`
}

// UpdateServiceResponse defines the response fields for the updated service
type UpdateServiceResponse struct {
	AccountSid                   string                             `json:"account_sid"`
	ConsumptionReportInterval    int                                `json:"consumption_report_interval"`
	DateCreated                  time.Time                          `json:"date_created"`
	DateUpdated                  *time.Time                         `json:"date_updated,omitempty"`
	DefaultChannelCreatorRoleSid string                             `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                             `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                             `json:"default_service_role_sid"`
	FriendlyName                 string                             `json:"friendly_name"`
	Limits                       UpdateServiceLimitsResponse        `json:"limits"`
	Media                        UpdateServiceMediaResponse         `json:"media"`
	Notifications                UpdateServiceNotificationsResponse `json:"notifications"`
	PostWebhookRetryCount        *int                               `json:"post_webhook_retry_count,omitempty"`
	PostWebhookURL               *string                            `json:"post_webhook_url,omitempty"`
	PreWebhookRetryCount         *int                               `json:"pre_webhook_retry_count,omitempty"`
	PreWebhookURL                *string                            `json:"pre_webhook_url,omitempty"`
	ReachabilityEnabled          bool                               `json:"reachability_enabled"`
	ReadStatusEnabled            bool                               `json:"read_status_enabled"`
	Sid                          string                             `json:"sid"`
	TypingIndicatorTimeout       int                                `json:"typing_indicator_timeout"`
	URL                          string                             `json:"url"`
	WebhookFilters               *[]string                          `json:"webhook_filters,omitempty"`
	WebhookMethod                *string                            `json:"webhook_method,omitempty"`
}

// Update modifies a service resource
// See https://www.twilio.com/docs/chat/rest/service-resource#update-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service resource
// See https://www.twilio.com/docs/chat/rest/service-resource#update-a-service-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	response := &UpdateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
