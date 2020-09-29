// Package execution contains auto-generated files. DO NOT MODIFY
package execution

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchExecutionResponse defines the response fields for the retrieved execution
type FetchExecutionResponse struct {
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

// Fetch retrieves a execution resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#fetch-a-single-execution for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchExecutionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a execution resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#fetch-a-single-execution for more details
func (c Client) FetchWithContext(context context.Context) (*FetchExecutionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions/{sid}",
		PathParams: map[string]string{
			"flowSid": c.flowSid,
			"sid":     c.sid,
		},
	}

	response := &FetchExecutionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
