// Package binding contains auto-generated files. DO NOT MODIFY
package binding

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchBindingResponse defines the response fields for the retrieved service binding
type FetchBindingResponse struct {
	AccountSid                  string     `json:"account_sid"`
	Address                     string     `json:"address"`
	BindingType                 string     `json:"binding_type"`
	CredentialSid               *string    `json:"credential_sid,omitempty"`
	DateCreated                 time.Time  `json:"date_created"`
	DateUpdated                 *time.Time `json:"date_updated,omitempty"`
	Identity                    string     `json:"identity"`
	NotificationProtocolVersion string     `json:"notification_protocol_version"`
	ServiceSid                  string     `json:"service_sid"`
	Sid                         string     `json:"sid"`
	Tags                        []string   `json:"tags"`
	URL                         string     `json:"url"`
}

// Fetch retrieves a service binding resource
// See https://www.twilio.com/docs/notify/api/binding-resource#fetch-a-binding-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchBindingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a service binding resource
// See https://www.twilio.com/docs/notify/api/binding-resource#fetch-a-binding-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchBindingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Bindings/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchBindingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
