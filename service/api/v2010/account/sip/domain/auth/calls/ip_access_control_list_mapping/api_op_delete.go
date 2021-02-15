// Package ip_access_control_list_mapping contains auto-generated files. DO NOT MODIFY
package ip_access_control_list_mapping

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a IP control access list mapping from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#delete-a-sip-ipaccesscontrollistmapping-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a IP control access list mapping from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource#delete-a-sip-ipaccesscontrollistmapping-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/SIP/Domains/{domainSid}/Auth/Calls/IpAccessControlListMappings/{sid}.json",
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
