// Package conversation contains auto-generated files. DO NOT MODIFY
package conversation

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a conversation resource from the account
// See https://www.twilio.com/docs/conversations/api/conversation-resource#delete-a-conversation-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a conversation resource from the account
// See https://www.twilio.com/docs/conversations/api/conversation-resource#delete-a-conversation-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Conversations/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
