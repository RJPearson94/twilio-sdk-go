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

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/flex"
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
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/web_channels"
	"github.com/RJPearson94/twilio-sdk-go/session"
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

	flexSession := flex.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V1

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

			resp, err := configurationClient.Fetch()
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

			resp, err := configurationClient.Fetch()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
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
				Expect(resp.Integration).To(Equal(&flex_flows.CreateFlexFlowIntegrationResponse{
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
				ExpectInternalServerError(err)
			})

			It("Then the create flex flow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of flex flows are successfully retrieved", func() {
			pageOptions := &flex_flows.FlexFlowsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flexFlowsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flexFlowsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the flex flows page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/FlexFlows?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/FlexFlows?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("flex_flows"))

				flexFlows := resp.FlexFlows
				Expect(flexFlows).ToNot(BeNil())
				Expect(len(flexFlows)).To(Equal(1))

				Expect(flexFlows[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(flexFlows[0].Sid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(flexFlows[0].FriendlyName).To(Equal("Test Flex Flow"))
				Expect(flexFlows[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(flexFlows[0].ChannelType).To(Equal("web"))
				Expect(flexFlows[0].ContactIdentity).To(Equal(utils.String("12345")))
				Expect(flexFlows[0].Enabled).To(Equal(true))
				Expect(flexFlows[0].IntegrationType).To(Equal(utils.String("studio")))
				Expect(flexFlows[0].Integration).To(Equal(&flex_flows.PageFlexFlowIntegrationResponse{
					RetryCount: utils.Int(1),
					FlowSid:    utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(flexFlows[0].LongLived).To(Equal(utils.Bool(true)))
				Expect(flexFlows[0].JanitorEnabled).To(Equal(utils.Bool(true)))
				Expect(flexFlows[0].DateUpdated).To(BeNil())
				Expect(flexFlows[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(flexFlows[0].URL).To(Equal("https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of flex flows api returns a 500 response", func() {
			pageOptions := &flex_flows.FlexFlowsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := flexFlowsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the flex flows page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated flex flows are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flexFlowsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flexFlowsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := flexFlowsClient.NewFlexFlowsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated flex flows current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated flex flows results should be returned", func() {
				Expect(len(paginator.FlexFlows)).To(Equal(3))
			})
		})

		Describe("When the flex flows api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/flexFlowsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/FlexFlows?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := flexFlowsClient.NewFlexFlowsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated flex flows current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := flexFlowClient.Fetch()
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
				Expect(resp.Integration).To(Equal(&flex_flow.FetchFlexFlowIntegrationResponse{
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

			resp, err := flexSession.FlexFlow("FO71").Fetch()
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
				FriendlyName: utils.String("New Flex Flow"),
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
				Expect(resp.Integration).To(Equal(&flex_flow.UpdateFlexFlowIntegrationResponse{
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
				ExpectInternalServerError(err)
			})

			It("Then the create channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of channels are successfully retrieved", func() {
			pageOptions := &channels.ChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the channels page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/Channels?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/Channels?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("flex_chat_channels"))

				channels := resp.Channels
				Expect(channels).ToNot(BeNil())
				Expect(len(channels)).To(Equal(1))

				Expect(channels[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].TaskSid).To(Equal(utils.String("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(channels[0].UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].DateUpdated).To(BeNil())
				Expect(channels[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(channels[0].URL).To(Equal("https://flex-api.twilio.com/v1/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of channels api returns a 500 response", func() {
			pageOptions := &channels.ChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the channels page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated channels are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := channelsClient.NewChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated channels current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated channels results should be returned", func() {
				Expect(len(paginator.Channels)).To(Equal(3))
			})
		})

		Describe("When the channels api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/Channels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := channelsClient.NewChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated channels current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := channelClient.Fetch()
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

			resp, err := flexSession.Channel("CH71").Fetch()
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
				ExpectInternalServerError(err)
			})

			It("Then the create web channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of web channels are successfully retrieved", func() {
			pageOptions := &web_channels.WebChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webChannelsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := webChannelsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the web channels page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/WebChannels?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/WebChannels?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("flex_chat_channels"))

				webChannels := resp.WebChannels
				Expect(webChannels).ToNot(BeNil())
				Expect(len(webChannels)).To(Equal(1))

				Expect(webChannels[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webChannels[0].Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webChannels[0].FlexFlowSid).To(Equal("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webChannels[0].DateUpdated).To(BeNil())
				Expect(webChannels[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(webChannels[0].URL).To(Equal("https://flex-api.twilio.com/v1/WebChannels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of web channels api returns a 500 response", func() {
			pageOptions := &web_channels.WebChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := webChannelsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the web channels page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated web channels are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webChannelsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := webChannelsClient.NewWebChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated web channels current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated web channels results should be returned", func() {
				Expect(len(paginator.WebChannels)).To(Equal(3))
			})
		})

		Describe("When the web channels api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/WebChannels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := webChannelsClient.NewWebChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated web channels current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := webChannelClient.Fetch()
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

			resp, err := flexSession.WebChannel("CH71").Fetch()
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

	Describe("Given the plugins client", func() {
		pluginsClient := flexSession.Plugins

		Describe("When the plugin is successfully created", func() {
			createInput := &plugins.CreatePluginInput{
				UniqueName: "test",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Plugins",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := pluginsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create plugin response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.Description).To(BeNil())
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the plugin request does not contain a unique name", func() {
			createInput := &plugins.CreatePluginInput{}

			resp, err := pluginsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create plugin response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the plugin api returns a 500 response", func() {
			createInput := &plugins.CreatePluginInput{
				UniqueName: "test",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Plugins",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := pluginsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create plugin response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of plugins are successfully retrieved", func() {
			pageOptions := &plugins.PluginsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := pluginsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the plugins page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("plugins"))

				plugins := resp.Plugins
				Expect(plugins).ToNot(BeNil())
				Expect(len(plugins)).To(Equal(1))

				Expect(plugins[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(plugins[0].Sid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(plugins[0].FriendlyName).To(Equal("test"))
				Expect(plugins[0].UniqueName).To(Equal("test"))
				Expect(plugins[0].Description).To(BeNil())
				Expect(plugins[0].Archived).To(Equal(false))
				Expect(plugins[0].DateUpdated).To(BeNil())
				Expect(plugins[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(plugins[0].URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of plugins api returns a 500 response", func() {
			pageOptions := &plugins.PluginsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := pluginsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the plugins page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated plugins are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := pluginsClient.NewPluginsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated plugins current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated plugins results should be returned", func() {
				Expect(len(paginator.Plugins)).To(Equal(3))
			})
		})

		Describe("When the plugins api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := pluginsClient.NewPluginsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated plugins current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a plugin sid", func() {
		pluginClient := flexSession.Plugin("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the plugin is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := pluginClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get plugin response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("test"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.Description).To(BeNil())
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get plugin response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.Plugin("FP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get plugin response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the plugin is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatePluginResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &plugin.UpdatePluginInput{
				FriendlyName: utils.String("test 2s"),
			}

			resp, err := pluginClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the updated plugin response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("test 2"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.Description).To(BeNil())
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2016-08-01T22:15:40Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update plugin response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Plugins/FP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &plugin.UpdatePluginInput{}

			resp, err := flexSession.Plugin("FP71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update plugin response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the plugin versions client", func() {
		versionsClient := flexSession.Plugin("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Versions

		Describe("When the plugin version is successfully created", func() {
			createInput := &versions.CreateVersionInput{
				Version:   "1.0.0",
				PluginURL: "https://example.com",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/versionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := versionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create plugin version response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PluginSid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PluginURL).To(Equal("https://example.com"))
				Expect(resp.Version).To(Equal("1.0.0"))
				Expect(resp.Changelog).To(BeNil())
				Expect(resp.Private).To(Equal(false))
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the plugin version request does not contain a version", func() {
			createInput := &versions.CreateVersionInput{
				PluginURL: "https://example.com",
			}

			resp, err := versionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create plugin version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the plugin version request does not contain a plugin url", func() {
			createInput := &versions.CreateVersionInput{
				Version: "1.0.0",
			}

			resp, err := versionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create plugin version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the plugin version api returns a 500 response", func() {
			createInput := &versions.CreateVersionInput{
				Version:   "1.0.0",
				PluginURL: "https://example.com",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := versionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create plugin version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of plugin versions are successfully retrieved", func() {
			pageOptions := &versions.VersionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/versionsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := versionsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the plugin versions page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("plugin_versions"))

				versions := resp.Versions
				Expect(versions).ToNot(BeNil())
				Expect(len(versions)).To(Equal(1))

				Expect(versions[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(versions[0].Sid).To(Equal("FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(versions[0].PluginSid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(versions[0].PluginURL).To(Equal("https://example.com"))
				Expect(versions[0].Version).To(Equal("1.0.0"))
				Expect(versions[0].Changelog).To(BeNil())
				Expect(versions[0].Private).To(Equal(false))
				Expect(versions[0].Archived).To(Equal(false))
				Expect(versions[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(versions[0].URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of plugin versions api returns a 500 response", func() {
			pageOptions := &versions.VersionsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := versionsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the plugin versions page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated plugin versions are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/versionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/versionsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := versionsClient.NewVersionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated plugin versions current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated plugin versions results should be returned", func() {
				Expect(len(paginator.Versions)).To(Equal(3))
			})
		})

		Describe("When the plugin versions api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/versionsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := versionsClient.NewVersionsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated plugin versions current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a plugin version sid", func() {
		versionClient := flexSession.Plugin("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the plugin version is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/versionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := versionClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get plugin version response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PluginSid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PluginURL).To(Equal("https://example.com"))
				Expect(resp.Version).To(Equal("1.0.0"))
				Expect(resp.Changelog).To(BeNil())
				Expect(resp.Private).To(Equal(false))
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get plugin version response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Versions/FV71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.Plugin("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Version("FV71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get plugin version response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the plugin configurations client", func() {
		pluginConfigurationsClient := flexSession.PluginConfigurations

		Describe("When the plugin configuration is successfully created", func() {
			createInput := &plugin_configurations.CreateConfigurationInput{
				Name: "test",
				Plugins: &[]string{
					`{"plugin_version":"FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}`,
				},
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Configurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginConfigurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := pluginConfigurationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create plugin configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.Description).To(BeNil())
				Expect(resp.Name).To(Equal("test"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the plugin configuration request does not contain a name", func() {
			createInput := &plugin_configurations.CreateConfigurationInput{}

			resp, err := pluginConfigurationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create plugin configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the plugin configuration api returns a 500 response", func() {
			createInput := &plugin_configurations.CreateConfigurationInput{
				Name: "test",
				Plugins: &[]string{
					`{"plugin_version":"FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}`,
				},
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Configurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := pluginConfigurationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create plugin configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of plugin configurations are successfully retrieved", func() {
			pageOptions := &plugin_configurations.ConfigurationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginConfigurationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := pluginConfigurationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the plugin configurations page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("configurations"))

				configurations := resp.Configurations
				Expect(configurations).ToNot(BeNil())
				Expect(len(configurations)).To(Equal(1))

				Expect(configurations[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(configurations[0].Sid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(configurations[0].Archived).To(Equal(false))
				Expect(configurations[0].Description).To(BeNil())
				Expect(configurations[0].Name).To(Equal("test"))
				Expect(configurations[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(configurations[0].URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of plugin configurations api returns a 500 response", func() {
			pageOptions := &plugin_configurations.ConfigurationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := pluginConfigurationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the plugin configurations page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated plugin configurations are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginConfigurationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginConfigurationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := pluginConfigurationsClient.NewConfigurationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated plugin configurations current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated plugin configurations results should be returned", func() {
				Expect(len(paginator.Configurations)).To(Equal(3))
			})
		})

		Describe("When the plugin configurations api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginConfigurationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := pluginConfigurationsClient.NewConfigurationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated plugin configurations current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a plugin configuration sid", func() {
		configurationClient := flexSession.PluginConfiguration("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the plugin configuration is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/pluginConfigurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := configurationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get plugin configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Archived).To(Equal(false))
				Expect(resp.Description).To(BeNil())
				Expect(resp.Name).To(Equal("test"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get plugin configuration response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.PluginConfiguration("FJ71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get plugin configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the configuration plugins client", func() {
		pluginsClient := flexSession.PluginConfiguration("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Plugins

		Describe("When the page of plugins are successfully retrieved", func() {
			pageOptions := &configurationPlugins.PluginsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationPluginsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := pluginsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the plugins page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("plugins"))

				plugins := resp.Plugins
				Expect(plugins).ToNot(BeNil())
				Expect(len(plugins)).To(Equal(1))

				Expect(plugins[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(plugins[0].PluginVersionSid).To(Equal("FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(plugins[0].PluginSid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(plugins[0].ConfigurationSid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(plugins[0].UniqueName).To(Equal("test"))
				Expect(plugins[0].Private).To(Equal(false))
				Expect(plugins[0].Version).To(Equal("1.0.1"))
				Expect(plugins[0].PluginURL).To(Equal("https://example.com"))
				Expect(plugins[0].Phase).To(Equal(3))
				Expect(plugins[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(plugins[0].URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of plugins api returns a 500 response", func() {
			pageOptions := &configurationPlugins.PluginsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := pluginsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the plugins page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated plugins are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationPluginsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationPluginsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := pluginsClient.NewPluginsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated plugins current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated plugins results should be returned", func() {
				Expect(len(paginator.Plugins)).To(Equal(3))
			})
		})

		Describe("When the plugins api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationPluginsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := pluginsClient.NewPluginsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated plugins current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a configuration plugin sid", func() {
		pluginClient := flexSession.PluginConfiguration("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Plugin("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the plugin is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationPluginResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := pluginClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get plugin response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PluginVersionSid).To(Equal("FVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PluginSid).To(Equal("FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConfigurationSid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("test"))
				Expect(resp.Private).To(Equal(false))
				Expect(resp.Version).To(Equal("1.0.1"))
				Expect(resp.PluginURL).To(Equal("https://example.com"))
				Expect(resp.Phase).To(Equal(3))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get plugin response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Plugins/FP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.PluginConfiguration("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Plugin("FP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get plugin response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the plugin releases client", func() {
		releasesClient := flexSession.PluginReleases

		Describe("When the plugin release is successfully created", func() {
			createInput := &plugin_releases.CreateReleaseInput{
				ConfigurationId: "FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Releases",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/releaseResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := releasesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create plugin release response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConfigurationSid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Releases/FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the plugin release request does not contain a name", func() {
			createInput := &plugin_releases.CreateReleaseInput{}

			resp, err := releasesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create plugin release response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the plugin release api returns a 500 response", func() {
			createInput := &plugin_releases.CreateReleaseInput{
				ConfigurationId: "FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://flex-api.twilio.com/v1/PluginService/Releases",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := releasesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create plugin release response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of plugin releases are successfully retrieved", func() {
			pageOptions := &plugin_releases.ReleasesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/releasesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := releasesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the plugin releases page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Releases?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Releases?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("releases"))

				releases := resp.Releases
				Expect(releases).ToNot(BeNil())
				Expect(len(releases)).To(Equal(1))

				Expect(releases[0].Sid).To(Equal("FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(releases[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(releases[0].ConfigurationSid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(releases[0].DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(releases[0].URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Releases/FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of plugin releases api returns a 500 response", func() {
			pageOptions := &plugin_releases.ReleasesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := releasesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the plugin releases page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated plugin releases are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/releasesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/releasesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := releasesClient.NewReleasesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated plugin releases current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated plugin releases results should be returned", func() {
				Expect(len(paginator.Releases)).To(Equal(3))
			})
		})

		Describe("When the plugin releases api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/releasesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := releasesClient.NewReleasesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated plugin releases current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a plugin release sid", func() {
		releaseClient := flexSession.PluginRelease("FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the plugin release is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases/FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/releaseResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := releaseClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get plugin release response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConfigurationSid).To(Equal("FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				Expect(resp.URL).To(Equal("https://flex-api.twilio.com/v1/PluginService/Releases/FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get plugin release response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://flex-api.twilio.com/v1/PluginService/Releases/FK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := flexSession.PluginRelease("FK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get plugin release response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})
})

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}

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
