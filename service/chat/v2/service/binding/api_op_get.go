// This is an autogenerated file. DO NOT MODIFY
package binding

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetBindingOutput struct {
	Sid           string     `json:"sid"`
	AccountSid    string     `json:"account_sid"`
	CredentialSid *string    `json:"credential_sid,omitempty"`
	ServiceSid    string     `json:"service_sid"`
	BindingType   *string    `json:"binding_type,omitempty"`
	Endpoint      *string    `json:"endpoint,omitempty"`
	Identity      *string    `json:"identity,omitempty"`
	MessageTypes  *[]string  `json:"message_types,omitempty"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	URL           string     `json:"url"`
}

func (c Client) Get() (*GetBindingOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetBindingOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Services/{serviceSid}/Bindings/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	output := &GetBindingOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}