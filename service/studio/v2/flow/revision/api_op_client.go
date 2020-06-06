// This is an autogenerated file. DO NOT MODIFY
package revision

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
)

type Client struct {
	client         *client.Client
	flowSid        string
	revisionNumber int
}

func New(client *client.Client, flowSid string, revisionNumber int) *Client {
	return &Client{
		client:         client,
		flowSid:        flowSid,
		revisionNumber: revisionNumber,
	}
}
