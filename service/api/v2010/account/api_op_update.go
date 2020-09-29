// Package account contains auto-generated files. DO NOT MODIFY
package account

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// UpdateAccountInput defines input fields for updating an account resource
type UpdateAccountInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
	Status       *string `form:"Status,omitempty"`
}

// UpdateAccountResponse defines the response fields for the updated account
type UpdateAccountResponse struct {
	AuthToken       string             `json:"auth_token"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	OwnerAccountSid string             `json:"owner_account_sid"`
	Sid             string             `json:"sid"`
	Status          string             `json:"status"`
	Type            string             `json:"type"`
}

// Update modifies a Twilio Account (parent or sub account) resource
// See https://www.twilio.com/docs/iam/api/account#update-an-account-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateAccountInput) (*UpdateAccountResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a Twilio Account (parent or sub account) resource
// See https://www.twilio.com/docs/iam/api/account#update-an-account-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateAccountInput) (*UpdateAccountResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateAccountInput{}
	}

	response := &UpdateAccountResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
