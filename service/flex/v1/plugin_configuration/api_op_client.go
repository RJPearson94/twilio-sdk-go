// Package plugin_configuration contains auto-generated files. DO NOT MODIFY
package plugin_configuration

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configuration/plugin"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configuration/plugins"
)

// Client for managing a specific plugin configuration resource
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	sid string

	Plugin  func(string) *plugin.Client
	Plugins *plugins.Client
}

// ClientProperties are the properties required to manage the plugin configuration resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the plugin configuration client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Plugin: func(pluginSid string) *plugin.Client {
			return plugin.New(client, plugin.ClientProperties{
				ConfigurationSid: properties.Sid,
				Sid:              pluginSid,
			})
		},
		Plugins: plugins.New(client, plugins.ClientProperties{
			ConfigurationSid: properties.Sid,
		}),
	}
}
