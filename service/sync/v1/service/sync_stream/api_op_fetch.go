// Package sync_stream contains auto-generated files. DO NOT MODIFY
package sync_stream

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchSyncStreamResponse defines the response fields for the retrieved stream
type FetchSyncStreamResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// Fetch retrieves an stream resource
// See https://www.twilio.com/docs/sync/api/stream-resource#fetch-a-sync-stream-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSyncStreamResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an stream resource
// See https://www.twilio.com/docs/sync/api/stream-resource#fetch-a-sync-stream-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSyncStreamResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Streams/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchSyncStreamResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
