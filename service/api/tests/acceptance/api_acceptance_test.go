package acceptance

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/addresses"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/application"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/applications"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/local"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/mobile"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/toll_free"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call/feedback"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/call/feedbacks"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/calls/feedback_summaries"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queues"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/tokens"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/accounts"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var accountSid = os.Getenv("TWILIO_ACCOUNT_SID")

var _ = Describe("API Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       accountSid,
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	apiSession := twilio.NewWithCredentials(creds).API.V2010

	Describe("Given the account clients", func() {
		It("Then the account is created, fetched and updated", func() {
			accountsClient := apiSession.Accounts

			createResp, createErr := accountsClient.Create(&accounts.CreateAccountInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := accountsClient.Page(&accounts.AccountsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Accounts)).Should(BeNumerically(">=", 1))

			paginator := accountsClient.NewAccountsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Accounts)).Should(BeNumerically(">=", 1))

			accountClient := apiSession.Account(createResp.Sid)

			fetchResp, fetchErr := accountClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := accountClient.Update(&account.UpdateAccountInput{
				Status: utils.String("closed"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the balance client", func() {
		It("Then the balance is fetched", func() {
			balanceClient := apiSession.Account(accountSid).Balance()

			fetchResp, fetchErr := balanceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the keys clients", func() {

		var accountSid string

		BeforeEach(func() {
			// Create sub account so if the key is leaked the parent account isn't compromised
			resp, err := apiSession.Accounts.Create(&accounts.CreateAccountInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create account. Error %s", err.Error()))
			}
			accountSid = resp.Sid
		})

		AfterEach(func() {
			if _, err := apiSession.Account(accountSid).Update(&account.UpdateAccountInput{
				Status: utils.String("closed"),
			}); err != nil {
				Fail(fmt.Sprintf("Failed to delete account. Error %s", err.Error()))
			}
		})

		It("Then the key is created, fetched, updated and deleted", func() {
			keysClient := apiSession.Account(accountSid).Keys

			createResp, createErr := keysClient.Create(&keys.CreateKeyInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := keysClient.Page(&keys.KeysPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Keys)).Should(BeNumerically(">=", 1))

			paginator := keysClient.NewKeysPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Keys)).Should(BeNumerically(">=", 1))

			keyClient := apiSession.Account(accountSid).Key(createResp.Sid)

			fetchResp, fetchErr := keyClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := keyClient.Update(&key.UpdateKeyInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := keyClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the message clients", func() {
		It("Then the message is created, fetched, updated and deleted", func() {
			messagesClient := apiSession.Account(accountSid).Messages

			createResp, createErr := messagesClient.Create(&messages.CreateMessageInput{
				To:   os.Getenv("DESTINATION_PHONE_NUMBER"),
				From: utils.String(os.Getenv("TWILIO_PHONE_NUMBER")),
				Body: utils.String("Hello World"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			poll(30, 1000, apiSession, accountSid, createResp.Sid)

			pageResp, pageErr := messagesClient.Page(&messages.MessagesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Messages)).Should(BeNumerically(">=", 1))

			paginator := messagesClient.NewMessagesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Messages)).Should(BeNumerically(">=", 1))

			messageClient := apiSession.Account(accountSid).Message(createResp.Sid)

			fetchResp, fetchErr := messageClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := messageClient.Update(&message.UpdateMessageInput{
				Body: "",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			poll(30, 1000, apiSession, accountSid, createResp.Sid)

			deleteErr := messageClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the token client", func() {
		It("Then the token is created", func() {
			tokensClient := apiSession.Account(accountSid).Tokens

			createResp, createErr := tokensClient.Create(&tokens.CreateTokenInput{
				Ttl: utils.Int(1),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
		})
	})

	Describe("Given the queue clients", func() {
		It("Then the queue is created, fetched, updated and deleted", func() {
			queuesClient := apiSession.Account(accountSid).Queues

			createResp, createErr := queuesClient.Create(&queues.CreateQueueInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := queuesClient.Page(&queues.QueuesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Queues)).Should(BeNumerically(">=", 1))

			paginator := queuesClient.NewQueuesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Queues)).Should(BeNumerically(">=", 1))

			queueClient := apiSession.Account(accountSid).Queue(createResp.Sid)

			fetchResp, fetchErr := queueClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := queueClient.Update(&queue.UpdateQueueInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := queueClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the call clients", func() {
		It("Then the call is created, fetched, updated and deleted", func() {
			callsClient := apiSession.Account(accountSid).Calls

			twiMLResponse := voice.New()
			twiMLResponse.Say("Hello World")
			twiML, _ := twiMLResponse.ToTwiML()

			createResp, createErr := callsClient.Create(&calls.CreateCallInput{
				To:    os.Getenv("DESTINATION_PHONE_NUMBER"),
				From:  os.Getenv("TWILIO_PHONE_NUMBER"),
				TwiML: twiML,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := callsClient.Page(&calls.CallsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Calls)).Should(BeNumerically(">=", 1))

			paginator := callsClient.NewCallsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Calls)).Should(BeNumerically(">=", 1))

			callClient := apiSession.Account(accountSid).Call(createResp.Sid)

			fetchResp, fetchErr := callClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := callClient.Update(&call.UpdateCallInput{
				Status: utils.String("completed"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := callClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the address clients", func() {
		It("Then the address is created, fetched, updated and deleted", func() {
			addressesClient := apiSession.Account(accountSid).Addresses

			createResp, createErr := addressesClient.Create(&addresses.CreateAddressInput{
				CustomerName:       os.Getenv("TWILIO_CUSTOMER_NAME"),
				Street:             os.Getenv("TWILIO_ADDRESS_STREET"),
				City:               os.Getenv("TWILIO_ADDRESS_CITY"),
				Region:             os.Getenv("TWILIO_ADDRESS_REGION"),
				PostalCode:         os.Getenv("TWILIO_ADDRESS_POSTAL_CODE"),
				IsoCountry:         os.Getenv("TWILIO_ADDRESS_ISO_COUNTRY"),
				AutoCorrectAddress: utils.Bool(true),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := addressesClient.Page(&addresses.AddressesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Addresses)).Should(BeNumerically(">=", 1))

			paginator := addressesClient.NewAddressesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Addresses)).Should(BeNumerically(">=", 1))

			addressClient := apiSession.Account(accountSid).Address(createResp.Sid)

			fetchResp, fetchErr := addressClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := addressClient.Update(&address.UpdateAddressInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := addressClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the feedback summary clients", func() {
		It("Then the feedback summary is created, fetched and deleted", func() {
			feedbackSummariesClient := apiSession.Account(accountSid).Calls.FeedbackSummaries

			createResp, createErr := feedbackSummariesClient.Create(&feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
				EndDate:   "2020-10-03",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			feedbackSummaryClient := apiSession.Account(accountSid).Calls.FeedbackSummary(createResp.Sid)

			fetchResp, fetchErr := feedbackSummaryClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := feedbackSummaryClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the call feedback clients", func() {
		var callSid string

		BeforeEach(func() {
			twiMLResponse := voice.New()
			twiMLResponse.Say("Hello World")
			twiML, _ := twiMLResponse.ToTwiML()

			resp, err := apiSession.Account(accountSid).Calls.Create(&calls.CreateCallInput{
				To:    os.Getenv("DESTINATION_PHONE_NUMBER"),
				From:  os.Getenv("TWILIO_PHONE_NUMBER"),
				TwiML: twiML,
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create call. Error %s", err.Error()))
			}
			callSid = resp.Sid

			_, endCallErr := apiSession.Account(accountSid).Call(callSid).Update(&call.UpdateCallInput{
				Status: utils.String("completed"),
			})
			if endCallErr != nil {
				Fail(fmt.Sprintf("Failed to update call. Error %s", endCallErr.Error()))
			}
		})

		AfterEach(func() {
			if err := apiSession.Account(accountSid).Call(callSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete call. Error %s", err.Error()))
			}
		})

		It("Then the feedback is created, fetched and updated", func() {
			feedbacksClient := apiSession.Account(accountSid).Call(callSid).Feedbacks

			createResp, createErr := feedbacksClient.Create(&feedbacks.CreateFeedbackInput{
				QualityScore: 5,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			feedbackClient := apiSession.Account(accountSid).Call(callSid).Feedback()

			fetchResp, fetchErr := feedbackClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := feedbackClient.Update(&feedback.UpdateFeedbackInput{
				QualityScore: 4,
				Issues:       &[]string{"audio-latency"},
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the application clients", func() {
		It("Then the application is created, fetched, updated and deleted", func() {
			applicationsClient := apiSession.Account(accountSid).Applications

			createResp, createErr := applicationsClient.Create(&applications.CreateApplicationInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := applicationsClient.Page(&applications.ApplicationsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Applications)).Should(BeNumerically(">=", 1))

			paginator := applicationsClient.NewApplicationsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Applications)).Should(BeNumerically(">=", 1))

			applicationClient := apiSession.Account(accountSid).Application(createResp.Sid)

			fetchResp, fetchErr := applicationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := applicationClient.Update(&application.UpdateApplicationInput{
				FriendlyName: utils.String("Test"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := applicationClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the available phone number countries clients", func() {
		It("Then the countries are fetched", func() {
			countriesClient := apiSession.Account(accountSid).AvailablePhoneNumbers

			pageResp, pageErr := countriesClient.Page()
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Countries)).Should(BeNumerically(">=", 1))

			countryClient := apiSession.Account(accountSid).AvailablePhoneNumber("GB")

			fetchResp, fetchErr := countryClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the available toll free phone numbers clients", func() {
		It("Then the available phone numbers are fetched", func() {
			availablePhoneNumbersClient := apiSession.Account(accountSid).AvailablePhoneNumber("GB").TollFree

			pageResp, pageErr := availablePhoneNumbersClient.Page(&toll_free.AvailablePhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.AvailablePhoneNumbers)).Should(BeNumerically(">=", 1))
		})
	})

	Describe("Given the available local phone numbers clients", func() {
		It("Then the available phone numbers are fetched", func() {
			availablePhoneNumbersClient := apiSession.Account(accountSid).AvailablePhoneNumber("GB").Local

			pageResp, pageErr := availablePhoneNumbersClient.Page(&local.AvailablePhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.AvailablePhoneNumbers)).Should(BeNumerically(">=", 1))
		})
	})

	Describe("Given the available mobile phone numbers clients", func() {
		It("Then the available phone numbers are fetched", func() {
			availablePhoneNumbersClient := apiSession.Account(accountSid).AvailablePhoneNumber("GB").Mobile

			pageResp, pageErr := availablePhoneNumbersClient.Page(&mobile.AvailablePhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.AvailablePhoneNumbers)).Should(BeNumerically(">=", 1))
		})
	})

	// TODO SMS Media tests
	// TODO Queue Members tests
	// TODO Conference tests
	// TODO Conference Participants tests
	// TODO Recording & Call Recording & conference recording tests
})

func poll(maxAttempts int, delay int, client *v2010.V2010, accountSid string, messageSid string) error {
	for i := 0; i < maxAttempts; i++ {
		resp, err := client.Account(accountSid).Message(messageSid).Fetch()
		if err != nil {
			return fmt.Errorf("Failed to poll message: %s", err)
		}

		if resp.Status == "failed" || resp.Status == "undelivered" {
			return fmt.Errorf("Mesage failed")
		}
		if resp.Status == "delivered" {
			return nil
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return fmt.Errorf("Reached max polling attempts without a completed message delivery")
}
