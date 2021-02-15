// Package credential_list_mapping contains auto-generated files. DO NOT MODIFY
package credential_list_mapping

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a credential list mapping from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallistmapping-resource#delete-a-sip-credentiallistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a credential list mapping from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallistmapping-resource#delete-a-sip-credentiallistmapping-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Calls/CredentialListMappings/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"domainSid":  c.domainSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
