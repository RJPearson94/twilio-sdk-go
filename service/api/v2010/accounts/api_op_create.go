// Package accounts contains auto-generated files. DO NOT MODIFY
package accounts

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CreateAccountInput defines the input fields for creating a new Twilio sub account resource
type CreateAccountInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// CreateAccountResponse defines the response fields for the created sub account
type CreateAccountResponse struct {
	AuthToken       string             `json:"auth_token"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	OwnerAccountSid string             `json:"owner_account_sid"`
	Sid             string             `json:"sid"`
	Status          string             `json:"status"`
	Type            string             `json:"type"`
}

// Create provisions a new Twilio sub account under the current account
// See https://www.twilio.com/docs/iam/api/account#create-an-account-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAccountInput) (*CreateAccountResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext provisions a new Twilio sub account under the current account
// See https://www.twilio.com/docs/iam/api/account#create-an-account-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAccountInput) (*CreateAccountResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts.json",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateAccountInput{}
	}

	response := &CreateAccountResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
