package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channels"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Flex struct {
	client        *client.Client
	Configuration func() *configuration.Client
	FlexFlows     *flex_flows.Client
	FlexFlow      func(string) *flex_flow.Client
	Channels      *channels.Client
	Channel       func(string) *channel.Client
	WebChannels   *web_channels.Client
	WebChannel    func(string) *web_channel.Client
}

// Used for testing purposes only
func (s Flex) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *Flex {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "flex-api"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *Flex {
	return &Flex{
		client:        client,
		Configuration: func() *configuration.Client { return configuration.New(client) },
		FlexFlows:     flex_flows.New(client),
		FlexFlow:      func(sid string) *flex_flow.Client { return flex_flow.New(client, sid) },
		Channels:      channels.New(client),
		Channel:       func(sid string) *channel.Client { return channel.New(client, sid) },
		WebChannels:   web_channels.New(client),
		WebChannel:    func(sid string) *web_channel.Client { return web_channel.New(client, sid) },
	}
}

func NewWithCredentials(creds *credentials.Credentials) *Flex {
	return New(session.New(creds))
}
