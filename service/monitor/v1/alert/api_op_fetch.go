// Package alert contains auto-generated files. DO NOT MODIFY
package alert

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchAlertResponse defines the response fields for the retrieved alert
type FetchAlertResponse struct {
	APIVersion       *string    `json:"api_version,omitempty"`
	AccountSid       string     `json:"account_sid"`
	AlertText        *string    `json:"alert_text,omitempty"`
	DateCreated      time.Time  `json:"date_created"`
	DateGenerated    time.Time  `json:"date_generated"`
	DateUpdated      *time.Time `json:"date_updated,omitempty"`
	ErrorCode        string     `json:"error_code"`
	LogLevel         string     `json:"log_level"`
	MoreInfo         string     `json:"more_info"`
	RequestHeaders   *string    `json:"request_headers,omitempty"`
	RequestMethod    *string    `json:"request_method,omitempty"`
	RequestURL       *string    `json:"request_url,omitempty"`
	RequestVariables *string    `json:"request_variables,omitempty"`
	ResourceSid      string     `json:"resource_sid"`
	ResponseBody     *string    `json:"response_body,omitempty"`
	ResponseHeaders  *string    `json:"response_headers,omitempty"`
	ServiceSid       string     `json:"service_sid"`
	Sid              string     `json:"sid"`
	URL              string     `json:"url"`
}

// Fetch retrieves an alert resource
// See https://www.twilio.com/docs/usage/monitor-alert#fetch-an-alert-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAlertResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an alert resource
// See https://www.twilio.com/docs/usage/monitor-alert#fetch-an-alert-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAlertResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Alerts/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchAlertResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
