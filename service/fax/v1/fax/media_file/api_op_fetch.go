// Package media_file contains auto-generated files. DO NOT MODIFY
package media_file

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchMediaResponse defines the response fields for the retrieved media
type FetchMediaResponse struct {
	AccountSid  string     `json:"account_sid"`
	ContentType string     `json:"content_type"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FaxSid      string     `json:"fax_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Fetch retrieves a media resource
// See https://www.twilio.com/docs/fax/api/fax-media-resource#fetch-a-fax-media-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchMediaResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a media resource
// See https://www.twilio.com/docs/fax/api/fax-media-resource#fetch-a-fax-media-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchMediaResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Faxes/{faxSid}/Media/{sid}",
		PathParams: map[string]string{
			"faxSid": c.faxSid,
			"sid":    c.sid,
		},
	}

	response := &FetchMediaResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
