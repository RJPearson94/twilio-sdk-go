// Package services contains auto-generated files. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateServicePushInput struct {
	ApnCredentialSid *string `form:"ApnCredentialSid,omitempty"`
	FcmCredentialSid *string `form:"FcmCredentialSid,omitempty"`
	IncludeDate      *bool   `form:"IncludeDate,omitempty"`
}

// CreateServiceInput defines the input fields for creating a new service
type CreateServiceInput struct {
	CodeLength               *int                    `form:"CodeLength,omitempty"`
	CustomCodeEnabled        *bool                   `form:"CustomCodeEnabled,omitempty"`
	DoNotShareWarningEnabled *bool                   `form:"DoNotShareWarningEnabled,omitempty"`
	DtmfInputRequired        *bool                   `form:"DtmfInputRequired,omitempty"`
	FriendlyName             string                  `validate:"required" form:"FriendlyName"`
	LookupEnabled            *bool                   `form:"LookupEnabled,omitempty"`
	Psd2Enabled              *bool                   `form:"Psd2Enabled,omitempty"`
	Push                     *CreateServicePushInput `form:"Push,omitempty"`
	SkipSmsToLandlines       *bool                   `form:"SkipSmsToLandlines,omitempty"`
	TtsName                  *string                 `form:"TtsName,omitempty"`
}

type CreateServicePushResponse struct {
	ApnCredentialSid *string `json:"apn_credential_sid,omitempty"`
	FcmCredentialSid *string `json:"fcm_credential_sid,omitempty"`
	IncludeDate      bool    `json:"include_date"`
}

// CreateServiceResponse defines the response fields for the created service
type CreateServiceResponse struct {
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
	Push                     CreateServicePushResponse `json:"push"`
	Sid                      string                    `json:"sid"`
	SkipSmsToLandlines       bool                      `json:"skip_sms_to_landlines"`
	TtsName                  *string                   `json:"tts_name,omitempty"`
	URL                      string                    `json:"url"`
}

// Create creates a new service
// See https://www.twilio.com/docs/verify/api/service#create-a-verification-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new service
// See https://www.twilio.com/docs/verify/api/service#create-a-verification-service for more details
func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
