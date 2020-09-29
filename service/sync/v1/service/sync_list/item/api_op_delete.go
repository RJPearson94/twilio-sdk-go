// Package item contains auto-generated files. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"strconv"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a list item resource from the account
// See https://www.twilio.com/docs/sync/api/listitem-resource#delete-a-listitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a list item resource from the account
// See https://www.twilio.com/docs/sync/api/listitem-resource#delete-a-listitem-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Lists/{syncListSid}/Items/{index}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
			"index":       strconv.Itoa(c.index),
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
