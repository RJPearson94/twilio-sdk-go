// Package conference contains auto-generated files. DO NOT MODIFY
package conference

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateConferenceInput defines input fields for updating a conference
type UpdateConferenceInput struct {
	AnnounceMethod *string `form:"AnnounceMethod,omitempty"`
	AnnounceURL    *string `form:"AnnounceUrl,omitempty"`
	Status         *string `form:"Status,omitempty"`
}

// UpdateConferenceResponse defines the response fields for the updated conference
type UpdateConferenceResponse struct {
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

// Update modifies a conference resource
// See https://www.twilio.com/docs/voice/api/conference-resource#update-a-conference-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateConferenceInput) (*UpdateConferenceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a conference resource
// See https://www.twilio.com/docs/voice/api/conference-resource#update-a-conference-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateConferenceInput) (*UpdateConferenceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Conferences/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateConferenceInput{}
	}

	response := &UpdateConferenceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
