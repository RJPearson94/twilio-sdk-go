// Package available_phone_number contains auto-generated files. DO NOT MODIFY
package available_phone_number

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchCountryResponse defines the response fields for retrieving available phone numbers for a specific country
type FetchCountryResponse struct {
	Beta        bool   `json:"beta"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// Fetch retrieves the available phone number resource for a specific country
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#fetch-a-specific-country for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCountryResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the available phone number resource for a specific country
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#fetch-a-specific-country for more details
func (c Client) FetchWithContext(context context.Context) (*FetchCountryResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/AvailablePhoneNumbers/{countryCode}.json",
		PathParams: map[string]string{
			"accountSid":  c.accountSid,
			"countryCode": c.countryCode,
		},
	}

	response := &FetchCountryResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
