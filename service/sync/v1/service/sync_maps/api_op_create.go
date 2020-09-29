// Package sync_maps contains auto-generated files. DO NOT MODIFY
package sync_maps

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateSyncMapInput defines the input fields for creating a new map resource
type CreateSyncMapInput struct {
	CollectionTtl *int    `form:"CollectionTtl,omitempty"`
	Ttl           *int    `form:"Ttl,omitempty"`
	UniqueName    *string `form:"UniqueName,omitempty"`
}

// CreateSyncMapResponse defines the response fields for the created map
type CreateSyncMapResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Revision    string     `json:"revision"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// Create creates a new map
// See https://www.twilio.com/docs/sync/api/map-resource#create-a-syncmap-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSyncMapInput) (*CreateSyncMapResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new map
// See https://www.twilio.com/docs/sync/api/map-resource#create-a-syncmap-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSyncMapInput) (*CreateSyncMapResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Maps",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateSyncMapInput{}
	}

	response := &CreateSyncMapResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
