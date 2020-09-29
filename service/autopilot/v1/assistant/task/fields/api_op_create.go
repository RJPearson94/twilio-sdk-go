// Package fields contains auto-generated files. DO NOT MODIFY
package fields

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateFieldInput defines the input fields for creating a new task field resource
type CreateFieldInput struct {
	FieldType  string `validate:"required" form:"FieldType"`
	UniqueName string `validate:"required" form:"UniqueName"`
}

// CreateFieldResponse defines the response fields for the created task field
type CreateFieldResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FieldType    string     `json:"field_type"`
	Sid          string     `json:"sid"`
	TaskSid      string     `json:"task_sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Create creates a new task field
// See https://www.twilio.com/docs/autopilot/api/task-field#create-a-field-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFieldInput) (*CreateFieldResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new task field
// See https://www.twilio.com/docs/autopilot/api/task-field#create-a-field-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFieldInput) (*CreateFieldResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks/{taskSid}/Fields",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	if input == nil {
		input = &CreateFieldInput{}
	}

	response := &CreateFieldResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
