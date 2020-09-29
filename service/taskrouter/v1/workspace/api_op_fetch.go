// Package workspace contains auto-generated files. DO NOT MODIFY
package workspace

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchWorkspaceResponse defines the response fields for the retrieved workspace
type FetchWorkspaceResponse struct {
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

// Fetch retrieves a workspace resource
// See https://www.twilio.com/docs/taskrouter/api/workspace#fetch-a-workspace-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWorkspaceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a workspace resource
// See https://www.twilio.com/docs/taskrouter/api/workspace#fetch-a-workspace-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchWorkspaceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchWorkspaceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
