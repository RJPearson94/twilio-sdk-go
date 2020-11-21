// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateChannelMessageInput defines the input fields for creating a new message resource
type CreateChannelMessageInput struct {
	Attributes    *string    `form:"Attributes,omitempty"`
	Body          *string    `form:"Body,omitempty"`
	DateCreated   *time.Time `form:"DateCreated,omitempty"`
	DateUpdated   *time.Time `form:"DateUpdated,omitempty"`
	From          *string    `form:"From,omitempty"`
	LastUpdatedBy *string    `form:"LastUpdatedBy,omitempty"`
	MediaSid      *string    `form:"MediaSid,omitempty"`
}

type CreateChannelMessageMediaResponse struct {
	ContentType string `json:"content_type"`
	FileName    string `json:"filename"`
	Sid         string `json:"sid"`
	Size        int    `json:"size"`
}

// CreateChannelMessageResponse defines the response fields for the created message
type CreateChannelMessageResponse struct {
	AccountSid    string                             `json:"account_sid"`
	Attributes    *string                            `json:"attributes,omitempty"`
	Body          *string                            `json:"body,omitempty"`
	ChannelSid    string                             `json:"channel_sid"`
	DateCreated   time.Time                          `json:"date_created"`
	DateUpdated   *time.Time                         `json:"date_updated,omitempty"`
	From          *string                            `json:"from,omitempty"`
	Index         *int                               `json:"index,omitempty"`
	LastUpdatedBy *string                            `json:"last_updated_by,omitempty"`
	Media         *CreateChannelMessageMediaResponse `json:"media,omitempty"`
	ServiceSid    string                             `json:"service_sid"`
	Sid           string                             `json:"sid"`
	To            *string                            `json:"to,omitempty"`
	Type          *string                            `json:"type,omitempty"`
	URL           string                             `json:"url"`
	WasEdited     *bool                              `json:"was_edited,omitempty"`
}

// Create creates a new message
// See https://www.twilio.com/docs/chat/rest/message-resource#create-a-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChannelMessageInput) (*CreateChannelMessageResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new message
// See https://www.twilio.com/docs/chat/rest/message-resource#create-a-message-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateChannelMessageInput) (*CreateChannelMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Messages",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
	}

	if input == nil {
		input = &CreateChannelMessageInput{}
	}

	response := &CreateChannelMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
