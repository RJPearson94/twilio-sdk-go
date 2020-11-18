// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdatePhoneNumberInput defines input fields for updating a phone number resource
type UpdatePhoneNumberInput struct {
	IsReserved *bool `form:"IsReserved,omitempty"`
}

type UpdatePhoneNumberCapabilitiesResponse struct {
	FaxInbound               *bool `json:"fax_inbound,omitempty"`
	FaxOutbound              *bool `json:"fax_outbound,omitempty"`
	MmsInbound               *bool `json:"mms_inbound,omitempty"`
	MmsOutbound              *bool `json:"mms_outbound,omitempty"`
	RestrictionFaxDomestic   *bool `json:"restriction_fax_domestic,omitempty"`
	RestrictionMmsDomestic   *bool `json:"restriction_mms_domestic,omitempty"`
	RestrictionSmsDomestic   *bool `json:"restriction_sms_domestic,omitempty"`
	RestrictionVoiceDomestic *bool `json:"restriction_voice_domestic,omitempty"`
	SipTrunking              *bool `json:"sip_trunking,omitempty"`
	SmsInbound               *bool `json:"sms_inbound,omitempty"`
	SmsOutbound              *bool `json:"sms_outbound,omitempty"`
	VoiceInbound             *bool `json:"voice_inbound,omitempty"`
	VoiceOutbound            *bool `json:"voice_outbound,omitempty"`
}

// UpdatePhoneNumberResponse defines the response fields for the updated phone number
type UpdatePhoneNumberResponse struct {
	AccountSid   string                                 `json:"account_sid"`
	Capabilities *UpdatePhoneNumberCapabilitiesResponse `json:"capabilities,omitempty"`
	DateCreated  time.Time                              `json:"date_created"`
	DateUpdated  *time.Time                             `json:"date_updated,omitempty"`
	FriendlyName *string                                `json:"friendly_name,omitempty"`
	InUse        *int                                   `json:"in_use,omitempty"`
	IsReserved   *bool                                  `json:"is_reserved,omitempty"`
	IsoCountry   *string                                `json:"iso_country,omitempty"`
	PhoneNumber  *string                                `json:"phone_number,omitempty"`
	ServiceSid   string                                 `json:"service_sid"`
	Sid          string                                 `json:"sid"`
	URL          string                                 `json:"url"`
}

// Update modifies a phone number resource
// See https://www.twilio.com/docs/proxy/api/phone-number#update-a-phonenumber-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdatePhoneNumberInput) (*UpdatePhoneNumberResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a phone number resource
// See https://www.twilio.com/docs/proxy/api/phone-number#update-a-phonenumber-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdatePhoneNumberInput) (*UpdatePhoneNumberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/PhoneNumbers/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdatePhoneNumberInput{}
	}

	response := &UpdatePhoneNumberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
