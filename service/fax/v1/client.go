package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/faxes"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Fax struct {
	client *client.Client
	Faxes  *faxes.Client
	Fax    func(string) *fax.Client
}

// Used for testing purposes only
func (s Fax) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Fax {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "fax"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Fax {
	return &Fax{
		client: client,
		Faxes:  faxes.New(client),
		Fax: func(sid string) *fax.Client {
			return fax.New(client, fax.ClientProperties{
				Sid: sid,
			})
		},
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Fax {
	return New(session.New(creds))
}
