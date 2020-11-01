// Package verification contains auto-generated files. DO NOT MODIFY
package verification

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchVerificationCarrierLookupResponse struct {
	ErrorCode         *string `json:"error_code,omitempty"`
	MobileCountryCode *string `json:"mobile_country_code,omitempty"`
	MobileNetworkCode *string `json:"mobile_network_code,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
}

type FetchVerificationLookupResponse struct {
	Carrier *FetchVerificationCarrierLookupResponse `json:"carrier,omitempty"`
}

type FetchVerificationSendCodeAttemptResponse struct {
	Channel   string    `json:"channel"`
	ChannelId *string   `json:"channel_id,omitempty"`
	Time      time.Time `json:"time"`
}

// FetchVerificationResponse defines the response fields for the retrieved verification
type FetchVerificationResponse struct {
	AccountSid       string                                     `json:"account_sid"`
	Amount           *string                                    `json:"amount,omitempty"`
	Channel          string                                     `json:"channel"`
	DateCreated      time.Time                                  `json:"date_created"`
	DateUpdated      *time.Time                                 `json:"date_updated,omitempty"`
	Lookup           FetchVerificationLookupResponse            `json:"lookup"`
	Payee            *string                                    `json:"payee,omitempty"`
	SendCodeAttempts []FetchVerificationSendCodeAttemptResponse `json:"send_code_attempts"`
	ServiceSid       string                                     `json:"service_sid"`
	Sid              string                                     `json:"sid"`
	Status           string                                     `json:"status"`
	To               string                                     `json:"to"`
	URL              string                                     `json:"url"`
	Valid            bool                                       `json:"valid"`
}

// Fetch retrieves a verification resource
// See https://www.twilio.com/docs/verify/api/verification#fetch-a-verification for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchVerificationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a verification resource
// See https://www.twilio.com/docs/verify/api/verification#fetch-a-verification for more details
func (c Client) FetchWithContext(context context.Context) (*FetchVerificationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Verifications/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchVerificationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
