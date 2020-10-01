// Package address contains auto-generated files. DO NOT MODIFY
package address

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchAddressResponse defines the response fields for retrieving an address
type FetchAddressResponse struct {
	AccountSid       string             `json:"account_sid"`
	City             string             `json:"City"`
	CustomerName     string             `json:"customer_name"`
	DateCreated      utils.RFC2822Time  `json:"date_created"`
	DateUpdated      *utils.RFC2822Time `json:"date_updated,omitempty"`
	EmergencyEnabled bool               `json:"emergency_enabled"`
	FriendlyName     *string            `json:"friendly_name,omitempty"`
	IsoCountry       string             `json:"iso_country"`
	PostalCode       string             `json:"postal_code"`
	Region           string             `json:"region"`
	Sid              string             `json:"sid"`
	Street           string             `json:"street"`
	StreetSecondary  *string            `json:"street_secondary,omitempty"`
	Validated        bool               `json:"validated"`
	Verified         bool               `json:"verified"`
}

// Fetch retrieves the address resource
// See https://www.twilio.com/docs/usage/api/address#fetch-an-address-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAddressResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the address resource
// See https://www.twilio.com/docs/usage/api/address#fetch-an-address-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAddressResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Addresses/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchAddressResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
