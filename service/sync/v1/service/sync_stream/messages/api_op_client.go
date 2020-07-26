// This is an autogenerated file. DO NOT MODIFY
package messages

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	serviceSid    string
	syncStreamSid string
}

type ClientProperties struct {
	ServiceSid    string
	SyncStreamSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid:    properties.ServiceSid,
		syncStreamSid: properties.SyncStreamSid,
	}
}
