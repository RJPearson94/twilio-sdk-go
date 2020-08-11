package acceptance

import (
	"fmt"
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
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	messagingSession := twilio.NewWithCredentials(creds).Messaging.V1

	Describe("Given the messaging service clients", func() {
		It("Then the service is created, fetched, updated and deleted", func() {
			servicesClient := messagingSession.Services

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
			phoneNumbersClient := messagingSession.Service(serviceSid).PhoneNumbers

			createResp, createErr := phoneNumbersClient.Create(&phone_numbers.CreatePhoneNumberInput{
				PhoneNumberSid: os.Getenv("TWILIO_PHONE_NUMBER_SID"),
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
