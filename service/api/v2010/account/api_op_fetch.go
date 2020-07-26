// This is an autogenerated file. DO NOT MODIFY
package account

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// Resource/ response properties for the retrieved account
type FetchAccountResponse struct {
	// The generated authorization token for the account
	AuthToken string `json:"auth_token"`
	// The date and time (in RFC2822 format) when the resource was created
	DateCreated utils.RFC2822Time `json:"date_created"`
	// The date and time (in RFC2822 format) when the resource was last updated
	DateUpdated *utils.RFC2822Time `json:"date_updated,omitempty"`
	// The human readable name of the account
	FriendlyName string `json:"friendly_name"`
	// The SID of the parent account. Can be the same as the SID of the account
	OwnerAccountSid string `json:"owner_account_sid"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The current status of the account
	Status string `json:"status"`
	// The current account type
	Type string `json:"type"`
}

// Retrieve a Twilio account (parent or sub account) resource
// See https://www.twilio.com/docs/iam/api/account#fetch-an-account-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAccountResponse, error) {
	return c.FetchWithContext(context.Background())
}

// Retrieve a Twilio account (parent or sub account) resource
// See https://www.twilio.com/docs/iam/api/account#fetch-an-account-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAccountResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{sid}.json",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchAccountResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}