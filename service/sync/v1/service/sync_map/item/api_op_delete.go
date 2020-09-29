// Package item contains auto-generated files. DO NOT MODIFY
package item

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a map item resource from the account
// See https://www.twilio.com/docs/sync/api/map-item-resource#delete-a-mapitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a map item resource from the account
// See https://www.twilio.com/docs/sync/api/map-item-resource#delete-a-mapitem-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Maps/{syncMapSid}/Items/{key}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
			"key":        c.key,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
