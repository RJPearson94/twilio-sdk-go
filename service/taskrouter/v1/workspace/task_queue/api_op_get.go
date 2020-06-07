// This is an autogenerated file. DO NOT MODIFY
package task_queue

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetTaskQueueOutput struct {
	Sid                     string     `json:"sid"`
	AccountSid              string     `json:"account_sid"`
	WorkspaceSid            string     `json:"workspace_sid"`
	FriendlyName            string     `json:"friendly_name"`
	EventCallbackURL        *string    `json:"event_callback_url,omitempty"`
	AssignmentActivityName  string     `json:"assignment_activity_name"`
	AssignmentActivitySid   string     `json:"assignment_activity_sid"`
	ReservationActivityName string     `json:"reservation_activity_name"`
	ReservationActivitySid  string     `json:"reservation_activity_sid"`
	TargetWorkers           string     `json:"target_workers"`
	TaskOrder               string     `json:"task_order"`
	MaxReservedWorkers      int        `json:"max_reserved_workers"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	URL                     string     `json:"url"`
}

func (c Client) Get() (*GetTaskQueueOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetTaskQueueOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Workspaces/{workspaceSid}/TaskQueues/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	output := &GetTaskQueueOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
