// Package verification contains auto-generated files. DO NOT MODIFY
package verification

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateVerificationInput defines input fields for updating a verification resource
type UpdateVerificationInput struct {
	Status string `validate:"required" form:"Status"`
}

type UpdateVerificationCarrierLookupResponse struct {
	ErrorCode         *string `json:"error_code,omitempty"`
	MobileCountryCode *string `json:"mobile_country_code,omitempty"`
	MobileNetworkCode *string `json:"mobile_network_code,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
}

type UpdateVerificationLookupResponse struct {
	Carrier *UpdateVerificationCarrierLookupResponse `json:"carrier,omitempty"`
}

type UpdateVerificationSendCodeAttemptResponse struct {
	Channel   string    `json:"channel"`
	ChannelId *string   `json:"channel_id,omitempty"`
	Time      time.Time `json:"time"`
}

// UpdateVerificationResponse defines the response fields for the updated verification
type UpdateVerificationResponse struct {
	AccountSid       string                                      `json:"account_sid"`
	Amount           *string                                     `json:"amount,omitempty"`
	Channel          string                                      `json:"channel"`
	DateCreated      time.Time                                   `json:"date_created"`
	DateUpdated      *time.Time                                  `json:"date_updated,omitempty"`
	Lookup           UpdateVerificationLookupResponse            `json:"lookup"`
	Payee            *string                                     `json:"payee,omitempty"`
	SendCodeAttempts []UpdateVerificationSendCodeAttemptResponse `json:"send_code_attempts"`
	ServiceSid       string                                      `json:"service_sid"`
	Sid              string                                      `json:"sid"`
	Status           string                                      `json:"status"`
	To               string                                      `json:"to"`
	URL              string                                      `json:"url"`
	Valid            bool                                        `json:"valid"`
}

// Update modifies a verification resource
// See https://www.twilio.com/docs/verify/api/verification#update-a-verification-status for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateVerificationInput) (*UpdateVerificationResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a verification resource
// See https://www.twilio.com/docs/verify/api/verification#update-a-verification-status for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateVerificationInput) (*UpdateVerificationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Verifications/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateVerificationInput{}
	}

	response := &UpdateVerificationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
