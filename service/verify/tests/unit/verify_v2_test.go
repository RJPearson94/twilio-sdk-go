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

	"github.com/RJPearson94/twilio-sdk-go/service/verify"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Verify V2", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	verifySession := verify.NewWithCredentials(creds).V2

	httpmock.ActivateNonDefault(verifySession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a services client", func() {
		servicesClient := verifySession.Services

		Describe("When the verify resource is successfully created", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Test Service",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services",
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
				Expect(resp.Sid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TtsName).To(BeNil())
				Expect(resp.MailerSid).To(BeNil())
				Expect(resp.Psd2Enabled).To(Equal(false))
				Expect(resp.DoNotShareWarningEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(Equal("Test Service"))
				Expect(resp.CodeLength).To(Equal(6))
				Expect(resp.CustomCodeEnabled).To(Equal(false))
				Expect(resp.DtmfInputRequired).To(Equal(true))
				Expect(resp.SkipSmsToLandlines).To(Equal(true))
				Expect(resp.LookupEnabled).To(Equal(true))
				Expect(resp.Push).To(Equal(services.CreateServicePushResponse{
					ApnCredentialSid: nil,
					IncludeDate:      true,
					FcmCredentialSid: nil,
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service request does not contain a friendly name", func() {
			createInput := &services.CreateServiceInput{}

			resp, err := servicesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create services api returns a 500 response", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Test Service",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services",
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
		serviceClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := serviceClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get service resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TtsName).To(BeNil())
				Expect(resp.MailerSid).To(BeNil())
				Expect(resp.Psd2Enabled).To(Equal(false))
				Expect(resp.DoNotShareWarningEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(Equal("Test Service"))
				Expect(resp.CodeLength).To(Equal(6))
				Expect(resp.CustomCodeEnabled).To(Equal(false))
				Expect(resp.DtmfInputRequired).To(Equal(true))
				Expect(resp.SkipSmsToLandlines).To(Equal(true))
				Expect(resp.LookupEnabled).To(Equal(true))
				Expect(resp.Push).To(Equal(service.FetchServicePushResponse{
					ApnCredentialSid: nil,
					IncludeDate:      true,
					FcmCredentialSid: nil,
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{}

			resp, err := serviceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TtsName).To(BeNil())
				Expect(resp.MailerSid).To(BeNil())
				Expect(resp.Psd2Enabled).To(Equal(false))
				Expect(resp.DoNotShareWarningEnabled).To(Equal(false))
				Expect(resp.FriendlyName).To(Equal("Test Service"))
				Expect(resp.CodeLength).To(Equal(6))
				Expect(resp.CustomCodeEnabled).To(Equal(false))
				Expect(resp.DtmfInputRequired).To(Equal(true))
				Expect(resp.SkipSmsToLandlines).To(Equal(true))
				Expect(resp.LookupEnabled).To(Equal(true))
				Expect(resp.Push).To(Equal(service.UpdateServicePushResponse{
					ApnCredentialSid: nil,
					IncludeDate:      true,
					FcmCredentialSid: nil,
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update service resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{}

			resp, err := verifySession.Service("VA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the service resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VA71").Delete()
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Services/VA71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
