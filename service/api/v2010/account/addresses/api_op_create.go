// Package addresses contains auto-generated files. DO NOT MODIFY
package addresses

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateAddressInput defines input fields for creating a new address
type CreateAddressInput struct {
	AutoCorrectAddress *bool   `form:"AutoCorrectAddress,omitempty"`
	City               string  `validate:"required" form:"City"`
	CustomerName       string  `validate:"required" form:"CustomerName"`
	EmergencyEnabled   *bool   `form:"EmergencyEnabled,omitempty"`
	FriendlyName       *string `form:"FriendlyName,omitempty"`
	IsoCountry         string  `validate:"required" form:"IsoCountry"`
	PostalCode         string  `validate:"required" form:"PostalCode"`
	Region             string  `validate:"required" form:"Region"`
	Street             string  `validate:"required" form:"Street"`
	StreetSecondary    *string `form:"StreetSecondary,omitempty"`
}

// CreateAddressResponse defines the response fields for creating a new address
type CreateAddressResponse struct {
	AccountSid       string             `json:"account_sid"`
	City             string             `json:"city"`
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

// Create creates a new address resource
// See https://www.twilio.com/docs/usage/api/address#create-an-address-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAddressInput) (*CreateAddressResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new address resource
// See https://www.twilio.com/docs/usage/api/address#create-an-address-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAddressInput) (*CreateAddressResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Addresses.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateAddressInput{}
	}

	response := &CreateAddressResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
