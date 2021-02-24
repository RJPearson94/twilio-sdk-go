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
	"github.com/RJPearson94/twilio-sdk-go/service/verify"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/access_tokens"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entities"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/challenge"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/challenges"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/factor"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/entity/factors"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configurations"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/bucket"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/buckets"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limits"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verification_check"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/verifications"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhooks"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
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

	verifySession := verify.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V2

	httpmock.ActivateNonDefault(verifySession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a services client", func() {
		servicesClient := verifySession.Services

		Describe("When the service resource is successfully created", func() {
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

		Describe("When the page of services are successfully retrieved", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := servicesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the services page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("services"))

				pushResponse := services.PageServicePushResponse{
					ApnCredentialSid: nil,
					IncludeDate:      true,
					FcmCredentialSid: nil,
				}

				services := resp.Services
				Expect(services).ToNot(BeNil())
				Expect(len(services)).To(Equal(1))

				Expect(services[0].Sid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].TtsName).To(BeNil())
				Expect(services[0].MailerSid).To(BeNil())
				Expect(services[0].Psd2Enabled).To(Equal(false))
				Expect(services[0].DoNotShareWarningEnabled).To(Equal(false))
				Expect(services[0].FriendlyName).To(Equal("Test Service"))
				Expect(services[0].CodeLength).To(Equal(6))
				Expect(services[0].CustomCodeEnabled).To(Equal(false))
				Expect(services[0].DtmfInputRequired).To(Equal(true))
				Expect(services[0].SkipSmsToLandlines).To(Equal(true))
				Expect(services[0].LookupEnabled).To(Equal(true))
				Expect(services[0].Push).To(Equal(pushResponse))
				Expect(services[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(services[0].DateUpdated).To(BeNil())
				Expect(services[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of services api returns a 500 response", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := servicesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the services page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated services are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := servicesClient.NewServicesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated services current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated services results should be returned", func() {
				Expect(len(paginator.Services)).To(Equal(3))
			})
		})

		Describe("When the services api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := servicesClient.NewServicesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated services current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

	Describe("Given I have a rate limits client", func() {
		rateLimitsClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimits

		Describe("When the rate limit resource is successfully created", func() {
			createInput := &rate_limits.CreateRateLimitInput{
				UniqueName: "Test Rate Limit",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rateLimitResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := rateLimitsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create rate limit response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("Test Rate Limit"))
				Expect(resp.Description).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the rate limit request does not contain a friendly name", func() {
			createInput := &rate_limits.CreateRateLimitInput{}

			resp, err := rateLimitsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create rate limit response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create rate limits api returns a 500 response", func() {
			createInput := &rate_limits.CreateRateLimitInput{
				UniqueName: "Test Rate Limit",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := rateLimitsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create rate limit response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of rate limits are successfully retrieved", func() {
			pageOptions := &rate_limits.RateLimitsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rateLimitsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := rateLimitsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the rate limits page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("rate_limits"))

				rateLimits := resp.RateLimits
				Expect(rateLimits).ToNot(BeNil())
				Expect(len(rateLimits)).To(Equal(1))

				Expect(rateLimits[0].Sid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rateLimits[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rateLimits[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rateLimits[0].UniqueName).To(Equal("Test Rate Limit"))
				Expect(rateLimits[0].Description).To(BeNil())
				Expect(rateLimits[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(rateLimits[0].DateUpdated).To(BeNil())
				Expect(rateLimits[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of rate limits api returns a 500 response", func() {
			pageOptions := &rate_limits.RateLimitsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := rateLimitsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the rate limits page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated rate limits are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rateLimitsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rateLimitsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := rateLimitsClient.NewRateLimitsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated rate limits current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated rate limits results should be returned", func() {
				Expect(len(paginator.RateLimits)).To(Equal(3))
			})
		})

		Describe("When the rate limits api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rateLimitsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := rateLimitsClient.NewRateLimitsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated rate limits current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a rate limit sid", func() {
		rateLimitClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the rate limit resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rateLimitResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := rateLimitClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get rate limit resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("Test Rate Limit"))
				Expect(resp.Description).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the rate limit resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get rate limit response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the rate limit resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRateLimitResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &rate_limit.UpdateRateLimitInput{}

			resp, err := rateLimitClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update rate limit response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("Test Rate Limit"))
				Expect(resp.Description).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update rate limit resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &rate_limit.UpdateRateLimitInput{}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update rate limit response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the rate limit resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := rateLimitClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the rate limit resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a buckets client", func() {
		bucketsClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Buckets

		Describe("When the bucket resource is successfully created", func() {
			createInput := &buckets.CreateBucketInput{
				Max:      4,
				Interval: 10,
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bucketResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := bucketsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create bucket response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RateLimitSid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Interval).To(Equal(10))
				Expect(resp.Max).To(Equal(4))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the bucket request does not contain a max", func() {
			createInput := &buckets.CreateBucketInput{
				Interval: 10,
			}

			resp, err := bucketsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create bucket response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the bucket request does not contain an interval", func() {
			createInput := &buckets.CreateBucketInput{
				Max: 4,
			}

			resp, err := bucketsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create bucket response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create buckets api returns a 500 response", func() {
			createInput := &buckets.CreateBucketInput{
				Max:      4,
				Interval: 10,
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := bucketsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create bucket response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of buckets are successfully retrieved", func() {
			pageOptions := &buckets.BucketsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bucketsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := bucketsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the buckets page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("buckets"))

				buckets := resp.Buckets
				Expect(buckets).ToNot(BeNil())
				Expect(len(buckets)).To(Equal(1))

				Expect(buckets[0].Sid).To(Equal("BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(buckets[0].RateLimitSid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(buckets[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(buckets[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(buckets[0].Interval).To(Equal(10))
				Expect(buckets[0].Max).To(Equal(4))
				Expect(buckets[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(buckets[0].DateUpdated).To(BeNil())
				Expect(buckets[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of services api returns a 500 response", func() {
			pageOptions := &buckets.BucketsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := bucketsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the buckets page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated buckets are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bucketsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bucketsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := bucketsClient.NewBucketsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated buckets current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated buckets results should be returned", func() {
				Expect(len(paginator.Buckets)).To(Equal(3))
			})
		})

		Describe("When the buckets api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bucketsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := bucketsClient.NewBucketsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated buckets current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a bucket sid", func() {
		bucketClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bucket("BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the bucket resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bucketResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := bucketClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get bucket resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RateLimitSid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Interval).To(Equal(10))
				Expect(resp.Max).To(Equal(4))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the bucket resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bucket("BL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get bucket response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the bucket resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateBucketResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &bucket.UpdateBucketInput{}

			resp, err := bucketClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update bucket response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RateLimitSid).To(Equal("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Interval).To(Equal(10))
				Expect(resp.Max).To(Equal(4))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update bucket resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &bucket.UpdateBucketInput{}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bucket("BL71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update bucket response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the bucket resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := bucketClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the bucket resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").RateLimit("RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bucket("BL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a verifications client", func() {
		verificationsClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Verifications

		Describe("When the verification resource is successfully created", func() {
			createInput := &verifications.CreateVerificationInput{
				To:      "+123456789",
				Channel: "sms",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/verificationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := verificationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create verification response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.To).To(Equal("+123456789"))
				Expect(resp.Valid).To(Equal(false))

				sendAttemptTime, _ := time.Parse(time.RFC3339, "2020-06-20T20:51:24Z")
				Expect(resp.SendCodeAttempts).To(Equal([]verifications.CreateVerificationSendCodeAttemptResponse{
					{
						Channel:   "sms",
						Time:      sendAttemptTime,
						ChannelId: nil,
					},
				}))
				Expect(resp.Lookup).To(Equal(verifications.CreateVerificationLookupResponse{
					Carrier: &verifications.CreateVerificationCarrierLookupResponse{
						MobileCountryCode: nil,
						Type:              nil,
						ErrorCode:         nil,
						MobileNetworkCode: nil,
						Name:              nil,
					},
				}))
				Expect(resp.Channel).To(Equal("sms"))
				Expect(resp.Status).To(Equal("pending"))
				Expect(resp.Payee).To(BeNil())
				Expect(resp.Amount).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the verification request does not contain a to", func() {
			createInput := &verifications.CreateVerificationInput{
				Channel: "sms",
			}

			resp, err := verificationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create bucket response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the verification request does not contain an channel", func() {
			createInput := &verifications.CreateVerificationInput{
				To: "+123456789",
			}

			resp, err := verificationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create verification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create verification api returns a 500 response", func() {
			createInput := &verifications.CreateVerificationInput{
				To:      "+123456789",
				Channel: "sms",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := verificationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create verification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a verification sid", func() {
		verificationClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Verification("VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the verification resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/verificationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := verificationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get verification resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.To).To(Equal("+123456789"))
				Expect(resp.Valid).To(Equal(false))

				sendAttemptTime, _ := time.Parse(time.RFC3339, "2020-06-20T20:51:24Z")
				Expect(resp.SendCodeAttempts).To(Equal([]verification.FetchVerificationSendCodeAttemptResponse{
					{
						Channel:   "sms",
						Time:      sendAttemptTime,
						ChannelId: nil,
					},
				}))
				Expect(resp.Lookup).To(Equal(verification.FetchVerificationLookupResponse{
					Carrier: &verification.FetchVerificationCarrierLookupResponse{
						MobileCountryCode: nil,
						Type:              nil,
						ErrorCode:         nil,
						MobileNetworkCode: nil,
						Name:              nil,
					},
				}))
				Expect(resp.Channel).To(Equal("sms"))
				Expect(resp.Status).To(Equal("pending"))
				Expect(resp.Payee).To(BeNil())
				Expect(resp.Amount).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the verification resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VE71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Verification("VE71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get verification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the verification resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateVerificationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &verification.UpdateVerificationInput{
				Status: "canceled",
			}

			resp, err := verificationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update verification response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.To).To(Equal("+123456789"))
				Expect(resp.Valid).To(Equal(false))

				sendAttemptTime, _ := time.Parse(time.RFC3339, "2020-06-20T20:51:24Z")
				Expect(resp.SendCodeAttempts).To(Equal([]verification.UpdateVerificationSendCodeAttemptResponse{
					{
						Channel:   "sms",
						Time:      sendAttemptTime,
						ChannelId: nil,
					},
				}))
				Expect(resp.Lookup).To(Equal(verification.UpdateVerificationLookupResponse{
					Carrier: &verification.UpdateVerificationCarrierLookupResponse{
						MobileCountryCode: nil,
						Type:              nil,
						ErrorCode:         nil,
						MobileNetworkCode: nil,
						Name:              nil,
					},
				}))
				Expect(resp.Channel).To(Equal("sms"))
				Expect(resp.Status).To(Equal("canceled"))
				Expect(resp.Payee).To(BeNil())
				Expect(resp.Amount).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update verification request does not contain a status", func() {
			updateInput := &verification.UpdateVerificationInput{}

			resp, err := verificationClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update verification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update verification resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Verifications/VE71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &verification.UpdateVerificationInput{
				Status: "canceled",
			}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Verification("VE71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update verification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a verification check client", func() {
		verificationCheckClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").VerificationCheck

		Describe("When the verification check resource is successfully created", func() {
			createInput := &verification_check.CreateVerificationCheckInput{
				To:   utils.String("+123456789"),
				Code: "9876",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/VerificationCheck",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/verificationCheckResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := verificationCheckClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create verification check response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("approved"))
				Expect(resp.To).To(Equal("+123456789"))
				Expect(resp.Channel).To(Equal("sms"))
				Expect(resp.Valid).To(Equal(true))
				Expect(resp.Payee).To(BeNil())
				Expect(resp.Amount).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
			})
		})

		Describe("When the verification check request does not contain a code", func() {
			createInput := &verification_check.CreateVerificationCheckInput{
				To: utils.String("+123456789"),
			}

			resp, err := verificationCheckClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create verification check response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create verification check api returns a 500 response", func() {
			createInput := &verification_check.CreateVerificationCheckInput{
				To:   utils.String("+123456789"),
				Code: "9876",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/VerificationCheck",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := verificationCheckClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create verification check response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a access token client", func() {
		accessTokensClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").AccessTokens

		Describe("When the access token resource is successfully created", func() {
			createInput := &access_tokens.CreateAccessTokenInput{
				Identity:   "Test User",
				FactorType: "push",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AccessTokens",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/accessTokenResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := accessTokensClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create access token response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Token).To(Equal("abc1234"))
			})
		})

		Describe("When the access token request does not contain a identity", func() {
			createInput := &access_tokens.CreateAccessTokenInput{
				FactorType: "push",
			}

			resp, err := accessTokensClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create access token  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the access token request does not contain a factor type", func() {
			createInput := &access_tokens.CreateAccessTokenInput{
				Identity: "Test User",
			}

			resp, err := accessTokensClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create access token  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create access token api returns a 500 response", func() {
			createInput := &access_tokens.CreateAccessTokenInput{
				Identity:   "Test User",
				FactorType: "push",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AccessTokens",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := accessTokensClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create access token response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a entities client", func() {
		entitiesClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entities

		Describe("When the entity resource is successfully created", func() {
			createInput := &entities.CreateEntityInput{
				Identity: "TestUser",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/entityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := entitiesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create entity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/TestUser"))
			})
		})

		Describe("When the entity request does not contain an identity", func() {
			createInput := &entities.CreateEntityInput{}

			resp, err := entitiesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create entity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create entity api returns a 500 response", func() {
			createInput := &entities.CreateEntityInput{
				Identity: "TestUser",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := entitiesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create entity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of entities are successfully retrieved", func() {
			pageOptions := &entities.EntitiesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/entitiesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := entitiesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the entities page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("entities"))

				entities := resp.Entities
				Expect(entities).ToNot(BeNil())
				Expect(len(entities)).To(Equal(1))

				Expect(entities[0].Sid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(entities[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(entities[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(entities[0].Identity).To(Equal("TestUser"))
				Expect(entities[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(entities[0].DateUpdated).To(BeNil())
				Expect(entities[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/TestUser"))
			})
		})

		Describe("When the page of entities api returns a 500 response", func() {
			pageOptions := &entities.EntitiesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := entitiesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the entities page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated entities are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/entitiesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/entitiesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := entitiesClient.NewEntitiesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated entities current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated entities results should be returned", func() {
				Expect(len(paginator.Entities)).To(Equal(3))
			})
		})

		Describe("When the entities api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/entitiesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := entitiesClient.NewEntitiesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated entities current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a entity identity", func() {
		entityClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("TestUser")

		Describe("When the entity resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/TestUser",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/entityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := entityClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get entity resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/TestUser"))
			})
		})

		Describe("When the entity resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/UnknownUser",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("UnknownUser").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get entity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the entity resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/TestUser", httpmock.NewStringResponder(204, ""))

			err := entityClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the entity resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/UnknownUser",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("UnknownUser").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a webhooks client", func() {
		webhooksClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhooks

		Describe("When the webhook resource is successfully created", func() {
			createInput := &webhooks.CreateWebhookInput{
				FriendlyName: "Test Webhook",
				EventTypes:   []string{"*"},
				WebhookURL:   "http://localhost/webhook",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := webhooksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("enabled"))
				Expect(resp.FriendlyName).To(Equal("Test Webhook"))
				Expect(resp.WebhookMethod).To(Equal("POST"))
				Expect(resp.WebhookURL).To(Equal("http://localhost/webhook"))
				Expect(resp.EventTypes).To(Equal([]string{"*"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the webhook request does not contain a friendly name", func() {
			createInput := &webhooks.CreateWebhookInput{
				EventTypes: []string{"*"},
				WebhookURL: "http://localhost/webhook",
			}

			resp, err := webhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the webhook request does not contain an event types array", func() {
			createInput := &webhooks.CreateWebhookInput{
				FriendlyName: "Test Webhook",
				WebhookURL:   "http://localhost/webhook",
			}

			resp, err := webhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the webhook request does not contain a webhook url", func() {
			createInput := &webhooks.CreateWebhookInput{
				FriendlyName: "Test Webhook",
				EventTypes:   []string{"*"},
			}

			resp, err := webhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create webhooks api returns a 500 response", func() {
			createInput := &webhooks.CreateWebhookInput{
				FriendlyName: "Test Webhook",
				EventTypes:   []string{"*"},
				WebhookURL:   "http://localhost/webhook",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := webhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of webhooks are successfully retrieved", func() {
			pageOptions := &webhooks.WebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webhooksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := webhooksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the webhooks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("webhooks"))

				webhooks := resp.Webhooks
				Expect(webhooks).ToNot(BeNil())
				Expect(len(webhooks)).To(Equal(1))

				Expect(webhooks[0].Sid).To(Equal("YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].Status).To(Equal("enabled"))
				Expect(webhooks[0].FriendlyName).To(Equal("Test Webhook"))
				Expect(webhooks[0].WebhookMethod).To(Equal("POST"))
				Expect(webhooks[0].WebhookURL).To(Equal("http://localhost/webhook"))
				Expect(webhooks[0].EventTypes).To(Equal([]string{"*"}))
				Expect(webhooks[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(webhooks[0].DateUpdated).To(BeNil())
				Expect(webhooks[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of webhooks api returns a 500 response", func() {
			pageOptions := &webhooks.WebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := webhooksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the webhooks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated webhooks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webhooksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := webhooksClient.NewWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated webhooks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated webhooks results should be returned", func() {
				Expect(len(paginator.Webhooks)).To(Equal(3))
			})
		})

		Describe("When the webhooks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := webhooksClient.NewWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated webhooks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a webhook sid", func() {
		webhookClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the webhook resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/webhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := webhookClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get webhook resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("enabled"))
				Expect(resp.FriendlyName).To(Equal("Test Webhook"))
				Expect(resp.WebhookMethod).To(Equal("POST"))
				Expect(resp.WebhookURL).To(Equal("http://localhost/webhook"))
				Expect(resp.EventTypes).To(Equal([]string{"*"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the webhook resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YW71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("YW71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the webhook resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &webhook.UpdateWebhookInput{}

			resp, err := webhookClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("enabled"))
				Expect(resp.FriendlyName).To(Equal("Test Webhook"))
				Expect(resp.WebhookMethod).To(Equal("POST"))
				Expect(resp.WebhookURL).To(Equal("http://localhost/webhook"))
				Expect(resp.EventTypes).To(Equal([]string{"*"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update webhook resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YW71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &webhook.UpdateWebhookInput{}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("YW71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the webhook resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := webhookClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the webhook resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YW71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("YW71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a factors client", func() {
		factorsClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Factors

		Describe("When the factor resource is successfully created", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/factorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := factorsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create factor response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.FriendlyName).To(Equal("test factor"))
				Expect(resp.Status).To(Equal("unverified"))
				Expect(resp.FactorType).To(Equal("push"))
				Expect(resp.Config).To(Equal(factors.CreateFactorConfigResponse{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
					SdkVersion:           nil,
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the factor request does not contain a binding alg", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a binding struct", func() {
			createInput := &factors.CreateFactorInput{
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a binding public key", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg: "ES256",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
		Describe("When the factor request does not contain a config app id", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a config notification platform", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:             "test",
					NotificationToken: "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a config notification token", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a config struct", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a factor type", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FriendlyName: "test factor",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor request does not contain a friendly name", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType: "push",
			}

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create factor api returns a 500 response", func() {
			createInput := &factors.CreateFactorInput{
				Binding: factors.CreateFactorBindingInput{
					Alg:       "ES256",
					PublicKey: "TestKey",
				},
				Config: factors.CreateFactorConfigInput{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
				},
				FactorType:   "push",
				FriendlyName: "test factor",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := factorsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of factors are successfully retrieved", func() {
			pageOptions := &factors.FactorsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/factorsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := factorsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the factors page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("factors"))

				factorConfig := factors.PageFactorConfigResponse{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
					SdkVersion:           nil,
				}

				factors := resp.Factors
				Expect(factors).ToNot(BeNil())
				Expect(len(factors)).To(Equal(1))

				Expect(factors[0].Sid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(factors[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(factors[0].EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(factors[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(factors[0].Identity).To(Equal("test"))
				Expect(factors[0].FriendlyName).To(Equal("test factor"))
				Expect(factors[0].Status).To(Equal("unverified"))
				Expect(factors[0].FactorType).To(Equal("push"))
				Expect(factors[0].Config).To(Equal(factorConfig))
				Expect(factors[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(factors[0].DateUpdated).To(BeNil())
				Expect(factors[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of factors api returns a 500 response", func() {
			pageOptions := &factors.FactorsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := factorsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the factors page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated factors are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/factorsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/factorsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := factorsClient.NewFactorsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated factors current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated factors results should be returned", func() {
				Expect(len(paginator.Factors)).To(Equal(3))
			})
		})

		Describe("When the factors api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/factorsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := factorsClient.NewFactorsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated factors current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a factor sid", func() {
		factorClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Factor("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the factor resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/factorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := factorClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get factor resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.FriendlyName).To(Equal("test factor"))
				Expect(resp.Status).To(Equal("unverified"))
				Expect(resp.FactorType).To(Equal("push"))
				Expect(resp.Config).To(Equal(factor.FetchFactorConfigResponse{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
					SdkVersion:           nil,
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the factor resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Factor("YF71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateFactorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &factor.UpdateFactorInput{}

			resp, err := factorClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update factor response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.FriendlyName).To(Equal("test factor"))
				Expect(resp.Status).To(Equal("unverified"))
				Expect(resp.FactorType).To(Equal("push"))
				Expect(resp.Config).To(Equal(factor.UpdateFactorConfigResponse{
					AppId:                "test",
					NotificationPlatform: "fcm",
					NotificationToken:    "notification_token",
					SdkVersion:           nil,
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update factor resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &factor.UpdateFactorInput{}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Factor("YF71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the factor resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := factorClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the factor resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Factors/YF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Factor("YF71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a challenges client", func() {
		challengesClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Challenges

		Describe("When the challenge resource is successfully created", func() {
			createInput := &challenges.CreateChallengeInput{
				FactorSid: "YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/challengeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := challengesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create challenge response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FactorSid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Status).To(Equal("pending"))
				Expect(resp.FactorType).To(Equal("push"))
				Expect(resp.RespondedReason).To(Equal(utils.String("none")))
				Expect(resp.Details).To(BeNil())
				Expect(resp.HiddenDetails).To(BeNil())
				Expect(resp.DateResponded.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.ExpirationDate.Format(time.RFC3339)).To(Equal("2020-06-20T20:51:24Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the challenge request does not contain a factor sid", func() {
			createInput := &challenges.CreateChallengeInput{}

			resp, err := challengesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create challenge response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create challenge api returns a 500 response", func() {
			createInput := &challenges.CreateChallengeInput{
				FactorSid: "YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := challengesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create factor response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of challenges are successfully retrieved", func() {
			pageOptions := &challenges.ChallengesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/challengesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := challengesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the challenges page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("challenges"))

				challenges := resp.Challenges
				Expect(challenges).ToNot(BeNil())
				Expect(len(challenges)).To(Equal(1))

				Expect(challenges[0].Sid).To(Equal("YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(challenges[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(challenges[0].EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(challenges[0].FactorSid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(challenges[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(challenges[0].Identity).To(Equal("test"))
				Expect(challenges[0].Status).To(Equal("pending"))
				Expect(challenges[0].FactorType).To(Equal("push"))
				Expect(challenges[0].RespondedReason).To(Equal(utils.String("none")))
				Expect(challenges[0].Details).To(BeNil())
				Expect(challenges[0].HiddenDetails).To(BeNil())
				Expect(challenges[0].DateResponded.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(challenges[0].ExpirationDate.Format(time.RFC3339)).To(Equal("2020-06-20T20:51:24Z"))
				Expect(challenges[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(challenges[0].DateUpdated).To(BeNil())
				Expect(challenges[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of challenges api returns a 500 response", func() {
			pageOptions := &challenges.ChallengesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := challengesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the challenges page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated challenges are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/challengesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/challengesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := challengesClient.NewChallengesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated challenges current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated challenges results should be returned", func() {
				Expect(len(paginator.Challenges)).To(Equal(3))
			})
		})

		Describe("When the challenges api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/challengesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := challengesClient.NewChallengesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated challenges current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a challenge sid", func() {
		challengeClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Challenge("YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the challenge resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/challengeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := challengeClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get challenge resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FactorSid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Status).To(Equal("pending"))
				Expect(resp.FactorType).To(Equal("push"))
				Expect(resp.RespondedReason).To(Equal(utils.String("none")))
				Expect(resp.Details).To(BeNil())
				Expect(resp.HiddenDetails).To(BeNil())
				Expect(resp.DateResponded.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.ExpirationDate.Format(time.RFC3339)).To(Equal("2020-06-20T20:51:24Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the challenge resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Challenge("YC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get challenge response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the challenge resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateChallengeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &challenge.UpdateChallengeInput{}

			resp, err := challengeClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update challenge response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.EntitySid).To(Equal("YEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FactorSid).To(Equal("YFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("test"))
				Expect(resp.Status).To(Equal("pending"))
				Expect(resp.FactorType).To(Equal("push"))
				Expect(resp.RespondedReason).To(Equal(utils.String("none")))
				Expect(resp.Details).To(BeNil())
				Expect(resp.HiddenDetails).To(BeNil())
				Expect(resp.DateResponded.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.ExpirationDate.Format(time.RFC3339)).To(Equal("2020-06-20T20:51:24Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update challenge resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Entities/test/Challenges/YC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &challenge.UpdateChallengeInput{}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Entity("test").Challenge("YC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update challenge response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a messaging configurations client", func() {
		messagingConfigurationsClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MessagingConfigurations

		Describe("When the messaging configuration resource is successfully created", func() {
			createInput := &messaging_configurations.CreateMessagingConfigurationInput{
				Country:             "all",
				MessagingServiceSid: "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagingConfigurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := messagingConfigurationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create messaging configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Country).To(Equal("all"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessagingServiceSid).To(Equal("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all"))
			})
		})

		Describe("When the messaging configuration request does not contain a country", func() {
			createInput := &messaging_configurations.CreateMessagingConfigurationInput{
				MessagingServiceSid: "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			resp, err := messagingConfigurationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create messaging configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the messaging configuration request does not contain a messaging service sid", func() {
			createInput := &messaging_configurations.CreateMessagingConfigurationInput{
				Country: "all",
			}

			resp, err := messagingConfigurationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create messaging configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create messaging configurations api returns a 500 response", func() {
			createInput := &messaging_configurations.CreateMessagingConfigurationInput{
				Country:             "all",
				MessagingServiceSid: "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagingConfigurationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create messaging configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of messaging configurations are successfully retrieved", func() {
			pageOptions := &messaging_configurations.MessagingConfigurationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagingConfigurationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messagingConfigurationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the messaging configurations page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("messaging_configurations"))

				messagingConfigurations := resp.MessagingConfigurations
				Expect(messagingConfigurations).ToNot(BeNil())
				Expect(len(messagingConfigurations)).To(Equal(1))

				Expect(messagingConfigurations[0].Country).To(Equal("all"))
				Expect(messagingConfigurations[0].ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messagingConfigurations[0].MessagingServiceSid).To(Equal("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messagingConfigurations[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messagingConfigurations[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(messagingConfigurations[0].DateUpdated).To(BeNil())
				Expect(messagingConfigurations[0].URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all"))
			})
		})

		Describe("When the page of messaging configurations api returns a 500 response", func() {
			pageOptions := &messaging_configurations.MessagingConfigurationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagingConfigurationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the messaging configurations page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated messaging configurations are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagingConfigurationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagingConfigurationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := messagingConfigurationsClient.NewMessagingConfigurationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated messaging configurations current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated messaging configurations results should be returned", func() {
				Expect(len(paginator.MessagingConfigurations)).To(Equal(3))
			})
		})

		Describe("When the messaging configurations api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagingConfigurationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := messagingConfigurationsClient.NewMessagingConfigurationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated messaging configurations current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a messaging configuration sid", func() {
		messagingConfigurationClient := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MessagingConfiguration("all")

		Describe("When the messaging configuration resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagingConfigurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messagingConfigurationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get messaging configuration resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Country).To(Equal("all"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessagingServiceSid).To(Equal("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all"))
			})
		})

		Describe("When the messaging configuration resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/unknown",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MessagingConfiguration("unknown").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get messaging configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the messaging configuration resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateMessagingConfigurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &messaging_configuration.UpdateMessagingConfigurationInput{
				MessagingServiceSid: "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
			}

			resp, err := messagingConfigurationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update messaging configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Country).To(Equal("all"))
				Expect(resp.ServiceSid).To(Equal("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessagingServiceSid).To(Equal("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all"))
			})
		})

		Describe("When the update messaging configuration request does not contain a messaging service sid", func() {
			updateInput := &messaging_configuration.UpdateMessagingConfigurationInput{}

			resp, err := messagingConfigurationClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update messaging configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update messaging configuration resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/unknown",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &messaging_configuration.UpdateMessagingConfigurationInput{
				MessagingServiceSid: "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
			}

			resp, err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MessagingConfiguration("unknown").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update messaging configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the messaging configuration resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/all", httpmock.NewStringResponder(204, ""))

			err := messagingConfigurationClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the messaging configuration resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/unknown",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := verifySession.Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MessagingConfiguration("unknown").Delete()
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
