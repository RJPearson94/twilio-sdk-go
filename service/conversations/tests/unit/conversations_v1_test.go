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

	"github.com/RJPearson94/twilio-sdk-go/service/conversations"
	conversationResource "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/message"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/message/delivery_receipts"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/participants"
	conversationWebhook "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhooks"
	conversationWebhooks "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhooks"
	conversationsResource "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/webhook"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Conversation V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	conversationsSession := conversations.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(conversationsSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the conversations client", func() {
		conversationsClient := conversationsSession.Conversations

		Describe("When the conversation resource is successfully created", func() {
			createInput := &conversationsResource.CreateConversationInput{
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
				Expect(resp.Timers).To(Equal(conversationsResource.CreateConversationResponseTimers{}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create conversation resource api returns a 500 response", func() {
			createInput := &conversationsResource.CreateConversationInput{
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
			pageOptions := &conversationsResource.ConversationsPageOptions{
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

				conversations := resp.Conversations
				Expect(conversations).ToNot(BeNil())
				Expect(len(conversations)).To(Equal(1))

				Expect(conversations[0].Sid).To(Equal("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conversations[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(conversations[0].ChatServiceSid).To(Equal(utils.String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(conversations[0].MessagingServiceSid).To(Equal(utils.String("MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(conversations[0].FriendlyName).To(Equal(utils.String("Test")))
				Expect(conversations[0].Attributes).To(Equal("{}"))
				Expect(conversations[0].State).To(Equal("active"))
				Expect(conversations[0].Timers).To(Equal(conversationsResource.PageConversationResponseTimers{}))
				Expect(conversations[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(conversations[0].DateUpdated).To(BeNil())
				Expect(conversations[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of conversations api returns a 500 response", func() {
			pageOptions := &conversationsResource.ConversationsPageOptions{
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
				Expect(resp.Timers).To(Equal(conversationResource.FetchConversationResponseTimers{}))
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

			updateInput := &conversationResource.UpdateConversationInput{
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
				Expect(resp.Timers).To(Equal(conversationResource.UpdateConversationResponseTimers{
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

			updateInput := &conversationResource.UpdateConversationInput{
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
					Address: "+10123456789",
					ProjectedAddress: nil,
					ProxyAddress: "+19876543210",
					Type: "sms",
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

				messagingBinding := participant.UpdateParticipantResponseMessageBinding{
					Type:         "sms",
					Address:      "+123456789",
					ProxyAddress: "+987654321",
				}
				Expect(resp.MessagingBinding).To(Equal(&messagingBinding))
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
			createInput := &conversationWebhooks.CreateConversationWebhookInput{
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
				Expect(resp.Configuration).To(Equal(conversationWebhooks.CreateConversationWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the conversation webhook does not contain a target", func() {
			createInput := &webhooks.CreateConversationWebhookInput{}

			resp, err := conversationWebhooksClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create conversation webhook response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the conversation webhook api returns a 500 response", func() {
			createInput := &webhooks.CreateConversationWebhookInput{
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
			pageOptions := &conversationWebhooks.ConversationWebhooksPageOptions{
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
				Expect(webhooks[0].Configuration).To(Equal(conversationWebhooks.PageConversationWebhookResponseConfiguration{
					FlowSid: utils.String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				}))
				Expect(webhooks[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T22:19:51Z"))
				Expect(webhooks[0].DateUpdated).To(BeNil())
				Expect(webhooks[0].URL).To(Equal("https://conversations.twilio.com/v1/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of conversation webhooks api returns a 500 response", func() {
			pageOptions := &conversationWebhooks.ConversationWebhooksPageOptions{
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
			paginator := conversationWebhooksClient.NewConversationWebhooksPaginator()

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
			paginator := conversationWebhooksClient.NewConversationWebhooksPaginator()

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
				Expect(resp.Configuration).To(Equal(conversationWebhook.FetchConversationWebhookResponseConfiguration{
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

			updateInput := &conversationWebhook.UpdateConversationWebhookInput{}

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
				Expect(resp.Configuration).To(Equal(conversationWebhook.UpdateConversationWebhookResponseConfiguration{
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

			updateInput := &conversationWebhook.UpdateConversationWebhookInput{}

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
