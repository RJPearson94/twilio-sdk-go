// Package document contains auto-generated files. DO NOT MODIFY
package document

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateDocumentInput defines input fields for updating a document resource
type UpdateDocumentInput struct {
	Data *string `form:"Data,omitempty"`
	Ttl  *int    `form:"Ttl,omitempty"`
}

// UpdateDocumentResponse defines the response fields for the updated document
type UpdateDocumentResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	Sid         string                 `json:"sid"`
	URL         string                 `json:"url"`
	UniqueName  *string                `json:"unique_name,omitempty"`
}

// Update modifies a document resource
// See https://www.twilio.com/docs/sync/api/document-resource#update-a-document-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateDocumentInput) (*UpdateDocumentResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a document resource
// See https://www.twilio.com/docs/sync/api/document-resource#update-a-document-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateDocumentInput) (*UpdateDocumentResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Documents/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateDocumentInput{}
	}

	response := &UpdateDocumentResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
