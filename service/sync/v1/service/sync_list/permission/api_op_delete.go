// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a list permission resource from the account
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#delete-a-sync-list-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a list permission resource from the account
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#delete-a-sync-list-permission-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Lists/{syncListSid}/Permissions/{identity}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
			"identity":    c.identity,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
