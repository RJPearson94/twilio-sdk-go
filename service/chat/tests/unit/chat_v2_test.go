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
	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	v2Credential "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credential"
	v2Credentials "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/bindings"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/invites"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/member"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/members"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/message"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhooks"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/role"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/roles"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user"
	v2UserBindings "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/bindings"
	v2UserChannel "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channel"
	v2UserChannels "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/users"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
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

	chatSession := chat.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V2

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
				Expect(resp.Limits).To(Equal(services.CreateServiceLimitsResponse{
					UserChannels:   250,
					ChannelMembers: 100,
				}))
				Expect(resp.Media).To(Equal(services.CreateServiceMediaResponse{
					CompatibilityMessage: "Media messages are not supported by your client",
					SizeLimitMB:          150,
				}))
				Expect(resp.Notifications).To(Equal(services.CreateServiceNotificationsResponse{
					LogEnabled: false,
					AddedToChannel: services.CreateServiceNotificationsActionResponse{
						Enabled: false,
					},
					InvitedToChannel: services.CreateServiceNotificationsActionResponse{
						Enabled: false,
					},
					RemovedFromChannel: services.CreateServiceNotificationsActionResponse{
						Enabled: false,
					},
					NewMessage: services.CreateServiceNotificationsNewMessageResponse{
						Enabled: false,
					},
				}))
				Expect(resp.PostWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PostWebhookURL).To(BeNil())
				Expect(resp.PreWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PreWebhookURL).To(BeNil())
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
				FriendlyName: "TestFriendlyName",
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

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("services"))

				serviceLimits := services.PageServiceLimitsResponse{
					UserChannels:   250,
					ChannelMembers: 100,
				}

				serviceMedia := services.PageServiceMediaResponse{
					CompatibilityMessage: "Media messages are not supported by your client",
					SizeLimitMB:          150,
				}

				serviceNotifications := services.PageServiceNotificationsResponse{
					LogEnabled: false,
					AddedToChannel: services.PageServiceNotificationsActionResponse{
						Enabled: false,
					},
					InvitedToChannel: services.PageServiceNotificationsActionResponse{
						Enabled: false,
					},
					RemovedFromChannel: services.PageServiceNotificationsActionResponse{
						Enabled: false,
					},
					NewMessage: services.PageServiceNotificationsNewMessageResponse{
						Enabled: false,
					},
				}

				services := resp.Services
				Expect(services).ToNot(BeNil())
				Expect(len(services)).To(Equal(1))

				Expect(services[0].Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].ConsumptionReportInterval).To(Equal(10))
				Expect(services[0].DefaultChannelCreatorRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].DefaultChannelRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].DefaultServiceRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].FriendlyName).To(Equal("Test"))
				Expect(services[0].Limits).To(Equal(serviceLimits))
				Expect(services[0].Media).To(Equal(serviceMedia))
				Expect(services[0].Notifications).To(Equal(serviceNotifications))
				Expect(services[0].PostWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(services[0].PostWebhookURL).To(BeNil())
				Expect(services[0].PreWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(services[0].PreWebhookURL).To(BeNil())
				Expect(services[0].ReachabilityEnabled).To(Equal(false))
				Expect(services[0].ReadStatusEnabled).To(Equal(true))
				Expect(services[0].TypingIndicatorTimeout).To(Equal(5))
				Expect(services[0].WebhookFilters).To(BeNil())
				Expect(services[0].WebhookMethod).To(BeNil())
				Expect(services[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(services[0].DateUpdated).To(BeNil())
				Expect(services[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of services api returns a 500 response", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services?Page=1&PageSize=50&PageToken=abc1234",
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

			resp, err := serviceClient.Fetch()
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
				Expect(resp.Limits).To(Equal(service.FetchServiceLimitsResponse{
					UserChannels:   250,
					ChannelMembers: 100,
				}))
				Expect(resp.Media).To(Equal(service.FetchServiceMediaResponse{
					CompatibilityMessage: "Media messages are not supported by your client",
					SizeLimitMB:          150,
				}))
				Expect(resp.Notifications).To(Equal(service.FetchServiceNotificationsResponse{
					LogEnabled: false,
					AddedToChannel: service.FetchServiceNotificationsActionResponse{
						Enabled: false,
					},
					InvitedToChannel: service.FetchServiceNotificationsActionResponse{
						Enabled: false,
					},
					RemovedFromChannel: service.FetchServiceNotificationsActionResponse{
						Enabled: false,
					},
					NewMessage: service.FetchServiceNotificationsNewMessageResponse{
						Enabled: false,
					},
				}))
				Expect(resp.PostWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PostWebhookURL).To(BeNil())
				Expect(resp.PreWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PreWebhookURL).To(BeNil())
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

			resp, err := chatSession.Service("IS71").Fetch()
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
				Expect(resp.Limits).To(Equal(service.UpdateServiceLimitsResponse{
					UserChannels:   250,
					ChannelMembers: 100,
				}))
				Expect(resp.Media).To(Equal(service.UpdateServiceMediaResponse{
					CompatibilityMessage: "Media messages are not supported by your client",
					SizeLimitMB:          150,
				}))
				Expect(resp.Notifications).To(Equal(service.UpdateServiceNotificationsResponse{
					LogEnabled: false,
					AddedToChannel: service.UpdateServiceNotificationsActionResponse{
						Enabled: false,
					},
					InvitedToChannel: service.UpdateServiceNotificationsActionResponse{
						Enabled: false,
					},
					RemovedFromChannel: service.UpdateServiceNotificationsActionResponse{
						Enabled: false,
					},
					NewMessage: service.UpdateServiceNotificationsNewMessageResponse{
						Enabled: false,
					},
				}))
				Expect(resp.PostWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PostWebhookURL).To(BeNil())
				Expect(resp.PreWebhookRetryCount).To(Equal(utils.Int(0)))
				Expect(resp.PreWebhookURL).To(BeNil())
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
				ExpectInternalServerError(err)
			})

			It("Then the create credentials response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credentials are successfully retrieved", func() {
			pageOptions := &v2Credentials.CredentialsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := credentialsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the credentials page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Credentials?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Credentials?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("credentials"))

				credentials := resp.Credentials
				Expect(credentials).ToNot(BeNil())
				Expect(len(credentials)).To(Equal(1))

				Expect(credentials[0].Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].FriendlyName).To(Equal(utils.String("Test")))
				Expect(credentials[0].Type).To(Equal("apn"))
				Expect(credentials[0].Sandbox).To(Equal(utils.String("false")))
				Expect(credentials[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(credentials[0].DateUpdated).To(BeNil())
				Expect(credentials[0].URL).To(Equal("https://chat.twilio.com/v2/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of credentials api returns a 500 response", func() {
			pageOptions := &v2Credentials.CredentialsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := credentialsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the credentials page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated credentials are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := credentialsClient.NewCredentialsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated credentials current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated credentials results should be returned", func() {
				Expect(len(paginator.Credentials)).To(Equal(3))
			})
		})

		Describe("When the credentials api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Credentials?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := credentialsClient.NewCredentialsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated credentials current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := credentialClient.Fetch()
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

			resp, err := chatSession.Credential("CR71").Fetch()
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

	Describe("Given the bindings client", func() {
		bindingsClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bindings

		Describe("When the page of bindings are successfully retrieved", func() {
			pageOptions := &bindings.BindingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := bindingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the bindings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("bindings"))

				bindings := resp.Bindings
				Expect(bindings).ToNot(BeNil())
				Expect(len(bindings)).To(Equal(1))

				Expect(bindings[0].Sid).To(Equal("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].CredentialSid).To(Equal(utils.String("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(bindings[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].BindingType).To(Equal(utils.String("gcm")))
				Expect(bindings[0].Endpoint).To(Equal(utils.String("Test")))
				Expect(bindings[0].Identity).To(Equal(utils.String("TestUser")))
				Expect(bindings[0].MessageTypes).To(BeNil())
				Expect(bindings[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(bindings[0].DateUpdated).To(BeNil())
				Expect(bindings[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of bindings api returns a 500 response", func() {
			pageOptions := &bindings.BindingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := bindingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the bindings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated bindings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := bindingsClient.NewBindingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated bindings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated bindings results should be returned", func() {
				Expect(len(paginator.Bindings)).To(Equal(3))
			})
		})

		Describe("When the bindings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := bindingsClient.NewBindingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated bindings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := bindingClient.Fetch()
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

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Fetch()
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

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("channels"))

				channels := resp.Channels
				Expect(channels).ToNot(BeNil())
				Expect(len(channels)).To(Equal(1))

				Expect(channels[0].Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].FriendlyName).To(Equal(utils.String("Test")))
				Expect(channels[0].UniqueName).To(BeNil())
				Expect(channels[0].Attributes).To(Equal(utils.String("{}")))
				Expect(channels[0].Type).To(Equal("public"))
				Expect(channels[0].CreatedBy).To(Equal("system"))
				Expect(channels[0].MembersCount).To(Equal(0))
				Expect(channels[0].MessagesCount).To(Equal(0))
				Expect(channels[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(channels[0].DateUpdated).To(BeNil())
				Expect(channels[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of channels api returns a 500 response", func() {
			pageOptions := &channels.ChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=1&PageSize=50&PageToken=abc1234",
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

			resp, err := channelClient.Fetch()
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

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CH71").Fetch()
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
				Permissions:  []string{"sendMessage"},
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
				Type:        "channel",
				Permissions: []string{"sendMessage"},
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
				Permissions:  []string{"sendMessage"},
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
				Permissions:  []string{"sendMessage"},
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
				ExpectInternalServerError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of roles are successfully retrieved", func() {
			pageOptions := &roles.RolesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := rolesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the roles page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("roles"))

				roles := resp.Roles
				Expect(roles).ToNot(BeNil())
				Expect(len(roles)).To(Equal(1))

				Expect(roles[0].Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].FriendlyName).To(Equal("Test"))
				Expect(roles[0].Type).To(Equal("channel"))
				Expect(roles[0].Permissions).To(Equal([]string{"sendMessage"}))
				Expect(roles[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(roles[0].DateUpdated).To(BeNil())
				Expect(roles[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of roles api returns a 500 response", func() {
			pageOptions := &roles.RolesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := rolesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the roles page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated roles are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := rolesClient.NewRolesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated roles current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated roles results should be returned", func() {
				Expect(len(paginator.Roles)).To(Equal(3))
			})
		})

		Describe("When the roles api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := rolesClient.NewRolesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated roles current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := roleClient.Fetch()
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

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Fetch()
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
				Permissions: []string{"sendMessage", "leaveChannel"},
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
				Permissions: []string{"sendMessage", "leaveChannel"},
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
				ExpectInternalServerError(err)
			})

			It("Then the create user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of users are successfully retrieved", func() {
			pageOptions := &users.UsersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := usersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the users page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("users"))

				users := resp.Users
				Expect(users).ToNot(BeNil())
				Expect(len(users)).To(Equal(1))

				Expect(users[0].Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].Identity).To(Equal("Test"))
				Expect(users[0].Attributes).To(BeNil())
				Expect(users[0].IsOnline).To(Equal(utils.Bool(true)))
				Expect(users[0].IsNotifiable).To(BeNil())
				Expect(users[0].FriendlyName).To(BeNil())
				Expect(users[0].JoinedChannelsCount).To(Equal(utils.Int(0)))
				Expect(users[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(users[0].DateUpdated).To(BeNil())
				Expect(users[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of users api returns a 500 response", func() {
			pageOptions := &users.UsersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := usersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the users page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated users are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := usersClient.NewUsersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated users current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated users results should be returned", func() {
				Expect(len(paginator.Users)).To(Equal(3))
			})
		})

		Describe("When the users api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := usersClient.NewUsersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated users current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := userClient.Fetch()
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

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Fetch()
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

	Describe("Given the user bindings client", func() {
		userBindingsClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bindings

		Describe("When the page of userBindings are successfully retrieved", func() {
			pageOptions := &v2UserBindings.UserBindingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userBindingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := userBindingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the user bindings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("bindings"))

				bindings := resp.Bindings
				Expect(bindings).ToNot(BeNil())
				Expect(len(bindings)).To(Equal(1))

				Expect(bindings[0].Sid).To(Equal("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].CredentialSid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].Identity).To(Equal(utils.String("TestUser")))
				Expect(bindings[0].Endpoint).To(Equal(utils.String("Test")))
				Expect(bindings[0].MessageTypes).To(Equal(&[]string{"removed_from_channel", "new_message"}))
				Expect(bindings[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(bindings[0].DateUpdated).To(BeNil())
				Expect(bindings[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of user bindings api returns a 500 response", func() {
			pageOptions := &v2UserBindings.UserBindingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := userBindingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the user bindings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated user bindings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userBindingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userBindingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := userBindingsClient.NewUserBindingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated user bindings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated user bindings results should be returned", func() {
				Expect(len(paginator.Bindings)).To(Equal(3))
			})
		})

		Describe("When the user bindings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userBindingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := userBindingsClient.NewUserBindingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated user bindings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
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

			resp, err := userBindingClient.Fetch()
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

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Fetch()
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

	Describe("Given the user channels client", func() {
		userChannelsClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channels

		Describe("When the page of user channels are successfully retrieved", func() {
			pageOptions := &v2UserChannels.UserChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userChannelsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := userChannelsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the user channels page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("channels"))

				userChannels := resp.Channels
				Expect(userChannels).ToNot(BeNil())
				Expect(len(userChannels)).To(Equal(1))

				Expect(userChannels[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(userChannels[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(userChannels[0].UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(userChannels[0].ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(userChannels[0].MemberSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(userChannels[0].Status).To(Equal("joined"))
				Expect(userChannels[0].LastConsumedMessageIndex).To(Equal(utils.Int(5)))
				Expect(userChannels[0].UnreadMessagesCount).To(Equal(utils.Int(5)))
				Expect(userChannels[0].NotificationLevel).To(Equal(utils.String("muted")))
				Expect(userChannels[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of user channels api returns a 500 response", func() {
			pageOptions := &v2UserChannels.UserChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := userChannelsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the user channels page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated user channels are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userChannelsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := userChannelsClient.NewUserChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated user channels current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated user channels results should be returned", func() {
				Expect(len(paginator.Channels)).To(Equal(3))
			})
		})

		Describe("When the user channels api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := userChannelsClient.NewUserChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated user channels current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a user channel sid", func() {
		userChannelClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the user channel is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/userChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := userChannelClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get user channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MemberSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("joined"))
				Expect(resp.LastConsumedMessageIndex).To(Equal(utils.Int(5)))
				Expect(resp.UnreadMessagesCount).To(Equal(utils.Int(5)))
				Expect(resp.NotificationLevel).To(Equal(utils.String("muted")))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user channel api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CH71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get user channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user channel is successfully retrieved", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateUserChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &v2UserChannel.UpdateUserChannelInput{
				NotificationLevel: utils.String("default"),
			}

			resp, err := userChannelClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get user channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UserSid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MemberSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("joined"))
				Expect(resp.LastConsumedMessageIndex).To(Equal(utils.Int(5)))
				Expect(resp.UnreadMessagesCount).To(Equal(utils.Int(5)))
				Expect(resp.NotificationLevel).To(Equal(utils.String("default")))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user channel api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &v2UserChannel.UpdateUserChannelInput{
				NotificationLevel: utils.String("default"),
			}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update user channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the channel invites client", func() {
		channelInvitesClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Invites

		Describe("When the channel invites is successfully created", func() {
			createInput := &invites.CreateChannelInviteInput{
				Identity: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelInviteResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := channelInvitesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create channel invite response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.CreatedBy).To(Equal(utils.String("created_by")))
				Expect(resp.Identity).To(Equal("identity"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel invite does not contain a identity", func() {
			createInput := &invites.CreateChannelInviteInput{}

			resp, err := channelInvitesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create channel invite  response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel invite api returns a 500 response", func() {
			createInput := &invites.CreateChannelInviteInput{
				Identity: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelInvitesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create channel invite response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of channel invites are successfully retrieved", func() {
			pageOptions := &invites.ChannelInvitesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelInvitesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelInvitesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the channel invites page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("invites"))

				invites := resp.Invites
				Expect(invites).ToNot(BeNil())
				Expect(len(invites)).To(Equal(1))

				Expect(invites[0].Sid).To(Equal("INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(invites[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(invites[0].ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(invites[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(invites[0].RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(invites[0].CreatedBy).To(Equal(utils.String("created_by")))
				Expect(invites[0].Identity).To(Equal("identity"))
				Expect(invites[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(invites[0].DateUpdated).To(BeNil())
				Expect(invites[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of channel invites api returns a 500 response", func() {
			pageOptions := &invites.ChannelInvitesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelInvitesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the channel invites page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated channel invites are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelInvitesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelInvitesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := channelInvitesClient.NewChannelInvitesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated channel invites current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated channel invites results should be returned", func() {
				Expect(len(paginator.Invites)).To(Equal(3))
			})
		})

		Describe("When the channel invites api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelInvitesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := channelInvitesClient.NewChannelInvitesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated channel invites current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a channel invite sid", func() {
		channelInviteClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Invite("INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the channel invite is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelInviteResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelInviteClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get channel invite response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.CreatedBy).To(Equal(utils.String("created_by")))
				Expect(resp.Identity).To(Equal("identity"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel invite api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/IN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Invite("IN71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get channel invite response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel invite is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/INXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := channelInviteClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the channel invite api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Invites/IN71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Invite("IN71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the channel members client", func() {
		channelMembersClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Members

		Describe("When the channel invite is successfully created", func() {
			createInput := &members.CreateChannelMemberInput{
				Identity: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMemberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := channelMembersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create channel member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.LastConsumedMessageIndex).To(Equal(utils.Int(20)))
				Expect(resp.LastConsumedTimestamp.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.Attributes).To(Equal(utils.String("{}")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel member does not contain a identity", func() {
			createInput := &members.CreateChannelMemberInput{}

			resp, err := channelMembersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create channel member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel member api returns a 500 response", func() {
			createInput := &members.CreateChannelMemberInput{
				Identity: "Test",
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelMembersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create channel member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of channel members are successfully retrieved", func() {
			pageOptions := &members.ChannelMembersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMembersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelMembersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the channel members page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("members"))

				members := resp.Members
				Expect(members).ToNot(BeNil())
				Expect(len(members)).To(Equal(1))

				Expect(members[0].Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(members[0].RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(members[0].Identity).To(Equal("Test"))
				Expect(members[0].LastConsumedMessageIndex).To(Equal(utils.Int(20)))
				Expect(members[0].LastConsumedTimestamp.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(members[0].Attributes).To(Equal(utils.String("{}")))
				Expect(members[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(members[0].DateUpdated).To(BeNil())
				Expect(members[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of channel members api returns a 500 response", func() {
			pageOptions := &members.ChannelMembersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelMembersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the channel members page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated channel members are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMembersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMembersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := channelMembersClient.NewChannelMembersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated channel members current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated channel members results should be returned", func() {
				Expect(len(paginator.Members)).To(Equal(3))
			})
		})

		Describe("When the channel members api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMembersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := channelMembersClient.NewChannelMembersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated channel members current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a channel member sid", func() {
		channelMemberClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the channel member is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMemberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelMemberClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get channel member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.LastConsumedMessageIndex).To(Equal(utils.Int(20)))
				Expect(resp.LastConsumedTimestamp.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.Attributes).To(Equal(utils.String("{}")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel member api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("MB71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get channel member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel member is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateChannelMemberResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &member.UpdateChannelMemberInput{}

			resp, err := channelMemberClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update channel member response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.LastConsumedMessageIndex).To(Equal(utils.Int(20)))
				Expect(resp.LastConsumedTimestamp.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.Attributes).To(Equal(utils.String("{}")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel member api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &member.UpdateChannelMemberInput{}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("MB71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update channel member response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel member is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := channelMemberClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the channel member api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Member("MB71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the channel messages client", func() {
		channelMessagesClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Messages

		Describe("When the channel message is successfully created", func() {
			createInput := &messages.CreateChannelMessageInput{
				Body: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := channelMessagesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create channel message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.To).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Attributes).To(BeNil())
				Expect(resp.LastUpdatedBy).To(Equal(utils.String("system")))
				Expect(resp.WasEdited).To(Equal(utils.Bool(false)))
				Expect(resp.From).To(Equal(utils.String("system")))
				Expect(resp.Body).To(Equal(utils.String("Test")))
				Expect(resp.Index).To(Equal(utils.Int(0)))
				Expect(resp.Type).To(Equal(utils.String("text")))
				Expect(resp.Media).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel message api returns a 500 response", func() {
			createInput := &messages.CreateChannelMessageInput{
				Body: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelMessagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create channel message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of channel messages are successfully retrieved", func() {
			pageOptions := &messages.ChannelMessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMessagesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelMessagesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the channel messages page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("messages"))

				messages := resp.Messages
				Expect(messages).ToNot(BeNil())
				Expect(len(messages)).To(Equal(1))

				Expect(messages[0].Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].To).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(messages[0].Attributes).To(BeNil())
				Expect(messages[0].LastUpdatedBy).To(Equal(utils.String("system")))
				Expect(messages[0].WasEdited).To(Equal(utils.Bool(false)))
				Expect(messages[0].From).To(Equal(utils.String("system")))
				Expect(messages[0].Body).To(Equal(utils.String("Test")))
				Expect(messages[0].Index).To(Equal(utils.Int(0)))
				Expect(messages[0].Type).To(Equal(utils.String("text")))
				Expect(messages[0].Media).To(BeNil())
				Expect(messages[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(messages[0].DateUpdated).To(BeNil())
				Expect(messages[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of channel messages api returns a 500 response", func() {
			pageOptions := &messages.ChannelMessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelMessagesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the channel messages page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated channel messages are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMessagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMessagesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := channelMessagesClient.NewChannelMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated channel messages current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated channel messages results should be returned", func() {
				Expect(len(paginator.Messages)).To(Equal(3))
			})
		})

		Describe("When the channel messages api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMessagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := channelMessagesClient.NewChannelMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated channel messages current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a channel message sid", func() {
		channelMessageClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the channel message is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelMessageClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get channel message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.To).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Attributes).To(BeNil())
				Expect(resp.LastUpdatedBy).To(Equal(utils.String("system")))
				Expect(resp.WasEdited).To(Equal(utils.Bool(false)))
				Expect(resp.From).To(Equal(utils.String("system")))
				Expect(resp.Body).To(Equal(utils.String("Test")))
				Expect(resp.Index).To(Equal(utils.Int(0)))
				Expect(resp.Type).To(Equal(utils.String("text")))
				Expect(resp.Media).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel message api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get channel message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel message is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateChannelMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &message.UpdateChannelMessageInput{
				Body: utils.String("Hello World"),
			}

			resp, err := channelMessageClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update channel message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.To).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Attributes).To(BeNil())
				Expect(resp.LastUpdatedBy).To(Equal(utils.String("system")))
				Expect(resp.WasEdited).To(Equal(utils.Bool(true)))
				Expect(resp.From).To(Equal(utils.String("system")))
				Expect(resp.Body).To(Equal(utils.String("Hello World")))
				Expect(resp.Index).To(Equal(utils.Int(0)))
				Expect(resp.Type).To(Equal(utils.String("text")))
				Expect(resp.Media).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel message api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &message.UpdateChannelMessageInput{}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update channel message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel message is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := channelMessageClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the channel member api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the channel webhooks client", func() {
		channelWebhooksClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhooks

		Describe("When the channel webhook is successfully created", func() {
			createInput := &webhooks.CreateChannelWebhookInput{
				Type: "studio",
				Configuration: &webhooks.CreateChannelWebhookConfigurationInput{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				},
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := channelWebhooksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create channel webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(webhooks.CreateChannelWebhookConfigurationResponse{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel webhook does not contain a type", func() {
			createInput := &webhooks.CreateChannelWebhookInput{}

			resp, err := channelWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create channel webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel webhook api returns a 500 response", func() {
			createInput := &webhooks.CreateChannelWebhookInput{
				Type: "studio",
				Configuration: &webhooks.CreateChannelWebhookConfigurationInput{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				},
			}

			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create channel webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of channel webhooks are successfully retrieved", func() {
			pageOptions := &webhooks.ChannelWebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelWebhooksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelWebhooksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the channelWebhooks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("webhooks"))

				channelWebhooks := resp.Webhooks
				Expect(channelWebhooks).ToNot(BeNil())
				Expect(len(channelWebhooks)).To(Equal(1))

				Expect(channelWebhooks[0].Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channelWebhooks[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channelWebhooks[0].ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channelWebhooks[0].ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channelWebhooks[0].Type).To(Equal("studio"))
				Expect(channelWebhooks[0].Configuration).To(Equal(webhooks.PageChannelWebhookConfigurationResponse{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(channelWebhooks[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(channelWebhooks[0].DateUpdated).To(BeNil())
				Expect(channelWebhooks[0].URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of channel webhooks api returns a 500 response", func() {
			pageOptions := &webhooks.ChannelWebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelWebhooksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the channel webhooks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated channel webhooks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelWebhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelWebhooksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := channelWebhooksClient.NewChannelWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated channel webhooks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated channel webhooks results should be returned", func() {
				Expect(len(paginator.Webhooks)).To(Equal(3))
			})
		})

		Describe("When the channel webhooks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelWebhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := channelWebhooksClient.NewChannelWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated channel webhooks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a channel webhook sid", func() {
		channelWebhookClient := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the channel webhook is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/channelWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelWebhookClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get channel webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(webhook.FetchChannelWebhookConfigurationResponse{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel webhook api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get channel webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel webhook is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateChannelWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &webhook.UpdateChannelWebhookInput{}

			resp, err := channelWebhookClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update channel webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(webhook.UpdateChannelWebhookConfigurationResponse{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the channel webhook api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &webhook.UpdateChannelWebhookInput{}

			resp, err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update channel webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the channel webhook is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := channelWebhookClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the channel webhook api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := chatSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Delete()
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
