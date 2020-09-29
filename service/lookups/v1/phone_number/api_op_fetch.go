// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchPhoneNumberOptions defines the query options for the api operation
type FetchPhoneNumberOptions struct {
	CountryCode *string
	Type        *[]string
	AddOns      *[]string
	AddOnsData  *map[string]interface{}
}

type FetchCallerNameResponse struct {
	CallerName *string `json:"caller_name,omitempty"`
	CallerType *string `json:"caller_type,omitempty"`
	ErrorCode  *string `json:"error_code,omitempty"`
}

type FetchCarrierResponse struct {
	ErrorCode         *string `json:"error_code,omitempty"`
	MobileCountryCode *string `json:"mobile_country_code,omitempty"`
	MobileNetworkCode *string `json:"mobile_network_code,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
}

// FetchPhoneNumberResponse defines the response fields for the retrieved phone number
type FetchPhoneNumberResponse struct {
	AddOns         *map[string]interface{}  `json:"add_ons,omitempty"`
	CallerName     *FetchCallerNameResponse `json:"caller_name,omitempty"`
	Carrier        *FetchCarrierResponse    `json:"carrier,omitempty"`
	CountryCode    string                   `json:"country_code"`
	NationalFormat string                   `json:"national_format"`
	PhoneNumber    string                   `json:"phone_number"`
	URL            string                   `json:"url"`
}

// Fetch retrieves a phone number details
// See https://www.twilio.com/docs/lookup/api#lookup-a-phone-number for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchPhoneNumberOptions) (*FetchPhoneNumberResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves a phone number details
// See https://www.twilio.com/docs/lookup/api#lookup-a-phone-number for more details
func (c Client) FetchWithContext(context context.Context, options *FetchPhoneNumberOptions) (*FetchPhoneNumberResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/PhoneNumbers/{phoneNumber}",
		PathParams: map[string]string{
			"phoneNumber": c.phoneNumber,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchPhoneNumberResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
