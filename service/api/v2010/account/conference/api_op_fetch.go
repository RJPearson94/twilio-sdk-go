// Package conference contains auto-generated files. DO NOT MODIFY
package conference

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchConferenceResponse defines the response fields for retrieving a conference
type FetchConferenceResponse struct {
	APIVersion              string             `json:"api_version"`
	AccountSid              string             `json:"account_sid"`
	CallSidEndingConference *string            `json:"call_sid_ending_conference,omitempty"`
	DateCreated             utils.RFC2822Time  `json:"date_created"`
	DateUpdated             *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName            string             `json:"friendly_name"`
	ReasonConferenceEnded   *string            `json:"reason_conference_ended,omitempty"`
	Region                  string             `json:"region"`
	Sid                     string             `json:"sid"`
	Status                  string             `json:"status"`
}

// Fetch retrieves a conference resource
// See https://www.twilio.com/docs/voice/api/conference-resource#fetch-a-conference-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchConferenceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a conference resource
// See https://www.twilio.com/docs/voice/api/conference-resource#fetch-a-conference-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchConferenceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Conferences/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchConferenceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
