// Package entity contains auto-generated files. DO NOT MODIFY
package entity

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a entity resource from the account
// See https://www.twilio.com/docs/verify/api/entity#delete-an-entity-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a entity resource from the account
// See https://www.twilio.com/docs/verify/api/entity#delete-an-entity-resource for more details
// This resource is currently in beta and subject to change. Please use with caution
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Entities/{identity}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"identity":   c.identity,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
