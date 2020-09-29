// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateSyncListPermissionsInput defines input fields for updating a list permission resource
type UpdateSyncListPermissionsInput struct {
	Manage bool `form:"Manage"`
	Read   bool `form:"Read"`
	Write  bool `form:"Write"`
}

// UpdateSyncListPermissionsResponse defines the response fields for the updated list permission
type UpdateSyncListPermissionsResponse struct {
	AccountSid string `json:"account_sid"`
	Identity   string `json:"identity"`
	ListSid    string `json:"list_sid"`
	Manage     bool   `json:"manage"`
	Read       bool   `json:"read"`
	ServiceSid string `json:"service_sid"`
	URL        string `json:"url"`
	Write      bool   `json:"write"`
}

// Update modifies a list permission resource
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#update-a-sync-list-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncListPermissionsInput) (*UpdateSyncListPermissionsResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a list permission resource
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#update-a-sync-list-permission-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncListPermissionsInput) (*UpdateSyncListPermissionsResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists/{syncListSid}/Permissions/{identity}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
			"identity":    c.identity,
		},
	}

	if input == nil {
		input = &UpdateSyncListPermissionsInput{}
	}

	response := &UpdateSyncListPermissionsResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
