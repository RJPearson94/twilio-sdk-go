// Package credential_list contains auto-generated files. DO NOT MODIFY
package credential_list

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a credential list resource from the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource#delete-a-credentiallist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a credential list resource from the SIP trunk
// See https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource#delete-a-credentiallist-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Trunks/{trunkSid}/CredentialLists/{sid}",
		PathParams: map[string]string{
			"trunkSid": c.trunkSid,
			"sid":      c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
