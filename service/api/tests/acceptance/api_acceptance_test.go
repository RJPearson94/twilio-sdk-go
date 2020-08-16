package acceptance

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/key"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/keys"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/message"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/messages"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/accounts"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
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

	// TODO SMS Media tests
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
