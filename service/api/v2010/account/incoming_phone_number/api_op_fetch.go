// Package incoming_phone_number contains auto-generated files. DO NOT MODIFY
package incoming_phone_number

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type FetchIncomingPhoneNumberCapabilitiesResponse struct {
	Fax   bool `json:"fax"`
	Mms   bool `json:"MMS"`
	Sms   bool `json:"SMS"`
	Voice bool `json:"voice"`
}

// FetchIncomingPhoneNumberResponse defines the response fields for retrieving a phone number
type FetchIncomingPhoneNumberResponse struct {
	APIVersion           string                                       `json:"api_version"`
	AccountSid           string                                       `json:"account_sid"`
	AddressRequirements  string                                       `json:"address_requirements"`
	AddressSid           *string                                      `json:"address_sid,omitempty"`
	Beta                 bool                                         `json:"beta"`
	BundleSid            *string                                      `json:"bundle_sid,omitempty"`
	Capabilities         FetchIncomingPhoneNumberCapabilitiesResponse `json:"capabilities"`
	DateCreated          utils.RFC2822Time                            `json:"date_created"`
	DateUpdated          *utils.RFC2822Time                           `json:"date_updated,omitempty"`
	EmergencyAddressSid  *string                                      `json:"emergency_address_sid,omitempty"`
	EmergencyStatus      string                                       `json:"emergency_status"`
	FriendlyName         *string                                      `json:"friendly_name,omitempty"`
	IdentitySid          *string                                      `json:"identity_sid,omitempty"`
	Origin               string                                       `json:"origin"`
	PhoneNumber          string                                       `json:"phone_number"`
	Sid                  string                                       `json:"sid"`
	SmsApplicationSid    *string                                      `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    string                                       `json:"sms_fallback_method"`
	SmsFallbackURL       *string                                      `json:"sms_fallback_url,omitempty"`
	SmsMethod            string                                       `json:"sms_method"`
	SmsURL               *string                                      `json:"sms_url,omitempty"`
	Status               string                                       `json:"status"`
	StatusCallback       *string                                      `json:"status_callback,omitempty"`
	StatusCallbackMethod string                                       `json:"status_callback_method"`
	TrunkSid             *string                                      `json:"trunk_sid,omitempty"`
	VoiceApplicationSid  *string                                      `json:"voice_application_sid,omitempty"`
	VoiceCallerIDLookup  bool                                         `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string                                       `json:"voice_fallback_method"`
	VoiceFallbackURL     *string                                      `json:"voice_fallback_url,omitempty"`
	VoiceMethod          string                                       `json:"voice_method"`
	VoiceReceiveMode     string                                       `json:"voice_receive_mode"`
	VoiceURL             *string                                      `json:"voice_url,omitempty"`
}

// Fetch retrieves a phone number resource
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource#fetch-an-incomingphonenumber-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchIncomingPhoneNumberResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a phone number resource
// See https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource#fetch-an-incomingphonenumber-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchIncomingPhoneNumberResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/IncomingPhoneNumbers/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchIncomingPhoneNumberResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
