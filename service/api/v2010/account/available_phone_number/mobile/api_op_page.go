// Package mobile contains auto-generated files. DO NOT MODIFY
package mobile

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// AvailablePhoneNumbersPageOptions defines the query options for the api operation
type AvailablePhoneNumbersPageOptions struct {
	PageSize                      *int
	AreaCode                      *int
	Contains                      *string
	SmsEnabled                    *bool
	MmsEnabled                    *bool
	VoiceEnabled                  *bool
	ExcludeAllAddressRequired     *bool
	ExcludeLocalAddressRequired   *bool
	ExcludeForeignAddressRequired *bool
	Beta                          *bool
	NearNumber                    *string
	NearLatLong                   *string
	Distance                      *int
	InPostalCode                  *string
	InRegion                      *string
	InRateCenter                  *string
	InLata                        *string
	InLocality                    *string
	FaxEnabled                    *bool
}

type PageAvailablePhoneNumberResponse struct {
	AddressRequirements string                                       `json:"address_requirements"`
	Beta                bool                                         `json:"beta"`
	Capabilities        PageAvailablePhoneNumberResponseCapabilities `json:"capabilities"`
	FriendlyName        string                                       `json:"friendly_name"`
	IsoCountry          string                                       `json:"iso_country"`
	Lata                *string                                      `json:"lata,omitempty"`
	Latitude            string                                       `json:"latitude"`
	Locality            *string                                      `json:"locality,omitempty"`
	Longitude           string                                       `json:"longitude"`
	PhoneNumber         string                                       `json:"phone_number"`
	PostalCode          *string                                      `json:"postal_code,omitempty"`
	RateCenter          *string                                      `json:"rate_center,omitempty"`
	Region              *string                                      `json:"region,omitempty"`
}

type PageAvailablePhoneNumberResponseCapabilities struct {
	Fax   bool `json:"fax"`
	Mms   bool `json:"MMS"`
	Sms   bool `json:"SMS"`
	Voice bool `json:"voice"`
}

// AvailablePhoneNumbersPageResponse defines the response fields for the available mobile phone numbers page
type AvailablePhoneNumbersPageResponse struct {
	AvailablePhoneNumbers []PageAvailablePhoneNumberResponse `json:"available_phone_numbers"`
	URI                   string                             `json:"uri"`
}

// Page retrieves a page of available mobile phone numbers
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-mobile-resource#read-multiple-availablephonenumbermobile-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AvailablePhoneNumbersPageOptions) (*AvailablePhoneNumbersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of available mobile phone numbers
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-mobile-resource#read-multiple-availablephonenumbermobile-resources for more details
func (c Client) PageWithContext(context context.Context, options *AvailablePhoneNumbersPageOptions) (*AvailablePhoneNumbersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/AvailablePhoneNumbers/{countryCode}/Mobile.json",
		PathParams: map[string]string{
			"accountSid":  c.accountSid,
			"countryCode": c.countryCode,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AvailablePhoneNumbersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
