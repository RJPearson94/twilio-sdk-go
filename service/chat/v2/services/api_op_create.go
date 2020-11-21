// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateServiceInput defines the input fields for creating a new service resource
type CreateServiceInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

type CreateServiceLimitsResponse struct {
	ChannelMembers int `json:"channel_members"`
	UserChannels   int `json:"user_channels"`
}

type CreateServiceMediaResponse struct {
	CompatibilityMessage string `json:"compatibility_message"`
	SizeLimitMB          int    `json:"size_limit_mb"`
}

type CreateServiceNotificationsActionResponse struct {
	Enabled  bool    `json:"enabled"`
	Sound    *string `json:"sound,omitempty"`
	Template *string `json:"template,omitempty"`
}

type CreateServiceNotificationsNewMessageResponse struct {
	BadgeCountEnabled *bool   `json:"badge_count_enabled,omitempty"`
	Enabled           bool    `json:"enabled"`
	Sound             *string `json:"sound,omitempty"`
	Template          *string `json:"template,omitempty"`
}

type CreateServiceNotificationsResponse struct {
	AddedToChannel     CreateServiceNotificationsActionResponse     `json:"added_to_channel"`
	InvitedToChannel   CreateServiceNotificationsActionResponse     `json:"invited_to_channel"`
	LogEnabled         bool                                         `json:"log_enabled"`
	NewMessage         CreateServiceNotificationsNewMessageResponse `json:"new_message"`
	RemovedFromChannel CreateServiceNotificationsActionResponse     `json:"removed_from_channel"`
}

// CreateServiceResponse defines the response fields for the created service
type CreateServiceResponse struct {
	AccountSid                   string                             `json:"account_sid"`
	ConsumptionReportInterval    int                                `json:"consumption_report_interval"`
	DateCreated                  time.Time                          `json:"date_created"`
	DateUpdated                  *time.Time                         `json:"date_updated,omitempty"`
	DefaultChannelCreatorRoleSid string                             `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                             `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                             `json:"default_service_role_sid"`
	FriendlyName                 string                             `json:"friendly_name"`
	Limits                       CreateServiceLimitsResponse        `json:"limits"`
	Media                        CreateServiceMediaResponse         `json:"media"`
	Notifications                CreateServiceNotificationsResponse `json:"notifications"`
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

// Create creates a new service
// See https://www.twilio.com/docs/chat/rest/service-resource#create-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new service
// See https://www.twilio.com/docs/chat/rest/service-resource#create-a-service-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
