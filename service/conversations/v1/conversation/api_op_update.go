// Package conversation contains auto-generated files. DO NOT MODIFY
package conversation

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type UpdateConversationTimersInput struct {
	Closed   *string `form:"Closed,omitempty"`
	Inactive *string `form:"Inactive,omitempty"`
}

// UpdateConversationInput defines input fields for updating a conversation resource
type UpdateConversationInput struct {
	Attributes          *string                        `form:"Attributes,omitempty"`
	DateCreated         *utils.RFC2822Time             `form:"DateCreated,omitempty"`
	DateUpdated         *utils.RFC2822Time             `form:"DateUpdated,omitempty"`
	FriendlyName        *string                        `form:"FriendlyName,omitempty"`
	MessagingServiceSid *string                        `form:"MessagingServiceSid,omitempty"`
	State               *string                        `form:"State,omitempty"`
	Timers              *UpdateConversationTimersInput `form:"Timers,omitempty"`
	UniqueName          *string                        `form:"UniqueName,omitempty"`
}

type UpdateConversationTimersResponse struct {
	DateClosed   *time.Time `json:"date_closed,omitempty"`
	DateInactive *time.Time `json:"date_inactive,omitempty"`
}

// UpdateConversationResponse defines the response fields for the updated conversation
type UpdateConversationResponse struct {
	AccountSid          string                           `json:"account_sid"`
	Attributes          string                           `json:"attributes"`
	ChatServiceSid      *string                          `json:"chat_service_sid,omitempty"`
	DateCreated         time.Time                        `json:"date_created"`
	DateUpdated         *time.Time                       `json:"date_updated,omitempty"`
	FriendlyName        *string                          `json:"friendly_name,omitempty"`
	MessagingServiceSid *string                          `json:"messaging_service_sid,omitempty"`
	Sid                 string                           `json:"sid"`
	State               string                           `json:"state"`
	Timers              UpdateConversationTimersResponse `json:"timers"`
	URL                 string                           `json:"url"`
	UniqueName          *string                          `json:"unique_name,omitempty"`
}

// Update modifies a conversation resource
// See https://www.twilio.com/docs/conversations/api/conversation-resource#update-conversation for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateConversationInput) (*UpdateConversationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a conversation resource
// See https://www.twilio.com/docs/conversations/api/conversation-resource#update-conversation for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateConversationInput) (*UpdateConversationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateConversationInput{}
	}

	response := &UpdateConversationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
