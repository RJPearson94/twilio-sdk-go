// This is an autogenerated file. DO NOT MODIFY
package version

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/version/content"
)

type Client struct {
	client      *client.Client
	functionSid string
	serviceSid  string
	sid         string
	Content     *content.Client
}

func New(client *client.Client, functionSid string, serviceSid string, sid string) *Client {
	return &Client{
		client:      client,
		functionSid: functionSid,
		serviceSid:  serviceSid,
		sid:         sid,
		Content:     content.New(client, functionSid, serviceSid, sid),
	}
}
