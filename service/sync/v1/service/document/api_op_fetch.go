// Package document contains auto-generated files. DO NOT MODIFY
package document

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchDocumentResponse defines the response fields for the retrieved document
type FetchDocumentResponse struct {
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

// Fetch retrieves an document resource
// See https://www.twilio.com/docs/sync/api/document-resource#fetch-a-document-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDocumentResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an document resource
// See https://www.twilio.com/docs/sync/api/document-resource#fetch-a-document-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDocumentResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Documents/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchDocumentResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
