// This is an autogenerated file. DO NOT MODIFY
package member

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateChannelMemberInput struct {
	RoleSid                  *string    `form:"RoleSid,omitempty"`
	LastConsumedMessageIndex *int       `form:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp *time.Time `form:"LastConsumptionTimestamp,omitempty"`
	DateCreated              *time.Time `form:"DateCreated,omitempty"`
	DateUpdated              *time.Time `form:"DateUpdated,omitempty"`
	Attributes               *string    `form:"Attributes,omitempty"`
}

type UpdateChannelMemberOutput struct {
	Sid                      string     `json:"sid"`
	AccountSid               string     `json:"account_sid"`
	ServiceSid               string     `json:"service_sid"`
	ChannelSid               string     `json:"channel_sid"`
	RoleSid                  *string    `json:"role_sid,omitempty"`
	Identity                 string     `json:"identity"`
	LastConsumedMessageIndex *int       `json:"last_consumed_message_index,omitempty"`
	LastConsumedTimestamp    *time.Time `json:"last_consumption_timestamp,omitempty"`
	Attributes               *string    `json:"attributes,omitempty"`
	DateCreated              time.Time  `json:"date_created"`
	DateUpdated              *time.Time `json:"date_updated,omitempty"`
	URL                      string     `json:"url"`
}

func (c Client) Update(input *UpdateChannelMemberInput) (*UpdateChannelMemberOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateChannelMemberInput) (*UpdateChannelMemberOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services/{serviceSid}/Channels/{channelSid}/Members/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	output := &UpdateChannelMemberOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}