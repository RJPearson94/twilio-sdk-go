// Package short_code contains auto-generated files. DO NOT MODIFY
package short_code

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a short code resource from the the proxy service
// See https://www.twilio.com/docs/proxy/api/short-code#remove-a-short-code-from-a-proxy-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a short code resource from the the proxy service
// See https://www.twilio.com/docs/proxy/api/short-code#remove-a-short-code-from-a-proxy-service for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/ShortCodes/{sid}",
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
