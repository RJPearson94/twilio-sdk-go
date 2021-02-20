// Package composition_hook contains auto-generated files. DO NOT MODIFY
package composition_hook

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a composition hook resource from the account
// See https://www.twilio.com/docs/video/api/composition-hooks#hk-delete for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a composition hook resource from the account
// See https://www.twilio.com/docs/video/api/composition-hooks#hk-delete for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/CompositionHooks/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
