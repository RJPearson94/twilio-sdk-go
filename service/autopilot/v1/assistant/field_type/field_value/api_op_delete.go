// Package field_value contains auto-generated files. DO NOT MODIFY
package field_value

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a field type resource from the account
// See https://www.twilio.com/docs/autopilot/api/field-value#delete-a-fieldvalue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a field type resource from the account
// See https://www.twilio.com/docs/autopilot/api/field-value#delete-a-fieldvalue-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"fieldTypeSid": c.fieldTypeSid,
			"sid":          c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
