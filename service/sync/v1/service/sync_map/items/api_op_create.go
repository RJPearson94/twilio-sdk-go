// Package items contains auto-generated files. DO NOT MODIFY
package items

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateSyncMapItemInput defines the input fields for creating a new map item resource
type CreateSyncMapItemInput struct {
	CollectionTtl *int   `form:"CollectionTtl,omitempty"`
	Data          string `validate:"required" form:"Data"`
	ItemTtl       *int   `form:"ItemTtl,omitempty"`
	Key           string `validate:"required" form:"Key"`
	Ttl           *int   `form:"Ttl,omitempty"`
}

// CreateSyncMapItemResponse defines the response fields for the created map item
type CreateSyncMapItemResponse struct {
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

// Create creates a new map item
// See https://www.twilio.com/docs/sync/api/map-item-resource#create-a-mapitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSyncMapItemInput) (*CreateSyncMapItemResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new map item
// See https://www.twilio.com/docs/sync/api/map-item-resource#create-a-mapitem-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSyncMapItemInput) (*CreateSyncMapItemResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Maps/{syncMapSid}/Items",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
		},
	}

	if input == nil {
		input = &CreateSyncMapItemInput{}
	}

	response := &CreateSyncMapItemResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
