// Package web_channels contains auto-generated files. DO NOT MODIFY
package web_channels

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateWebChannelInput defines the input fields for creating a new web channel resource
type CreateWebChannelInput struct {
	ChatFriendlyName     string  `validate:"required" form:"ChatFriendlyName"`
	ChatUniqueName       *string `form:"ChatUniqueName,omitempty"`
	CustomerFriendlyName string  `validate:"required" form:"CustomerFriendlyName"`
	FlexFlowSid          string  `validate:"required" form:"FlexFlowSid"`
	Identity             string  `validate:"required" form:"Identity"`
	PreEngagementData    *string `form:"PreEngagementData,omitempty"`
}

// CreateWebChannelResponse defines the response fields for the created web channel
type CreateWebChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Create creates a new web channel
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateWebChannelInput) (*CreateWebChannelResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new web channel
func (c Client) CreateWithContext(context context.Context, input *CreateWebChannelInput) (*CreateWebChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/WebChannels",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateWebChannelInput{}
	}

	response := &CreateWebChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
