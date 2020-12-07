// Package trunk contains auto-generated files. DO NOT MODIFY
package trunk

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateTrunkInput defines input fields for updating a trunk resource
type UpdateTrunkInput struct {
	CnamLookupEnabled      *bool   `form:"CnamLookupEnabled,omitempty"`
	DisasterRecoveryMethod *string `form:"DisasterRecoveryMethod,omitempty"`
	DisasterRecoveryURL    *string `form:"DisasterRecoveryUrl,omitempty"`
	DomainName             *string `form:"DomainName,omitempty"`
	FriendlyName           *string `form:"FriendlyName,omitempty"`
	Secure                 *bool   `form:"Secure,omitempty"`
	TransferMode           *string `form:"TransferMode,omitempty"`
}

type UpdateTrunkRecordingResponse struct {
	Mode string `json:"mode"`
	Trim string `json:"trim"`
}

// UpdateTrunkResponse defines the response fields for the updated trunk
type UpdateTrunkResponse struct {
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
	Recording              UpdateTrunkRecordingResponse `json:"recording"`
	Secure                 bool                         `json:"secure"`
	Sid                    string                       `json:"sid"`
	TransferMode           string                       `json:"transfer_mode"`
	URL                    string                       `json:"url"`
}

// Update modifies a trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#update-a-trunk-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateTrunkInput) (*UpdateTrunkResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a trunk resource
// See https://www.twilio.com/docs/sip-trunking/api/trunk-resource#update-a-trunk-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateTrunkInput) (*UpdateTrunkResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Trunks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateTrunkInput{}
	}

	response := &UpdateTrunkResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
