// This is an autogenerated file. DO NOT MODIFY
package permissions

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a document permission resource from the account
// See https://www.twilio.com/docs/sync/api/document-permission-resource#delete-a-document-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a document permission resource from the account
// See https://www.twilio.com/docs/sync/api/document-permission-resource#delete-a-document-permission-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Documents/{documentSid}/Permissions/{identity}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"documentSid": c.documentSid,
			"identity":    c.identity,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
