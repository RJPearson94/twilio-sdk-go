// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchDocumentPermissionsResponse defines the response fields for the retrieved document permission
type FetchDocumentPermissionsResponse struct {
	AccountSid  string `json:"account_sid"`
	DocumentSid string `json:"document_sid"`
	Identity    string `json:"identity"`
	Manage      bool   `json:"manage"`
	Read        bool   `json:"read"`
	ServiceSid  string `json:"service_sid"`
	URL         string `json:"url"`
	Write       bool   `json:"write"`
}

// Fetch retrieves an document permission resource
// See https://www.twilio.com/docs/sync/api/document-permission-resource#fetch-a-document-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDocumentPermissionsResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an document permission resource
// See https://www.twilio.com/docs/sync/api/document-permission-resource#fetch-a-document-permission-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDocumentPermissionsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Documents/{documentSid}/Permissions/{identity}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"documentSid": c.documentSid,
			"identity":    c.identity,
		},
	}

	response := &FetchDocumentPermissionsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
