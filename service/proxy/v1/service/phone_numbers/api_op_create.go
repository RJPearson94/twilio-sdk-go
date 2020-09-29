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
	IsReserved  *bool   `form:"IsReserved,omitempty"`
	PhoneNumber *string `form:"PhoneNumber,omitempty"`
	Sid         *string `form:"Sid,omitempty"`
}

type CreatePhoneNumberResponseCapabilities struct {
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

// CreatePhoneNumberResponse defines the response fields for the created phone number
type CreatePhoneNumberResponse struct {
	AccountSid   string                                 `json:"account_sid"`
	Capabilities *CreatePhoneNumberResponseCapabilities `json:"capabilities,omitempty"`
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

// Create adds a new phone number resource to the proxy service
// See https://www.twilio.com/docs/proxy/api/phone-number#add-a-phone-number-to-a-proxy-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreatePhoneNumberInput) (*CreatePhoneNumberResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext adds a new phone number resource to the proxy service
// See https://www.twilio.com/docs/proxy/api/phone-number#add-a-phone-number-to-a-proxy-service for more details
func (c Client) CreateWithContext(context context.Context, input *CreatePhoneNumberInput) (*CreatePhoneNumberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/PhoneNumbers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
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
