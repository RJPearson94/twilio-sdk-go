// Package notification contains auto-generated files. DO NOT MODIFY
package notification

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateNotificationConversationActionInput struct {
	Enabled  *bool   `form:"Enabled,omitempty"`
	Sound    *string `form:"Sound,omitempty"`
	Template *string `form:"Template,omitempty"`
}

type UpdateNotificationNewMessageInput struct {
	BadgeCountEnabled *bool   `form:"BadgeCountEnabled,omitempty"`
	Enabled           *bool   `form:"Enabled,omitempty"`
	Sound             *string `form:"Sound,omitempty"`
	Template          *string `form:"Template,omitempty"`
}

// UpdateNotificationInput defines input fields for updating a service notification resource
type UpdateNotificationInput struct {
	AddedToConversation     *UpdateNotificationConversationActionInput `form:"AddedToConversation,omitempty"`
	LogEnabled              *bool                                      `form:"LogEnabled,omitempty"`
	NewMessage              *UpdateNotificationNewMessageInput         `form:"NewMessage,omitempty"`
	RemovedFromConversation *UpdateNotificationConversationActionInput `form:"RemovedFromConversation,omitempty"`
}

type UpdateNotificationResponseConversationAction struct {
	Enabled  bool    `json:"enabled"`
	Sound    *string `json:"sound,omitempty"`
	Template *string `json:"template,omitempty"`
}

type UpdateNotificationResponseNewMessage struct {
	BadgeCountEnabled *bool   `json:"badge_count_enabled,omitempty"`
	Enabled           bool    `json:"enabled"`
	Sound             *string `json:"sound,omitempty"`
	Template          *string `json:"template,omitempty"`
}

// UpdateNotificationResponse defines the response fields for the updated service notification
type UpdateNotificationResponse struct {
	AccountSid              string                                       `json:"account_sid"`
	AddedToConversation     UpdateNotificationResponseConversationAction `json:"added_to_conversation"`
	ChatServiceSid          string                                       `json:"chat_service_sid"`
	LogEnabled              bool                                         `json:"log_enabled"`
	NewMessage              UpdateNotificationResponseNewMessage         `json:"new_message"`
	RemovedFromConversation UpdateNotificationResponseConversationAction `json:"removed_from_conversation"`
	URL                     string                                       `json:"url"`
}

// Update modifies a service notification resource
// See https://www.twilio.com/docs/conversations/api/service-notification-resource#update-a-servicenotification-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateNotificationInput) (*UpdateNotificationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service notification resource
// See https://www.twilio.com/docs/conversations/api/service-notification-resource#update-a-servicenotification-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateNotificationInput) (*UpdateNotificationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Configuration/Notifications",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &UpdateNotificationInput{}
	}

	response := &UpdateNotificationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
