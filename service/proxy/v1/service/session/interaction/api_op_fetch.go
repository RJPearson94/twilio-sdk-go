// Package interaction contains auto-generated files. DO NOT MODIFY
package interaction

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchInteractionResponse defines the response fields for the retrieved interaction
type FetchInteractionResponse struct {
	AccountSid             string     `json:"account_sid"`
	Data                   *string    `json:"data,omitempty"`
	DateCreated            time.Time  `json:"date_created"`
	DateUpdated            *time.Time `json:"date_updated,omitempty"`
	InboundParticipantSid  *string    `json:"inbound_participant_sid,omitempty"`
	InboundResourceSid     *string    `json:"inbound_resource_sid,omitempty"`
	InboundResourceStatus  *string    `json:"inbound_resource_status,omitempty"`
	InboundResourceType    *string    `json:"inbound_resource_type,omitempty"`
	InboundResourceURL     *string    `json:"inbound_resource_url,omitempty"`
	OutboundParticipantSid *string    `json:"outbound_participant_sid,omitempty"`
	OutboundResourceSid    *string    `json:"outbound_resource_sid,omitempty"`
	OutboundResourceStatus *string    `json:"outbound_resource_status,omitempty"`
	OutboundResourceType   *string    `json:"outbound_resource_type,omitempty"`
	OutboundResourceURL    *string    `json:"outbound_resource_url,omitempty"`
	ServiceSid             string     `json:"service_sid"`
	SessionSid             string     `json:"session_sid"`
	Sid                    string     `json:"sid"`
	Type                   *string    `json:"type,omitempty"`
	URL                    string     `json:"url"`
}

// Fetch retrieves a interaction resource
// See https://www.twilio.com/docs/proxy/api/interaction#fetch-an-interaction-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchInteractionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a interaction resource
// See https://www.twilio.com/docs/proxy/api/interaction#fetch-an-interaction-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchInteractionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Sessions/{sessionSid}/Interactions/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sessionSid": c.sessionSid,
			"sid":        c.sid,
		},
	}

	response := &FetchInteractionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
