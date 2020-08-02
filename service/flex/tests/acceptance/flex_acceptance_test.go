package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Flex Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	flexSession := twilio.NewWithCredentials(creds).Flex.V1

	Describe("Given the flex configuration client", func() {
		It("Then the configuration is fetched and updated", func() {
			configurationClient := flexSession.Configuration()

			fetchResp, fetchErr := configurationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := configurationClient.Update(&configuration.UpdateConfigurationInput{
				AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the flex flow clients", func() {

		It("Then the flex flow is created, fetched, updated and deleted", func() {
			createResp, createErr := flexSession.FlexFlows.Create(&flex_flows.CreateFlexFlowInput{
				FriendlyName:    uuid.New().String(),
				ChatServiceSid:  os.Getenv("TWILIO_FLEX_CHANNEL_SERVICE_SID"),
				ChannelType:     "web",
				IntegrationType: utils.String("external"),
				IntegrationURL:  utils.String("https://test.com/external"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			flexFlowClient := flexSession.FlexFlow(createResp.Sid)

			fetchResp, fetchErr := flexFlowClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := flexFlowClient.Update(&flex_flow.UpdateFlexFlowInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := flexFlowClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the flex channel clients", func() {

		var flexFlowSid string

		BeforeEach(func() {
			flexFlowResp, flexFlowErr := flexSession.FlexFlows.Create(&flex_flows.CreateFlexFlowInput{
				FriendlyName:    uuid.New().String(),
				ChatServiceSid:  os.Getenv("TWILIO_FLEX_CHANNEL_SERVICE_SID"),
				ChannelType:     "web",
				IntegrationType: utils.String("external"),
				IntegrationURL:  utils.String("https://test.com/external"),
			})
			if flexFlowErr != nil {
				Fail(fmt.Sprintf("Failed to create flex flow. Error %s", flexFlowErr.Error()))
			}
			flexFlowSid = flexFlowResp.Sid
		})

		AfterEach(func() {
			if err := flexSession.FlexFlow(flexFlowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete flex flow. Error %s", err.Error()))
			}
		})

		It("Then the channel is created, fetched and deleted", func() {
			createResp, createErr := flexSession.Channels.Create(&channels.CreateChannelInput{
				FlexFlowSid:          flexFlowSid,
				Identity:             uuid.New().String(),
				ChatUserFriendlyName: uuid.New().String(),
				ChatFriendlyName:     uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			channelClient := flexSession.Channel(createResp.Sid)

			fetchResp, fetchErr := channelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := channelClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	// TODO add web channel support
})
