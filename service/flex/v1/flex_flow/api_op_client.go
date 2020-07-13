// This is an autogenerated file. DO NOT MODIFY
package flex_flow

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	sid string
}

type ClientProperties struct {
	Sid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
