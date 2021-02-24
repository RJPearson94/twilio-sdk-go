// Package messaging_configuration contains auto-generated files. DO NOT MODIFY
package messaging_configuration

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a messaging configuration resource from the account
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a messaging configuration resource from the account
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/MessagingConfigurations/{countryCode}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"countryCode": c.countryCode,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
