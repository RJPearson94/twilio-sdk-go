// This is an autogenerated file. DO NOT MODIFY
package workflow

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
)

type Client struct {
	client       *client.Client
	sid          string
	workspaceSid string
}

func New(client *client.Client, sid string, workspaceSid string) *Client {
	return &Client{
		client:       client,
		sid:          sid,
		workspaceSid: workspaceSid,
	}
}