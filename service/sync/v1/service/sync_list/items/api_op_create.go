// Package items contains auto-generated files. DO NOT MODIFY
package items

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateSyncListItemInput defines the input fields for creating a new list item resource
type CreateSyncListItemInput struct {
	CollectionTtl *int   `form:"CollectionTtl,omitempty"`
	Data          string `validate:"required" form:"Data"`
	ItemTtl       *int   `form:"ItemTtl,omitempty"`
	Ttl           *int   `form:"Ttl,omitempty"`
}

// CreateSyncListItemResponse defines the response fields for the created list item
type CreateSyncListItemResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Index       int                    `json:"index"`
	ListSid     string                 `json:"list_sid"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	URL         string                 `json:"url"`
}

// Create creates a new list item
// See https://www.twilio.com/docs/sync/api/listitem-resource#create-a-listitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSyncListItemInput) (*CreateSyncListItemResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new list item
// See https://www.twilio.com/docs/sync/api/listitem-resource#create-a-listitem-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSyncListItemInput) (*CreateSyncListItemResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists/{syncListSid}/Items",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
		},
	}

	if input == nil {
		input = &CreateSyncListItemInput{}
	}

	response := &CreateSyncListItemResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
