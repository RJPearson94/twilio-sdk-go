// Package verification_check contains auto-generated files. DO NOT MODIFY
package verification_check

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateVerificationCheckInput defines the input fields for performing a new verification check
type CreateVerificationCheckInput struct {
	Amount          *string `form:"Amount,omitempty"`
	Code            string  `validate:"required" form:"Code"`
	Payee           *string `form:"Payee,omitempty"`
	To              *string `form:"To,omitempty"`
	VerificationSid *string `form:"VerificationSid,omitempty"`
}

// CreateVerificationCheckResponse defines the response fields for the performed verification check
type CreateVerificationCheckResponse struct {
	AccountSid  string     `json:"account_sid"`
	Amount      *string    `json:"amount,omitempty"`
	Channel     string     `json:"channel"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Payee       *string    `json:"payee,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	To          string     `json:"to"`
	Valid       bool       `json:"valid"`
}

// Create creates/ performs a verification check
// See https://www.twilio.com/docs/verify/api/verification-check#check-a-verification for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateVerificationCheckInput) (*CreateVerificationCheckResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates/ performs a verification check
// See https://www.twilio.com/docs/verify/api/verification-check#check-a-verification for more details
func (c Client) CreateWithContext(context context.Context, input *CreateVerificationCheckInput) (*CreateVerificationCheckResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/VerificationCheck",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateVerificationCheckInput{}
	}

	response := &CreateVerificationCheckResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
