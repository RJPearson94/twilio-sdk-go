// Package item contains auto-generated files. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchSyncListItemResponse defines the response fields for the retrieved list item
type FetchSyncListItemResponse struct {
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

// Fetch retrieves an list item resource
// See https://www.twilio.com/docs/sync/api/listitem-resource#fetch-a-listitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSyncListItemResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an list item resource
// See https://www.twilio.com/docs/sync/api/listitem-resource#fetch-a-listitem-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSyncListItemResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Lists/{syncListSid}/Items/{index}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
			"index":       strconv.Itoa(c.index),
		},
	}

	response := &FetchSyncListItemResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
