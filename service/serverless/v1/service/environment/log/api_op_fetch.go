// Package log contains auto-generated files. DO NOT MODIFY
package log

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchLogResponse defines the response fields for the retrieved log entry
type FetchLogResponse struct {
	AccountSid     string    `json:"account_sid"`
	BuildSid       string    `json:"build_sid"`
	DateCreated    time.Time `json:"date_created"`
	DeploymentSid  string    `json:"deployment_sid"`
	EnvironmentSid string    `json:"environment_sid"`
	FunctionSid    string    `json:"function_sid"`
	Level          string    `json:"level"`
	Message        string    `json:"message"`
	RequestSid     string    `json:"request_sid"`
	ServiceSid     string    `json:"service_sid"`
	Sid            string    `json:"sid"`
	URL            string    `json:"url"`
}

// Fetch retrieves a log resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/logs#fetch-a-log-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchLogResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a log resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/logs#fetch-a-log-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchLogResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Logs/{sid}",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
			"sid":            c.sid,
		},
	}

	response := &FetchLogResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
