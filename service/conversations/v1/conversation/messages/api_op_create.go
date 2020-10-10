// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateMessageInput defines the input fields for creating a new message resource
type CreateMessageInput struct {
	Attributes            *string            `form:"Attributes,omitempty"`
	Author                *string            `form:"Author,omitempty"`
	Body                  *string            `form:"Body,omitempty"`
	DateCreated           *utils.RFC2822Time `form:"DateCreated,omitempty"`
	DateUpdated           *utils.RFC2822Time `form:"DateUpdated,omitempty"`
	MediaSid              *string            `form:"MediaSid,omitempty"`
	XTwilioWebhookEnabled *string            `form:"X-Twilio-Webhook-Enabled,omitempty"`
}

type CreateMessageResponseDelivery struct {
	Delivered   string `json:"delivered"`
	Failed      string `json:"failed"`
	Read        string `json:"read"`
	Sent        string `json:"sent"`
	Total       int    `json:"total"`
	Undelivered string `json:"undelivered"`
}

type CreateMessageResponseMedia struct {
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
}

// CreateMessageResponse defines the response fields for the created message
type CreateMessageResponse struct {
	AccountSid      string                         `json:"account_sid"`
	Attributes      string                         `json:"attributes"`
	Author          string                         `json:"author"`
	Body            *string                        `json:"body,omitempty"`
	ConversationSid string                         `json:"conversation_sid"`
	DateCreated     time.Time                      `json:"date_created"`
	DateUpdated     *time.Time                     `json:"date_updated,omitempty"`
	Delivery        *CreateMessageResponseDelivery `json:"delivery,omitempty"`
	Index           int                            `json:"index"`
	Media           *[]CreateMessageResponseMedia  `json:"media,omitempty"`
	ParticipantSid  *string                        `json:"participant_sid,omitempty"`
	Sid             string                         `json:"sid"`
	URL             string                         `json:"url"`
}

// Create creates a new message
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#create-a-conversationmessage-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateMessageInput) (*CreateMessageResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new message
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#create-a-conversationmessage-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateMessageInput) (*CreateMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{conversationSid}/Messages",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
		},
	}

	if input == nil {
		input = &CreateMessageInput{}
	}

	response := &CreateMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
