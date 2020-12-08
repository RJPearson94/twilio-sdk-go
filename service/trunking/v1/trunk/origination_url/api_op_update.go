// Package origination_url contains auto-generated files. DO NOT MODIFY
package origination_url

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateOriginationURLInput defines input fields for updating an origination url resource
type UpdateOriginationURLInput struct {
	Enabled      *bool   `form:"Enabled,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	Priority     *int    `form:"Priority,omitempty"`
	SipURL       *string `form:"SipUrl,omitempty"`
	Weight       *int    `form:"Weight,omitempty"`
}

// UpdateOriginationURLResponse defines the response fields for the updated origination url
type UpdateOriginationURLResponse struct {
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

// Update modifies an origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#update-an-originationurl-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateOriginationURLInput) (*UpdateOriginationURLResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#update-an-originationurl-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateOriginationURLInput) (*UpdateOriginationURLResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{trunkSid}/OriginationUrls/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
			"sid":      c.sid,
		},
	}

	if input == nil {
		input = &UpdateOriginationURLInput{}
	}

	response := &UpdateOriginationURLResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
