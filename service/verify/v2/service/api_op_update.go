// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateServiceInput defines input fields for updating a service resource
type UpdateServiceInput struct {
	CodeLength               *int    `form:"CodeLength,omitempty"`
	CustomCodeEnabled        *bool   `form:"CustomCodeEnabled,omitempty"`
	DoNotShareWarningEnabled *bool   `form:"DoNotShareWarningEnabled,omitempty"`
	DtmfInputRequired        *bool   `form:"DtmfInputRequired,omitempty"`
	FriendlyName             *string `form:"FriendlyName,omitempty"`
	LookupEnabled            *bool   `form:"LookupEnabled,omitempty"`
	Psd2Enabled              *bool   `form:"Psd2Enabled,omitempty"`
	PushApnCredentialSid     *string `form:"Push.ApnCredentialSid,omitempty"`
	PushFcmCredentialSid     *string `form:"Push.FcmCredentialSid,omitempty"`
	PushIncludeDate          *bool   `form:"Push.IncludeDate,omitempty"`
	SkipSmsToLandlines       *bool   `form:"SkipSmsToLandlines,omitempty"`
	TtsName                  *string `form:"TtsName,omitempty"`
}

type UpdateServicePushResponse struct {
	ApnCredentialSid *string `json:"apn_credential_sid,omitempty"`
	FcmCredentialSid *string `json:"fcm_credential_sid,omitempty"`
	IncludeDate      bool    `json:"include_date"`
}

// UpdateServiceResponse defines the response fields for the updated service
type UpdateServiceResponse struct {
	AccountSid               string                    `json:"account_sid"`
	CodeLength               int                       `json:"code_length"`
	CustomCodeEnabled        bool                      `json:"custom_code_enabled"`
	DateCreated              time.Time                 `json:"date_created"`
	DateUpdated              *time.Time                `json:"date_updated,omitempty"`
	DoNotShareWarningEnabled bool                      `json:"do_not_share_warning_enabled"`
	DtmfInputRequired        bool                      `json:"dtmf_input_required"`
	FriendlyName             string                    `json:"friendly_name"`
	LookupEnabled            bool                      `json:"lookup_enabled"`
	MailerSid                *string                   `json:"mailer_sid,omitempty"`
	Psd2Enabled              bool                      `json:"psd2_enabled"`
	Push                     UpdateServicePushResponse `json:"push"`
	Sid                      string                    `json:"sid"`
	SkipSmsToLandlines       bool                      `json:"skip_sms_to_landlines"`
	TtsName                  *string                   `json:"tts_name,omitempty"`
	URL                      string                    `json:"url"`
}

// Update modifies a service resource
// See https://www.twilio.com/docs/verify/api/service#update-a-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a service resource
// See https://www.twilio.com/docs/verify/api/service#update-a-service for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	response := &UpdateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
