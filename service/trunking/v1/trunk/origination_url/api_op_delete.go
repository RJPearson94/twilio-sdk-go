// Package origination_url contains auto-generated files. DO NOT MODIFY
package origination_url

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a origination url resource from the account
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#delete-an-originationurl-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a origination url resource from the account
// See https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#delete-an-originationurl-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Trunks/{trunkSid}/OriginationUrls/{sid}",
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
