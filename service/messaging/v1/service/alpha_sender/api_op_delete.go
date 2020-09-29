// Package alpha_sender contains auto-generated files. DO NOT MODIFY
package alpha_sender

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a alpha sender resource from the account
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#delete-an-alphasender-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a alpha sender resource from the account
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#delete-an-alphasender-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/AlphaSenders/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
