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
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	conversationsResource "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/role"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/roles"
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
			conversationsClient := conversationsSession.Conversations

			createResp, createErr := conversationsClient.Create(&conversationsResource.CreateConversationInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := conversationsClient.Page(&conversations.ConversationsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Conversations)).Should(BeNumerically(">=", 1))

			paginator := conversationsClient.NewConversationsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Conversations)).Should(BeNumerically(">=", 1))

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
			messagesClient := conversationsSession.Conversation(conversationSid).Messages

			createResp, createErr := messagesClient.Create(&messages.CreateMessageInput{
				Body: utils.String("Hello World"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := messagesClient.Page(&messages.MessagesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Messages)).Should(BeNumerically(">=", 1))

			paginator := messagesClient.NewMessagesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Messages)).Should(BeNumerically(">=", 1))

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
			participantsClient := conversationsSession.Conversation(conversationSid).Participants

			createResp, createErr := participantsClient.Create(&participants.CreateParticipantInput{
				Identity: utils.String(uuid.New().String()),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := participantsClient.Page(&participants.ParticipantsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Participants)).Should(BeNumerically(">=", 1))

			paginator := participantsClient.NewParticipantsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Participants)).Should(BeNumerically(">=", 1))

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
			webhooksClient := conversationsSession.Conversation(conversationSid).Webhooks

			createResp, createErr := webhooksClient.Create(&conversationWebhooks.CreateConversationWebhookInput{
				Target:               "webhook",
				ConfigurationURL:     utils.String("https://localhost.com/webhook"),
				ConfigurationFilters: &[]string{"onMessageAdded"},
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := webhooksClient.Page(&conversationWebhooks.ConversationWebhooksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Webhooks)).Should(BeNumerically(">=", 1))

			paginator := webhooksClient.NewConversationWebhooksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Webhooks)).Should(BeNumerically(">=", 1))

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

	Describe("Given the conversations role clients", func() {
		It("Then the role is created, fetched, updated and deleted", func() {
			rolesClient := conversationsSession.Roles

			createResp, createErr := rolesClient.Create(&roles.CreateRoleInput{
				FriendlyName: uuid.New().String(),
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
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := rolesClient.Page(&roles.RolesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Roles)).Should(BeNumerically(">=", 1))

			paginator := rolesClient.NewRolesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Roles)).Should(BeNumerically(">=", 1))

			roleClient := conversationsSession.Role(createResp.Sid)

			fetchResp, fetchErr := roleClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := roleClient.Update(&role.UpdateRoleInput{
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
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := roleClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	// TODO Add delivery receipt tests
})
