// Package task_queue contains auto-generated files. DO NOT MODIFY
package task_queue

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchTaskQueueResponse defines the response fields for the retrieved task queue
type FetchTaskQueueResponse struct {
	AccountSid              string     `json:"account_sid"`
	AssignmentActivityName  *string    `json:"assignment_activity_name,omitempty"`
	AssignmentActivitySid   *string    `json:"assignment_activity_sid,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	EventCallbackURL        *string    `json:"event_callback_url,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	MaxReservedWorkers      int        `json:"max_reserved_workers"`
	ReservationActivityName *string    `json:"reservation_activity_name,omitempty"`
	ReservationActivitySid  *string    `json:"reservation_activity_sid,omitempty"`
	Sid                     string     `json:"sid"`
	TargetWorkers           *string    `json:"target_workers,omitempty"`
	TaskOrder               string     `json:"task_order"`
	URL                     string     `json:"url"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// Fetch retrieves a task queue resource
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchTaskQueueResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task queue resource
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-get for more details
func (c Client) FetchWithContext(context context.Context) (*FetchTaskQueueResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/TaskQueues/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &FetchTaskQueueResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
