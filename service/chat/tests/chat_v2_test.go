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

	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Chat V2", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	chatSession := chat.NewWithCredentials(creds).V2

	httpmock.ActivateNonDefault(chatSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the services client", func() {
		servicesClient := chatSession.Services

		Describe("When the service is successfully created", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := servicesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConsumptionReportInterval).To(Equal(10))
				Expect(resp.DefaultChannelCreatorRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultChannelRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultServiceRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))

				limitsFixture, _ := ioutil.ReadFile("testdata/limitsResponse.json")
				limitsResp := make(map[string]interface{})
				json.Unmarshal(limitsFixture, &limitsResp)
				Expect(resp.Limits).To(Equal(limitsResp))

				mediaFixture, _ := ioutil.ReadFile("testdata/mediaResponse.json")
				mediaResp := make(map[string]interface{})
				json.Unmarshal(mediaFixture, &mediaResp)
				Expect(resp.Media).To(Equal(mediaResp))

				notificationFixture, _ := ioutil.ReadFile("testdata/notificationResponse.json")
				notificationResp := make(map[string]interface{})
				json.Unmarshal(notificationFixture, &notificationResp)
				Expect(resp.Notifications).To(Equal(notificationResp))

				Expect(resp.PostWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PostWebhookUrl).To(BeNil())
				Expect(resp.PreWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PreWebhookUrl).To(BeNil())
				Expect(resp.ReachabilityEnabled).To(Equal(false))
				Expect(resp.ReadStatusEnabled).To(Equal(true))
				Expect(resp.TypingIndicatorTimeout).To(Equal(5))
				Expect(resp.WebhookFilters).To(BeNil())
				Expect(resp.WebhookMethod).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service does not contain a friendly name", func() {
			createInput := &services.CreateServiceInput{}

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create service api returns a 500 response", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "TesFriendlyNamet",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a service sid", func() {
		serviceClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := serviceClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("IS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := serviceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConsumptionReportInterval).To(Equal(10))
				Expect(resp.DefaultChannelCreatorRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultChannelRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultServiceRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))

				limitsFixture, _ := ioutil.ReadFile("testdata/limitsResponse.json")
				limitsResp := make(map[string]interface{})
				json.Unmarshal(limitsFixture, &limitsResp)
				Expect(resp.Limits).To(Equal(limitsResp))

				mediaFixture, _ := ioutil.ReadFile("testdata/mediaResponse.json")
				mediaResp := make(map[string]interface{})
				json.Unmarshal(mediaFixture, &mediaResp)
				Expect(resp.Media).To(Equal(mediaResp))

				notificationFixture, _ := ioutil.ReadFile("testdata/notificationResponse.json")
				notificationResp := make(map[string]interface{})
				json.Unmarshal(notificationFixture, &notificationResp)
				Expect(resp.Notifications).To(Equal(notificationResp))

				Expect(resp.PostWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PostWebhookUrl).To(BeNil())
				Expect(resp.PreWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PreWebhookUrl).To(BeNil())
				Expect(resp.ReachabilityEnabled).To(Equal(false))
				Expect(resp.ReadStatusEnabled).To(Equal(true))
				Expect(resp.TypingIndicatorTimeout).To(Equal(5))
				Expect(resp.WebhookFilters).To(BeNil())
				Expect(resp.WebhookMethod).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T22:50:24Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update service response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := chatSession.Service("IS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("IS71").Delete()
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

	code := 20404
	Expect(twilioErr.Code).To(Equal(&code))
	Expect(twilioErr.Message).To(Equal("The requested resource /Services/KS71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
