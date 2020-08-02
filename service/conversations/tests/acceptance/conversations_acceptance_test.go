package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	conversationResource "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/message"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/participant"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/participants"
	conversationWebhook "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhook"
	conversationWebhooks "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversation/webhooks"
	conversationsResource "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/webhook"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Conversations Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	conversationsSession := twilio.NewWithCredentials(creds).Conversations.V1

	Describe("Given the conversations conversation clients", func() {
		It("Then the conversation is created, fetched, updated and deleted", func() {
			createResp, createErr := conversationsSession.Conversations.Create(&conversationsResource.CreateConversationInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			conversationClient := conversationsSession.Conversation(createResp.Sid)

			fetchResp, fetchErr := conversationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := conversationClient.Update(&conversationResource.UpdateConversationInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := conversationClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the conversations webhook client", func() {
		It("Then the webhook is fetched and updated", func() {
			webhookClient := conversationsSession.Webhook()

			fetchResp, fetchErr := webhookClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := webhookClient.Update(&webhook.UpdateWebhookInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the conversation message clients", func() {

		var conversationSid string

		BeforeEach(func() {
			resp, err := conversationsSession.Conversations.Create(&conversationsResource.CreateConversationInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create conversation. Error %s", err.Error()))
			}
			conversationSid = resp.Sid
		})

		AfterEach(func() {
			if err := conversationsSession.Conversation(conversationSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete conversation. Error %s", err.Error()))
			}
		})

		It("Then the message is created, fetched, updated and deleted", func() {
			createResp, createErr := conversationsSession.Conversation(conversationSid).Messages.Create(&messages.CreateMessageInput{
				Body: utils.String("Hello World"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			messageClient := conversationsSession.Conversation(conversationSid).Message(createResp.Sid)

			fetchResp, fetchErr := messageClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := messageClient.Update(&message.UpdateMessageInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := messageClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the conversation participant clients", func() {

		var conversationSid string

		BeforeEach(func() {
			resp, err := conversationsSession.Conversations.Create(&conversationsResource.CreateConversationInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create conversation. Error %s", err.Error()))
			}
			conversationSid = resp.Sid
		})

		AfterEach(func() {
			if err := conversationsSession.Conversation(conversationSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete conversation. Error %s", err.Error()))
			}
		})

		It("Then the participant is created, fetched, updated and deleted", func() {
			createResp, createErr := conversationsSession.Conversation(conversationSid).Participants.Create(&participants.CreateParticipantInput{
				Identity: utils.String(uuid.New().String()),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			participantClient := conversationsSession.Conversation(conversationSid).Participant(createResp.Sid)

			fetchResp, fetchErr := participantClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := participantClient.Update(&participant.UpdateParticipantInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := participantClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the conversation webhook clients", func() {

		var conversationSid string

		BeforeEach(func() {
			resp, err := conversationsSession.Conversations.Create(&conversationsResource.CreateConversationInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create conversation. Error %s", err.Error()))
			}
			conversationSid = resp.Sid
		})

		AfterEach(func() {
			if err := conversationsSession.Conversation(conversationSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete conversation. Error %s", err.Error()))
			}
		})

		It("Then the webhook is created, fetched, updated and deleted", func() {
			createResp, createErr := conversationsSession.Conversation(conversationSid).Webhooks.Create(&conversationWebhooks.CreateConversationWebhookInput{
				Target:               "webhook",
				ConfigurationURL:     utils.String("https://localhost.com/webhook"),
				ConfigurationFilters: &[]string{"onMessageAdded"},
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			webhookClient := conversationsSession.Conversation(conversationSid).Webhook(createResp.Sid)

			fetchResp, fetchErr := webhookClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := webhookClient.Update(&conversationWebhook.UpdateConversationWebhookInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := webhookClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

})
