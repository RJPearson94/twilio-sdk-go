// Package available_phone_numbers contains auto-generated files. DO NOT MODIFY
package available_phone_numbers

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type PageCountryResponse struct {
	Beta        bool   `json:"beta"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// CountriesPageResponse defines the response fields for the available phone number countries page
type CountriesPageResponse struct {
	Countries []PageCountryResponse `json:"countries"`
	URI       string                `json:"uri"`
}

// Page retrieves a page of countries with available phone numbers
// See https://www.twilio.com/docs/iam/keys/api-key-resource#read-a-key-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page() (*CountriesPageResponse, error) {
	return c.PageWithContext(context.Background())
}

// PageWithContext retrieves a page of countries with available phone numbers
// See https://www.twilio.com/docs/iam/keys/api-key-resource#read-a-key-resource for more details
func (c Client) PageWithContext(context context.Context) (*CountriesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/AvailablePhoneNumbers.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	response := &CountriesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
