// This is an autogenerated file. DO NOT MODIFY
package workflow

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	sid          string
	workspaceSid string
}

type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,
	}
}
