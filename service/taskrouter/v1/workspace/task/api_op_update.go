// Package task contains auto-generated files. DO NOT MODIFY
package task

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateTaskInput defines input fields for updating a task resource
type UpdateTaskInput struct {
	AssignmentStatus *string `form:"AssignmentStatus,omitempty"`
	Attributes       *string `form:"Attributes,omitempty"`
	Priority         *int    `form:"Priority,omitempty"`
	Reason           *string `form:"Reason,omitempty"`
	TaskChannel      *string `form:"TaskChannel,omitempty"`
}

// UpdateTaskResponse defines the response fields for the updated task
type UpdateTaskResponse struct {
	AccountSid            string      `json:"account_sid"`
	Age                   int         `json:"age"`
	AssignmentStatus      string      `json:"assignment_status"`
	Attributes            interface{} `json:"attributes"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	Priority              *int        `json:"priority,omitempty"`
	Reason                *string     `json:"reason,omitempty"`
	Sid                   string      `json:"sid"`
	TaskChannelSid        *string     `json:"task_channel_sid,omitempty"`
	TaskChannelUniqueName *string     `json:"task_channel_unique_name,omitempty"`
	TaskQueueEnteredDate  *time.Time  `json:"task_queue_entered_date,omitempty"`
	TaskQueueFriendlyName *string     `json:"task_queue_friendly_name,omitempty"`
	TaskQueueSid          *string     `json:"task_queue_sid,omitempty"`
	Timeout               int         `json:"timeout"`
	URL                   string      `json:"url"`
	WorkflowFriendlyName  *string     `json:"workflow_friendly_name,omitempty"`
	WorkflowSid           *string     `json:"workflow_sid,omitempty"`
	WorkspaceSid          string      `json:"workspace_sid"`
}

// Update modifies a task resource
// See https://www.twilio.com/docs/taskrouter/api/task#update-a-task-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateTaskInput) (*UpdateTaskResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a task resource
// See https://www.twilio.com/docs/taskrouter/api/task#update-a-task-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateTaskInput) (*UpdateTaskResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Tasks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateTaskInput{}
	}

	response := &UpdateTaskResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
