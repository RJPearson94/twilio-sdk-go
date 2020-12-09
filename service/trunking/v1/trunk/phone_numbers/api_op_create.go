// Package phone_numbers contains auto-generated files. DO NOT MODIFY
package phone_numbers

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreatePhoneNumberInput defines the input fields for creating a new phone number resource
type CreatePhoneNumberInput struct {
	PhoneNumberSid string `validate:"required" form:"PhoneNumberSid"`
}

type CreatePhoneNumberCapabilitiesResponse struct {
	Fax   *bool `json:"fax,omitempty"`
	Mms   bool  `json:"MMS"`
	Sms   bool  `json:"SMS"`
	Voice bool  `json:"voice"`
}

// CreatePhoneNumberResponse defines the response fields for the created phone number resource
type CreatePhoneNumberResponse struct {
	APIVersion           string                                `json:"api_version"`
	AccountSid           string                                `json:"account_sid"`
	AddressRequirements  string                                `json:"address_requirements"`
	Beta                 bool                                  `json:"beta"`
	Capabilities         CreatePhoneNumberCapabilitiesResponse `json:"capabilities"`
	DateCreated          time.Time                             `json:"date_created"`
	DateUpdated          *time.Time                            `json:"date_updated,omitempty"`
	FriendlyName         *string                               `json:"friendly_name,omitempty"`
	PhoneNumber          string                                `json:"phone_number"`
	Sid                  string                                `json:"sid"`
	SmsApplicationSid    *string                               `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    string                                `json:"sms_fallback_method"`
	SmsFallbackURL       *string                               `json:"sms_fallback_url,omitempty"`
	SmsMethod            string                                `json:"sms_method"`
	SmsURL               *string                               `json:"sms_url,omitempty"`
	StatusCallback       *string                               `json:"status_callback,omitempty"`
	StatusCallbackMethod string                                `json:"status_callback_method"`
	TrunkSid             string                                `json:"trunk_sid"`
	URL                  string                                `json:"url"`
	VoiceApplicationSid  *string                               `json:"voice_application_sid,omitempty"`
	VoiceCallerIDLookup  bool                                  `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string                                `json:"voice_fallback_method"`
	VoiceFallbackURL     *string                               `json:"voice_fallback_url,omitempty"`
	VoiceMethod          string                                `json:"voice_method"`
	VoiceReceiveMode     *string                               `json:"voice_receive_mode,omitempty"`
	VoiceURL             *string                               `json:"voice_url,omitempty"`
}

// Create adds a phone number resource to the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource#create-a-phonenumber-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreatePhoneNumberInput) (*CreatePhoneNumberResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext adds a phone number resource to the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource#create-a-phonenumber-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreatePhoneNumberInput) (*CreatePhoneNumberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{trunkSid}/PhoneNumbers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
	}

	if input == nil {
		input = &CreatePhoneNumberInput{}
	}

	response := &CreatePhoneNumberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
