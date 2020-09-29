// Package session contains auto-generated files. DO NOT MODIFY
package session

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchSessionResponse defines the response fields for the retrieved session
type FetchSessionResponse struct {
	AccountSid          string     `json:"account_sid"`
	ClosedReason        *string    `json:"closed_reason,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateEnded           *time.Time `json:"date_ended,omitempty"`
	DateExpiry          *time.Time `json:"date_expiry,omitempty"`
	DateLastInteraction *time.Time `json:"date_last_interaction,omitempty"`
	DateStarted         *time.Time `json:"date_started,omitempty"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	Mode                *string    `json:"mode,omitempty"`
	ServiceSid          string     `json:"service_sid"`
	Sid                 string     `json:"sid"`
	Status              *string    `json:"status,omitempty"`
	Ttl                 *int       `json:"ttl,omitempty"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
}

// Fetch retrieves a session resource
// See https://www.twilio.com/docs/proxy/api/session#fetch-a-session-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSessionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a session resource
// See https://www.twilio.com/docs/proxy/api/session#fetch-a-session-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSessionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Sessions/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchSessionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
