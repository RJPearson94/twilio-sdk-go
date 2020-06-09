// This is an autogenerated file. DO NOT MODIFY
package workflows

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateWorkflowInput struct {
	FriendlyName                  string `validate:"required" form:"FriendlyName"`
	Configuration                 string `validate:"required" form:"Configuration"`
	AssignmentCallbackURL         string `form:"AssignmentCallbackUrl,omitempty"`
	FallbackAssignmentCallbackURL string `form:"fallbackAssignmentCallbackUrl,omitempty"`
	TaskReservationTimeout        *int   `form:"TaskReservationTimeout,omitempty"`
}

type CreateWorkflowOutput struct {
	Sid                           string      `json:"sid"`
	AccountSid                    string      `json:"account_sid"`
	WorkspaceSid                  string      `json:"workspace_sid"`
	FriendlyName                  string      `json:"friendly_name"`
	FallbackAssignmentCallbackURL *string     `json:"fallback_assignment_callback_url,omitempty"`
	AssignmentCallbackURL         *string     `json:"assignment_callback_url,omitempty"`
	TaskReservationTimeout        int         `json:"task_reservation_timeout"`
	DocumentContentType           string      `json:"document_content_type"`
	Configuration                 interface{} `json:"configuration"`
	DateCreated                   time.Time   `json:"date_created"`
	DateUpdated                   *time.Time  `json:"date_updated,omitempty"`
	URL                           string      `json:"url"`
}

func (c Client) Create(input *CreateWorkflowInput) (*CreateWorkflowOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateWorkflowInput) (*CreateWorkflowOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Workspaces/{workspaceSid}/Workflows",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	output := &CreateWorkflowOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
