// Package ip_access_control_list contains auto-generated files. DO NOT MODIFY
package ip_access_control_list

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a IP control access list from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#delete-a-sip-ipaccesscontrollist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a IP control access list from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#delete-a-sip-ipaccesscontrollist-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/SIP/IpAccessControlLists/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
