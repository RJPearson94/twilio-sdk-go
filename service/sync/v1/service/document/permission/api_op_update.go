// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateDocumentPermissionsInput defines input fields for updating a document permission resource
type UpdateDocumentPermissionsInput struct {
	Manage bool `form:"Manage"`
	Read   bool `form:"Read"`
	Write  bool `form:"Write"`
}

// UpdateDocumentPermissionsResponse defines the response fields for the updated document permission
type UpdateDocumentPermissionsResponse struct {
	AccountSid  string `json:"account_sid"`
	DocumentSid string `json:"document_sid"`
	Identity    string `json:"identity"`
	Manage      bool   `json:"manage"`
	Read        bool   `json:"read"`
	ServiceSid  string `json:"service_sid"`
	URL         string `json:"url"`
	Write       bool   `json:"write"`
}

// Update modifies a document permission resource
// See https://www.twilio.com/docs/sync/api/document-permission-resource#update-a-document-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateDocumentPermissionsInput) (*UpdateDocumentPermissionsResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a document permission resource
// See https://www.twilio.com/docs/sync/api/document-permission-resource#update-a-document-permission-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateDocumentPermissionsInput) (*UpdateDocumentPermissionsResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Documents/{documentSid}/Permissions/{identity}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"documentSid": c.documentSid,
			"identity":    c.identity,
		},
	}

	if input == nil {
		input = &UpdateDocumentPermissionsInput{}
	}

	response := &UpdateDocumentPermissionsResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
