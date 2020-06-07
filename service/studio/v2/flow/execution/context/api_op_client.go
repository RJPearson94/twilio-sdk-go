// This is an autogenerated file. DO NOT MODIFY
package context

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
)

type Client struct {
	client       *client.Client
	executionSid string
	flowSid      string
}

func New(client *client.Client, executionSid string, flowSid string) *Client {
	return &Client{
		client:       client,
		executionSid: executionSid,
		flowSid:      flowSid,
	}
}