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
				Expect(resp.Integration).To(Equal(&flex_flows.CreateFlexFlowResponseIntegration{
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
				Expect(flexFlows[0].Integration).To(Equal(&flex_flows.PageFlexFlowResponseIntegration{
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
				Expect(resp.Integration).To(Equal(&flex_flow.FetchFlexFlowResponseIntegration{
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
				Expect(resp.Integration).To(Equal(&flex_flow.UpdateFlexFlowResponseIntegration{
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
