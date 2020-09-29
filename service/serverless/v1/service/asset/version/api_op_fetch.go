// Package version contains auto-generated files. DO NOT MODIFY
package version

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchVersionResponse defines the response fields for the retrieved asset version
type FetchVersionResponse struct {
	AccountSid  string    `json:"account_sid"`
	AssetSid    string    `json:"asset_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

// Fetch retrieves a asset version resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#fetch-an-asset-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchVersionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a asset version resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#fetch-an-asset-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchVersionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Assets/{assetSid}/Versions/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"assetSid":   c.assetSid,
			"sid":        c.sid,
		},
	}

	response := &FetchVersionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
