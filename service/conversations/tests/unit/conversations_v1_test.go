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

	conversationsClient "github.com/RJPearson94/twilio-sdk-go/service/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/message"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/message/delivery_receipts"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/participants"
	conversationWebhook "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhook"
	conversationWebhooks "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhooks"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/credential"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/role"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/roles"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/bindings"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration/notification"
	serviceConversation "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation"
	serviceConversationMessage "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/message"
	serviceDeliveryReceipts "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/message/delivery_receipts"
	serviceConversationMessages "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/messages"
	serviceParticipant "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/participant"
	serviceParticipants "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/participants"
	serviceConversationWebhook "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/webhook"
	serviceConversationWebhooks "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/webhooks"
	serviceConversations "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversations"
	serviceRole "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/role"
	serviceRoles "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/roles"
	serviceUser "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/user"
	serviceUsers "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/users"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/user"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/users"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/webhook"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Conversation V1", func() {
	creds, err := sessionCredentials.New(sessionCredentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	conversationsSession := conversationsClient.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(conversationsSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the conversations client", func() {
		conversationsClient := conversationsSession.Conversations

		Describe("When the conversation resource is successfully created", func() {
			createInput := &conversations.CreateConversationInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := conversationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create conversation resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.State).To(Equal("active"))
				Expect(resp.Timers).To(Equal(conversations.CreateConversationResponseTimers{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create conversation resource api returns a 500 response", func() {
			createInput := &conversations.CreateConversationInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create conversation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of conversations are successfully retrieved", func() {
			pageOptions := &conversations.ConversationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the conversations page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Conversations?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Conversations?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("conversations"))

				conversationsResp := resp.Conversations
				Expect(conversationsResp).ToNot(BeNil())
				Expect(len(conversationsResp)).To(Equal(1))

				Expect(conversationsResp[0].Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conversationsResp[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conversationsResp[0].ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(conversationsResp[0].MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(conversationsResp[0].FriendlyName).To(Equal(utils.String("Test")))
				Expect(conversationsResp[0].Attributes).To(Equal("{}"))
				Expect(conversationsResp[0].State).To(Equal("active"))
				Expect(conversationsResp[0].Timers).To(Equal(conversations.PageConversationResponseTimers{}))
				Expect(conversationsResp[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(conversationsResp[0].DateUpdated).To(BeNil())
				Expect(conversationsResp[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of conversations api returns a 500 response", func() {
			pageOptions := &conversations.ConversationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the conversations page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated conversations are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := conversationsClient.NewConversationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated conversations current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated conversations results should be returned", func() {
				Expect(len(paginator.Conversations)).To(Equal(3))
			})
		})

		Describe("When the conversations api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := conversationsClient.NewConversationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated conversations current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a conversation sid", func() {
		conversationClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the conversation resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get conversation resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.State).To(Equal("active"))
				Expect(resp.Timers).To(Equal(conversation.FetchConversationResponseTimers{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Conversation("CH71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get conversation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateConversationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &conversation.UpdateConversationInput{
				FriendlyName:   utils.String("Test 2"),
				TimersClosed:   utils.String("PT10M"),
				TimersInactive: utils.String("PT1M"),
			}

			resp, err := conversationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update conversation response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test 2")))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.State).To(Equal("active"))

				dateClosed, _ := time.Parse(time.RFC3339, "2020-06-20T21:00:24Z")
				dateInactive, _ := time.Parse(time.RFC3339, "2020-06-20T20:51:24Z")
				Expect(resp.Timers).To(Equal(conversation.UpdateConversationResponseTimers{
					DateClosed:   utils.Time(dateClosed),
					DateInactive: utils.Time(dateInactive),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update conversation resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &conversation.UpdateConversationInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := conversationsSession.Conversation("CH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update conversation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := conversationClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the conversation resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Conversation("CH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a webhook client", func() {
		webhookClient := conversationsSession.Webhook()

		Describe("When the webhook resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/Webhooks",
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

			It("Then the get webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("webhook"))
				Expect(resp.Method).To(Equal("POST"))
				Expect(resp.PreWebhookURL).To(BeNil())
				Expect(resp.PostWebhookURL).To(BeNil())
				Expect(len(resp.Filters)).To(Equal(0))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/Webhooks"))
			})
		})

		Describe("When the webhook api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := webhookClient.Fetch()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the webhook is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			filters := []string{"onMessageAdded"}
			updateInput := &webhook.UpdateWebhookInput{
				PostWebhookURL: utils.String("http://localhost/pre"),
				Filters:        &filters,
			}

			resp, err := webhookClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("webhook"))
				Expect(resp.Method).To(Equal("POST"))
				Expect(resp.PreWebhookURL).To(Equal(utils.String("http://localhost/pre")))
				Expect(resp.PostWebhookURL).To(BeNil())
				Expect(len(resp.Filters)).To(Equal(1))
				Expect(resp.Filters[0]).To(Equal("onMessageAdded"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/Webhooks"))
			})
		})

		Describe("When the update webhook response returns a 500", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			filters := []string{"onMessageAdded"}
			updateInput := &webhook.UpdateWebhookInput{
				PostWebhookURL: utils.String("http://localhost/pre"),
				Filters:        &filters,
			}

			resp, err := webhookClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the update webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the messages client", func() {
		messagesClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Messages

		Describe("When the message is successfully created", func() {
			createInput := &messages.CreateMessageInput{
				Body: utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create message api returns a 500 response", func() {
			createInput := &messages.CreateMessageInput{
				Body: utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of messages are successfully retrieved", func() {
			pageOptions := &messages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the messages page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("messages"))

				messages := resp.Messages
				Expect(messages).ToNot(BeNil())
				Expect(len(messages)).To(Equal(1))

				Expect(messages[0].Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ParticipantSid).To(BeNil())
				Expect(messages[0].Body).To(Equal(utils.String("Hello World")))
				Expect(messages[0].Author).To(Equal("system"))
				Expect(messages[0].Index).To(Equal(0))
				Expect(messages[0].Attributes).To(Equal("{}"))
				Expect(messages[0].Media).To(BeNil())
				Expect(messages[0].Delivery).To(BeNil())
				Expect(messages[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(messages[0].DateUpdated).To(BeNil())
				Expect(messages[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of messages api returns a 500 response", func() {
			pageOptions := &messages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the messages page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated messages are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated messages current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated messages results should be returned", func() {
				Expect(len(paginator.Messages)).To(Equal(3))
			})
		})

		Describe("When the messages api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated messages current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a message sid", func() {
		messageClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the message is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messageClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the message with delivery is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/messageWithDeliveryResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messageClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(Equal(&message.FetchMessageResponseDelivery{
					Delivered:   "all",
					Failed:      "none",
					Read:        "none",
					Sent:        "all",
					Total:       1,
					Undelivered: "none",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &message.UpdateMessageInput{
				Body: utils.String("Hello World Updated"),
			}

			resp, err := messageClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World Updated")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update message api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &message.UpdateMessageInput{
				Body: utils.String("Hello World Updated"),
			}

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := messageClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the participants client", func() {
		participantsClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participants

		Describe("When the participant is successfully created", func() {
			createInput := &participants.CreateParticipantInput{
				Identity: utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
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
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingBinding).To(BeNil())
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant is successfully created with messaging binding", func() {
			createInput := &participants.CreateParticipantInput{
				MessagingBindingAddress:      utils.String("+123456789"),
				MessagingBindingProxyAddress: utils.String("+987654321"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantResponseWithMessagingBinding.json")
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
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(BeNil())

				messagingBinding := participants.CreateParticipantResponseMessageBinding{
					Type:         "sms",
					Address:      "+123456789",
					ProxyAddress: "+987654321",
				}
				Expect(resp.MessagingBinding).To(Equal(&messagingBinding))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create participant api returns a 500 response", func() {
			createInput := &participants.CreateParticipantInput{
				Identity: utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of participants are successfully retrieved", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the participants page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("participants"))

				participants := resp.Participants
				Expect(participants).ToNot(BeNil())
				Expect(len(participants)).To(Equal(1))

				Expect(participants[0].Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(participants[0].Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(participants[0].MessagingBinding).To(BeNil())
				Expect(participants[0].Attributes).To(Equal("{}"))
				Expect(participants[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(participants[0].DateUpdated).To(BeNil())
				Expect(participants[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of participants api returns a 500 response", func() {
			pageOptions := &participants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the participants page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated participants are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated participants current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated participants results should be returned", func() {
				Expect(len(paginator.Participants)).To(Equal(3))
			})
		})

		Describe("When the participants api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated participants current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a participant sid", func() {
		participantClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the participant is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingBinding).To(BeNil())
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant with message binding is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/participantWithMessageBindingResponse.json")
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
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(BeNil())
				Expect(resp.MessagingBinding).To(Equal(&participant.FetchParticipantResponseMessageBinding{
					Address:          "+10123456789",
					ProjectedAddress: nil,
					ProxyAddress:     "+19876543210",
					Type:             "sms",
				}))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MB71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateParticipantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Attributes: utils.String("{\"Test\": \"Test\"}"),
			}

			resp, err := participantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingBinding).To(BeNil())
				Expect(resp.Attributes).To(Equal("{\"Test\": \"Test\"}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant is successfully updated with messaging binding", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateParticipantResponseWithMessagingBinding.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Attributes: utils.String("{\"Test\": \"Test\"}"),
			}

			resp, err := participantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(BeNil())
				Expect(resp.MessagingBinding).To(Equal(&participant.UpdateParticipantResponseMessageBinding{
					Type:         "sms",
					Address:      "+123456789",
					ProxyAddress: "+987654321",
				}))
				Expect(resp.Attributes).To(Equal("{\"Test\": \"Test\"}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update participant api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &participant.UpdateParticipantInput{
				Attributes: utils.String("{\"Test\": \"Test\"}"),
			}

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MB71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := participantClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MB71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the conversation webhooks client", func() {
		conversationWebhooksClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhooks

		Describe("When the conversation webhook is successfully created", func() {
			createInput := &conversationWebhooks.CreateWebhookInput{
				Target:               "studio",
				ConfigurationFlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create conversation webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(conversationWebhooks.CreateWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook does not contain a target", func() {
			createInput := &conversationWebhooks.CreateWebhookInput{}

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook api returns a 500 response", func() {
			createInput := &conversationWebhooks.CreateWebhookInput{
				Target:               "studio",
				ConfigurationFlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of conversation webhooks are successfully retrieved", func() {
			pageOptions := &conversationWebhooks.WebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationWebhooksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationWebhooksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the conversationWebhooks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("webhooks"))

				webhooks := resp.Webhooks
				Expect(webhooks).ToNot(BeNil())
				Expect(len(webhooks)).To(Equal(1))

				Expect(webhooks[0].Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].Target).To(Equal("studio"))
				Expect(webhooks[0].Configuration).To(Equal(conversationWebhooks.PageWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(webhooks[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(webhooks[0].DateUpdated).To(BeNil())
				Expect(webhooks[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of conversation webhooks api returns a 500 response", func() {
			pageOptions := &conversationWebhooks.WebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationWebhooksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the conversation webhooks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated conversation webhooks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationWebhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationWebhooksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := conversationWebhooksClient.NewWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated conversation webhooks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated conversation webhooks results should be returned", func() {
				Expect(len(paginator.Webhooks)).To(Equal(3))
			})
		})

		Describe("When the conversation webhooks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationWebhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := conversationWebhooksClient.NewWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated conversation webhooks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a conversation webhook sid", func() {
		conversationWebhookClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the conversation webhook is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationWebhookClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get conversation webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(conversationWebhook.FetchWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateConversationWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &conversationWebhook.UpdateWebhookInput{}

			resp, err := conversationWebhookClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update conversation webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(conversationWebhook.UpdateWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &conversationWebhook.UpdateWebhookInput{}

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := conversationWebhookClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the conversation webhook api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the delivery receipts client", func() {
		deliveryReceiptsClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").DeliveryReceipts

		Describe("When the page of delivery receipts are successfully retrieved", func() {
			pageOptions := &delivery_receipts.DeliveryReceiptsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deliveryReceiptsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deliveryReceiptsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the delivery receipts page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("delivery_receipts"))

				deliveryReceipts := resp.DeliveryReceipts
				Expect(deliveryReceipts).ToNot(BeNil())
				Expect(len(deliveryReceipts)).To(Equal(1))

				Expect(deliveryReceipts[0]).ToNot(BeNil())
				Expect(deliveryReceipts[0].Sid).To(Equal("DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].AccountSid).To(BeNil())
				Expect(deliveryReceipts[0].MessageSid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ParticipantSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ChannelMessageSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].Status).To(Equal("sent"))
				Expect(deliveryReceipts[0].ErrorCode).To(Equal(utils.Int(0)))
				Expect(deliveryReceipts[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(deliveryReceipts[0].DateUpdated).To(BeNil())
				Expect(deliveryReceipts[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of messages api returns a 500 response", func() {
			pageOptions := &delivery_receipts.DeliveryReceiptsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := deliveryReceiptsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the delivery receipts page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated delivery receipts are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deliveryReceiptsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deliveryReceiptsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := deliveryReceiptsClient.NewDeliveryReceiptsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated delivery receipts current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated delivery receipts results should be returned", func() {
				Expect(len(paginator.DeliveryReceipts)).To(Equal(3))
			})
		})

		Describe("When the delivery receipts api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deliveryReceiptsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := deliveryReceiptsClient.NewDeliveryReceiptsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated delivery receipts current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a delivery receipt sid", func() {
		deliveryReceiptClient := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").DeliveryReceipt("DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the delivery receipt is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/deliveryReceiptResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deliveryReceiptClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get delivery receipt response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(BeNil())
				Expect(resp.MessageSid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelMessageSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("sent"))
				Expect(resp.ErrorCode).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the delivery receipt api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DY71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").DeliveryReceipt("DY71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get delivery receipt response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the roles client", func() {
		rolesClient := conversationsSession.Roles

		Describe("When the role resource is successfully created", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "channel admin",
				Type:         "conversation",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Roles",
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

			It("Then the create role resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("conversation"))
				Expect(resp.FriendlyName).To(Equal("channel admin"))
				Expect(resp.Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role does not contain a friendly name", func() {
			createInput := &roles.CreateRoleInput{
				Type: "conversation",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
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
				FriendlyName: "channel admin",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role does not contain permissions", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "channel admin",
				Type:         "conversation",
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create role resource api returns a 500 response", func() {
			createInput := &roles.CreateRoleInput{
				FriendlyName: "channel admin",
				Type:         "conversation",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Roles",
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

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Roles?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Roles?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("roles"))

				roles := resp.Roles
				Expect(roles).ToNot(BeNil())
				Expect(len(roles)).To(Equal(1))

				Expect(roles[0].Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].Type).To(Equal("conversation"))
				Expect(roles[0].FriendlyName).To(Equal("channel admin"))
				Expect(roles[0].Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				}))
				Expect(roles[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(roles[0].DateUpdated).To(BeNil())
				Expect(roles[0].URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of roles api returns a 500 response", func() {
			pageOptions := &roles.RolesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles?Page=1&PageSize=50&PageToken=abc1234",
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
		roleClient := conversationsSession.Role("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the role resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get role resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("conversation"))
				Expect(resp.FriendlyName).To(Equal("channel admin"))
				Expect(resp.Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Role("RL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRoleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &role.UpdateRoleInput{
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
				},
			}

			resp, err := roleClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update role response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("conversation"))
				Expect(resp.FriendlyName).To(Equal("channel admin"))
				Expect(resp.Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role does not contain permissions", func() {
			updateInput := &role.UpdateRoleInput{}

			resp, err := roleClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update role resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &role.UpdateRoleInput{
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
				},
			}

			resp, err := conversationsSession.Role("RL71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := roleClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the role resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Role("RL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the users client", func() {
		usersClient := conversationsSession.Users

		Describe("When the user resource is successfully created", func() {
			createInput := &users.CreateUserInput{
				Identity: "TestUser",
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Users",
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

			It("Then the create user resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.IsOnline).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user does not contain identity", func() {
			createInput := &users.CreateUserInput{}

			resp, err := usersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create user resource api returns a 500 response", func() {
			createInput := &users.CreateUserInput{
				Identity: "TestUser",
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Users",
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

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Users?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Users?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("users"))

				users := resp.Users
				Expect(users).ToNot(BeNil())
				Expect(len(users)).To(Equal(1))

				Expect(users[0].Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].Identity).To(Equal("TestUser"))
				Expect(users[0].Attributes).To(Equal("{}"))
				Expect(users[0].FriendlyName).To(BeNil())
				Expect(users[0].IsOnline).To(BeNil())
				Expect(users[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(users[0].DateUpdated).To(BeNil())
				Expect(users[0].URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of users api returns a 500 response", func() {
			pageOptions := &users.UsersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users?Page=1&PageSize=50&PageToken=abc1234",
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
		userClient := conversationsSession.User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the user resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get user resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.IsOnline).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.User("US71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateUserResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &user.UpdateUserInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := userClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update user response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.IsOnline).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update role resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &user.UpdateUserInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := conversationsSession.User("US71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := userClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the user resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.User("US71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the services client", func() {
		servicesClient := conversationsSession.Services

		Describe("When the service resource is successfully created", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Flex Chat Service",
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services",
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

			It("Then the create service resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("Flex Chat Service"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
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

		Describe("When the create service resource api returns a 500 response", func() {
			createInput := &services.CreateServiceInput{
				FriendlyName: "Flex Chat Service",
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services",
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

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Services?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Services?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("services"))

				services := resp.Services
				Expect(services).ToNot(BeNil())
				Expect(len(services)).To(Equal(1))

				Expect(services[0].Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(services[0].AccountSid).To(BeNil())
				Expect(services[0].FriendlyName).To(Equal("Flex Chat Service"))
				Expect(services[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(services[0].DateUpdated).To(BeNil())
				Expect(services[0].URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of services api returns a 500 response", func() {
			pageOptions := &services.ServicesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/servicesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services?Page=1&PageSize=50&PageToken=abc1234",
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
		serviceClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the service resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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
				Expect(resp.Sid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("Flex Chat Service"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the service resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("IS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the service resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := serviceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the service resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/IS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("IS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a service configuration client", func() {
		configurationClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Configuration()

		Describe("When the configuration resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration",
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

			It("Then the get configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultConversationRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2"))
				Expect(resp.DefaultConversationCreatorRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX3"))
				Expect(resp.DefaultChatServiceRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration"))
			})
		})

		Describe("When the configuration api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration",
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

		Describe("When the configuration is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/configurationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &configuration.UpdateConfigurationInput{}

			resp, err := configurationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update configuration response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultConversationRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2"))
				Expect(resp.DefaultConversationCreatorRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX3"))
				Expect(resp.DefaultChatServiceRoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration"))
			})
		})

		Describe("When the update configuration response returns a 500", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			updateInput := &configuration.UpdateConfigurationInput{}

			resp, err := configurationClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the update configuration response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a service notification client", func() {
		notificationClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Configuration().Notification()

		Describe("When the notification resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration/Notifications",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notificationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := notificationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get notification response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration/Notifications"))
			})
		})

		Describe("When the notification api returns a 500", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration/Notifications",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := notificationClient.Fetch()
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the get notification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the notification is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration/Notifications",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notificationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &notification.UpdateNotificationInput{}

			resp, err := notificationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update notification response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.NewMessage).To(Equal(notification.UpdateNotificationResponseNewMessage{
					Enabled:           false,
					BadgeCountEnabled: nil,
					Sound:             nil,
					Template:          nil,
				}))
				Expect(resp.AddedToConversation).To(Equal(notification.UpdateNotificationResponseConversationAction{
					Enabled:  false,
					Sound:    nil,
					Template: nil,
				}))
				Expect(resp.RemovedFromConversation).To(Equal(notification.UpdateNotificationResponseConversationAction{
					Enabled:  false,
					Sound:    nil,
					Template: nil,
				}))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration/Notifications"))
			})
		})

		Describe("When the update notification response returns a 500", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Configuration/Notifications",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			updateInput := &notification.UpdateNotificationInput{}

			resp, err := notificationClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the update notification response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the credentials client", func() {
		credentialsClient := conversationsSession.Credentials

		Describe("When the credential resource is successfully created", func() {
			createInput := &credentials.CreateCredentialInput{
				Type:   "fcm",
				Secret: utils.String("test"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Credentials",
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

			It("Then the create credential resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("fcm"))
				Expect(resp.FriendlyName).To(Equal(utils.String("TestCreds")))
				Expect(resp.Sandbox).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the credential does not contain a type", func() {
			createInput := &credentials.CreateCredentialInput{}

			resp, err := credentialsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create credential resource api returns a 500 response", func() {
			createInput := &credentials.CreateCredentialInput{
				Type:   "fcm",
				Secret: utils.String("test"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Credentials",
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

			It("Then the create credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of credentials are successfully retrieved", func() {
			pageOptions := &credentials.CredentialsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Credentials?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Credentials?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("credentials"))

				credentials := resp.Credentials
				Expect(credentials).ToNot(BeNil())
				Expect(len(credentials)).To(Equal(1))

				Expect(credentials[0].Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(credentials[0].Type).To(Equal("fcm"))
				Expect(credentials[0].FriendlyName).To(Equal(utils.String("TestCreds")))
				Expect(credentials[0].Sandbox).To(BeNil())
				Expect(credentials[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(credentials[0].DateUpdated).To(BeNil())
				Expect(credentials[0].URL).To(Equal("https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of credentials api returns a 500 response", func() {
			pageOptions := &credentials.CredentialsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/credentialsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials?Page=1&PageSize=50&PageToken=abc1234",
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
		credentialClient := conversationsSession.Credential("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the credential resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get credential resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("fcm"))
				Expect(resp.FriendlyName).To(Equal(utils.String("TestCreds")))
				Expect(resp.Sandbox).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the credential resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Credentials/CR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Credential("CR71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateCredentialResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &credential.UpdateCredentialInput{
				FriendlyName: utils.String("New TestCreds"),
			}

			resp, err := credentialClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update credential response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("fcm"))
				Expect(resp.FriendlyName).To(Equal(utils.String("New TestCreds")))
				Expect(resp.Sandbox).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update credential resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Credentials/CR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &credential.UpdateCredentialInput{
				FriendlyName: utils.String("New TestCreds"),
			}

			resp, err := conversationsSession.Credential("CR71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update credential response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the credential resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := credentialClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the credential resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Credentials/CR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Credential("CR71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the bindings client", func() {
		bindingsClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Bindings

		Describe("When the page of bindings are successfully retrieved", func() {
			pageOptions := &bindings.BindingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("bindings"))

				bindings := resp.Bindings
				Expect(bindings).ToNot(BeNil())
				Expect(len(bindings)).To(Equal(1))

				Expect(bindings[0].Sid).To(Equal("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].Endpoint).To(Equal("Endpoint"))
				Expect(bindings[0].Identity).To(Equal("Test"))
				Expect(bindings[0].BindingType).To(Equal("fcm"))
				Expect(bindings[0].CredentialSid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(bindings[0].MessageTypes).To(Equal([]string{
					"new_message",
				}))
				Expect(bindings[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(bindings[0].DateUpdated).To(BeNil())
				Expect(bindings[0].URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of bindings api returns a 500 response", func() {
			pageOptions := &bindings.BindingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/bindingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings?Page=1&PageSize=50&PageToken=abc1234",
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
		bindingClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the binding resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get binding resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Endpoint).To(Equal("Endpoint"))
				Expect(resp.Identity).To(Equal("Test"))
				Expect(resp.BindingType).To(Equal("fcm"))
				Expect(resp.CredentialSid).To(Equal("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessageTypes).To(Equal([]string{
					"new_message",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the binding resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get binding response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the binding resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := bindingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the binding resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings/BS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Binding("BS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service users client", func() {
		usersClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Users

		Describe("When the user resource is successfully created", func() {
			createInput := &serviceUsers.CreateUserInput{
				Identity: "TestUser",
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
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

			It("Then the create user resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.IsOnline).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user does not contain identity", func() {
			createInput := &serviceUsers.CreateUserInput{}

			resp, err := usersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create user resource api returns a 500 response", func() {
			createInput := &serviceUsers.CreateUserInput{
				Identity: "TestUser",
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
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
			pageOptions := &serviceUsers.UsersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Users?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Users?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("users"))

				users := resp.Users
				Expect(users).ToNot(BeNil())
				Expect(len(users)).To(Equal(1))

				Expect(users[0].Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(users[0].Identity).To(Equal("TestUser"))
				Expect(users[0].Attributes).To(Equal("{}"))
				Expect(users[0].FriendlyName).To(BeNil())
				Expect(users[0].IsOnline).To(BeNil())
				Expect(users[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(users[0].DateUpdated).To(BeNil())
				Expect(users[0].URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of users api returns a 500 response", func() {
			pageOptions := &serviceUsers.UsersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/usersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users?Page=1&PageSize=50&PageToken=abc1234",
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

	Describe("Given I have a service user sid", func() {
		userClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the user resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get user resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.FriendlyName).To(BeNil())
				Expect(resp.IsOnline).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the user resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateUserResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &serviceUser.UpdateUserInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := userClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update user response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Identity).To(Equal("TestUser"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.IsOnline).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update role resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &serviceUser.UpdateUserInput{
				FriendlyName: utils.String("Test"),
			}

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update user response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the user resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := userClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the user resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/US71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").User("US71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service roles client", func() {
		rolesClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Roles

		Describe("When the role resource is successfully created", func() {
			createInput := &serviceRoles.CreateRoleInput{
				FriendlyName: "channel admin",
				Type:         "conversation",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
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

			It("Then the create role resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("conversation"))
				Expect(resp.FriendlyName).To(Equal("channel admin"))
				Expect(resp.Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role does not contain a friendly name", func() {
			createInput := &serviceRoles.CreateRoleInput{
				Type: "conversation",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
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
			createInput := &serviceRoles.CreateRoleInput{
				FriendlyName: "channel admin",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role does not contain permissions", func() {
			createInput := &serviceRoles.CreateRoleInput{
				FriendlyName: "channel admin",
				Type:         "conversation",
			}

			resp, err := rolesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create role resource api returns a 500 response", func() {
			createInput := &serviceRoles.CreateRoleInput{
				FriendlyName: "channel admin",
				Type:         "conversation",
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				},
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
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
			pageOptions := &serviceRoles.RolesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=0&PageSize=50",
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
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Roles?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Roles?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("roles"))

				roles := resp.Roles
				Expect(roles).ToNot(BeNil())
				Expect(len(roles)).To(Equal(1))

				Expect(roles[0].Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(roles[0].Type).To(Equal("conversation"))
				Expect(roles[0].FriendlyName).To(Equal("channel admin"))
				Expect(roles[0].Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				}))
				Expect(roles[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(roles[0].DateUpdated).To(BeNil())
				Expect(roles[0].URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of roles api returns a 500 response", func() {
			pageOptions := &serviceRoles.RolesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=0&PageSize=50",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=1&PageSize=50&PageToken=abc1234",
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
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/rolesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?Page=1&PageSize=50&PageToken=abc1234",
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

	Describe("Given I have a service role sid", func() {
		roleClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the role resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
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

			It("Then the get role resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("conversation"))
				Expect(resp.FriendlyName).To(Equal("channel admin"))
				Expect(resp.Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
					"editNotificationLevel",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRoleResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &serviceRole.UpdateRoleInput{
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
				},
			}

			resp, err := roleClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update role response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Type).To(Equal("conversation"))
				Expect(resp.FriendlyName).To(Equal("channel admin"))
				Expect(resp.Permissions).To(Equal([]string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the role does not contain permissions", func() {
			updateInput := &serviceRole.UpdateRoleInput{}

			resp, err := roleClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update role resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &serviceRole.UpdateRoleInput{
				Permissions: []string{
					"deleteConversation",
					"removeParticipant",
					"editConversationName",
					"editConversationAttributes",
					"addParticipant",
					"sendMessage",
					"sendMediaMessage",
					"leaveConversation",
					"editAnyMessage",
					"editAnyMessageAttributes",
					"editAnyParticipantAttributes",
					"deleteAnyMessage",
				},
			}

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update role response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the role resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := roleClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the role resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RL71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Role("RL71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service conversations client", func() {
		conversationsClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversations

		Describe("When the conversation resource is successfully created", func() {
			createInput := &serviceConversations.CreateConversationInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := conversationsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create conversation resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.State).To(Equal("active"))
				Expect(resp.Timers).To(Equal(serviceConversations.CreateConversationResponseTimers{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create conversation resource api returns a 500 response", func() {
			createInput := &serviceConversations.CreateConversationInput{
				FriendlyName: utils.String("Test"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create conversation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of conversations are successfully retrieved", func() {
			pageOptions := &serviceConversations.ConversationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the conversations page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Conversations?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Conversations?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("conversations"))

				conversationsResp := resp.Conversations
				Expect(conversationsResp).ToNot(BeNil())
				Expect(len(conversationsResp)).To(Equal(1))

				Expect(conversationsResp[0].Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conversationsResp[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conversationsResp[0].ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(conversationsResp[0].MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(conversationsResp[0].FriendlyName).To(Equal(utils.String("Test")))
				Expect(conversationsResp[0].Attributes).To(Equal("{}"))
				Expect(conversationsResp[0].State).To(Equal("active"))
				Expect(conversationsResp[0].Timers).To(Equal(serviceConversations.PageConversationResponseTimers{}))
				Expect(conversationsResp[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(conversationsResp[0].DateUpdated).To(BeNil())
				Expect(conversationsResp[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of conversations api returns a 500 response", func() {
			pageOptions := &serviceConversations.ConversationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the conversations page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated conversations are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := conversationsClient.NewConversationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated conversations current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated conversations results should be returned", func() {
				Expect(len(paginator.Conversations)).To(Equal(3))
			})
		})

		Describe("When the conversations api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := conversationsClient.NewConversationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated conversations current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a service conversation sid", func() {
		conversationClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the conversation resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/conversationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get conversation resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test")))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.State).To(Equal("active"))
				Expect(resp.Timers).To(Equal(serviceConversation.FetchConversationResponseTimers{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CH71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get conversation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateConversationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &serviceConversation.UpdateConversationInput{
				FriendlyName:   utils.String("Test 2"),
				TimersClosed:   utils.String("PT10M"),
				TimersInactive: utils.String("PT1M"),
			}

			resp, err := conversationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update conversation response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.FriendlyName).To(Equal(utils.String("Test 2")))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.State).To(Equal("active"))

				dateClosed, _ := time.Parse(time.RFC3339, "2020-06-20T21:00:24Z")
				dateInactive, _ := time.Parse(time.RFC3339, "2020-06-20T20:51:24Z")
				Expect(resp.Timers).To(Equal(serviceConversation.UpdateConversationResponseTimers{
					DateClosed:   utils.Time(dateClosed),
					DateInactive: utils.Time(dateInactive),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update conversation resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &serviceConversation.UpdateConversationInput{
				FriendlyName: utils.String("Test 2"),
			}

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update conversation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := conversationClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the conversation resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service conversation webhooks client", func() {
		conversationWebhooksClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhooks

		Describe("When the conversation webhook is successfully created", func() {
			createInput := &serviceConversationWebhooks.CreateWebhookInput{
				Target:               "studio",
				ConfigurationFlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create conversation webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(serviceConversationWebhooks.CreateWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook does not contain a target", func() {
			createInput := &serviceConversationWebhooks.CreateWebhookInput{}

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook api returns a 500 response", func() {
			createInput := &serviceConversationWebhooks.CreateWebhookInput{
				Target:               "studio",
				ConfigurationFlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of conversation webhooks are successfully retrieved", func() {
			pageOptions := &serviceConversationWebhooks.WebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationWebhooksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationWebhooksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the conversation webhooks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("webhooks"))

				webhooks := resp.Webhooks
				Expect(webhooks).ToNot(BeNil())
				Expect(len(webhooks)).To(Equal(1))

				Expect(webhooks[0].Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(webhooks[0].Target).To(Equal("studio"))
				Expect(webhooks[0].Configuration).To(Equal(serviceConversationWebhooks.PageWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(webhooks[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(webhooks[0].DateUpdated).To(BeNil())
				Expect(webhooks[0].URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of conversation webhooks api returns a 500 response", func() {
			pageOptions := &serviceConversationWebhooks.WebhooksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := conversationWebhooksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the conversation webhooks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated conversation webhooks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationWebhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationWebhooksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := conversationWebhooksClient.NewWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated conversation webhooks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated conversation webhooks results should be returned", func() {
				Expect(len(paginator.Webhooks)).To(Equal(3))
			})
		})

		Describe("When the conversation webhooks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Service/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationWebhooksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := conversationWebhooksClient.NewWebhooksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated conversation webhooks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a service conversation webhook sid", func() {
		conversationWebhookClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the conversation webhook is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := conversationWebhookClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get conversation webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(serviceConversationWebhook.FetchWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceConversationWebhookResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &serviceConversationWebhook.UpdateWebhookInput{}

			resp, err := conversationWebhookClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update conversation webhook response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Target).To(Equal("studio"))
				Expect(resp.Configuration).To(Equal(serviceConversationWebhook.UpdateWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T23:19:51Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &serviceConversationWebhook.UpdateWebhookInput{}

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := conversationWebhookClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the conversation webhook api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WH71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Webhook("WH71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service conversation participants client", func() {
		participantsClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participants

		Describe("When the participant is successfully created", func() {
			createInput := &serviceParticipants.CreateParticipantInput{
				Identity: utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceParticipantResponse.json")
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
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessagingBinding).To(BeNil())
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create participant api returns a 500 response", func() {
			createInput := &serviceParticipants.CreateParticipantInput{
				Identity: utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of participants are successfully retrieved", func() {
			pageOptions := &serviceParticipants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceParticipantsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the participants page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("participants"))

				participants := resp.Participants
				Expect(participants).ToNot(BeNil())
				Expect(len(participants)).To(Equal(1))

				Expect(participants[0].Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(participants[0].Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(participants[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(participants[0].MessagingBinding).To(BeNil())
				Expect(participants[0].Attributes).To(Equal("{}"))
				Expect(participants[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(participants[0].DateUpdated).To(BeNil())
				Expect(participants[0].URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of participants api returns a 500 response", func() {
			pageOptions := &serviceParticipants.ParticipantsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := participantsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the participants page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated participants are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceParticipantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceParticipantsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated participants current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated participants results should be returned", func() {
				Expect(len(paginator.Participants)).To(Equal(3))
			})
		})

		Describe("When the participants api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceParticipantsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := participantsClient.NewParticipantsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated participants current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a service conversation participant sid", func() {
		participantClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the participant is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceParticipantResponse.json")
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
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessagingBinding).To(BeNil())
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MB71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceParticipantResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &serviceParticipant.UpdateParticipantInput{
				Attributes: utils.String("{\"Test\": \"Test\"}"),
			}

			resp, err := participantClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update participant response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.RoleSid).To(Equal(utils.String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Identity).To(Equal(utils.String("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.MessagingBinding).To(BeNil())
				Expect(resp.Attributes).To(Equal("{\"Test\": \"Test\"}"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update participant api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &serviceParticipant.UpdateParticipantInput{
				Attributes: utils.String("{\"Test\": \"Test\"}"),
			}

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MB71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update participant response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the participant is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := participantClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the participant api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Participants/MB71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Participant("MB71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service conversation messages client", func() {
		messagesClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Messages

		Describe("When the message is successfully created", func() {
			createInput := &serviceConversationMessages.CreateMessageInput{
				Body: utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create message api returns a 500 response", func() {
			createInput := &serviceConversationMessages.CreateMessageInput{
				Body: utils.String("Hello World"),
			}

			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of messages are successfully retrieved", func() {
			pageOptions := &serviceConversationMessages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessagesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the messages page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("messages"))

				messages := resp.Messages
				Expect(messages).ToNot(BeNil())
				Expect(len(messages)).To(Equal(1))

				Expect(messages[0].Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(messages[0].ParticipantSid).To(BeNil())
				Expect(messages[0].Body).To(Equal(utils.String("Hello World")))
				Expect(messages[0].Author).To(Equal("system"))
				Expect(messages[0].Index).To(Equal(0))
				Expect(messages[0].Attributes).To(Equal("{}"))
				Expect(messages[0].Media).To(BeNil())
				Expect(messages[0].Delivery).To(BeNil())
				Expect(messages[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(messages[0].DateUpdated).To(BeNil())
				Expect(messages[0].URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of messages api returns a 500 response", func() {
			pageOptions := &serviceConversationMessages.MessagesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := messagesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the messages page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated messages are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessagesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated messages current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated messages results should be returned", func() {
				Expect(len(paginator.Messages)).To(Equal(3))
			})
		})

		Describe("When the messages api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessagesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := messagesClient.NewMessagesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated messages current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a service conversation message sid", func() {
		messageClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the message is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := messageClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateServiceConversationMessageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &serviceConversationMessage.UpdateMessageInput{
				Body: utils.String("Hello World Updated"),
			}

			resp, err := messageClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update message response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(BeNil())
				Expect(resp.Body).To(Equal(utils.String("Hello World Updated")))
				Expect(resp.Author).To(Equal("system"))
				Expect(resp.Index).To(Equal(0))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Media).To(BeNil())
				Expect(resp.Delivery).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update message api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &serviceConversationMessage.UpdateMessageInput{
				Body: utils.String("Hello World Updated"),
			}

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update message response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the message is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := messageClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the message api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IM71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the service conversation message delivery receipts client", func() {
		deliveryReceiptsClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").DeliveryReceipts

		Describe("When the page of delivery receipts are successfully retrieved", func() {
			pageOptions := &serviceDeliveryReceipts.DeliveryReceiptsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageDeliveryReceiptsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deliveryReceiptsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the delivery receipts page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("delivery_receipts"))

				deliveryReceipts := resp.DeliveryReceipts
				Expect(deliveryReceipts).ToNot(BeNil())
				Expect(len(deliveryReceipts)).To(Equal(1))

				Expect(deliveryReceipts[0]).ToNot(BeNil())
				Expect(deliveryReceipts[0].Sid).To(Equal("DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].AccountSid).To(BeNil())
				Expect(deliveryReceipts[0].MessageSid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ParticipantSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ChannelMessageSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(deliveryReceipts[0].Status).To(Equal("sent"))
				Expect(deliveryReceipts[0].ErrorCode).To(Equal(utils.Int(0)))
				Expect(deliveryReceipts[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(deliveryReceipts[0].DateUpdated).To(BeNil())
				Expect(deliveryReceipts[0].URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of messages api returns a 500 response", func() {
			pageOptions := &serviceDeliveryReceipts.DeliveryReceiptsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := deliveryReceiptsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the delivery receipts page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated delivery receipts are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageDeliveryReceiptsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageDeliveryReceiptsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := deliveryReceiptsClient.NewDeliveryReceiptsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated delivery receipts current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated delivery receipts results should be returned", func() {
				Expect(len(paginator.DeliveryReceipts)).To(Equal(3))
			})
		})

		Describe("When the delivery receipts api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageDeliveryReceiptsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := deliveryReceiptsClient.NewDeliveryReceiptsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated delivery receipts current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a delivery receipt sid", func() {
		deliveryReceiptClient := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").DeliveryReceipt("DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the delivery receipt is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/serviceConversationMessageDeliveryReceiptResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := deliveryReceiptClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get delivery receipt response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(BeNil())
				Expect(resp.MessageSid).To(Equal("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ConversationSid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ParticipantSid).To(Equal("MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChannelMessageSid).To(Equal("SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ChatServiceSid).To(Equal("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("sent"))
				Expect(resp.ErrorCode).To(Equal(utils.Int(0)))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DYXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the delivery receipt api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://conversations.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Receipts/DY71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := conversationsSession.Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Message("IMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").DeliveryReceipt("DY71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get delivery receipt response should be nil", func() {
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Conversations/CH71 was not found"))

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
