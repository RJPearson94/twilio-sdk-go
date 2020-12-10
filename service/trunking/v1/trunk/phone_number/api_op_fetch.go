// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchPhoneNumberCapabilitiesResponse struct {
	Fax   *bool `json:"fax,omitempty"`
	Mms   bool  `json:"MMS"`
	Sms   bool  `json:"SMS"`
	Voice bool  `json:"voice"`
}

// FetchPhoneNumberResponse defines the response fields for the retrieved phone number resource
type FetchPhoneNumberResponse struct {
	APIVersion           string                               `json:"api_version"`
	AccountSid           string                               `json:"account_sid"`
	AddressRequirements  string                               `json:"address_requirements"`
	Beta                 bool                                 `json:"beta"`
	Capabilities         FetchPhoneNumberCapabilitiesResponse `json:"capabilities"`
	DateCreated          time.Time                            `json:"date_created"`
	DateUpdated          *time.Time                           `json:"date_updated,omitempty"`
	FriendlyName         *string                              `json:"friendly_name,omitempty"`
	PhoneNumber          string                               `json:"phone_number"`
	Sid                  string                               `json:"sid"`
	SmsApplicationSid    *string                              `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    string                               `json:"sms_fallback_method"`
	SmsFallbackURL       *string                              `json:"sms_fallback_url,omitempty"`
	SmsMethod            string                               `json:"sms_method"`
	SmsURL               *string                              `json:"sms_url,omitempty"`
	StatusCallback       *string                              `json:"status_callback,omitempty"`
	StatusCallbackMethod string                               `json:"status_callback_method"`
	TrunkSid             string                               `json:"trunk_sid"`
	URL                  string                               `json:"url"`
	VoiceApplicationSid  *string                              `json:"voice_application_sid,omitempty"`
	VoiceCallerIDLookup  bool                                 `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string                               `json:"voice_fallback_method"`
	VoiceFallbackURL     *string                              `json:"voice_fallback_url,omitempty"`
	VoiceMethod          string                               `json:"voice_method"`
	VoiceReceiveMode     *string                              `json:"voice_receive_mode,omitempty"`
	VoiceURL             *string                              `json:"voice_url,omitempty"`
}

// Fetch retrieves a phone number resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchPhoneNumberResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a phone number resource
func (c Client) FetchWithContext(context context.Context) (*FetchPhoneNumberResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/PhoneNumbers/{sid}",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
			"sid":      c.sid,
		},
	}

	response := &FetchPhoneNumberResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
