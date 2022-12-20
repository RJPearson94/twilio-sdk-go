// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchServiceResponse defines the response fields for the retrieved service
type FetchServiceResponse struct {
	AccountSid         string     `json:"account_sid"`
	DateCreated        time.Time  `json:"date_created"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	DomainBase         string     `json:"domain_base"`
	FriendlyName       string     `json:"friendly_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	Sid                string     `json:"sid"`
	URL                string     `json:"url"`
	UiEditable         bool       `json:"ui_editable"`
	UniqueName         string     `json:"unique_name"`
}

// Fetch retrieves a service resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service#fetch-a-service-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a service resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service#fetch-a-service-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
