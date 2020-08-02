package acceptance

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Messaging Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	messagingSession := twilio.NewWithCredentials(creds).Messaging.V1

	Describe("Given the messaging service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			createResp, createErr := messagingSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			serviceClient := messagingSession.Service(createResp.Sid)

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

	Describe("Given the messaging phone number clients", func() {

		var serviceSid string

		BeforeEach(func() {
			resp, err := messagingSession.Services.Create(&services.CreateServiceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create service. Error %s", err.Error()))
			}
			serviceSid = resp.Sid
		})

		AfterEach(func() {
			if err := messagingSession.Service(serviceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete service. Error %s", err.Error()))
			}
		})

		It("Then the phone number is created, fetched and deleted", func() {
			createResp, createErr := messagingSession.Service(serviceSid).PhoneNumbers.Create(&phone_numbers.CreatePhoneNumberInput{
				PhoneNumberSid: os.Getenv("TWILIO_PHONE_NUMBER_SID"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			phoneNumberClient := messagingSession.Service(serviceSid).PhoneNumber(createResp.Sid)

			fetchResp, fetchErr := phoneNumberClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := phoneNumberClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	// TODO Add short code and alpha numeric tests
})
