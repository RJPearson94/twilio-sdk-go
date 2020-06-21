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
	v2Credential "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credential"
	v2Credentials "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/role"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/roles"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/users"
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

	Describe("Given the credentials client", func() {
		credentialsClient := chatSession.Credentials

		Describe("When the credential is successfully created", func() {
			createInput := &v2Credentials.CreateCredentialInput{
				Type: "apn",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Credentials",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := credentialsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Type).To(Equal("apn"))
				Expect(resp.Sandbox).To(Equal(utils.String("false")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the credential does not contain a type", func() {
			createInput := &v2Credentials.CreateCredentialInput{}

			resp, err := credentialsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential api returns a 500 response", func() {
			createInput := &v2Credentials.CreateCredentialInput{
				Type: "apn",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Credentials",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialsClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create credentials response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a credential sid", func() {
		credentialClient := chatSession.Credential("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Type).To(Equal("apn"))
				Expect(resp.Sandbox).To(Equal(utils.String("false")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the credential api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials/CR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Credential("CR71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCredentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &v2Credential.UpdateCredentialInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := credentialClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test 2")))
				Expect(resp.Type).To(Equal("apn"))
				Expect(resp.Sandbox).To(Equal(utils.String("false")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T22:50:24Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update credential response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Credentials/CR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &v2Credential.UpdateCredentialInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := chatSession.Credential("CR71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := credentialClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the credential api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Credentials/CR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Credential("CR71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a binding sid", func() {
		bindingClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the binding is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := bindingClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get binding response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CredentialSid).To(Equal(utils.String("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.BindingType).To(Equal(utils.String("gcm")))
				Expect(resp.Endpoint).To(Equal(utils.String("Test")))
				Expect(resp.Identity).To(Equal(utils.String("TestUser")))
				Expect(resp.MessageTypes).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the binding api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get binding response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the binding is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := bindingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the binding api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the channels client", func() {
		channelsClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channels

		Describe("When the channel is successfully created", func() {
			createInput := &channels.CreateChannelInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
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
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Attributes).To(Equal(utils.String("{}")))
				Expect(resp.Type).To(Equal("public"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.MembersCount).To(Equal(0))
				Expect(resp.MessagesCount).To(Equal(0))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel api returns a 500 response", func() {
			createInput := &channels.CreateChannelInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
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
		channelClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the channel is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Attributes).To(Equal(utils.String("{}")))
				Expect(resp.Type).To(Equal("public"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.MembersCount).To(Equal(0))
				Expect(resp.MessagesCount).To(Equal(0))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CH71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &channel.UpdateChannelInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := channelClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test 2")))
				Expect(resp.UniqueName).To(BeNil())
				Expect(resp.Attributes).To(Equal(utils.String("{}")))
				Expect(resp.Type).To(Equal("public"))
				Expect(resp.CreatedBy).To(Equal("system"))
				Expect(resp.MembersCount).To(Equal(0))
				Expect(resp.MessagesCount).To(Equal(0))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update channel response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &channel.UpdateChannelInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := channelClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the channel api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the roles client", func() {
		rolesClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Roles

		Describe("When the role is successfully created", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "Test",
				Type:         "channel",
				Permission:   []string{"sendMessage"},
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := rolesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create role response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Type).To(Equal("channel"))
				Expect(resp.Permissions).To(Equal([]string{"sendMessage"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role does not contain a friendly name", func() {
			createInput := &roles.CreateRoleInput{
				Type:       "channel",
				Permission: []string{"sendMessage"},
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role does not contain a type", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "Test",
				Permission:   []string{"sendMessage"},
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role does not contain permission", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "Test",
				Type:         "channel",
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role api returns a 500 response", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "Test",
				Type:         "channel",
				Permission:   []string{"sendMessage"},
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a role sid", func() {
		roleClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the role is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roleClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get role response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Type).To(Equal("channel"))
				Expect(resp.Permissions).To(Equal([]string{"sendMessage"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRoleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &role.UpdateRoleInput{
				Permission: []string{"sendMessage", "leaveChannel"},
			}

			resp, err := roleClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update role response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test"))
				Expect(resp.Type).To(Equal("channel"))
				Expect(resp.Permissions).To(Equal([]string{"sendMessage", "leaveChannel"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update role response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &role.UpdateRoleInput{
				Permission: []string{"sendMessage", "leaveChannel"},
			}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := roleClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the role api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the users client", func() {
		usersClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Users

		Describe("When the role is successfully created", func() {
			createInput := &users.CreateUserInput{
				Identity: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := usersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create user response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.Attributes).To(BeNil())
				Expect(resp.IsOnline).To(Equal(utils.Bool(true)))
				Expect(resp.IsNotifiable).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.JoinedChannelsCount).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user does not contain a identity", func() {
			createInput := &users.CreateUserInput{}

			resp, err := usersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user api returns a 500 response", func() {
			createInput := &users.CreateUserInput{
				Identity: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := usersClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a user sid", func() {
		userClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the user is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := userClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get user response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.Attributes).To(BeNil())
				Expect(resp.IsOnline).To(Equal(utils.Bool(true)))
				Expect(resp.IsNotifiable).To(BeNil())
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.JoinedChannelsCount).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateUserResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &user.UpdateUserInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := userClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update user response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.Attributes).To(BeNil())
				Expect(resp.IsOnline).To(Equal(utils.Bool(true)))
				Expect(resp.IsNotifiable).To(BeNil())
				Expect(resp.FriendlyName).To(Equal(utils.String("Test 2")))
				Expect(resp.JoinedChannelsCount).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update user response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &user.UpdateUserInput{}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := userClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the user api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a user binding sid", func() {
		userBindingClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the user is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userBindingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := userBindingClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get user binding response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.CredentialSid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal(utils.String("TestUser")))
				Expect(resp.Endpoint).To(Equal(utils.String("Test")))
				Expect(resp.MessageTypes).To(Equal(&[]string{"removed_from_channel", "new_message"}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user binding api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get user binding response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user binding is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := userBindingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the user binding api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Delete()
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
