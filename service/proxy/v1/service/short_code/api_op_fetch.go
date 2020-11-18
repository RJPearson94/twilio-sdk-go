// Package short_code contains auto-generated files. DO NOT MODIFY
package short_code

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchShortCodeCapabilitiesResponse struct {
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

// FetchShortCodeResponse defines the response fields for the retrieved short code
type FetchShortCodeResponse struct {
	AccountSid   string                              `json:"account_sid"`
	Capabilities *FetchShortCodeCapabilitiesResponse `json:"capabilities,omitempty"`
	DateCreated  time.Time                           `json:"date_created"`
	DateUpdated  *time.Time                          `json:"date_updated,omitempty"`
	IsReserved   *bool                               `json:"is_reserved,omitempty"`
	IsoCountry   *string                             `json:"iso_country,omitempty"`
	ServiceSid   string                              `json:"service_sid"`
	ShortCode    *string                             `json:"short_code,omitempty"`
	Sid          string                              `json:"sid"`
	URL          string                              `json:"url"`
}

// Fetch retrieves a short code resource
// See https://www.twilio.com/docs/proxy/api/short-code#fetch-a-shortcode-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchShortCodeResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a short code resource
// See https://www.twilio.com/docs/proxy/api/short-code#fetch-a-shortcode-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchShortCodeResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/ShortCodes/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchShortCodeResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
