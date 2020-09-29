// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateUserChannelInput defines input fields for updating a user channel resource
type UpdateUserChannelInput struct {
	LastConsumedMessageIndex *int       `form:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp *time.Time `form:"LastConsumptionTimestamp,omitempty"`
	NotificationLevel        *string    `form:"NotificationLevel,omitempty"`
}

// UpdateUserChannelResponse defines the response fields for the updated user channel
type UpdateUserChannelResponse struct {
	AccountSid               string  `json:"account_sid"`
	ChannelSid               string  `json:"channel_sid"`
	LastConsumedMessageIndex *int    `json:"last_consumed_message_index,omitempty"`
	MemberSid                string  `json:"member_sid"`
	NotificationLevel        *string `json:"notification_level,omitempty"`
	ServiceSid               string  `json:"service_sid"`
	Status                   string  `json:"status"`
	URL                      string  `json:"url"`
	UnreadMessagesCount      *int    `json:"unread_messages_count,omitempty"`
	UserSid                  string  `json:"user_sid"`
}

// Update modifies a user channel resource
// See https://www.twilio.com/docs/chat/rest/user-channel-resource#set-the-notificationlevel for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateUserChannelInput) (*UpdateUserChannelResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a user channel resource
// See https://www.twilio.com/docs/chat/rest/user-channel-resource#set-the-notificationlevel for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateUserChannelInput) (*UpdateUserChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Users/{userSid}/Channels/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateUserChannelInput{}
	}

	response := &UpdateUserChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
