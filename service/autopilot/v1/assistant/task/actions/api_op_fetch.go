// Package actions contains auto-generated files. DO NOT MODIFY
package actions

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchActionResponse defines the response fields for the retrieved task action>
type FetchActionResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	TaskSid      string                 `json:"task_sid"`
	URL          string                 `json:"url"`
}

// Fetch retrieves a task action resource
// See https://www.twilio.com/docs/autopilot/api/task-action#fetch-a-taskactions-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchActionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task action resource
// See https://www.twilio.com/docs/autopilot/api/task-action#fetch-a-taskactions-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchActionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Actions",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	response := &FetchActionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
