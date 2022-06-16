// Package verifications contains auto-generated files. DO NOT MODIFY
package verifications

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateVerificationInput defines the input fields for creating a new verification
type CreateVerificationInput struct {
	Amount                      *string `form:"Amount,omitempty"`
	AppHash                     *string `form:"AppHash,omitempty"`
	Channel                     string  `validate:"required" form:"Channel"`
	ChannelConfiguration        *string `form:"ChannelConfiguration,omitempty"`
	CustomCode                  *string `form:"CustomCode,omitempty"`
	CustomFriendlyName          *string `form:"CustomFriendlyName,omitempty"`
	Locale                      *string `form:"Locale,omitempty"`
	Payee                       *string `form:"Payee,omitempty"`
	RateLimits                  *string `form:"RateLimits,omitempty"`
	SendDigits                  *string `form:"SendDigits,omitempty"`
	TemplateCustomSubstitutions *string `form:"TemplateCustomSubstitutions,omitempty"`
	TemplateSid                 *string `form:"TemplateSid,omitempty"`
	To                          string  `validate:"required" form:"To"`
}

type CreateVerificationCarrierLookupResponse struct {
	ErrorCode         *string `json:"error_code,omitempty"`
	MobileCountryCode *string `json:"mobile_country_code,omitempty"`
	MobileNetworkCode *string `json:"mobile_network_code,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
}

type CreateVerificationLookupResponse struct {
	Carrier *CreateVerificationCarrierLookupResponse `json:"carrier,omitempty"`
}

type CreateVerificationSendCodeAttemptResponse struct {
	Channel   string    `json:"channel"`
	ChannelId *string   `json:"channel_id,omitempty"`
	Time      time.Time `json:"time"`
}

// CreateVerificationResponse defines the response fields for the created verification
type CreateVerificationResponse struct {
	AccountSid       string                                      `json:"account_sid"`
	Amount           *string                                     `json:"amount,omitempty"`
	Channel          string                                      `json:"channel"`
	DateCreated      time.Time                                   `json:"date_created"`
	DateUpdated      *time.Time                                  `json:"date_updated,omitempty"`
	Lookup           CreateVerificationLookupResponse            `json:"lookup"`
	Payee            *string                                     `json:"payee,omitempty"`
	SendCodeAttempts []CreateVerificationSendCodeAttemptResponse `json:"send_code_attempts"`
	ServiceSid       string                                      `json:"service_sid"`
	Sid              string                                      `json:"sid"`
	Status           string                                      `json:"status"`
	To               string                                      `json:"to"`
	URL              string                                      `json:"url"`
	Valid            bool                                        `json:"valid"`
}

// Create creates/ starts a new verification
// See https://www.twilio.com/docs/verify/api/verification#start-new-verification for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateVerificationInput) (*CreateVerificationResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates/ starts a new verification
// See https://www.twilio.com/docs/verify/api/verification#start-new-verification for more details
func (c Client) CreateWithContext(context context.Context, input *CreateVerificationInput) (*CreateVerificationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Verifications",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateVerificationInput{}
	}

	response := &CreateVerificationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
