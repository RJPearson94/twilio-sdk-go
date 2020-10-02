// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchParticipantResponse defines the response fields for retrieving a participant
type FetchParticipantResponse struct {
	AccountSid             string             `json:"account_sid"`
	CallSid                string             `json:"call_sid"`
	CallSidToCoach         *string            `json:"call_sid_to_coach,omitempty"`
	Coaching               bool               `json:"coaching"`
	ConferenceSid          string             `json:"conference_sid"`
	DateCreated            utils.RFC2822Time  `json:"date_created"`
	DateUpdated            *utils.RFC2822Time `json:"date_updated,omitempty"`
	EndConferenceOnExit    bool               `json:"end_conference_on_exit"`
	Hold                   bool               `json:"hold"`
	Label                  *string            `json:"label,omitempty"`
	Muted                  bool               `json:"muted"`
	StartConferenceOnEnter bool               `json:"start_conference_on_enter"`
	Status                 string             `json:"status"`
}

// Fetch retrieves the participant resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#fetch-a-participant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchParticipantResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the participant resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#fetch-a-participant-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchParticipantResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Conferences/{conferenceSid}/Participants/{sid}.json",
		PathParams: map[string]string{
			"accountSid":    c.accountSid,
			"conferenceSid": c.conferenceSid,
			"sid":           c.sid,
		},
	}

	response := &FetchParticipantResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
