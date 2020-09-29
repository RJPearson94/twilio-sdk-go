// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchSyncMapPermissionsResponse defines the response fields for the retrieved map item permission
type FetchSyncMapPermissionsResponse struct {
	AccountSid string `json:"account_sid"`
	Identity   string `json:"identity"`
	Manage     bool   `json:"manage"`
	MapSid     string `json:"map_sid"`
	Read       bool   `json:"read"`
	ServiceSid string `json:"service_sid"`
	URL        string `json:"url"`
	Write      bool   `json:"write"`
}

// Fetch retrieves an map item permission resource
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource#fetch-a-sync-map-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSyncMapPermissionsResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an map item permission resource
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource#fetch-a-sync-map-permission-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSyncMapPermissionsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Maps/{syncMapSid}/Permissions/{identity}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
			"identity":   c.identity,
		},
	}

	response := &FetchSyncMapPermissionsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
