// Package binding contains auto-generated files. DO NOT MODIFY
package binding

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchUserBindingResponse defines the response fields for the retrieved user binding
type FetchUserBindingResponse struct {
	AccountSid    string     `json:"account_sid"`
	BindingType   *string    `json:"binding_type,omitempty"`
	CredentialSid string     `json:"credential_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Endpoint      *string    `json:"endpoint,omitempty"`
	Identity      *string    `json:"identity,omitempty"`
	MessageTypes  *[]string  `json:"message_types,omitempty"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	URL           string     `json:"url"`
	UserSid       string     `json:"user_sid"`
}

// Fetch retrieves a user binding resource
// See https://www.twilio.com/docs/chat/rest/user-binding-resource#fetch-a-user-binding-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchUserBindingResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a user binding resource
// See https://www.twilio.com/docs/chat/rest/user-binding-resource#fetch-a-user-binding-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchUserBindingResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Users/{userSid}/Bindings/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
			"sid":        c.sid,
		},
	}

	response := &FetchUserBindingResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
