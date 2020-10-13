// Package available_phone_number contains auto-generated files. DO NOT MODIFY
package available_phone_number

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/local"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/mobile"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/toll_free"
)

// Client for managing available phone numbers for a specific country
// See https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource for more details
type Client struct {
	client *client.Client

	accountSid  string
	countryCode string

	Local    *local.Client
	Mobile   *mobile.Client
	TollFree *toll_free.Client
}

// ClientProperties are the properties required to manage the available phone number resources
type ClientProperties struct {
	AccountSid  string
	CountryCode string
}

// New creates a new instance of the available phone number client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:  properties.AccountSid,
		countryCode: properties.CountryCode,

		Local: local.New(client, local.ClientProperties{
			AccountSid:  properties.AccountSid,
			CountryCode: properties.CountryCode,
		}),
		Mobile: mobile.New(client, mobile.ClientProperties{
			AccountSid:  properties.AccountSid,
			CountryCode: properties.CountryCode,
		}),
		TollFree: toll_free.New(client, toll_free.ClientProperties{
			AccountSid:  properties.AccountSid,
			CountryCode: properties.CountryCode,
		}),
	}
}
