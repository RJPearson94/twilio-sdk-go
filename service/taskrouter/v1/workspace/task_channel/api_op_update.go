// Package task_channel contains auto-generated files. DO NOT MODIFY
package task_channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateTaskChannelInput defines input fields for updating a task channel resource
type UpdateTaskChannelInput struct {
	ChannelOptimizedRouting *bool   `form:"ChannelOptimizedRouting,omitempty"`
	FriendlyName            *string `form:"FriendlyName,omitempty"`
}

// UpdateTaskChannelResponse defines the response fields for the updated task channel
type UpdateTaskChannelResponse struct {
	AccountSid              string     `json:"account_sid"`
	ChannelOptimizedRouting *bool      `json:"channel_optimized_routing,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// Update modifies a task channel resource
// See https://www.twilio.com/docs/taskrouter/api/task-channel#update-a-taskchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateTaskChannelInput) (*UpdateTaskChannelResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a task channel resource
// See https://www.twilio.com/docs/taskrouter/api/task-channel#update-a-taskchannel-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateTaskChannelInput) (*UpdateTaskChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/TaskChannels/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateTaskChannelInput{}
	}

	response := &UpdateTaskChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
