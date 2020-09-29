// Package sync_streams contains auto-generated files. DO NOT MODIFY
package sync_streams

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateSyncStreamInput defines the input fields for creating a new stream resource
type CreateSyncStreamInput struct {
	Ttl        *int    `form:"Ttl,omitempty"`
	UniqueName *string `form:"UniqueName,omitempty"`
}

// CreateSyncStreamResponse defines the response fields for the created stream
type CreateSyncStreamResponse struct {
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

// Create creates a new stream
// See https://www.twilio.com/docs/sync/api/stream-resource#create-a-sync-stream-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSyncStreamInput) (*CreateSyncStreamResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new stream
// See https://www.twilio.com/docs/sync/api/stream-resource#create-a-sync-stream-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSyncStreamInput) (*CreateSyncStreamResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Streams",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateSyncStreamInput{}
	}

	response := &CreateSyncStreamResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
