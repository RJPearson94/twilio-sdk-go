// This is an autogenerated file. DO NOT MODIFY
package web_channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetWebChannelResponse struct {
	Sid         string     `json:"sid"`
	AccountSid  string     `json:"account_sid"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	URL         string     `json:"url"`
}

func (c Client) Get() (*GetWebChannelResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetWebChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/WebChannels/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &GetWebChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
