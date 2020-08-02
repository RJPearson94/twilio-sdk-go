package acceptance

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session"
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
		log.Panicf("%s", err)
	}

	proxySession := twilio.NewWithCredentials(creds).Proxy.V1

	Describe("Given the proxy service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			createResp, createErr := proxySession.Services.Create(&services.CreateServiceInput{
				UniqueName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

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
			createResp, createErr := proxySession.Service(serviceSid).PhoneNumbers.Create(&phone_numbers.CreatePhoneNumberInput{
				Sid: utils.String(os.Getenv("TWILIO_PHONE_NUMBER_SID")),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

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
			createResp, createErr := proxySession.Service(serviceSid).Sessions.Create(&sessions.CreateSessionInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

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
			createResp, createErr := proxySession.Service(serviceSid).Session(sessionSid).Participants.Create(&participants.CreateParticipantInput{
				Identifier: os.Getenv("DESTINATION_PHONE_NUMBER"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

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
