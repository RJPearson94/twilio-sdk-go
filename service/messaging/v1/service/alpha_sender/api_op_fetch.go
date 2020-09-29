// Package alpha_sender contains auto-generated files. DO NOT MODIFY
package alpha_sender

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchAlphaSenderResponse defines the response fields for the retrieved alpha sender
type FetchAlphaSenderResponse struct {
	AccountSid   string     `json:"account_sid"`
	AlphaSender  string     `json:"alpha_sender"`
	Capabilities []string   `json:"capabilities"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a alpha sender resource
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#fetch-an-alphasender-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAlphaSenderResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a alpha sender resource
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#fetch-an-alphasender-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAlphaSenderResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/AlphaSenders/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchAlphaSenderResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
