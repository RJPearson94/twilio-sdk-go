// This is an autogenerated file. DO NOT MODIFY
package keys

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client     *client.Client
	accountSid string
}

func New(client *client.Client, accountSid string) *Client {
	return &Client{
		client:     client,
		accountSid: accountSid,
	}
}