// Package variable contains auto-generated files. DO NOT MODIFY
package variable

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchVariableResponse defines the response fields for the retrieved environment variable
type FetchVariableResponse struct {
	AccountSid     string     `json:"account_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string     `json:"environment_sid"`
	Key            string     `json:"key"`
	ServiceSid     string     `json:"service_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
	Value          string     `json:"value"`
}

// Fetch retrieves a environment variable resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#fetch-a-variable-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchVariableResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a environment variable resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#fetch-a-variable-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchVariableResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Variables/{sid}",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
			"sid":            c.sid,
		},
	}

	response := &FetchVariableResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
