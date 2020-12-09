// Package incoming_phone_numbers contains auto-generated files. DO NOT MODIFY
package incoming_phone_numbers

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateIncomingPhoneNumberInput defines input fields for creating a new phone number
type CreateIncomingPhoneNumberInput struct {
	APIVersion           *string `form:"ApiVersion,omitempty"`
	AddressSid           *string `form:"AddressSid,omitempty"`
	AreaCode             *string `form:"AreaCode,omitempty"`
	BundleSid            *string `form:"BundleSid,omitempty"`
	EmergencyAddressSid  *string `form:"EmergencyAddressSid,omitempty"`
	EmergencyStatus      *string `form:"EmergencyStatus,omitempty"`
	FriendlyName         *string `form:"FriendlyName,omitempty"`
	IdentitySid          *string `form:"IdentitySid,omitempty"`
	PhoneNumber          *string `form:"PhoneNumber,omitempty"`
	SmsApplicationSid    *string `form:"SmsApplicationSid,omitempty"`
	SmsFallbackMethod    *string `form:"SmsFallbackMethod,omitempty"`
	SmsFallbackURL       *string `form:"SmsFallbackUrl,omitempty"`
	SmsMethod            *string `form:"SmsMethod,omitempty"`
	SmsURL               *string `form:"SmsUrl,omitempty"`
	StatusCallback       *string `form:"StatusCallback,omitempty"`
	StatusCallbackMethod *string `form:"StatusCallbackMethod,omitempty"`
	TrunkSid             *string `form:"TrunkSid,omitempty"`
	VoiceApplicationSid  *string `form:"VoiceApplicationSid,omitempty"`
	VoiceCallerIDLookup  *bool   `form:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod  *string `form:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackURL     *string `form:"VoiceFallbackUrl,omitempty"`
	VoiceMethod          *string `form:"VoiceMethod,omitempty"`
	VoiceReceiveMode     *string `form:"VoiceReceiveMode,omitempty"`
	VoiceURL             *string `form:"VoiceUrl,omitempty"`
}

type CreateIncomingPhoneNumberCapabilitiesResponse struct {
	Fax   *bool `json:"fax,omitempty"`
	Mms   bool  `json:"MMS"`
	Sms   bool  `json:"SMS"`
	Voice bool  `json:"voice"`
}

// CreateIncomingPhoneNumberResponse defines the response fields for creating a new phone number
type CreateIncomingPhoneNumberResponse struct {
	APIVersion           string                                        `json:"api_version"`
	AccountSid           string                                        `json:"account_sid"`
	AddressRequirements  string                                        `json:"address_requirements"`
	AddressSid           *string                                       `json:"address_sid,omitempty"`
	Beta                 bool                                          `json:"beta"`
	BundleSid            *string                                       `json:"bundle_sid,omitempty"`
	Capabilities         CreateIncomingPhoneNumberCapabilitiesResponse `json:"capabilities"`
	DateCreated          utils.RFC2822Time                             `json:"date_created"`
	DateUpdated          *utils.RFC2822Time                            `json:"date_updated,omitempty"`
	EmergencyAddressSid  *string                                       `json:"emergency_address_sid,omitempty"`
	EmergencyStatus      string                                        `json:"emergency_status"`
	FriendlyName         *string                                       `json:"friendly_name,omitempty"`
	IdentitySid          *string                                       `json:"identity_sid,omitempty"`
	Origin               string                                        `json:"origin"`
	PhoneNumber          string                                        `json:"phone_number"`
	Sid                  string                                        `json:"sid"`
	SmsApplicationSid    *string                                       `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    string                                        `json:"sms_fallback_method"`
	SmsFallbackURL       *string                                       `json:"sms_fallback_url,omitempty"`
	SmsMethod            string                                        `json:"sms_method"`
	SmsURL               *string                                       `json:"sms_url,omitempty"`
	Status               string                                        `json:"status"`
	StatusCallback       *string                                       `json:"status_callback,omitempty"`
	StatusCallbackMethod string                                        `json:"status_callback_method"`
	TrunkSid             *string                                       `json:"trunk_sid,omitempty"`
	VoiceApplicationSid  *string                                       `json:"voice_application_sid,omitempty"`
	VoiceCallerIDLookup  bool                                          `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string                                        `json:"voice_fallback_method"`
	VoiceFallbackURL     *string                                       `json:"voice_fallback_url,omitempty"`
	VoiceMethod          string                                        `json:"voice_method"`
	VoiceReceiveMode     *string                                       `json:"voice_receive_mode,omitempty"`
	VoiceURL             *string                                       `json:"voice_url,omitempty"`
}

// Create creates a new phone number resource
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource#create-an-incomingphonenumber-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateIncomingPhoneNumberInput) (*CreateIncomingPhoneNumberResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new phone number resource
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource#create-an-incomingphonenumber-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateIncomingPhoneNumberInput) (*CreateIncomingPhoneNumberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/IncomingPhoneNumbers.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateIncomingPhoneNumberInput{}
	}

	response := &CreateIncomingPhoneNumberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
