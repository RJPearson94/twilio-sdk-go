// Package trunks contains auto-generated files. DO NOT MODIFY
package trunks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateTrunkInput defines the input fields for creating a new trunk resource
type CreateTrunkInput struct {
	CnamLookupEnabled      *bool   `form:"CnamLookupEnabled,omitempty"`
	DisasterRecoveryMethod *string `form:"DisasterRecoveryMethod,omitempty"`
	DisasterRecoveryURL    *string `form:"DisasterRecoveryUrl,omitempty"`
	DomainName             *string `form:"DomainName,omitempty"`
	FriendlyName           *string `form:"FriendlyName,omitempty"`
	Secure                 *bool   `form:"Secure,omitempty"`
	TransferMode           *string `form:"TransferMode,omitempty"`
}

type CreateTrunkRecordingResponse struct {
	Mode string `json:"mode"`
	Trim string `json:"trim"`
}

// CreateTrunkResponse defines the response fields for the created trunk resource
type CreateTrunkResponse struct {
	AccountSid             string                       `json:"account_sid"`
	AuthType               *string                      `json:"auth_type,omitempty"`
	AuthTypeSet            *[]string                    `json:"auth_type_set,omitempty"`
	CnamLookupEnabled      bool                         `json:"cnam_lookup_enabled"`
	DateCreated            time.Time                    `json:"date_created"`
	DateUpdated            *time.Time                   `json:"date_updated,omitempty"`
	DisasterRecoveryMethod *string                      `json:"disaster_recovery_method,omitempty"`
	DisasterRecoveryURL    *string                      `json:"disaster_recovery_url,omitempty"`
	DomainName             *string                      `json:"domain_name,omitempty"`
	FriendlyName           *string                      `json:"friendly_name,omitempty"`
	Recording              CreateTrunkRecordingResponse `json:"recording"`
	Secure                 bool                         `json:"secure"`
	Sid                    string                       `json:"sid"`
	TransferMode           string                       `json:"transfer_mode"`
	URL                    string                       `json:"url"`
}

// Create creates a new trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#create-a-trunk-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateTrunkInput) (*CreateTrunkResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#create-a-trunk-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateTrunkInput) (*CreateTrunkResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateTrunkInput{}
	}

	response := &CreateTrunkResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
