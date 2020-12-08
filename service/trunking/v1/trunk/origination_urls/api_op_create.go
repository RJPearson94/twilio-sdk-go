// Package origination_urls contains auto-generated files. DO NOT MODIFY
package origination_urls

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateOriginationURLInput defines the input fields for creating a new origination url resource
type CreateOriginationURLInput struct {
	Enabled      bool   `form:"Enabled"`
	FriendlyName string `validate:"required" form:"FriendlyName"`
	Priority     int    `form:"Priority"`
	SipURL       string `validate:"required" form:"SipUrl"`
	Weight       int    `form:"Weight"`
}

// CreateOriginationURLResponse defines the response fields for the created origination url resource
type CreateOriginationURLResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Enabled      bool       `json:"enabled"`
	FriendlyName string     `json:"friendly_name"`
	Priority     int        `json:"priority"`
	Sid          string     `json:"sid"`
	SipURL       string     `json:"sip_url"`
	TrunkSid     string     `json:"trunk_sid"`
	URL          string     `json:"url"`
	Weight       int        `json:"weight"`
}

// Create creates a new origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#create-an-originationurl-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateOriginationURLInput) (*CreateOriginationURLResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#create-an-originationurl-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateOriginationURLInput) (*CreateOriginationURLResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{trunkSid}/OriginationUrls",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
		},
	}

	if input == nil {
		input = &CreateOriginationURLInput{}
	}

	response := &CreateOriginationURLResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
