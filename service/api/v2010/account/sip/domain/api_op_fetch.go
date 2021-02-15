// Package domain contains auto-generated files. DO NOT MODIFY
package domain

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchDomainResponse defines the response fields for retrieving a SIP domain
type FetchDomainResponse struct {
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

// Fetch retrieves a SIP domain resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#fetch-a-sipdomain-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchDomainResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a SIP domain resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#fetch-a-sipdomain-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchDomainResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchDomainResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
