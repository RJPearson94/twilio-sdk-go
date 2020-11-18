// Package flex_flows contains auto-generated files. DO NOT MODIFY
package flex_flows

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateFlexFlowIntegrationInput struct {
	Channel           *string `form:"Channel,omitempty"`
	CreationOnMessage *bool   `form:"CreationOnMessage,omitempty"`
	FlowSid           *string `form:"FlowSid,omitempty"`
	Priority          *int    `form:"Priority,omitempty"`
	RetryCount        *int    `form:"RetryCount,omitempty"`
	Timeout           *int    `form:"Timeout,omitempty"`
	URL               *string `form:"Url,omitempty"`
	WorkflowSid       *string `form:"WorkflowSid,omitempty"`
	WorkspaceSid      *string `form:"WorkspaceSid,omitempty"`
}

// CreateFlexFlowInput defines the input fields for creating a new flex flow resource
type CreateFlexFlowInput struct {
	ChannelType     string                          `validate:"required" form:"ChannelType"`
	ChatServiceSid  string                          `validate:"required" form:"ChatServiceSid"`
	ContactIdentity *string                         `form:"ContactIdentity,omitempty"`
	Enabled         *bool                           `form:"Enabled,omitempty"`
	FriendlyName    string                          `validate:"required" form:"FriendlyName"`
	Integration     *CreateFlexFlowIntegrationInput `form:"Integration,omitempty"`
	IntegrationType *string                         `form:"IntegrationType,omitempty"`
	JanitorEnabled  *bool                           `form:"JanitorEnabled,omitempty"`
	LongLived       *bool                           `form:"LongLived,omitempty"`
}

type CreateFlexFlowIntegrationResponse struct {
	Channel           *string `json:"channel,omitempty"`
	CreationOnMessage *bool   `json:"creation_on_message,omitempty"`
	FlowSid           *string `json:"flow_sid,omitempty"`
	Priority          *int    `json:"priority,omitempty"`
	RetryCount        *int    `json:"retry_count,omitempty"`
	Timeout           *int    `json:"timeout,omitempty"`
	URL               *string `json:"url,omitempty"`
	WorkflowSid       *string `json:"workflow_sid,omitempty"`
	WorkspaceSid      *string `json:"workspace_sid,omitempty"`
}

// CreateFlexFlowResponse defines the response fields for the created flex flow
type CreateFlexFlowResponse struct {
	AccountSid      string                             `json:"account_sid"`
	ChannelType     string                             `json:"channel_type"`
	ChatServiceSid  string                             `json:"chat_service_sid"`
	ContactIdentity *string                            `json:"contact_identity,omitempty"`
	DateCreated     time.Time                          `json:"date_created"`
	DateUpdated     *time.Time                         `json:"date_updated,omitempty"`
	Enabled         bool                               `json:"enabled"`
	FriendlyName    string                             `json:"friendly_name"`
	Integration     *CreateFlexFlowIntegrationResponse `json:"integration,omitempty"`
	IntegrationType *string                            `json:"integration_type,omitempty"`
	JanitorEnabled  *bool                              `json:"janitor_enabled,omitempty"`
	LongLived       *bool                              `json:"long_lived,omitempty"`
	Sid             string                             `json:"sid"`
	URL             string                             `json:"url"`
}

// Create creates a new flex flow
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFlexFlowInput) (*CreateFlexFlowResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new flex flow
func (c Client) CreateWithContext(context context.Context, input *CreateFlexFlowInput) (*CreateFlexFlowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/FlexFlows",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateFlexFlowInput{}
	}

	response := &CreateFlexFlowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
