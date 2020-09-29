// Package workflow contains auto-generated files. DO NOT MODIFY
package workflow

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchWorkflowResponse defines the response fields for the retrieved workflow
type FetchWorkflowResponse struct {
	AccountSid                    string     `json:"account_sid"`
	AssignmentCallbackURL         *string    `json:"assignment_callback_url,omitempty"`
	Configuration                 string     `json:"configuration"`
	DateCreated                   time.Time  `json:"date_created"`
	DateUpdated                   *time.Time `json:"date_updated,omitempty"`
	DocumentContentType           string     `json:"document_content_type"`
	FallbackAssignmentCallbackURL *string    `json:"fallback_assignment_callback_url,omitempty"`
	FriendlyName                  string     `json:"friendly_name"`
	Sid                           string     `json:"sid"`
	TaskReservationTimeout        int        `json:"task_reservation_timeout"`
	URL                           string     `json:"url"`
	WorkspaceSid                  string     `json:"workspace_sid"`
}

// Fetch retrieves an workflow resource
// See https://www.twilio.com/docs/taskrouter/api/workflow#fetch-a-workflow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWorkflowResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an workflow resource
// See https://www.twilio.com/docs/taskrouter/api/workflow#fetch-a-workflow-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchWorkflowResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workflows/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &FetchWorkflowResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
