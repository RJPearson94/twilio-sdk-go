// This is an autogenerated file. DO NOT MODIFY
package account

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
)

type Client struct {
	client *client.Client
	sid    string
	Keys   *keys.Client
	Key    func(string) *key.Client
}

func New(client *client.Client, sid string) *Client {
	return &Client{
		client: client,
		sid:    sid,
		Keys:   keys.New(client, sid),
		Key:    func(keySid string) *key.Client { return key.New(client, sid, keySid) },
	}
}