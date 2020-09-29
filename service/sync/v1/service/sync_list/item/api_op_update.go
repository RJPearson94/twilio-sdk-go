// Package item contains auto-generated files. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateSyncListItemInput defines input fields for updating a list item resource
type UpdateSyncListItemInput struct {
	CollectionTtl *int    `form:"CollectionTtl,omitempty"`
	Data          *string `form:"Data,omitempty"`
	ItemTtl       *int    `form:"ItemTtl,omitempty"`
	Ttl           *int    `form:"Ttl,omitempty"`
}

// UpdateSyncListItemResponse defines the response fields for the updated list item
type UpdateSyncListItemResponse struct {
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

// Update modifies a list item resource
// See https://www.twilio.com/docs/sync/api/listitem-resource#update-a-listitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncListItemInput) (*UpdateSyncListItemResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a list item resource
// See https://www.twilio.com/docs/sync/api/listitem-resource#update-a-listitem-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncListItemInput) (*UpdateSyncListItemResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists/{syncListSid}/Items/{index}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
			"index":       strconv.Itoa(c.index),
		},
	}

	if input == nil {
		input = &UpdateSyncListItemInput{}
	}

	response := &UpdateSyncListItemResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
