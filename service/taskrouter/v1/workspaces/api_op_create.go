// Package workspaces contains auto-generated files. DO NOT MODIFY
package workspaces

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateWorkspaceInput defines the input fields for creating a new workspace resource
type CreateWorkspaceInput struct {
	EventCallbackURL     *string `form:"EventCallbackUrl,omitempty"`
	EventsFilter         *string `form:"EventsFilter,omitempty"`
	FriendlyName         string  `validate:"required" form:"FriendlyName"`
	MultiTaskEnabled     *bool   `form:"MultiTaskEnabled,omitempty"`
	PrioritizeQueueOrder *string `form:"PrioritizeQueueOrder,omitempty"`
	Template             *string `form:"Template,omitempty"`
}

// CreateWorkspaceResponse defines the response fields for the created workspace
type CreateWorkspaceResponse struct {
	AccountSid           string     `json:"account_sid"`
	DateCreated          time.Time  `json:"date_created"`
	DateUpdated          *time.Time `json:"date_updated,omitempty"`
	DefaultActivityName  string     `json:"default_activity_name"`
	DefaultActivitySid   string     `json:"default_activity_sid"`
	EventCallbackURL     *string    `json:"event_callback_url,omitempty"`
	EventsFilter         *string    `json:"events_filter,omitempty"`
	FriendlyName         string     `json:"friendly_name"`
	MultiTaskEnabled     bool       `json:"multi_task_enabled"`
	PrioritizeQueueOrder string     `json:"prioritize_queue_order"`
	Sid                  string     `json:"sid"`
	TimeoutActivityName  string     `json:"timeout_activity_name"`
	TimeoutActivitySid   string     `json:"timeout_activity_sid"`
	URL                  string     `json:"url"`
}

// Create creates a new workspace
// See https://www.twilio.com/docs/taskrouter/api/workspace#create-a-workspace-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateWorkspaceInput) (*CreateWorkspaceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new workspace
// See https://www.twilio.com/docs/taskrouter/api/workspace#create-a-workspace-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateWorkspaceInput) (*CreateWorkspaceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateWorkspaceInput{}
	}

	response := &CreateWorkspaceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
