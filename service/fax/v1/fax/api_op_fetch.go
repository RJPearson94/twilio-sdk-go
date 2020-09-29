// Package fax contains auto-generated files. DO NOT MODIFY
package fax

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFaxResponse defines the response fields for the retrieved fax
type FetchFaxResponse struct {
	APIVersion  string     `json:"api_version"`
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Direction   string     `json:"direction"`
	Duration    *int       `json:"duration,omitempty"`
	From        string     `json:"from"`
	MediaSid    *string    `json:"media_sid,omitempty"`
	MediaURL    *string    `json:"media_url,omitempty"`
	NumPages    *int       `json:"num_pages,omitempty"`
	Price       *string    `json:"price,omitempty"`
	PriceUnit   *string    `json:"price_unit,omitempty"`
	Quality     string     `json:"quality"`
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	To          string     `json:"to"`
	URL         string     `json:"url"`
}

// Fetch retrieves a fax resource
// See https://www.twilio.com/docs/fax/api/fax-resource#fetch-a-fax-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFaxResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a fax resource
// See https://www.twilio.com/docs/fax/api/fax-resource#fetch-a-fax-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFaxResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Faxes/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchFaxResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
