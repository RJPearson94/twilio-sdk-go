// Package trunk contains auto-generated files. DO NOT MODIFY
package trunk

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchTrunkRecordingResponse struct {
	Mode string `json:"mode"`
	Trim string `json:"trim"`
}

// FetchTrunkResponse defines the response fields for the retrieved trunk resource
type FetchTrunkResponse struct {
	AccountSid             string                      `json:"account_sid"`
	AuthType               *string                     `json:"auth_type,omitempty"`
	AuthTypeSet            *[]string                   `json:"auth_type_set,omitempty"`
	CnamLookupEnabled      bool                        `json:"cnam_lookup_enabled"`
	DateCreated            time.Time                   `json:"date_created"`
	DateUpdated            *time.Time                  `json:"date_updated,omitempty"`
	DisasterRecoveryMethod *string                     `json:"disaster_recovery_method,omitempty"`
	DisasterRecoveryURL    *string                     `json:"disaster_recovery_url,omitempty"`
	DomainName             *string                     `json:"domain_name,omitempty"`
	FriendlyName           *string                     `json:"friendly_name,omitempty"`
	Recording              FetchTrunkRecordingResponse `json:"recording"`
	Secure                 bool                        `json:"secure"`
	Sid                    string                      `json:"sid"`
	TransferMode           string                      `json:"transfer_mode"`
	URL                    string                      `json:"url"`
}

// Fetch retrieves a trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#fetch-a-trunk-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchTrunkResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#fetch-a-trunk-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchTrunkResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchTrunkResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
