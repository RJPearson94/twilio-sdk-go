// Package address contains auto-generated files. DO NOT MODIFY
package address

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateAddressInput defines input fields for updating an addresses
type UpdateAddressInput struct {
	AutoCorrectAddress *bool   `form:"AutoCorrectAddress,omitempty"`
	City               *string `form:"City,omitempty"`
	CustomerName       *string `form:"CustomerName,omitempty"`
	EmergencyEnabled   *bool   `form:"EmergencyEnabled,omitempty"`
	FriendlyName       *string `form:"FriendlyName,omitempty"`
	PostalCode         *string `form:"PostalCode,omitempty"`
	Region             *string `form:"Region,omitempty"`
	Street             *string `form:"Street,omitempty"`
	StreetSecondary    *string `form:"StreetSecondary,omitempty"`
}

// UpdateAddressResponse defines the response fields for the updated addresses
type UpdateAddressResponse struct {
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

// Update modifies an address resource
// See https://www.twilio.com/docs/usage/api/address#update-an-address-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateAddressInput) (*UpdateAddressResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an address resource
// See https://www.twilio.com/docs/usage/api/address#update-an-address-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateAddressInput) (*UpdateAddressResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Addresses/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateAddressInput{}
	}

	response := &UpdateAddressResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
