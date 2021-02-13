// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a SIP credential from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#delete-a-sip-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a SIP credential from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-credential-resource#delete-a-sip-credential-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/SIP/CredentialLists/{credentialListSid}/Credentials/{sid}.json",
		PathParams: map[string]string{
			"accountSid":        c.accountSid,
			"credentialListSid": c.credentialListSid,
			"sid":               c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
