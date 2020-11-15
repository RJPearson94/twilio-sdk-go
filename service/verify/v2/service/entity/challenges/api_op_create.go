// Package challenges contains auto-generated files. DO NOT MODIFY
package challenges

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateChallengeInput defines the input fields for creating a new challenge
type CreateChallengeInput struct {
	DetailsFields  *[]string  `form:"Details.Fields,omitempty"`
	DetailsMessage *string    `form:"Details.Message,omitempty"`
	ExpirationDate *time.Time `form:"ExpirationDate,omitempty"`
	FactorSid      string     `validate:"required" form:"FactorSid"`
	HiddenDetails  *string    `form:"HiddenDetails,omitempty"`
}

type CreateChallengeDetailsResponse struct {
	Date    *time.Time     `json:"date,omitempty"`
	Fields  *[]interface{} `json:"fields,omitempty"`
	Message *string        `json:"message,omitempty"`
}

// CreateChallengeResponse defines the response fields for the created challenge
type CreateChallengeResponse struct {
	AccountSid      string                          `json:"account_sid"`
	DateCreated     time.Time                       `json:"date_created"`
	DateResponded   *time.Time                      `json:"date_responded,omitempty"`
	DateUpdated     *time.Time                      `json:"date_updated,omitempty"`
	Details         *CreateChallengeDetailsResponse `json:"details,omitempty"`
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

// Create creates a new challenge
// See https://www.twilio.com/docs/verify/api/challenge#create-a-challenge-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChallengeInput) (*CreateChallengeResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new challenge
// See https://www.twilio.com/docs/verify/api/challenge#create-a-challenge-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateChallengeInput) (*CreateChallengeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Entities/{identity}/Challenges",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
		},
	}

	if input == nil {
		input = &CreateChallengeInput{}
	}

	response := &CreateChallengeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
