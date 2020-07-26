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

	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Sync V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	syncSession := sync.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(syncSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the services client", func() {
		servicesClient := syncSession.Services

		Describe("When the service is successfully created", func() {
			createInput := &services.CreateServiceInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services",
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
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(resp.ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(resp.ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(resp.WebhookURL).To(BeNil())
				Expect(resp.WebhooksFromRestEnabled).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create service api returns a 500 response", func() {
			createInput := &services.CreateServiceInput{}

			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a service sid", func() {
		serviceClient := syncSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(resp.ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(resp.ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(resp.WebhookURL).To(BeNil())
				Expect(resp.WebhooksFromRestEnabled).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://sync.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := syncSession.Service("IS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := serviceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.FriendlyName).To(Equal(utils.String("test")))
				Expect(resp.ReachabilityDebouncingWindow).To(Equal(5000))
				Expect(resp.ReachabilityDebouncingEnabled).To(Equal(false))
				Expect(resp.ReachabilityWebhooksEnabled).To(Equal(false))
				Expect(resp.WebhookURL).To(BeNil())
				Expect(resp.WebhooksFromRestEnabled).To(Equal(false))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T21:00:24Z"))
				Expect(resp.URL).To(Equal("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update service response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://sync.twilio.com/v1/Services/IS71",
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

			resp, err := syncSession.Service("IS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://sync.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := syncSession.Service("IS71").Delete()
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Services/IS71 was not found"))
	Expect(twilioErr.MoreInfo).To(Equal(utils.String("https://www.twilio.com/docs/errors/20404")))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
