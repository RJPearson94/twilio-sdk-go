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

	"github.com/RJPearson94/twilio-sdk-go/service/proxy"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session/participant/message_interactions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session/participants"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/sessions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Proxy V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	proxySession := proxy.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(proxySession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the services client", func() {
		servicesClient := proxySession.Services

		Describe("When the service is successfully created", func() {
			createInput := &services.CreateServiceInput{
				UniqueName: "Test",
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services",
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
				Expect(resp.Sid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatInstanceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.UniqueName).To(Equal("Test"))
				Expect(resp.DefaultTtl).To(Equal(utils.Int(3600)))
				Expect(resp.CallbackURL).To(Equal(utils.String("http://www.callback_url.com")))
				Expect(resp.GeoMatchLevel).To(Equal(utils.String("country")))
				Expect(resp.NumberSelectionBehavior).To(Equal(utils.String("prefer_sticky")))
				Expect(resp.InterceptCallbackURL).To(Equal(utils.String("http://www.intercept_callback_url.com")))
				Expect(resp.OutOfSessionCallbackURL).To(Equal(utils.String("http://www.out_of_session_callback_url.com")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service does not contain a unique name", func() {
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
				UniqueName: "Tet",
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services",
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
		serviceClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatInstanceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.UniqueName).To(Equal("Test"))
				Expect(resp.DefaultTtl).To(Equal(utils.Int(3600)))
				Expect(resp.CallbackURL).To(Equal(utils.String("http://www.callback_url.com")))
				Expect(resp.GeoMatchLevel).To(Equal(utils.String("country")))
				Expect(resp.NumberSelectionBehavior).To(Equal(utils.String("prefer_sticky")))
				Expect(resp.InterceptCallbackURL).To(Equal(utils.String("http://www.intercept_callback_url.com")))
				Expect(resp.OutOfSessionCallbackURL).To(Equal(utils.String("http://www.out_of_session_callback_url.com")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := proxySession.Service("KS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				UniqueName: utils.String("Test 2"),
			}

			resp, err := serviceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update service response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatInstanceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.UniqueName).To(Equal("Test 2"))
				Expect(resp.DefaultTtl).To(Equal(utils.Int(3600)))
				Expect(resp.CallbackURL).To(Equal(utils.String("http://www.callback_url.com")))
				Expect(resp.GeoMatchLevel).To(Equal(utils.String("country")))
				Expect(resp.NumberSelectionBehavior).To(Equal(utils.String("prefer_sticky")))
				Expect(resp.InterceptCallbackURL).To(Equal(utils.String("http://www.intercept_callback_url.com")))
				Expect(resp.OutOfSessionCallbackURL).To(Equal(utils.String("http://www.out_of_session_callback_url.com")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-08-01T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update service response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &service.UpdateServiceInput{
				UniqueName: utils.String("Test 2"),
			}

			resp, err := proxySession.Service("KS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the service api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := proxySession.Service("KS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the phone numbers client", func() {
		phoneNumbersClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumbers

		Describe("When the phone number is successfully created", func() {
			createInput := &phone_numbers.CreatePhoneNumberInput{
				Sid: utils.String("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := phoneNumbersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create phone number response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PhoneNumber).To(Equal(utils.String("+123456789")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Capabilities).To(Equal(&phone_numbers.CreatePhoneNumberResponseCapabilities{
					VoiceInbound:             utils.Bool(true),
					SmsInbound:               utils.Bool(true),
					RestrictionVoiceDomestic: utils.Bool(false),
					FaxOutbound:              utils.Bool(false),
					RestrictionSmsDomestic:   utils.Bool(false),
					FaxInbound:               utils.Bool(false),
					SmsOutbound:              utils.Bool(true),
					RestrictionMmsDomestic:   utils.Bool(false),
					MmsOutbound:              utils.Bool(true),
					VoiceOutbound:            utils.Bool(true),
					MmsInbound:               utils.Bool(true),
					RestrictionFaxDomestic:   utils.Bool(false),
					SipTrunking:              utils.Bool(true),
				}))
				Expect(resp.IsoCountry).To(Equal(utils.String("US")))
				Expect(resp.IsReserved).To(Equal(utils.Bool(false)))
				Expect(resp.InUse).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the phone number api returns a 500 response", func() {
			createInput := &phone_numbers.CreatePhoneNumberInput{
				Sid: utils.String("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := phoneNumbersClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a phone number sid", func() {
		phoneNumberClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the phone number is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/phoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := phoneNumberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get phone number response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PhoneNumber).To(Equal(utils.String("+123456789")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Capabilities).To(Equal(&phone_number.FetchPhoneNumberResponseCapabilities{
					VoiceInbound:             utils.Bool(true),
					SmsInbound:               utils.Bool(true),
					RestrictionVoiceDomestic: utils.Bool(false),
					FaxOutbound:              utils.Bool(false),
					RestrictionSmsDomestic:   utils.Bool(false),
					FaxInbound:               utils.Bool(false),
					SmsOutbound:              utils.Bool(true),
					RestrictionMmsDomestic:   utils.Bool(false),
					MmsOutbound:              utils.Bool(true),
					VoiceOutbound:            utils.Bool(true),
					MmsInbound:               utils.Bool(true),
					RestrictionFaxDomestic:   utils.Bool(false),
					SipTrunking:              utils.Bool(true),
				}))
				Expect(resp.IsoCountry).To(Equal(utils.String("US")))
				Expect(resp.IsReserved).To(Equal(utils.Bool(false)))
				Expect(resp.InUse).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the phone number api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the phone number is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatePhoneNumberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &phone_number.UpdatePhoneNumberInput{
				IsReserved: utils.Bool(true),
			}

			resp, err := phoneNumberClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update phone number response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.PhoneNumber).To(Equal(utils.String("+123456789")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Capabilities).To(Equal(&phone_number.UpdatePhoneNumberResponseCapabilities{
					VoiceInbound:             utils.Bool(true),
					SmsInbound:               utils.Bool(true),
					RestrictionVoiceDomestic: utils.Bool(false),
					FaxOutbound:              utils.Bool(false),
					RestrictionSmsDomestic:   utils.Bool(false),
					FaxInbound:               utils.Bool(false),
					SmsOutbound:              utils.Bool(true),
					RestrictionMmsDomestic:   utils.Bool(false),
					MmsOutbound:              utils.Bool(true),
					VoiceOutbound:            utils.Bool(true),
					MmsInbound:               utils.Bool(true),
					RestrictionFaxDomestic:   utils.Bool(false),
					SipTrunking:              utils.Bool(true),
				}))
				Expect(resp.IsoCountry).To(Equal(utils.String("US")))
				Expect(resp.IsReserved).To(Equal(utils.Bool(true)))
				Expect(resp.InUse).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-08-01T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the phone number api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &phone_number.UpdatePhoneNumberInput{
				IsReserved: utils.Bool(true),
			}

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update phone number response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the phone number is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := phoneNumberClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the phone number api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").PhoneNumber("PN71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the short codes client", func() {
		shortCodesClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ShortCodes

		Describe("When the short code is successfully created", func() {
			createInput := &short_codes.CreateShortCodeInput{
				Sid: "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/shortCodeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := shortCodesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create short code response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ShortCode).To(Equal(utils.String("12345")))
				Expect(resp.Capabilities).To(Equal(&short_codes.CreateShortCodeResponseCapabilities{
					VoiceInbound: utils.Bool(false),
					SmsOutbound:  utils.Bool(true),
				}))
				Expect(resp.IsoCountry).To(Equal(utils.String("US")))
				Expect(resp.IsReserved).To(Equal(utils.Bool(false)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the short code does not contain a sid", func() {
			createInput := &short_codes.CreateShortCodeInput{}

			resp, err := shortCodesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create short code response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the short code api returns a 500 response", func() {
			createInput := &short_codes.CreateShortCodeInput{
				Sid: "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := shortCodesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create short code response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a short code sid", func() {
		shortCodeClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ShortCode("SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the short code is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/shortCodeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := shortCodeClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get short code response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ShortCode).To(Equal(utils.String("12345")))
				Expect(resp.Capabilities).To(Equal(&short_code.FetchShortCodeResponseCapabilities{
					VoiceInbound: utils.Bool(false),
					SmsOutbound:  utils.Bool(true),
				}))
				Expect(resp.IsoCountry).To(Equal(utils.String("US")))
				Expect(resp.IsReserved).To(Equal(utils.Bool(false)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the short code api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ShortCode("SC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get short code response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the short code is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateShortCodeResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &short_code.UpdateShortCodeInput{
				IsReserved: utils.Bool(true),
			}

			resp, err := shortCodeClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update short code response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ShortCode).To(Equal(utils.String("12345")))
				Expect(resp.Capabilities).To(Equal(&short_code.UpdateShortCodeResponseCapabilities{
					VoiceInbound: utils.Bool(false),
					SmsOutbound:  utils.Bool(true),
				}))
				Expect(resp.IsoCountry).To(Equal(utils.String("US")))
				Expect(resp.IsReserved).To(Equal(utils.Bool(true)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-08-01T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the short code api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &short_code.UpdateShortCodeInput{
				IsReserved: utils.Bool(true),
			}

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ShortCode("SC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update short code response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the short code is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := shortCodeClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the short codes api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes/SC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").ShortCode("SC71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the session client", func() {
		sessionsClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Sessions

		Describe("When the session is successfully created", func() {
			createInput := &sessions.CreateSessionInput{
				UniqueName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sessionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := sessionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create session response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal(utils.String("open")))
				Expect(resp.UniqueName).To(Equal("Test"))
				Expect(resp.DateStarted.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateEnded).To(BeNil())
				Expect(resp.DateLastInteraction.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateExpiry.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.Ttl).To(Equal(utils.Int(3600)))
				Expect(resp.Mode).To(Equal(utils.String("voice-and-message")))
				Expect(resp.ClosedReason).To(Equal(utils.String("")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the session api returns a 500 response", func() {
			createInput := &sessions.CreateSessionInput{
				UniqueName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := sessionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create session response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a session sid", func() {
		sessionClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the session is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/sessionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := sessionClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get session response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal(utils.String("open")))
				Expect(resp.UniqueName).To(Equal("Test"))
				Expect(resp.DateStarted.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateEnded).To(BeNil())
				Expect(resp.DateLastInteraction.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateExpiry.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.Ttl).To(Equal(utils.Int(3600)))
				Expect(resp.Mode).To(Equal(utils.String("voice-and-message")))
				Expect(resp.ClosedReason).To(Equal(utils.String("")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the session api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get session response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the session is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateSessionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &session.UpdateSessionInput{
				Status: utils.String("in-progress"),
			}

			resp, err := sessionClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update session response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal(utils.String("in-progress")))
				Expect(resp.UniqueName).To(Equal("Test"))
				Expect(resp.DateStarted.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateEnded).To(BeNil())
				Expect(resp.DateLastInteraction.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateExpiry.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.Ttl).To(Equal(utils.Int(3600)))
				Expect(resp.Mode).To(Equal(utils.String("voice-and-message")))
				Expect(resp.ClosedReason).To(Equal(utils.String("")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the session api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &session.UpdateSessionInput{
				Status: utils.String("in-progress"),
			}

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update session response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the session is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := sessionClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the session api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KC71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a interaction sid", func() {
		interactionClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Interaction("KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the interaction is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Interactions/KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/interactionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := interactionClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get interaction response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.SessionSid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Data).To(Equal(utils.String("{ \"body\": \"some message\" }")))
				Expect(resp.InboundParticipantSid).To(Equal(utils.String("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.InboundResourceSid).To(Equal(utils.String("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.InboundResourceStatus).To(Equal(utils.String("sent")))
				Expect(resp.InboundResourceType).To(Equal(utils.String("Message")))
				Expect(resp.InboundResourceURL).To(BeNil())
				Expect(resp.OutboundParticipantSid).To(Equal(utils.String("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.OutboundResourceSid).To(Equal(utils.String("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.OutboundResourceStatus).To(Equal(utils.String("sent")))
				Expect(resp.OutboundResourceType).To(Equal(utils.String("Message")))
				Expect(resp.OutboundResourceURL).To(BeNil())
				Expect(resp.Type).To(Equal(utils.String("message")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Interactions/KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the interaction api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Interactions/KI71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Interaction("KI71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get interaction response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the interaction is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Interactions/KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := interactionClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the interaction api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Interactions/KI71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Interaction("KI71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the participant client", func() {
		participantsClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participants

		Describe("When the participant is successfully created", func() {
			createInput := &participants.CreateParticipantInput{
				Identifier: "+123456789",
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.SessionSid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identifier).To(Equal("+123456789"))
				Expect(resp.ProxyIdentifier).To(Equal(utils.String("+123456788")))
				Expect(resp.ProxyIdentifierSid).To(Equal(utils.String("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateDeleted).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant does not contain an identifier", func() {
			createInput := &participants.CreateParticipantInput{}

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant api returns a 500 response", func() {
			createInput := &participants.CreateParticipantInput{
				Identifier: "+123456789",
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a participant sid", func() {
		participantClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the interaction is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.SessionSid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identifier).To(Equal("+123456789"))
				Expect(resp.ProxyIdentifier).To(Equal(utils.String("+123456788")))
				Expect(resp.ProxyIdentifierSid).To(Equal(utils.String("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateDeleted).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("KP71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := participantClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KP71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("KP71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the message interactions client", func() {
		messageInteractionsClient := proxySession.Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").MessageInteractions

		Describe("When the message interaction is successfully created", func() {
			interactionData, _ := ioutil.ReadFile("testdata/interactionData.json")

			createInput := &message_interactions.CreateMessageInteractionInput{
				Body: utils.String(string(interactionData)),
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessageInteractions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageInteractionResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := messageInteractionsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create message interaction response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.SessionSid).To(Equal("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Data).To(Equal(utils.String("{ \"body\": \"some message\" }")))
				Expect(resp.InboundParticipantSid).To(BeNil())
				Expect(resp.InboundResourceSid).To(BeNil())
				Expect(resp.InboundResourceStatus).To(BeNil())
				Expect(resp.InboundResourceType).To(BeNil())
				Expect(resp.InboundResourceURL).To(BeNil())
				Expect(resp.OutboundParticipantSid).To(Equal(utils.String("KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.OutboundResourceSid).To(Equal(utils.String("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.OutboundResourceStatus).To(Equal(utils.String("sent")))
				Expect(resp.OutboundResourceType).To(Equal(utils.String("Message")))
				Expect(resp.OutboundResourceURL).To(BeNil())
				Expect(resp.Type).To(Equal(utils.String("message")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-07-30T20:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessageInteractions/KIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the message interaction api returns a 500 response", func() {
			interactionData, _ := ioutil.ReadFile("testdata/interactionData.json")

			createInput := &message_interactions.CreateMessageInteractionInput{
				Body: utils.String(string(interactionData)),
			}

			httpmock.RegisterResponder("POST", "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions/KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/KPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessageInteractions",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messageInteractionsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create message interaction response should be nil", func() {
				Expect(resp).To(BeNil())
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
