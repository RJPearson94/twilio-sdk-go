package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin/versions"
	configurationPlugins "github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configuration/plugins"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configurations"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_releases"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugins"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Flex Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	flexSession := twilio.NewWithCredentials(creds).Flex.V1

	Describe("Given the flex configuration client", func() {
		It("Then the configuration is fetched and updated", func() {
			configurationClient := flexSession.Configuration()

			fetchResp, fetchErr := configurationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := configurationClient.Update(&configuration.UpdateConfigurationInput{
				AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the flex flow clients", func() {

		It("Then the flex flow is created, fetched, updated and deleted", func() {
			flexFlowsClient := flexSession.FlexFlows

			createResp, createErr := flexFlowsClient.Create(&flex_flows.CreateFlexFlowInput{
				FriendlyName:    uuid.New().String(),
				ChatServiceSid:  os.Getenv("TWILIO_FLEX_CHANNEL_SERVICE_SID"),
				ChannelType:     "web",
				IntegrationType: utils.String("external"),
				Integration: &flex_flows.CreateFlexFlowIntegrationInput{
					URL: utils.String("https://test.com/external"),
				},
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := flexFlowsClient.Page(&flex_flows.FlexFlowsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.FlexFlows)).Should(BeNumerically(">=", 1))

			paginator := flexFlowsClient.NewFlexFlowsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.FlexFlows)).Should(BeNumerically(">=", 1))

			flexFlowClient := flexSession.FlexFlow(createResp.Sid)

			fetchResp, fetchErr := flexFlowClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := flexFlowClient.Update(&flex_flow.UpdateFlexFlowInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := flexFlowClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the flex channel clients", func() {

		var flexFlowSid string

		BeforeEach(func() {
			flexFlowResp, flexFlowErr := flexSession.FlexFlows.Create(&flex_flows.CreateFlexFlowInput{
				FriendlyName:    uuid.New().String(),
				ChatServiceSid:  os.Getenv("TWILIO_FLEX_CHANNEL_SERVICE_SID"),
				ChannelType:     "web",
				IntegrationType: utils.String("external"),
				Integration: &flex_flows.CreateFlexFlowIntegrationInput{
					URL: utils.String("https://test.com/external"),
				},
			})
			if flexFlowErr != nil {
				Fail(fmt.Sprintf("Failed to create flex flow. Error %s", flexFlowErr.Error()))
			}
			flexFlowSid = flexFlowResp.Sid
		})

		AfterEach(func() {
			if err := flexSession.FlexFlow(flexFlowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete flex flow. Error %s", err.Error()))
			}
		})

		It("Then the channel is created, fetched and deleted", func() {
			channelsClient := flexSession.Channels

			createResp, createErr := channelsClient.Create(&channels.CreateChannelInput{
				FlexFlowSid:          flexFlowSid,
				Identity:             uuid.New().String(),
				ChatUserFriendlyName: uuid.New().String(),
				ChatFriendlyName:     uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := channelsClient.Page(&channels.ChannelsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Channels)).Should(BeNumerically(">=", 1))

			paginator := channelsClient.NewChannelsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Channels)).Should(BeNumerically(">=", 1))

			channelClient := flexSession.Channel(createResp.Sid)

			fetchResp, fetchErr := channelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := channelClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the plugin clients", func() {
		It("Then the plugin is created, fetched and updated", func() {
			pluginsClient := flexSession.Plugins

			createResp, createErr := pluginsClient.Create(&plugins.CreatePluginInput{
				UniqueName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := pluginsClient.Page(&plugins.PluginsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Plugins)).Should(BeNumerically(">=", 1))

			paginator := pluginsClient.NewPluginsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Plugins)).Should(BeNumerically(">=", 1))

			pluginClient := flexSession.Plugin(createResp.Sid)

			fetchResp, fetchErr := pluginClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := pluginClient.Update(&plugin.UpdatePluginInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the flex plugin version clients", func() {

		var pluginSid string

		BeforeEach(func() {
			pluginResp, pluginErr := flexSession.Plugins.Create(&plugins.CreatePluginInput{
				UniqueName: uuid.New().String(),
			})
			if pluginErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin. Error %s", pluginErr.Error()))
			}
			pluginSid = pluginResp.Sid
		})

		It("Then the version is created and fetched", func() {
			versionsClient := flexSession.Plugin(pluginSid).Versions

			createResp, createErr := versionsClient.Create(&versions.CreateVersionInput{
				Version:   "1.0.0",
				PluginURL: "https://example.com",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := versionsClient.Page(&versions.VersionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Versions)).Should(BeNumerically(">=", 1))

			paginator := versionsClient.NewVersionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Versions)).Should(BeNumerically(">=", 1))

			versionClient := flexSession.Plugin(pluginSid).Version(createResp.Sid)

			fetchResp, fetchErr := versionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the flex plugin configuration clients", func() {

		var pluginSid string
		var pluginVersionSid string

		BeforeEach(func() {
			pluginResp, pluginErr := flexSession.Plugins.Create(&plugins.CreatePluginInput{
				UniqueName: uuid.New().String(),
			})
			if pluginErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin. Error %s", pluginErr.Error()))
			}
			pluginSid = pluginResp.Sid

			versionResp, versionErr := flexSession.Plugin(pluginSid).Versions.Create(&versions.CreateVersionInput{
				Version:   "1.0.0",
				PluginURL: "https://example.com",
			})
			if versionErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin version. Error %s", versionErr.Error()))
			}
			pluginVersionSid = versionResp.Sid
		})

		It("Then the plugin configuration is created and fetched", func() {
			pluginConfigurationsClient := flexSession.PluginConfigurations

			createResp, createErr := pluginConfigurationsClient.Create(&plugin_configurations.CreateConfigurationInput{
				Name: uuid.New().String(),
				Plugins: &[]string{
					fmt.Sprintf(`{"plugin_version": "%s"}`, pluginVersionSid),
				},
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := pluginConfigurationsClient.Page(&plugin_configurations.ConfigurationsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Configurations)).Should(BeNumerically(">=", 1))

			paginator := pluginConfigurationsClient.NewConfigurationsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Configurations)).Should(BeNumerically(">=", 1))

			pluginConfigurationClient := flexSession.PluginConfiguration(createResp.Sid)

			fetchResp, fetchErr := pluginConfigurationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the flex configuration plugin clients", func() {

		var pluginSid string
		var pluginVersionSid string
		var configurationSid string

		BeforeEach(func() {
			pluginResp, pluginErr := flexSession.Plugins.Create(&plugins.CreatePluginInput{
				UniqueName: uuid.New().String(),
			})
			if pluginErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin. Error %s", pluginErr.Error()))
			}
			pluginSid = pluginResp.Sid

			versionResp, versionErr := flexSession.Plugin(pluginSid).Versions.Create(&versions.CreateVersionInput{
				Version:   "1.0.0",
				PluginURL: "https://example.com",
			})
			if versionErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin version. Error %s", versionErr.Error()))
			}
			pluginVersionSid = versionResp.Sid

			configurationResp, configurationErr := flexSession.PluginConfigurations.Create(&plugin_configurations.CreateConfigurationInput{
				Name: uuid.New().String(),
				Plugins: &[]string{
					fmt.Sprintf(`{"plugin_version": "%s"}`, pluginVersionSid),
				},
			})
			if configurationErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin configuration. Error %s", configurationErr.Error()))
			}
			configurationSid = configurationResp.Sid
		})

		It("Then the plugin is fetched", func() {
			pluginsClient := flexSession.PluginConfiguration(configurationSid).Plugins

			pageResp, pageErr := pluginsClient.Page(&configurationPlugins.PluginsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Plugins)).Should(BeNumerically(">=", 1))

			paginator := pluginsClient.NewPluginsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Plugins)).Should(BeNumerically(">=", 1))

			pluginClient := flexSession.PluginConfiguration(configurationSid).Plugin(paginator.Plugins[0].PluginSid)

			fetchResp, fetchErr := pluginClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the flex plugin release clients", func() {

		var pluginSid string
		var pluginVersionSid string
		var configurationSid string

		BeforeEach(func() {
			pluginResp, pluginErr := flexSession.Plugins.Create(&plugins.CreatePluginInput{
				UniqueName: uuid.New().String(),
			})
			if pluginErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin. Error %s", pluginErr.Error()))
			}
			pluginSid = pluginResp.Sid

			versionResp, versionErr := flexSession.Plugin(pluginSid).Versions.Create(&versions.CreateVersionInput{
				Version:   "1.0.0",
				PluginURL: "https://example.com",
			})
			if versionErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin version. Error %s", versionErr.Error()))
			}
			pluginVersionSid = versionResp.Sid

			configurationResp, configurationErr := flexSession.PluginConfigurations.Create(&plugin_configurations.CreateConfigurationInput{
				Name: uuid.New().String(),
				Plugins: &[]string{
					fmt.Sprintf(`{"plugin_version": "%s"}`, pluginVersionSid),
				},
			})
			if configurationErr != nil {
				Fail(fmt.Sprintf("Failed to create flex plugin configuration. Error %s", configurationErr.Error()))
			}
			configurationSid = configurationResp.Sid
		})

		It("Then the release is created and fetched", func() {
			pluginReleasesClient := flexSession.PluginReleases

			createResp, createErr := pluginReleasesClient.Create(&plugin_releases.CreateReleaseInput{
				ConfigurationId: configurationSid,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := pluginReleasesClient.Page(&plugin_releases.ReleasesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Releases)).Should(BeNumerically(">=", 1))

			paginator := pluginReleasesClient.NewReleasesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Releases)).Should(BeNumerically(">=", 1))

			pluginReleaseClient := flexSession.PluginRelease(createResp.Sid)

			fetchResp, fetchErr := pluginReleaseClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	// TODO add web channel support
})
