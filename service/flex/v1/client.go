// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configurations"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_release"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_releases"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugins"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channels"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Flex client is used to manage resources for Twilio Flex
// See https://www.twilio.com/docs/flex for more details
type Flex struct {
	client *client.Client

	Channel              func(string) *channel.Client
	Channels             *channels.Client
	Configuration        func() *configuration.Client
	FlexFlow             func(string) *flex_flow.Client
	FlexFlows            *flex_flows.Client
	Plugin               func(string) *plugin.Client
	PluginConfiguration  func(string) *plugin_configuration.Client
	PluginConfigurations *plugin_configurations.Client
	PluginRelease        func(string) *plugin_release.Client
	PluginReleases       *plugin_releases.Client
	Plugins              *plugins.Client
	WebChannel           func(string) *web_channel.Client
	WebChannels          *web_channels.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Flex {
	return &Flex{
		client: client,

		Channel: func(channelSid string) *channel.Client {
			return channel.New(client, channel.ClientProperties{
				Sid: channelSid,
			})
		},
		Channels:      channels.New(client),
		Configuration: func() *configuration.Client { return configuration.New(client) },
		FlexFlow: func(flexFlowSid string) *flex_flow.Client {
			return flex_flow.New(client, flex_flow.ClientProperties{
				Sid: flexFlowSid,
			})
		},
		FlexFlows: flex_flows.New(client),
		Plugin: func(pluginSid string) *plugin.Client {
			return plugin.New(client, plugin.ClientProperties{
				Sid: pluginSid,
			})
		},
		PluginConfiguration: func(configurationSid string) *plugin_configuration.Client {
			return plugin_configuration.New(client, plugin_configuration.ClientProperties{
				Sid: configurationSid,
			})
		},
		PluginConfigurations: plugin_configurations.New(client),
		PluginRelease: func(releaseSid string) *plugin_release.Client {
			return plugin_release.New(client, plugin_release.ClientProperties{
				Sid: releaseSid,
			})
		},
		PluginReleases: plugin_releases.New(client),
		Plugins:        plugins.New(client),
		WebChannel: func(webChannelSid string) *web_channel.Client {
			return web_channel.New(client, web_channel.ClientProperties{
				Sid: webChannelSid,
			})
		},
		WebChannels: web_channels.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Flex) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Flex {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "flex-api"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
