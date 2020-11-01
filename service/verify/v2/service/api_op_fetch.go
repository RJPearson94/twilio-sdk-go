// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchServicePushResponse struct {
	ApnCredentialSid *string `json:"apn_credential_sid,omitempty"`
	FcmCredentialSid *string `json:"fcm_credential_sid,omitempty"`
	IncludeDate      bool    `json:"include_date"`
}

// FetchServiceResponse defines the response fields for the retrieved service
type FetchServiceResponse struct {
	AccountSid               string                   `json:"account_sid"`
	CodeLength               int                      `json:"code_length"`
	CustomCodeEnabled        bool                     `json:"custom_code_enabled"`
	DateCreated              time.Time                `json:"date_created"`
	DateUpdated              *time.Time               `json:"date_updated,omitempty"`
	DoNotShareWarningEnabled bool                     `json:"do_not_share_warning_enabled"`
	DtmfInputRequired        bool                     `json:"dtmf_input_required"`
	FriendlyName             string                   `json:"friendly_name"`
	LookupEnabled            bool                     `json:"lookup_enabled"`
	MailerSid                *string                  `json:"mailer_sid,omitempty"`
	Psd2Enabled              bool                     `json:"psd2_enabled"`
	Push                     FetchServicePushResponse `json:"push"`
	Sid                      string                   `json:"sid"`
	SkipSmsToLandlines       bool                     `json:"skip_sms_to_landlines"`
	TtsName                  *string                  `json:"tts_name,omitempty"`
	URL                      string                   `json:"url"`
}

// Fetch retrieves a service resource
// See https://www.twilio.com/docs/verify/api/service#fetch-a-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a service resource
// See https://www.twilio.com/docs/verify/api/service#fetch-a-service for more details
func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
