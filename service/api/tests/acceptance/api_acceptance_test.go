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
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list"
	sipCredential "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credential"
	sipCredentials "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/credential_list_mappings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/calls/ip_access_control_list_mappings"
	registrationCredentialListMappings "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/registrations/credential_list_mappings"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domains"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_addresses"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_lists"
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

	apiClient := twilio.NewWithCredentials(creds).API.V2010

	Describe("Given the account clients", func() {
		It("Then the account is created, fetched and updated", func() {
			accountsClient := apiClient.Accounts

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

			accountClient := apiClient.Account(createResp.Sid)

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
			balanceClient := apiClient.Account(accountSid).Balance()

			fetchResp, fetchErr := balanceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the keys clients", func() {

		var accountSid string

		BeforeEach(func() {
			// Create sub account so if the key is leaked the parent account isn't compromised
			resp, err := apiClient.Accounts.Create(&accounts.CreateAccountInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create account. Error %s", err.Error()))
			}
			accountSid = resp.Sid
		})

		AfterEach(func() {
			if _, err := apiClient.Account(accountSid).Update(&account.UpdateAccountInput{
				Status: utils.String("closed"),
			}); err != nil {
				Fail(fmt.Sprintf("Failed to delete account. Error %s", err.Error()))
			}
		})

		It("Then the key is created, fetched, updated and deleted", func() {
			keysClient := apiClient.Account(accountSid).Keys

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

			keyClient := apiClient.Account(accountSid).Key(createResp.Sid)

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
			messagesClient := apiClient.Account(accountSid).Messages

			createResp, createErr := messagesClient.Create(&messages.CreateMessageInput{
				To:   os.Getenv("DESTINATION_PHONE_NUMBER"),
				From: utils.String(os.Getenv("TWILIO_PHONE_NUMBER")),
				Body: utils.String("Hello World"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			poll(30, 1000, apiClient, accountSid, createResp.Sid)

			pageResp, pageErr := messagesClient.Page(&messages.MessagesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Messages)).Should(BeNumerically(">=", 1))

			paginator := messagesClient.NewMessagesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Messages)).Should(BeNumerically(">=", 1))

			messageClient := apiClient.Account(accountSid).Message(createResp.Sid)

			fetchResp, fetchErr := messageClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := messageClient.Update(&message.UpdateMessageInput{
				Body: "",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			poll(30, 1000, apiClient, accountSid, createResp.Sid)

			deleteErr := messageClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the token client", func() {
		It("Then the token is created", func() {
			tokensClient := apiClient.Account(accountSid).Tokens

			createResp, createErr := tokensClient.Create(&tokens.CreateTokenInput{
				Ttl: utils.Int(1),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
		})
	})

	Describe("Given the queue clients", func() {
		It("Then the queue is created, fetched, updated and deleted", func() {
			queuesClient := apiClient.Account(accountSid).Queues

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

			queueClient := apiClient.Account(accountSid).Queue(createResp.Sid)

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
			callsClient := apiClient.Account(accountSid).Calls

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

			callClient := apiClient.Account(accountSid).Call(createResp.Sid)

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
			addressesClient := apiClient.Account(accountSid).Addresses

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

			addressClient := apiClient.Account(accountSid).Address(createResp.Sid)

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
			feedbackSummariesClient := apiClient.Account(accountSid).Calls.FeedbackSummaries

			createResp, createErr := feedbackSummariesClient.Create(&feedback_summaries.CreateFeedbackSummaryInput{
				StartDate: "2019-10-03",
				EndDate:   "2020-10-03",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			feedbackSummaryClient := apiClient.Account(accountSid).Calls.FeedbackSummary(createResp.Sid)

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

			resp, err := apiClient.Account(accountSid).Calls.Create(&calls.CreateCallInput{
				To:    os.Getenv("DESTINATION_PHONE_NUMBER"),
				From:  os.Getenv("TWILIO_PHONE_NUMBER"),
				TwiML: twiML,
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create call. Error %s", err.Error()))
			}
			callSid = resp.Sid

			_, endCallErr := apiClient.Account(accountSid).Call(callSid).Update(&call.UpdateCallInput{
				Status: utils.String("completed"),
			})
			if endCallErr != nil {
				Fail(fmt.Sprintf("Failed to update call. Error %s", endCallErr.Error()))
			}
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Call(callSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete call. Error %s", err.Error()))
			}
		})

		It("Then the feedback is created, fetched and updated", func() {
			feedbacksClient := apiClient.Account(accountSid).Call(callSid).Feedbacks

			createResp, createErr := feedbacksClient.Create(&feedbacks.CreateFeedbackInput{
				QualityScore: 5,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			feedbackClient := apiClient.Account(accountSid).Call(callSid).Feedback()

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
			applicationsClient := apiClient.Account(accountSid).Applications

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

			applicationClient := apiClient.Account(accountSid).Application(createResp.Sid)

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
			countriesClient := apiClient.Account(accountSid).AvailablePhoneNumbers

			pageResp, pageErr := countriesClient.Page()
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Countries)).Should(BeNumerically(">=", 1))

			countryClient := apiClient.Account(accountSid).AvailablePhoneNumber("GB")

			fetchResp, fetchErr := countryClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the available toll free phone numbers clients", func() {
		It("Then the available phone numbers are fetched", func() {
			availablePhoneNumbersClient := apiClient.Account(accountSid).AvailablePhoneNumber("GB").TollFree

			pageResp, pageErr := availablePhoneNumbersClient.Page(&toll_free.AvailablePhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.AvailablePhoneNumbers)).Should(BeNumerically(">=", 1))
		})
	})

	Describe("Given the available local phone numbers clients", func() {
		It("Then the available phone numbers are fetched", func() {
			availablePhoneNumbersClient := apiClient.Account(accountSid).AvailablePhoneNumber("GB").Local

			pageResp, pageErr := availablePhoneNumbersClient.Page(&local.AvailablePhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.AvailablePhoneNumbers)).Should(BeNumerically(">=", 1))
		})
	})

	Describe("Given the available mobile phone numbers clients", func() {
		It("Then the available phone numbers are fetched", func() {
			availablePhoneNumbersClient := apiClient.Account(accountSid).AvailablePhoneNumber("GB").Mobile

			pageResp, pageErr := availablePhoneNumbersClient.Page(&mobile.AvailablePhoneNumbersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.AvailablePhoneNumbers)).Should(BeNumerically(">=", 1))
		})
	})

	Describe("Given the SIP credential list clients", func() {
		It("Then the credential list is created, fetched, updated and deleted", func() {
			credentialListsClient := apiClient.Account(accountSid).Sip.CredentialLists

			createResp, createErr := credentialListsClient.Create(&credential_lists.CreateCredentialListInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := credentialListsClient.Page(&credential_lists.CredentialListsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.CredentialLists)).Should(BeNumerically(">=", 1))

			paginator := credentialListsClient.NewCredentialListsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.CredentialLists)).Should(BeNumerically(">=", 1))

			credentialListClient := apiClient.Account(accountSid).Sip.CredentialList(createResp.Sid)

			fetchResp, fetchErr := credentialListClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := credentialListClient.Update(&credential_list.UpdateCredentialListInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := credentialListClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the SIP credential clients", func() {

		var credentialListSid string

		BeforeEach(func() {
			resp, err := apiClient.Account(accountSid).Sip.CredentialLists.Create(&credential_lists.CreateCredentialListInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create credential list. Error %s", err.Error()))
			}
			credentialListSid = resp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.CredentialList(credentialListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete credential list. Error %s", err.Error()))
			}
		})

		It("Then the credential is created, fetched, updated and deleted", func() {
			credentialsClient := apiClient.Account(accountSid).Sip.CredentialList(credentialListSid).Credentials

			createResp, createErr := credentialsClient.Create(&sipCredentials.CreateCredentialInput{
				Username: uuid.New().String()[0:32],
				Password: "A" + uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := credentialsClient.Page(&sipCredentials.CredentialsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Credentials)).Should(BeNumerically(">=", 1))

			paginator := credentialsClient.NewCredentialsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Credentials)).Should(BeNumerically(">=", 1))

			credentialClient := apiClient.Account(accountSid).Sip.CredentialList(credentialListSid).Credential(createResp.Sid)

			fetchResp, fetchErr := credentialClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := credentialClient.Update(&sipCredential.UpdateCredentialInput{
				Password: "B" + uuid.New().String(),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := credentialClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the IP Access Control list clients", func() {
		It("Then the IP Access Control list is created, fetched, updated and deleted", func() {
			ipAccessControlListsClient := apiClient.Account(accountSid).Sip.IpAccessControlLists

			createResp, createErr := ipAccessControlListsClient.Create(&ip_access_control_lists.CreateIpAccessControlListInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := ipAccessControlListsClient.Page(&ip_access_control_lists.IpAccessControlListsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.IpAccessControlLists)).Should(BeNumerically(">=", 1))

			paginator := ipAccessControlListsClient.NewIpAccessControlListsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.IpAccessControlLists)).Should(BeNumerically(">=", 1))

			ipAccessControlListClient := apiClient.Account(accountSid).Sip.IpAccessControlList(createResp.Sid)

			fetchResp, fetchErr := ipAccessControlListClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := ipAccessControlListClient.Update(&ip_access_control_list.UpdateIpAccessControlListInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := ipAccessControlListClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the IP address clients", func() {
		var ipAccessControlListSid string

		BeforeEach(func() {
			resp, err := apiClient.Account(accountSid).Sip.IpAccessControlLists.Create(&ip_access_control_lists.CreateIpAccessControlListInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create IP access control list. Error %s", err.Error()))
			}
			ipAccessControlListSid = resp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete IP access control list. Error %s", err.Error()))
			}
		})

		It("Then the IP address is created, fetched, updated and deleted", func() {
			ipAddressesClient := apiClient.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).IpAddresses

			createResp, createErr := ipAddressesClient.Create(&ip_addresses.CreateIpAddressInput{
				FriendlyName: uuid.New().String(),
				IpAddress:    "127.0.0.1",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := ipAddressesClient.Page(&ip_addresses.IpAddressesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.IpAddresses)).Should(BeNumerically(">=", 1))

			paginator := ipAddressesClient.NewIpAddressesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.IpAddresses)).Should(BeNumerically(">=", 1))

			ipAddressClient := apiClient.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).IpAddress(createResp.Sid)

			fetchResp, fetchErr := ipAddressClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := ipAddressClient.Update(&ip_address.UpdateIpAddressInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := ipAddressClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the SIP domain clients", func() {
		It("Then the domain is created, fetched, updated and deleted", func() {
			domainsClient := apiClient.Account(accountSid).Sip.Domains

			createResp, createErr := domainsClient.Create(&domains.CreateDomainInput{
				DomainName: uuid.New().String() + ".sip.twilio.com",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := domainsClient.Page(&domains.DomainsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Domains)).Should(BeNumerically(">=", 1))

			paginator := domainsClient.NewDomainsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Domains)).Should(BeNumerically(">=", 1))

			domainClient := apiClient.Account(accountSid).Sip.Domain(createResp.Sid)

			fetchResp, fetchErr := domainClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := domainClient.Update(&domain.UpdateDomainInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := domainClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the IP access control list mapping clients", func() {
		var domainSid string
		var ipAccessControlListSid string

		BeforeEach(func() {
			resp, err := apiClient.Account(accountSid).Sip.Domains.Create(&domains.CreateDomainInput{
				DomainName: uuid.New().String() + ".sip.twilio.com",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create SIP domain. Error %s", err.Error()))
			}
			domainSid = resp.Sid

			ipAccessControlListsResp, ipAccessControlListsErr := apiClient.Account(accountSid).Sip.IpAccessControlLists.Create(&ip_access_control_lists.CreateIpAccessControlListInput{
				FriendlyName: uuid.New().String(),
			})
			if ipAccessControlListsErr != nil {
				Fail(fmt.Sprintf("Failed to create IP access control list. Error %s", ipAccessControlListsErr.Error()))
			}
			ipAccessControlListSid = ipAccessControlListsResp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.Domain(domainSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete SIP domain. Error %s", err.Error()))
			}
			if err := apiClient.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete IP access control list. Error %s", err.Error()))
			}
		})

		It("Then the IP access control list mapping is created, fetched and deleted", func() {
			ipAccessControlListMappingsClient := apiClient.Account(accountSid).Sip.Domain(domainSid).Auth.Calls.IpAccessControlListMappings

			createResp, createErr := ipAccessControlListMappingsClient.Create(&ip_access_control_list_mappings.CreateIpAccessControlListMappingInput{
				IpAccessControlListSid: ipAccessControlListSid,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := ipAccessControlListMappingsClient.Page(&ip_access_control_list_mappings.IpAccessControlListMappingsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.IpAccessControlListMappings)).Should(BeNumerically(">=", 1))

			paginator := ipAccessControlListMappingsClient.NewIpAccessControlListMappingsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.IpAccessControlListMappings)).Should(BeNumerically(">=", 1))

			ipAccessControlListMappingClient := apiClient.Account(accountSid).Sip.Domain(domainSid).Auth.Calls.IpAccessControlListMapping(createResp.Sid)

			fetchResp, fetchErr := ipAccessControlListMappingClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := ipAccessControlListMappingClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the credential list mapping clients", func() {
		var domainSid string
		var credentialListSid string

		BeforeEach(func() {
			resp, err := apiClient.Account(accountSid).Sip.Domains.Create(&domains.CreateDomainInput{
				DomainName: uuid.New().String() + ".sip.twilio.com",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create SIP domain. Error %s", err.Error()))
			}
			domainSid = resp.Sid

			credentialListResp, credentialListErr := apiClient.Account(accountSid).Sip.CredentialLists.Create(&credential_lists.CreateCredentialListInput{
				FriendlyName: uuid.New().String(),
			})
			if credentialListErr != nil {
				Fail(fmt.Sprintf("Failed to create credential list. Error %s", credentialListErr.Error()))
			}
			credentialListSid = credentialListResp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.Domain(domainSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete SIP domain. Error %s", err.Error()))
			}
			if err := apiClient.Account(accountSid).Sip.CredentialList(credentialListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete credential list. Error %s", err.Error()))
			}
		})

		It("Then the credential list mapping is created, fetched and deleted", func() {
			credentialListMappingsClient := apiClient.Account(accountSid).Sip.Domain(domainSid).Auth.Calls.CredentialListMappings

			createResp, createErr := credentialListMappingsClient.Create(&credential_list_mappings.CreateCredentialListMappingInput{
				CredentialListSid: credentialListSid,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := credentialListMappingsClient.Page(&credential_list_mappings.CredentialListMappingsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.CredentialListMappings)).Should(BeNumerically(">=", 1))

			paginator := credentialListMappingsClient.NewCredentialListMappingsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.CredentialListMappings)).Should(BeNumerically(">=", 1))

			credentialListMappingClient := apiClient.Account(accountSid).Sip.Domain(domainSid).Auth.Calls.CredentialListMapping(createResp.Sid)

			fetchResp, fetchErr := credentialListMappingClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := credentialListMappingClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the SIP registration credential list mapping clients", func() {
		var domainSid string
		var credentialListSid string

		BeforeEach(func() {
			resp, err := apiClient.Account(accountSid).Sip.Domains.Create(&domains.CreateDomainInput{
				DomainName: uuid.New().String() + ".sip.twilio.com",
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create SIP domain. Error %s", err.Error()))
			}
			domainSid = resp.Sid

			credentialListResp, credentialListErr := apiClient.Account(accountSid).Sip.CredentialLists.Create(&credential_lists.CreateCredentialListInput{
				FriendlyName: uuid.New().String(),
			})
			if credentialListErr != nil {
				Fail(fmt.Sprintf("Failed to create credential list. Error %s", credentialListErr.Error()))
			}
			credentialListSid = credentialListResp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.Domain(domainSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete SIP domain. Error %s", err.Error()))
			}
			if err := apiClient.Account(accountSid).Sip.CredentialList(credentialListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete credential list. Error %s", err.Error()))
			}
		})

		It("Then the credential list mapping is created, fetched and deleted", func() {
			credentialListMappingsClient := apiClient.Account(accountSid).Sip.Domain(domainSid).Auth.Registrations.CredentialListMappings

			createResp, createErr := credentialListMappingsClient.Create(&registrationCredentialListMappings.CreateCredentialListMappingInput{
				CredentialListSid: credentialListSid,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := credentialListMappingsClient.Page(&registrationCredentialListMappings.CredentialListMappingsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.CredentialListMappings)).Should(BeNumerically(">=", 1))

			paginator := credentialListMappingsClient.NewCredentialListMappingsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.CredentialListMappings)).Should(BeNumerically(">=", 1))

			credentialListMappingClient := apiClient.Account(accountSid).Sip.Domain(domainSid).Auth.Registrations.CredentialListMapping(createResp.Sid)

			fetchResp, fetchErr := credentialListMappingClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := credentialListMappingClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	// TODO SMS Media tests
	// TODO Incoming Phone Number tests
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
