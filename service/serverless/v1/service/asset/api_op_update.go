// Package asset contains auto-generated files. DO NOT MODIFY
package asset

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateAssetInput defines input fields for updating a asset resource
type UpdateAssetInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// UpdateAssetResponse defines the response fields for the updated asset
type UpdateAssetResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Update modifies a asset resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#update-an-asset-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateAssetInput) (*UpdateAssetResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a asset resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#update-an-asset-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateAssetInput) (*UpdateAssetResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Assets/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateAssetInput{}
	}

	response := &UpdateAssetResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
