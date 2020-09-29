// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateChannelInput defines the input fields for creating a new channel resource
type CreateChannelInput struct {
	ChatFriendlyName     string  `validate:"required" form:"ChatFriendlyName"`
	ChatUniqueName       *string `form:"ChatUniqueName,omitempty"`
	ChatUserFriendlyName string  `validate:"required" form:"ChatUserFriendlyName"`
	FlexFlowSid          string  `validate:"required" form:"FlexFlowSid"`
	Identity             string  `validate:"required" form:"Identity"`
	LongLived            *bool   `form:"LongLived,omitempty"`
	PreEngagementData    *string `form:"PreEngagementData,omitempty"`
	Target               *string `form:"Target,omitempty"`
	TaskAttributes       *string `form:"TaskAttributes,omitempty"`
	TaskSid              *string `form:"TaskSid,omitempty"`
}

// CreateChannelResponse defines the response fields for the created channel
type CreateChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	TaskSid     *string    `json:"task_sid,omitempty"`
	URL         string     `json:"url"`
	UserSid     string     `json:"user_sid"`
}

// Create creates a new channel
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChannelInput) (*CreateChannelResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new channel
func (c Client) CreateWithContext(context context.Context, input *CreateChannelInput) (*CreateChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Channels",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateChannelInput{}
	}

	response := &CreateChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
