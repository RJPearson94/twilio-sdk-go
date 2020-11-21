// Package message contains auto-generated files. DO NOT MODIFY
package message

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateChannelMessageInput defines input fields for updating a message resource
type UpdateChannelMessageInput struct {
	Attributes    *string    `form:"Attributes,omitempty"`
	Body          *string    `form:"Body,omitempty"`
	DateCreated   *time.Time `form:"DateCreated,omitempty"`
	DateUpdated   *time.Time `form:"DateUpdated,omitempty"`
	From          *string    `form:"From,omitempty"`
	LastUpdatedBy *string    `form:"LastUpdatedBy,omitempty"`
}

type UpdateChannelMessageMediaResponse struct {
	ContentType string `json:"content_type"`
	FileName    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
}

// UpdateChannelMessageResponse defines the response fields for the updated message
type UpdateChannelMessageResponse struct {
	AccountSid    string                             `json:"account_sid"`
	Attributes    *string                            `json:"attributes,omitempty"`
	Body          *string                            `json:"body,omitempty"`
	ChannelSid    string                             `json:"channel_sid"`
	DateCreated   time.Time                          `json:"date_created"`
	DateUpdated   *time.Time                         `json:"date_updated,omitempty"`
	From          *string                            `json:"from,omitempty"`
	Index         *int                               `json:"index,omitempty"`
	LastUpdatedBy *string                            `json:"last_updated_by,omitempty"`
	Media         *UpdateChannelMessageMediaResponse `json:"media,omitempty"`
	ServiceSid    string                             `json:"service_sid"`
	Sid           string                             `json:"sid"`
	To            *string                            `json:"to,omitempty"`
	Type          *string                            `json:"type,omitempty"`
	URL           string                             `json:"url"`
	WasEdited     *bool                              `json:"was_edited,omitempty"`
}

// Update modifies a message resource
// See https://www.twilio.com/docs/chat/rest/message-resource#update-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateChannelMessageInput) (*UpdateChannelMessageResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a message resource
// See https://www.twilio.com/docs/chat/rest/message-resource#update-a-message-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateChannelMessageInput) (*UpdateChannelMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Messages/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateChannelMessageInput{}
	}

	response := &UpdateChannelMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
