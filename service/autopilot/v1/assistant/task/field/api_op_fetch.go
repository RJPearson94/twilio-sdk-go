// Package field contains auto-generated files. DO NOT MODIFY
package field

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFieldResponse defines the response fields for the retrieved task field
type FetchFieldResponse struct {
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

// Fetch retrieves a task field resource
// See https://www.twilio.com/docs/autopilot/api/task-field#fetch-a-field-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFieldResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task field resource
// See https://www.twilio.com/docs/autopilot/api/task-field#fetch-a-field-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFieldResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Fields/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
			"sid":          c.sid,
		},
	}

	response := &FetchFieldResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
