// Package origination_url contains auto-generated files. DO NOT MODIFY
package origination_url

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchOriginationURLResponse defines the response fields for the retrieved origination url resource
type FetchOriginationURLResponse struct {
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

// Fetch retrieves an origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#fetch-an-originationurl-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchOriginationURLResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an origination url resource
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#fetch-an-originationurl-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchOriginationURLResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/OriginationUrls/{sid}",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
			"sid":      c.sid,
		},
	}

	response := &FetchOriginationURLResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
