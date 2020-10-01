// Package status contains auto-generated files. DO NOT MODIFY
package status

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchStatusResponse defines the response fields for the retrieved build status
type FetchStatusResponse struct {
	AccountSid string `json:"account_sid"`
	ServiceSid string `json:"service_sid"`
	Sid        string `json:"sid"`
	Status     string `json:"status"`
	URL        string `json:"url"`
}

// Fetch retrieves a build status resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchStatusResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a build status resource
func (c Client) FetchWithContext(context context.Context) (*FetchStatusResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Builds/{buildSid}/Status",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"buildSid":   c.buildSid,
		},
	}

	response := &FetchStatusResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
