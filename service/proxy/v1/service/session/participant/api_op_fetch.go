// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchParticipantResponse defines the response fields for the retrieved participant
type FetchParticipantResponse struct {
	AccountSid         string     `json:"account_sid"`
	DateCreated        time.Time  `json:"date_created"`
	DateDeleted        *time.Time `json:"date_deleted,omitempty"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	FriendlyName       *string    `json:"friendly_name,omitempty"`
	Identifier         string     `json:"identifier"`
	ProxyIdentifier    *string    `json:"proxy_identifier,omitempty"`
	ProxyIdentifierSid *string    `json:"proxy_identifier_sid,omitempty"`
	ServiceSid         string     `json:"service_sid"`
	SessionSid         string     `json:"session_sid"`
	Sid                string     `json:"sid"`
	URL                string     `json:"url"`
}

// Fetch retrieves a participant resource
// See https://www.twilio.com/docs/proxy/api/participant#fetch-a-participant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchParticipantResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a participant resource
// See https://www.twilio.com/docs/proxy/api/participant#fetch-a-participant-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchParticipantResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Sessions/{sessionSid}/Participants/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sessionSid": c.sessionSid,
			"sid":        c.sid,
		},
	}

	response := &FetchParticipantResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
