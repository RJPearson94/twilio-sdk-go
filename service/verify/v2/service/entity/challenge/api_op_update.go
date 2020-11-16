// Package challenge contains auto-generated files. DO NOT MODIFY
package challenge

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateChallengeInput defines input fields for updating a challenge resource
type UpdateChallengeInput struct {
	AuthPayload *string `form:"AuthPayload,omitempty"`
}

type UpdateChallengeDetailsResponse struct {
	Date    *time.Time     `json:"date,omitempty"`
	Fields  *[]interface{} `json:"fields,omitempty"`
	Message *string        `json:"message,omitempty"`
}

// UpdateChallengeResponse defines the response fields for the updated challenge
type UpdateChallengeResponse struct {
	AccountSid      string                          `json:"account_sid"`
	DateCreated     time.Time                       `json:"date_created"`
	DateResponded   *time.Time                      `json:"date_responded,omitempty"`
	DateUpdated     *time.Time                      `json:"date_updated,omitempty"`
	Details         *UpdateChallengeDetailsResponse `json:"details,omitempty"`
	EntitySid       string                          `json:"entity_sid"`
	ExpirationDate  time.Time                       `json:"expiration_date"`
	FactorSid       string                          `json:"factor_sid"`
	FactorType      string                          `json:"factor_type"`
	HiddenDetails   *map[string]interface{}         `json:"hidden_details,omitempty"`
	Identity        string                          `json:"identity"`
	RespondedReason *string                         `json:"responded_reason,omitempty"`
	ServiceSid      string                          `json:"service_sid"`
	Sid             string                          `json:"sid"`
	Status          string                          `json:"status"`
	URL             string                          `json:"url"`
}

// Update modifies a challenge resource
// See https://www.twilio.com/docs/verify/api/challenge#update-a-challenge-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Update(input *UpdateChallengeInput) (*UpdateChallengeResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a challenge resource
// See https://www.twilio.com/docs/verify/api/challenge#update-a-challenge-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) UpdateWithContext(context context.Context, input *UpdateChallengeInput) (*UpdateChallengeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Entities/{identity}/Challenges/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateChallengeInput{}
	}

	response := &UpdateChallengeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
