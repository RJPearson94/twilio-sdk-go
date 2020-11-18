// Package conversations contains auto-generated files. DO NOT MODIFY
package conversations

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type CreateConversationTimersInput struct {
	Closed   *string `form:"Closed,omitempty"`
	Inactive *string `form:"Inactive,omitempty"`
}

// CreateConversationInput defines the input fields for creating a new conversation resource
type CreateConversationInput struct {
	Attributes          *string                        `form:"Attributes,omitempty"`
	DateCreated         *utils.RFC2822Time             `form:"DateCreated,omitempty"`
	DateUpdated         *utils.RFC2822Time             `form:"DateUpdated,omitempty"`
	FriendlyName        *string                        `form:"FriendlyName,omitempty"`
	MessagingServiceSid *string                        `form:"MessagingServiceSid,omitempty"`
	State               *string                        `form:"State,omitempty"`
	Timers              *CreateConversationTimersInput `form:"Timers,omitempty"`
}

type CreateConversationTimersResponse struct {
	DateClosed   *time.Time `json:"date_closed,omitempty"`
	DateInactive *time.Time `json:"date_inactive,omitempty"`
}

// CreateConversationResponse defines the response fields for the created conversation
type CreateConversationResponse struct {
	AccountSid          string                           `json:"account_sid"`
	Attributes          string                           `json:"attributes"`
	ChatServiceSid      *string                          `json:"chat_service_sid,omitempty"`
	DateCreated         time.Time                        `json:"date_created"`
	DateUpdated         *time.Time                       `json:"date_updated,omitempty"`
	FriendlyName        *string                          `json:"friendly_name,omitempty"`
	MessagingServiceSid *string                          `json:"messaging_service_sid,omitempty"`
	Sid                 string                           `json:"sid"`
	State               string                           `json:"state"`
	Timers              CreateConversationTimersResponse `json:"timers"`
	URL                 string                           `json:"url"`
}

// Create creates a new conversation
// See https://www.twilio.com/docs/conversations/api/conversation-resource#create-a-conversation-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateConversationInput) (*CreateConversationResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new conversation
// See https://www.twilio.com/docs/conversations/api/conversation-resource#create-a-conversation-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateConversationInput) (*CreateConversationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Conversations",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateConversationInput{}
	}

	response := &CreateConversationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
