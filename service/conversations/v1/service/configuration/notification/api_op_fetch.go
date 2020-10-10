// Package notification contains auto-generated files. DO NOT MODIFY
package notification

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchNotificationResponseConversationAction struct {
	Enabled  bool    `json:"enabled"`
	Sound    *string `json:"sound,omitempty"`
	Template *string `json:"template,omitempty"`
}

type FetchNotificationResponseNewMessage struct {
	BadgeCountEnabled *bool   `json:"badge_count_enabled,omitempty"`
	Enabled           bool    `json:"enabled"`
	Sound             *string `json:"sound,omitempty"`
	Template          *string `json:"template,omitempty"`
}

// FetchNotificationResponse defines the response fields for the retrieved service notification
type FetchNotificationResponse struct {
	AccountSid              string                                      `json:"account_sid"`
	AddedToConversation     FetchNotificationResponseConversationAction `json:"added_to_conversation"`
	ChatServiceSid          string                                      `json:"chat_service_sid"`
	LogEnabled              bool                                        `json:"log_enabled"`
	NewMessage              FetchNotificationResponseNewMessage         `json:"new_message"`
	RemovedFromConversation FetchNotificationResponseConversationAction `json:"removed_from_conversation"`
	URL                     string                                      `json:"url"`
}

// Fetch retrieves service notification resource
// See https://www.twilio.com/docs/conversations/api/service-notification-resource#fetch-a-servicenotification-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchNotificationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves service notification resource
// See https://www.twilio.com/docs/conversations/api/service-notification-resource#fetch-a-servicenotification-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchNotificationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Configuration/Notifications",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	response := &FetchNotificationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
