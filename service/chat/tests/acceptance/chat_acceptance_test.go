package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	v2Credential "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credential"
	v2Credentials "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
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
	v2UserChannel "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channel"
	v2UserChannels "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/users"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Chat Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	chatSession := twilio.NewWithCredentials(creds).Chat.V2

	Describe("Given the chat service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			servicesClient := chatSession.Services

			createResp, createErr := servicesClient.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := servicesClient.Page(&services.ServicesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Services)).Should(BeNumerically(">=", 1))

			paginator := servicesClient.NewServicesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Services)).Should(BeNumerically(">=", 1))

			serviceClient := chatSession.Service(createResp.Sid)

			fetchResp, fetchErr := serviceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := serviceClient.Update(&service.UpdateServiceInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := serviceClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat credential clients", func() {
		It("Then the credential is created, fetched, updated and deleted", func() {
			credentialsClient := chatSession.Credentials

			createResp, createErr := credentialsClient.Create(&v2Credentials.CreateCredentialInput{
				Type:   "fcm",
				Secret: utils.String("secret"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := credentialsClient.Page(&v2Credentials.CredentialsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Credentials)).Should(BeNumerically(">=", 1))

			paginator := credentialsClient.NewCredentialsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Credentials)).Should(BeNumerically(">=", 1))

			credentialClient := chatSession.Credential(createResp.Sid)

			fetchResp, fetchErr := credentialClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := credentialClient.Update(&v2Credential.UpdateCredentialInput{
				Secret: utils.String("new secret"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := credentialClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat channel clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the channel is created, fetched, updated and deleted", func() {
			channelsClient := chatSession.Service(serviceSid).Channels

			createResp, createErr := channelsClient.Create(&channels.CreateChannelInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := channelsClient.Page(&channels.ChannelsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Channels)).Should(BeNumerically(">=", 1))

			paginator := channelsClient.NewChannelsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Channels)).Should(BeNumerically(">=", 1))

			channelClient := chatSession.Service(serviceSid).Channel(createResp.Sid)

			fetchResp, fetchErr := channelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := channelClient.Update(&channel.UpdateChannelInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := channelClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat channel invite clients", func() {

		var identity string
		var serviceSid string
		var userSid string
		var channelSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			userResp, userErr := chatSession.Service(serviceSid).Users.Create(&users.CreateUserInput{
				Identity: uuid.New().String(),
			})
			if userErr != nil {
				Fail(fmt.Sprintf("Failed to create user. Error %s", userErr.Error()))
			}
			userSid = userResp.Sid
			identity = userResp.Identity

			channelResp, channelErr := chatSession.Service(serviceSid).Channels.Create(&channels.CreateChannelInput{})
			if channelErr != nil {
				Fail(fmt.Sprintf("Failed to create channel. Error %s", channelErr.Error()))
			}
			channelSid = channelResp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Channel(channelSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).User(userSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel user. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the channel invite is created, fetched and deleted", func() {
			invitesClient := chatSession.Service(serviceSid).Channel(channelSid).Invites

			createResp, createErr := invitesClient.Create(&invites.CreateChannelInviteInput{
				Identity: identity,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := invitesClient.Page(&invites.ChannelInvitesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Invites)).Should(BeNumerically(">=", 1))

			paginator := invitesClient.NewChannelInvitesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Invites)).Should(BeNumerically(">=", 1))

			inviteClient := chatSession.Service(serviceSid).Channel(channelSid).Invite(createResp.Sid)

			fetchResp, fetchErr := inviteClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := inviteClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat channel member clients", func() {

		var serviceSid string
		var channelSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			channelResp, channelErr := chatSession.Service(serviceSid).Channels.Create(&channels.CreateChannelInput{})
			if channelErr != nil {
				Fail(fmt.Sprintf("Failed to create channel. Error %s", channelErr.Error()))
			}
			channelSid = channelResp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Channel(channelSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the channel member is created, fetched, updated and deleted", func() {
			membersClient := chatSession.Service(serviceSid).Channel(channelSid).Members

			createResp, createErr := membersClient.Create(&members.CreateChannelMemberInput{
				Identity: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := membersClient.Page(&members.ChannelMembersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Members)).Should(BeNumerically(">=", 1))

			paginator := membersClient.NewChannelMembersPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Members)).Should(BeNumerically(">=", 1))

			memberClient := chatSession.Service(serviceSid).Channel(channelSid).Member(createResp.Sid)

			fetchResp, fetchErr := memberClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := memberClient.Update(&member.UpdateChannelMemberInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := memberClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat channel message clients", func() {

		var serviceSid string
		var channelSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			channelResp, channelErr := chatSession.Service(serviceSid).Channels.Create(&channels.CreateChannelInput{})
			if channelErr != nil {
				Fail(fmt.Sprintf("Failed to create channel. Error %s", channelErr.Error()))
			}
			channelSid = channelResp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Channel(channelSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the channel message is created, fetched, updated and deleted", func() {
			messagesClient := chatSession.Service(serviceSid).Channel(channelSid).Messages

			createResp, createErr := messagesClient.Create(&messages.CreateChannelMessageInput{
				Body: utils.String("Test"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := messagesClient.Page(&messages.ChannelMessagesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Messages)).Should(BeNumerically(">=", 1))

			paginator := messagesClient.NewChannelMessagesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Messages)).Should(BeNumerically(">=", 1))

			messageClient := chatSession.Service(serviceSid).Channel(channelSid).Message(createResp.Sid)

			fetchResp, fetchErr := messageClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := messageClient.Update(&message.UpdateChannelMessageInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := messageClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat channel webhook clients", func() {

		var serviceSid string
		var channelSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			channelResp, channelErr := chatSession.Service(serviceSid).Channels.Create(&channels.CreateChannelInput{})
			if channelErr != nil {
				Fail(fmt.Sprintf("Failed to create channel. Error %s", channelErr.Error()))
			}
			channelSid = channelResp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Channel(channelSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the channel webhook is created, fetched, updated and deleted", func() {
			webhooksClient := chatSession.Service(serviceSid).Channel(channelSid).Webhooks

			createResp, createErr := webhooksClient.Create(&webhooks.CreateChannelWebhookInput{
				Type:                 "webhook",
				ConfigurationURL:     utils.String("https://localhost.com/webhook"),
				ConfigurationFilters: &[]string{"onMessageSent"},
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := webhooksClient.Page(&webhooks.ChannelWebhooksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Webhooks)).Should(BeNumerically(">=", 1))

			paginator := webhooksClient.NewChannelWebhooksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Webhooks)).Should(BeNumerically(">=", 1))

			webhookClient := chatSession.Service(serviceSid).Channel(channelSid).Webhook(createResp.Sid)

			fetchResp, fetchErr := webhookClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := webhookClient.Update(&webhook.UpdateChannelWebhookInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := webhookClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat role clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the role is created, fetched, updated and deleted", func() {
			rolesClient := chatSession.Service(serviceSid).Roles

			createResp, createErr := rolesClient.Create(&roles.CreateRoleInput{
				FriendlyName: uuid.New().String(),
				Type:         "channel",
				Permission:   []string{"sendMessage"},
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

			roleClient := chatSession.Service(serviceSid).Role(createResp.Sid)

			fetchResp, fetchErr := roleClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := roleClient.Update(&role.UpdateRoleInput{
				Permission: []string{"sendMessage", "leaveChannel"},
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := roleClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat user clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the user is created, fetched, updated and deleted", func() {
			usersClient := chatSession.Service(serviceSid).Users

			createResp, createErr := usersClient.Create(&users.CreateUserInput{
				Identity: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := usersClient.Page(&users.UsersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Users)).Should(BeNumerically(">=", 1))

			paginator := usersClient.NewUsersPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Users)).Should(BeNumerically(">=", 1))

			userClient := chatSession.Service(serviceSid).User(createResp.Sid)

			fetchResp, fetchErr := userClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := userClient.Update(&user.UpdateUserInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := userClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the chat user channel clients", func() {

		var identity string
		var serviceSid string
		var userSid string
		var inviteSid string
		var channelSid string

		BeforeEach(func() {
			resp, err := chatSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			userResp, userErr := chatSession.Service(serviceSid).Users.Create(&users.CreateUserInput{
				Identity: uuid.New().String(),
			})
			if userErr != nil {
				Fail(fmt.Sprintf("Failed to create user. Error %s", userErr.Error()))
			}
			userSid = userResp.Sid
			identity = userResp.Identity

			channelResp, channelErr := chatSession.Service(serviceSid).Channels.Create(&channels.CreateChannelInput{})
			if channelErr != nil {
				Fail(fmt.Sprintf("Failed to create channel. Error %s", channelErr.Error()))
			}
			channelSid = channelResp.Sid

			inviteResp, inviteErr := chatSession.Service(serviceSid).Channel(channelSid).Invites.Create(&invites.CreateChannelInviteInput{
				Identity: identity,
			})
			if inviteErr != nil {
				Fail(fmt.Sprintf("Failed to create channel invite. Error %s", inviteErr.Error()))
			}
			inviteSid = inviteResp.Sid
		})

		AfterEach(func() {
			if err := chatSession.Service(serviceSid).Channel(channelSid).Invite(inviteSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel invite. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).Channel(channelSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).User(userSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete channel user. Error %s", err.Error()))
			}

			if err := chatSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the user channel is fetched and updated", func() {
			userChannelsClient := chatSession.Service(serviceSid).User(userSid).Channels

			pageResp, pageErr := userChannelsClient.Page(&v2UserChannels.UserChannelsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Channels)).Should(BeNumerically(">=", 1))

			paginator := userChannelsClient.NewUserChannelsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Channels)).Should(BeNumerically(">=", 1))

			userChannelClient := chatSession.Service(serviceSid).User(userSid).Channel(channelSid)

			fetchResp, fetchErr := userChannelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := userChannelClient.Update(&v2UserChannel.UpdateUserChannelInput{
				NotificationLevel: utils.String("default"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	// TODO Add binfing & user binding tests
})
