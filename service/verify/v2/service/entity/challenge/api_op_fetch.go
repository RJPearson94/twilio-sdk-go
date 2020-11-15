// Package challenge contains auto-generated files. DO NOT MODIFY
package challenge

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchChallengeDetailsResponse struct {
	Date    *time.Time     `json:"date,omitempty"`
	Fields  *[]interface{} `json:"fields,omitempty"`
	Message *string        `json:"message,omitempty"`
}

// FetchChallengeResponse defines the response fields for the retrieved challenge
type FetchChallengeResponse struct {
	AccountSid      string                         `json:"account_sid"`
	DateCreated     time.Time                      `json:"date_created"`
	DateResponded   *time.Time                     `json:"date_responded,omitempty"`
	DateUpdated     *time.Time                     `json:"date_updated,omitempty"`
	Details         *FetchChallengeDetailsResponse `json:"details,omitempty"`
	EntitySid       string                         `json:"entity_sid"`
	ExpirationDate  time.Time                      `json:"expiration_date"`
	FactorSid       string                         `json:"factor_sid"`
	FactorType      string                         `json:"factor_type"`
	HiddenDetails   *map[string]interface{}        `json:"hidden_details,omitempty"`
	Identity        string                         `json:"identity"`
	RespondedReason *string                        `json:"responded_reason,omitempty"`
	ServiceSid      string                         `json:"service_sid"`
	Sid             string                         `json:"sid"`
	Status          string                         `json:"status"`
	URL             string                         `json:"url"`
}

// Fetch retrieves a challenge resource
// See https://www.twilio.com/docs/verify/api/challenge#fetch-a-challenge-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChallengeResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a challenge resource
// See https://www.twilio.com/docs/verify/api/challenge#fetch-a-challenge-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChallengeResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Entities/{identity}/Challenges/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
			"sid":        c.sid,
		},
	}

	response := &FetchChallengeResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
