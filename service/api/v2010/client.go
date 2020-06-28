package v2010

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type V2010 struct {
	client  *client.Client
	Account func(string) *account.Client
}

// Used for testing purposes only
func (s V2010) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *V2010 {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "api"
	config.APIVersion = "2010-04-01"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *V2010 {
	return &V2010{
		client:  client,
		Account: func(accountSid string) *account.Client { return account.New(client, accountSid) },
	}
}

func NewWithCredentials(creds *credentials.Credentials) *V2010 {
	return New(session.New(creds))
}
