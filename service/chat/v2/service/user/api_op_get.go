// This is an autogenerated file. DO NOT MODIFY
package user

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetUserOutput struct {
	Sid                 string     `json:"sid"`
	AccountSid          string     `json:"account_sid"`
	ServiceSid          string     `json:"service_sid"`
	RoleSid             string     `json:"role_sid"`
	Identity            string     `json:"identity"`
	Attributes          *string    `json:"attributes,omitempty"`
	FriendlyName        *string    `json:"friendly_name,omitempty"`
	IsOnline            *bool      `json:"is_online,omitempty"`
	IsNotifiable        *bool      `json:"is_notifiable,omitempty"`
	JoinedChannelsCount *int       `json:"joined_channels_count,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	URL                 string     `json:"url"`
}

func (c Client) Get() (*GetUserOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetUserOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Services/{serviceSid}/Users/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	output := &GetUserOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}