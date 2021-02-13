// Package ip_access_control_list contains auto-generated files. DO NOT MODIFY
package ip_access_control_list

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchIpAccessControlListResponse defines the response fields for the retrieved IP access control list resource
type FetchIpAccessControlListResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	TrunkSid     string     `json:"trunk_sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a IP access control list resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchIpAccessControlListResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a IP access control list resource
func (c Client) FetchWithContext(context context.Context) (*FetchIpAccessControlListResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Trunks/{trunkSid}/IpAccessControlLists/{sid}",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
			"sid":      c.sid,
		},
	}

	response := &FetchIpAccessControlListResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
