// Package application contains auto-generated files. DO NOT MODIFY
package application

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchApplicationResponse defines the response fields for retrieving an application
type FetchApplicationResponse struct {
	APIVersion            string             `json:"api_version"`
	AccountSid            string             `json:"account_sid"`
	DateCreated           utils.RFC2822Time  `json:"date_created"`
	DateUpdated           *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName          *string            `json:"friendly_name,omitempty"`
	MessageStatusCallback *string            `json:"message_status_callback,omitempty"`
	Sid                   string             `json:"sid"`
	SmsFallbackMethod     string             `json:"sms_fallback_method"`
	SmsFallbackURL        *string            `json:"sms_fallback_url,omitempty"`
	SmsMethod             string             `json:"sms_method"`
	SmsStatusCallback     *string            `json:"sms_status_callback,omitempty"`
	SmsURL                *string            `json:"sms_url,omitempty"`
	StatusCallback        *string            `json:"status_callback,omitempty"`
	StatusCallbackMethod  string             `json:"status_callback_method"`
	VoiceCallerIDLookup   bool               `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod   string             `json:"voice_fallback_method"`
	VoiceFallbackURL      *string            `json:"voice_fallback_url,omitempty"`
	VoiceMethod           string             `json:"voice_method"`
	VoiceURL              *string            `json:"voice_url,omitempty"`
}

// Fetch retrieves the application resource
// See https://www.twilio.com/docs/usage/api/applications#fetch-an-application-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchApplicationResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the application resource
// See https://www.twilio.com/docs/usage/api/applications#fetch-an-application-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchApplicationResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Applications/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchApplicationResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
