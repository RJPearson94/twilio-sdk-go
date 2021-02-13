// Package credential_list contains auto-generated files. DO NOT MODIFY
package credential_list

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchCredentialListResponse defines the response fields for the retrieved credential list resource
type FetchCredentialListResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	TrunkSid     string     `json:"trunk_sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a credential list resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchCredentialListResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a credential list resource
func (c Client) FetchWithContext(context context.Context) (*FetchCredentialListResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/CredentialLists/{sid}",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
			"sid":      c.sid,
		},
	}

	response := &FetchCredentialListResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
