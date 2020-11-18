// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateMessageInput defines input fields for updating a message resource
type UpdateMessageInput struct {
	Attributes  *string            `form:"Attributes,omitempty"`
	Author      *string            `form:"Author,omitempty"`
	Body        *string            `form:"Body,omitempty"`
	DateCreated *utils.RFC2822Time `form:"DateCreated,omitempty"`
	DateUpdated *utils.RFC2822Time `form:"DateUpdated,omitempty"`
}

type UpdateMessageDeliveryResponse struct {
	Delivered   string `json:"delivered"`
	Failed      string `json:"failed"`
	Read        string `json:"read"`
	Sent        string `json:"sent"`
	Total       int    `json:"total"`
	Undelivered string `json:"undelivered"`
}

type UpdateMessageMediaResponse struct {
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
}

// UpdateMessageResponse defines the response fields for the updated message
type UpdateMessageResponse struct {
	AccountSid      string                         `json:"account_sid"`
	Attributes      string                         `json:"attributes"`
	Author          string                         `json:"author"`
	Body            *string                        `json:"body,omitempty"`
	ChatServiceSid  string                         `json:"chat_service_sid"`
	ConversationSid string                         `json:"conversation_sid"`
	DateCreated     time.Time                      `json:"date_created"`
	DateUpdated     *time.Time                     `json:"date_updated,omitempty"`
	Delivery        *UpdateMessageDeliveryResponse `json:"delivery,omitempty"`
	Index           int                            `json:"index"`
	Media           *[]UpdateMessageMediaResponse  `json:"media,omitempty"`
	ParticipantSid  *string                        `json:"participant_sid,omitempty"`
	Sid             string                         `json:"sid"`
	URL             string                         `json:"url"`
}

// Update modifies a message resource
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#update-a-conversationmessage-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateMessageInput) (*UpdateMessageResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a message resource
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource#update-a-conversationmessage-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateMessageInput) (*UpdateMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Conversations/{conversationSid}/Messages/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	if input == nil {
		input = &UpdateMessageInput{}
	}

	response := &UpdateMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
