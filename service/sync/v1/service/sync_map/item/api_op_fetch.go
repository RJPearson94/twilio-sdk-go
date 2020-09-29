// Package item contains auto-generated files. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchSyncMapItemResponse defines the response fields for the retrieved map item
type FetchSyncMapItemResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Key         string                 `json:"key"`
	MapSid      string                 `json:"map_sid"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	URL         string                 `json:"url"`
}

// Fetch retrieves an map item resource
// See https://www.twilio.com/docs/sync/api/map-item-resource#fetch-a-mapitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSyncMapItemResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an map item resource
// See https://www.twilio.com/docs/sync/api/map-item-resource#fetch-a-mapitem-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSyncMapItemResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Maps/{syncMapSid}/Items/{key}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
			"key":        c.key,
		},
	}

	response := &FetchSyncMapItemResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
