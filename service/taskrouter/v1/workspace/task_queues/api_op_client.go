// This is an autogenerated file. DO NOT MODIFY
package task_queues

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
)

type Client struct {
	client       *client.Client
	workspaceSid string
}

func New(client *client.Client, workspaceSid string) *Client {
	return &Client{
		client:       client,
		workspaceSid: workspaceSid,
	}
}
