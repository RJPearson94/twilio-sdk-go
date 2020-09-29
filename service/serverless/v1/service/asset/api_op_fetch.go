// Package asset contains auto-generated files. DO NOT MODIFY
package asset

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchAssetResponse defines the response fields for the retrieved asset
type FetchAssetResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a asset resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#fetch-an-asset-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAssetResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a asset resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#fetch-an-asset-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAssetResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Assets/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchAssetResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
