// This is an autogenerated file. DO NOT MODIFY
package short_code

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateShortCodeInput struct {
	IsReserved *bool `form:"IsReserved,omitempty"`
}

type UpdateShortCodeResponseCapabilities struct {
	SmsInbound               *bool `json:"sms_inbound,omitempty"`
	SmsOutbound              *bool `json:"sms_outbound,omitempty"`
	RestrictionSmsDomestic   *bool `json:"restriction_sms_domestic,omitempty"`
	RestrictionVoiceDomestic *bool `json:"restriction_voice_domestic,omitempty"`
	VoiceOutbound            *bool `json:"voice_outbound,omitempty"`
	VoiceInbound             *bool `json:"voice_inbound,omitempty"`
	FaxInbound               *bool `json:"fax_inbound,omitempty"`
	FaxOutbound              *bool `json:"fax_outbound,omitempty"`
	RestrictionFaxDomestic   *bool `json:"restriction_fax_domestic,omitempty"`
	RestrictionMmsDomestic   *bool `json:"restriction_mms_domestic,omitempty"`
	MmsOutbound              *bool `json:"mms_outbound,omitempty"`
	MmsInbound               *bool `json:"mms_inbound,omitempty"`
	SipTrunking              *bool `json:"sip_trunking,omitempty"`
}

type UpdateShortCodeOutput struct {
	Sid          string                               `json:"sid"`
	AccountSid   string                               `json:"account_sid"`
	ServiceSid   string                               `json:"service_sid"`
	ShortCode    *string                              `json:"short_code,omitempty"`
	IsoCountry   *string                              `json:"iso_country,omitempty"`
	Capabilities *UpdateShortCodeResponseCapabilities `json:"capabilities,omitempty"`
	IsReserved   *bool                                `json:"is_reserved,omitempty"`
	DateCreated  time.Time                            `json:"date_created"`
	DateUpdated  *time.Time                           `json:"date_updated,omitempty"`
	URL          string                               `json:"url"`
}

func (c Client) Update(input *UpdateShortCodeInput) (*UpdateShortCodeOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateShortCodeInput) (*UpdateShortCodeOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services/{serviceSid}/ShortCodes/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	output := &UpdateShortCodeOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}