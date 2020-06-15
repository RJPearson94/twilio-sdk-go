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
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
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
})

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
