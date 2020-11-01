// Package rate_limit contains auto-generated files. DO NOT MODIFY
package rate_limit

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a rate limit resource from the account
// See https://www.twilio.com/docs/verify/api/service-rate-limits#delete-a-rate-limit for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a rate limit resource from the account
// See https://www.twilio.com/docs/verify/api/service-rate-limits#delete-a-rate-limit for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/RateLimits/{sid}",
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
