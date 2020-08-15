package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session/interactions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session/participant/message_interactions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session/participants"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/sessions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Proxy Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	proxySession := twilio.NewWithCredentials(creds).Proxy.V1

	Describe("Given the proxy service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			servicesClient := proxySession.Services

			createResp, createErr := servicesClient.Create(&services.CreateServiceInput{
				UniqueName: uuid.New().String(),
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

			serviceClient := proxySession.Service(createResp.Sid)

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

	Describe("Given the proxy phone number clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := proxySession.Services.Create(&services.CreateServiceInput{
				UniqueName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := proxySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the phone number is created, fetched, updated and deleted", func() {
			phoneNumbersClient := proxySession.Service(serviceSid).PhoneNumbers

			createResp, createErr := phoneNumbersClient.Create(&phone_numbers.CreatePhoneNumberInput{
				Sid: utils.String(os.Getenv("TWILIO_PHONE_NUMBER_SID")),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := phoneNumbersClient.Page(&phone_numbers.PhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.PhoneNumbers)).Should(BeNumerically(">=", 1))

			paginator := phoneNumbersClient.NewPhoneNumbersPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.PhoneNumbers)).Should(BeNumerically(">=", 1))

			phoneNumberClient := proxySession.Service(serviceSid).PhoneNumber(createResp.Sid)

			fetchResp, fetchErr := phoneNumberClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := phoneNumberClient.Update(&phone_number.UpdatePhoneNumberInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := phoneNumberClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the proxy session clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := proxySession.Services.Create(&services.CreateServiceInput{
				UniqueName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := proxySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the session is created, fetched, updated and deleted", func() {
			sessionsClient := proxySession.Service(serviceSid).Sessions

			createResp, createErr := sessionsClient.Create(&sessions.CreateSessionInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := sessionsClient.Page(&sessions.SessionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Sessions)).Should(BeNumerically(">=", 1))

			paginator := sessionsClient.NewSessionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Sessions)).Should(BeNumerically(">=", 1))

			sessionClient := proxySession.Service(serviceSid).Session(createResp.Sid)

			fetchResp, fetchErr := sessionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := sessionClient.Update(&session.UpdateSessionInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := sessionClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the proxy participant clients", func() {

		var serviceSid string
		var sessionSid string

		BeforeEach(func() {
			resp, err := proxySession.Services.Create(&services.CreateServiceInput{
				UniqueName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			_, phoneNumberErr := proxySession.Service(serviceSid).PhoneNumbers.Create(&phone_numbers.CreatePhoneNumberInput{
				Sid: utils.String(os.Getenv("TWILIO_PHONE_NUMBER_SID")),
			})
			if phoneNumberErr != nil {
				Fail(fmt.Sprintf("Failed to add phone number resource. Error %s", phoneNumberErr.Error()))
			}

			sessionResp, sessionErr := proxySession.Service(serviceSid).Sessions.Create(&sessions.CreateSessionInput{})
			if sessionErr != nil {
				Fail(fmt.Sprintf("Failed to create session. Error %s", sessionErr.Error()))
			}
			sessionSid = sessionResp.Sid
		})

		AfterEach(func() {
			if err := proxySession.Service(serviceSid).Session(sessionSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete session. Error %s", err.Error()))
			}

			if err := proxySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the participant is created, fetched and deleted", func() {
			participantsClient := proxySession.Service(serviceSid).Session(sessionSid).Participants

			createResp, createErr := participantsClient.Create(&participants.CreateParticipantInput{
				Identifier: os.Getenv("DESTINATION_PHONE_NUMBER"),
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

			participantClient := proxySession.Service(serviceSid).Session(sessionSid).Participant(createResp.Sid)

			fetchResp, fetchErr := participantClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := participantClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the proxy message interaction and interaction clients", func() {

		var serviceSid string
		var sessionSid string
		var participantSid string

		BeforeEach(func() {
			resp, err := proxySession.Services.Create(&services.CreateServiceInput{
				UniqueName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid

			_, phoneNumberErr := proxySession.Service(serviceSid).PhoneNumbers.Create(&phone_numbers.CreatePhoneNumberInput{
				Sid: utils.String(os.Getenv("TWILIO_PHONE_NUMBER_SID")),
			})
			if phoneNumberErr != nil {
				Fail(fmt.Sprintf("Failed to add phone number resource. Error %s", phoneNumberErr.Error()))
			}

			sessionResp, sessionErr := proxySession.Service(serviceSid).Sessions.Create(&sessions.CreateSessionInput{})
			if sessionErr != nil {
				Fail(fmt.Sprintf("Failed to create session. Error %s", sessionErr.Error()))
			}
			sessionSid = sessionResp.Sid

			participantResp, participantErr := proxySession.Service(serviceSid).Session(sessionSid).Participants.Create(&participants.CreateParticipantInput{
				Identifier: os.Getenv("DESTINATION_PHONE_NUMBER"),
			})
			if participantErr != nil {
				Fail(fmt.Sprintf("Failed to create participant. Error %s", participantErr.Error()))
			}
			participantSid = participantResp.Sid
		})

		AfterEach(func() {
			if err := proxySession.Service(serviceSid).Session(sessionSid).Participant(participantSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete participant. Error %s", err.Error()))
			}

			if err := proxySession.Service(serviceSid).Session(sessionSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete session. Error %s", err.Error()))
			}

			if err := proxySession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the message interaction is created, fetched and deleted", func() {
			createResp, createErr := proxySession.Service(serviceSid).Session(sessionSid).Participant(participantSid).MessageInteractions.Create(&message_interactions.CreateMessageInteractionInput{
				Body: utils.String(`{"message": "Hello World"}`),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			interactionsClient := proxySession.Service(serviceSid).Session(sessionSid).Interactions

			pageResp, pageErr := interactionsClient.Page(&interactions.InteractionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Interactions)).Should(BeNumerically(">=", 1))

			paginator := interactionsClient.NewInteractionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Interactions)).Should(BeNumerically(">=", 1))

			interactionClient := proxySession.Service(serviceSid).Session(sessionSid).Interaction(createResp.Sid)

			fetchResp, fetchErr := interactionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := interactionClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	// TODO Add short code tests
})
