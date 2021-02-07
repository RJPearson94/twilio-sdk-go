// Package entity contains auto-generated files. DO NOT MODIFY
package entity

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchEntityResponse defines the response fields for the retrieved entity
type FetchEntityResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    string     `json:"identity"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Fetch retrieves an entity resource
// See https://www.twilio.com/docs/verify/api/entity#fetch-an-entity-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Fetch() (*FetchEntityResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an entity resource
// See https://www.twilio.com/docs/verify/api/entity#fetch-an-entity-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) FetchWithContext(context context.Context) (*FetchEntityResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Entities/{identity}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
		},
	}

	response := &FetchEntityResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
