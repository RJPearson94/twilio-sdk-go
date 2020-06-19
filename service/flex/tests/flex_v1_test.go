package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/flex"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channels"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Flex V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	flexSession := flex.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(flexSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a Configuration Client", func() {
		configurationClient := flexSession.Configuration()

		Describe("When the configuration is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Configuration",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := configurationClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexServiceInstanceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceInstanceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceInstanceSid).To(BeNil())
				Expect(resp.CrmEnabled).To(BeNil())
				Expect(resp.CrmType).To(BeNil())
				Expect(resp.CrmCallbackURL).To(BeNil())
				Expect(resp.CrmFallbackURL).To(BeNil())
				Expect(resp.CrmAttributes).To(BeNil())
				Expect(resp.UiLanguage).To(BeNil())

				uiAttributesFixture, _ := ioutil.ReadFile("testdata/uiAttributes.json")
				uiAttributesResp := make(map[string]interface{})
				json.Unmarshal(uiAttributesFixture, &uiAttributesResp)

				Expect(resp.UiAttributes).To(Equal(utils.Interface(uiAttributesResp)))
				Expect(resp.UiDependencies).To(BeNil())
				Expect(resp.UiVersion).To(Equal("~1.19.0"))
				Expect(resp.TaskRouterWorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskRouterTargetWorkflowSid).To(Equal("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskRouterTargetTaskQueueSid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskRouterTaskQueues).To(BeNil())
				Expect(resp.TaskRouterSkills).To(BeNil())
				Expect(resp.TaskRouterWorkerChannels).To(BeNil())
				Expect(resp.TaskRouterWorkerAttributes).To(BeNil())
				Expect(resp.TaskRouterOfflineActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallRecordingEnabled).To(BeNil())
				Expect(resp.CallRecordingWebhookURL).To(BeNil())
				Expect(resp.PublicAttributes).To(BeNil())

				attributesFixture, _ := ioutil.ReadFile("testdata/attributes.json")
				attributesResp := make(map[string]interface{})
				json.Unmarshal(attributesFixture, &attributesResp)
				Expect(resp.Attributes).To(Equal(utils.Interface(attributesResp)))

				Expect(resp.Status).To(Equal("ok"))
				Expect(resp.RuntimeDomain).To(Equal("test.twil.io"))
				Expect(resp.ServiceVersion).To(BeNil())
				Expect(resp.PluginServiceEnabled).To(BeNil())
				Expect(resp.PluginServiceAttributes).To(BeNil())
				Expect(resp.Integrations).To(BeNil())
				Expect(resp.WfmIntegrations).To(BeNil())
				Expect(resp.OutboundCallFlows).To(BeNil())
				Expect(resp.QueueStatsConfiguration).To(BeNil())

				Expect(resp.ServerlessServiceSids).ToNot(BeNil())

				serverlessServiceSids := *resp.ServerlessServiceSids
				Expect(len(serverlessServiceSids)).To(Equal(1))
				Expect(serverlessServiceSids[0]).To(Equal("SZXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-05-10T20:17:41Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/Configuration"))
			})
		})

		Describe("When the get configuration response returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Configuration",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := configurationClient.Get()
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the get configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Configuration is successfully updated", func() {

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/Configuration",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateConfigurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &configuration.UpdateConfigurationInput{
				AccountSid:            "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				ServerlessServiceSids: nil,
			}

			resp, err := configurationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexServiceInstanceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceInstanceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceInstanceSid).To(BeNil())
				Expect(resp.CrmEnabled).To(BeNil())
				Expect(resp.CrmType).To(BeNil())
				Expect(resp.CrmCallbackURL).To(BeNil())
				Expect(resp.CrmFallbackURL).To(BeNil())
				Expect(resp.CrmAttributes).To(BeNil())
				Expect(resp.UiLanguage).To(BeNil())

				uiAttributesFixture, _ := ioutil.ReadFile("testdata/uiAttributes.json")
				uiAttributesResp := make(map[string]interface{})
				json.Unmarshal(uiAttributesFixture, &uiAttributesResp)

				Expect(resp.UiAttributes).To(Equal(utils.Interface(uiAttributesResp)))
				Expect(resp.UiDependencies).To(BeNil())
				Expect(resp.UiVersion).To(Equal("~1.19.0"))
				Expect(resp.TaskRouterWorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskRouterTargetWorkflowSid).To(Equal("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskRouterTargetTaskQueueSid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskRouterTaskQueues).To(BeNil())
				Expect(resp.TaskRouterSkills).To(BeNil())
				Expect(resp.TaskRouterWorkerChannels).To(BeNil())
				Expect(resp.TaskRouterWorkerAttributes).To(BeNil())
				Expect(resp.TaskRouterOfflineActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CallRecordingEnabled).To(BeNil())
				Expect(resp.CallRecordingWebhookURL).To(BeNil())
				Expect(resp.PublicAttributes).To(BeNil())

				attributesFixture, _ := ioutil.ReadFile("testdata/attributes.json")
				attributesResp := make(map[string]interface{})
				json.Unmarshal(attributesFixture, &attributesResp)
				Expect(resp.Attributes).To(Equal(utils.Interface(attributesResp)))

				Expect(resp.Status).To(Equal("ok"))
				Expect(resp.RuntimeDomain).To(Equal("test.twil.io"))
				Expect(resp.ServiceVersion).To(BeNil())
				Expect(resp.PluginServiceEnabled).To(BeNil())
				Expect(resp.PluginServiceAttributes).To(BeNil())
				Expect(resp.Integrations).To(BeNil())
				Expect(resp.WfmIntegrations).To(BeNil())
				Expect(resp.OutboundCallFlows).To(BeNil())
				Expect(resp.QueueStatsConfiguration).To(BeNil())
				Expect(resp.ServerlessServiceSids).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-05-10T20:17:41Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-05-11T20:17:41Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/Configuration"))
			})
		})

		Describe("When the update configuration request does not contain a account sid", func() {
			updateInput := &configuration.UpdateConfigurationInput{}

			resp, err := configurationClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the flex flows client", func() {
		flexFlowsClient := flexSession.FlexFlows

		Describe("When the flex flow is successfully created", func() {
			createInput := &flex_flows.CreateFlexFlowInput{
				FriendlyName:   "Test Flex Flow",
				ChatServiceSid: "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				ChannelType:    "web",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/FlexFlows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flexFlowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := flexFlowsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create flex flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test Flex Flow"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelType).To(Equal("web"))
				Expect(resp.ContactIdentity).To(Equal(utils.String("12345")))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.IntegrationType).To(Equal(utils.String("studio")))
				Expect(resp.Integration).To(Equal(&flex_flows.CreateFlexFlowOutputIntegration{
					RetryCount: utils.Int(1),
					FlowSid:    utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.LongLived).To(Equal(utils.Bool(true)))
				Expect(resp.JanitorEnabled).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the flex flow request does not contain a friendly name", func() {
			createInput := &flex_flows.CreateFlexFlowInput{
				ChatServiceSid: "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				ChannelType:    "web",
			}

			resp, err := flexFlowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the flex flow request does not contain a chat service sid", func() {
			createInput := &flex_flows.CreateFlexFlowInput{
				FriendlyName: "Test Flex Flow",
				ChannelType:  "web",
			}

			resp, err := flexFlowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the flex flow request does not contain a channel type", func() {
			createInput := &flex_flows.CreateFlexFlowInput{
				FriendlyName:   "Test Flex Flow",
				ChatServiceSid: "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			resp, err := flexFlowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the flex flow api returns a 500 response", func() {
			createInput := &flex_flows.CreateFlexFlowInput{
				FriendlyName:   "Test Flex Flow",
				ChatServiceSid: "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				ChannelType:    "web",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/FlexFlows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := flexFlowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a flex flow sid", func() {
		flexFlowClient := flexSession.FlexFlow("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the flex flow is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flexFlowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flexFlowClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get flex flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test Flex Flow"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelType).To(Equal("web"))
				Expect(resp.ContactIdentity).To(Equal(utils.String("12345")))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.IntegrationType).To(Equal(utils.String("studio")))
				Expect(resp.Integration).To(Equal(&flex_flow.GetFlexFlowOutputIntegration{
					RetryCount: utils.Int(1),
					FlowSid:    utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.LongLived).To(Equal(utils.Bool(true)))
				Expect(resp.JanitorEnabled).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get flex flow response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows/FO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.FlexFlow("FO71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the flex flow is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedFlexFlowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &flex_flow.UpdateFlexFlowInput{
				FriendlyName: "New Flex Flow",
			}

			resp, err := flexFlowClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the updated flex flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("New Flex Flow"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelType).To(Equal("web"))
				Expect(resp.ContactIdentity).To(Equal(utils.String("12345")))
				Expect(resp.Enabled).To(Equal(true))
				Expect(resp.IntegrationType).To(Equal(utils.String("studio")))
				Expect(resp.Integration).To(Equal(&flex_flow.UpdateFlexFlowOutputIntegration{
					RetryCount: utils.Int(1),
					FlowSid:    utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.LongLived).To(Equal(utils.Bool(true)))
				Expect(resp.JanitorEnabled).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2016-08-02T22:10:40Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update flex flow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/FlexFlows/FO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &flex_flow.UpdateFlexFlowInput{}

			resp, err := flexSession.FlexFlow("FO71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the flex flow is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := flexFlowClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete flex flow response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://flex-api.twilio.com/v1/FlexFlows/FO71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := flexSession.FlexFlow("FO71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the channels client", func() {
		channelsClient := flexSession.Channels

		Describe("When the channel is successfully created", func() {
			createInput := &channels.CreateChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:             "Test",
				ChatUserFriendlyName: "Test",
				ChatFriendlyName:     "Test",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := channelsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal(utils.String("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channels request does not contain a flex flow sid", func() {
			createInput := &channels.CreateChannelInput{
				Identity:             "Test",
				ChatUserFriendlyName: "Test",
				ChatFriendlyName:     "Test",
			}

			resp, err := channelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channels request does not contain a identity", func() {
			createInput := &channels.CreateChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				ChatUserFriendlyName: "Test",
				ChatFriendlyName:     "Test",
			}

			resp, err := channelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel request does not contain a chat user friendly name", func() {
			createInput := &channels.CreateChannelInput{
				FlexFlowSid:      "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:         "Test",
				ChatFriendlyName: "Test",
			}

			resp, err := channelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel request does not contain a chat friendly name", func() {
			createInput := &channels.CreateChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:             "Test",
				ChatUserFriendlyName: "Test",
			}

			resp, err := channelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel api returns a 500 response", func() {
			createInput := &channels.CreateChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:             "Test",
				ChatUserFriendlyName: "Test",
				ChatFriendlyName:     "Test",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a channel sid", func() {
		channelClient := flexSession.Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the channel is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal(utils.String("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get channel response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.Channel("CH71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://flex-api.twilio.com/v1/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := channelClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete channel response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://flex-api.twilio.com/v1/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := flexSession.Channel("CH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the web channels client", func() {
		webChannelsClient := flexSession.WebChannels

		Describe("When the web channel is successfully created", func() {
			createInput := &web_channels.CreateWebChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:             "Test",
				ChatFriendlyName:     "Test",
				CustomerFriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/WebChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := webChannelsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create web channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the web channels request does not contain a flex flow sid", func() {
			createInput := &web_channels.CreateWebChannelInput{
				Identity:             "Test",
				ChatFriendlyName:     "Test",
				CustomerFriendlyName: "Test",
			}

			resp, err := webChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the web channels request does not contain a identity", func() {
			createInput := &web_channels.CreateWebChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				ChatFriendlyName:     "Test",
				CustomerFriendlyName: "Test",
			}

			resp, err := webChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the web channel request does not contain a chat friendly name", func() {
			createInput := &web_channels.CreateWebChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:             "Test",
				CustomerFriendlyName: "Test",
			}

			resp, err := webChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the web channel request does not contain a customer friendly name", func() {
			createInput := &web_channels.CreateWebChannelInput{
				FlexFlowSid:      "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:         "Test",
				ChatFriendlyName: "Test",
			}

			resp, err := webChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the web channel api returns a 500 response", func() {
			createInput := &web_channels.CreateWebChannelInput{
				FlexFlowSid:          "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Identity:             "Test",
				ChatFriendlyName:     "Test",
				CustomerFriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/WebChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := webChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a web channel sid", func() {
		webChannelClient := flexSession.WebChannel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the web channel is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := webChannelClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get web channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the web channel api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.WebChannel("CH71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the web channel is successfully updated", func() {
			updateInput := &web_channel.UpdateWebChannelInput{}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateWebChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := webChannelClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update web channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2016-08-02T22:10:40Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the web channel api returns a 404", func() {
			updateInput := &web_channel.UpdateWebChannelInput{}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/WebChannels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.WebChannel("CH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the web channel is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := webChannelClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete web channel response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://flex-api.twilio.com/v1/WebChannels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := flexSession.WebChannel("CH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})
})

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectNotFoundError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(Equal(utils.Int(20404)))
	Expect(twilioErr.Message).To(Equal("The requested resource /FlexFlows/FO71 was not found"))
	Expect(twilioErr.MoreInfo).To(Equal(utils.String("https://www.twilio.com/docs/errors/20404")))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
