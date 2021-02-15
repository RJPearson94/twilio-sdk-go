// Package domain contains auto-generated files. DO NOT MODIFY
package domain

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateDomainInput defines input fields for updating a SIP domain
type UpdateDomainInput struct {
	ByocTrunkSid              *string `form:"ByocTrunkSid,omitempty"`
	DomainName                *string `form:"DomainName,omitempty"`
	EmergencyCallerSid        *string `form:"EmergencyCallerSid,omitempty"`
	EmergencyCallingEnabled   *bool   `form:"EmergencyCallingEnabled,omitempty"`
	FriendlyName              *string `form:"FriendlyName,omitempty"`
	Secure                    *bool   `form:"Secure,omitempty"`
	SipRegistration           *bool   `form:"SipRegistration,omitempty"`
	VoiceFallbackMethod       *string `form:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackURL          *string `form:"VoiceFallbackUrl,omitempty"`
	VoiceMethod               *string `form:"VoiceMethod,omitempty"`
	VoiceStatusCallbackMethod *string `form:"VoiceStatusCallbackMethod,omitempty"`
	VoiceStatusCallbackURL    *string `form:"VoiceStatusCallbackUrl,omitempty"`
	VoiceURL                  *string `form:"VoiceUrl,omitempty"`
}

// UpdateDomainResponse defines the response fields for the updated SIP domain
type UpdateDomainResponse struct {
	AccountSid                string             `json:"account_sid"`
	ApiVersion                string             `json:"api_version"`
	AuthType                  *string            `json:"auth_type,omitempty"`
	ByocTrunkSid              *string            `json:"byoc_trunk_sid,omitempty"`
	DateCreated               utils.RFC2822Time  `json:"date_created"`
	DateUpdated               *utils.RFC2822Time `json:"date_updated,omitempty"`
	DomainName                string             `json:"domain_name"`
	EmergencyCallerSid        *string            `json:"emergency_caller_sid,omitempty"`
	EmergencyCallingEnabled   bool               `json:"emergency_calling_enabled"`
	FriendlyName              *string            `json:"friendly_name,omitempty"`
	Secure                    bool               `json:"secure"`
	Sid                       string             `json:"sid"`
	SipRegistration           bool               `json:"sip_registration"`
	VoiceFallbackMethod       *string            `json:"voice_fallback_method,omitempty"`
	VoiceFallbackURL          *string            `json:"voice_fallback_url,omitempty"`
	VoiceMethod               *string            `json:"voice_method,omitempty"`
	VoiceStatusCallbackMethod *string            `json:"voice_status_callback_method,omitempty"`
	VoiceStatusCallbackURL    *string            `json:"voice_status_callback_url,omitempty"`
	VoiceURL                  *string            `json:"voice_url,omitempty"`
}

// Update modifies a SIP domain resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#update-a-sipdomain-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateDomainInput) (*UpdateDomainResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a SIP domain resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#update-a-sipdomain-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateDomainInput) (*UpdateDomainResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/Domains/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateDomainInput{}
	}

	response := &UpdateDomainResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
