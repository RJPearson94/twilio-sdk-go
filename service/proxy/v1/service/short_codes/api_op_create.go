// Package short_codes contains auto-generated files. DO NOT MODIFY
package short_codes

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateShortCodeInput defines the input fields for creating a new short code resource
type CreateShortCodeInput struct {
	Sid string `validate:"required" form:"Sid"`
}

type CreateShortCodeCapabilitiesResponse struct {
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

// CreateShortCodeResponse defines the response fields for the created short code
type CreateShortCodeResponse struct {
	AccountSid   string                               `json:"account_sid"`
	Capabilities *CreateShortCodeCapabilitiesResponse `json:"capabilities,omitempty"`
	DateCreated  time.Time                            `json:"date_created"`
	DateUpdated  *time.Time                           `json:"date_updated,omitempty"`
	IsReserved   *bool                                `json:"is_reserved,omitempty"`
	IsoCountry   *string                              `json:"iso_country,omitempty"`
	ServiceSid   string                               `json:"service_sid"`
	ShortCode    *string                              `json:"short_code,omitempty"`
	Sid          string                               `json:"sid"`
	URL          string                               `json:"url"`
}

// Create add a new short code to the proxy service
// See https://www.twilio.com/docs/proxy/api/short-code#add-a-short-code-to-a-proxy-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateShortCodeInput) (*CreateShortCodeResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext add a new short code to the proxy service
// See https://www.twilio.com/docs/proxy/api/short-code#add-a-short-code-to-a-proxy-service for more details
func (c Client) CreateWithContext(context context.Context, input *CreateShortCodeInput) (*CreateShortCodeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/ShortCodes",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateShortCodeInput{}
	}

	response := &CreateShortCodeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
