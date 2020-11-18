// Package short_code contains auto-generated files. DO NOT MODIFY
package short_code

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateShortCodeInput defines input fields for updating a short code resource
type UpdateShortCodeInput struct {
	IsReserved *bool `form:"IsReserved,omitempty"`
}

type UpdateShortCodeCapabilitiesResponse struct {
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

// UpdateShortCodeResponse defines the response fields for the updated short code
type UpdateShortCodeResponse struct {
	AccountSid   string                               `json:"account_sid"`
	Capabilities *UpdateShortCodeCapabilitiesResponse `json:"capabilities,omitempty"`
	DateCreated  time.Time                            `json:"date_created"`
	DateUpdated  *time.Time                           `json:"date_updated,omitempty"`
	IsReserved   *bool                                `json:"is_reserved,omitempty"`
	IsoCountry   *string                              `json:"iso_country,omitempty"`
	ServiceSid   string                               `json:"service_sid"`
	ShortCode    *string                              `json:"short_code,omitempty"`
	Sid          string                               `json:"sid"`
	URL          string                               `json:"url"`
}

// Update modifies a short code resource
// See https://www.twilio.com/docs/proxy/api/short-code#update-a-shortcode-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateShortCodeInput) (*UpdateShortCodeResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a short code resource
// See https://www.twilio.com/docs/proxy/api/short-code#update-a-shortcode-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateShortCodeInput) (*UpdateShortCodeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/ShortCodes/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateShortCodeInput{}
	}

	response := &UpdateShortCodeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
