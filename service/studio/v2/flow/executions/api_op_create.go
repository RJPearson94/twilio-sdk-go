// Package executions contains auto-generated files. DO NOT MODIFY
package executions

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateExecutionInput defines the input fields for creating a new execution resource
type CreateExecutionInput struct {
	From       string  `validate:"required" form:"From"`
	Parameters *string `form:"Parameters,omitempty"`
	To         string  `validate:"required" form:"To"`
}

// CreateExecutionResponse defines the response fields for the created execution
type CreateExecutionResponse struct {
	AccountSid            string      `json:"account_sid"`
	ContactChannelAddress string      `json:"contact_channel_address"`
	Context               interface{} `json:"context"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	FlowSid               string      `json:"flow_sid"`
	Sid                   string      `json:"sid"`
	Status                string      `json:"status"`
	URL                   string      `json:"url"`
}

// Create creates a new execution
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#create-a-new-execution for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateExecutionInput) (*CreateExecutionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new execution
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#create-a-new-execution for more details
func (c Client) CreateWithContext(context context.Context, input *CreateExecutionInput) (*CreateExecutionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Flows/{flowSid}/Executions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"flowSid": c.flowSid,
		},
	}

	if input == nil {
		input = &CreateExecutionInput{}
	}

	response := &CreateExecutionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
