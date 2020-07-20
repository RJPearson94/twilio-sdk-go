// This is an autogenerated file. DO NOT MODIFY
package media

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	faxSid string
	sid    string
}

type ClientProperties struct {
	FaxSid string
	Sid    string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		faxSid: properties.FaxSid,
		sid:    properties.Sid,
	}
}
