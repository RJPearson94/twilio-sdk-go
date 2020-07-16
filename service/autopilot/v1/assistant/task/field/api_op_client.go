// This is an autogenerated file. DO NOT MODIFY
package field

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	sid          string
	taskSid      string
	assistantSid string
}

type ClientProperties struct {
	Sid          string
	TaskSid      string
	AssistantSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		taskSid:      properties.TaskSid,
		assistantSid: properties.AssistantSid,
	}
}
