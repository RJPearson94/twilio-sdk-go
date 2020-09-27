package acceptance

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Lookups Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	lookupSession := twilio.NewWithCredentials(creds).Lookups.V1

	Describe("Given the lookups phone number client", func() {
		It("Then the phone number details are fetched", func() {
			phoneNumberClient := lookupSession.PhoneNumber(os.Getenv("DESTINATION_PHONE_NUMBER"))

			fetchResp, fetchErr := phoneNumberClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})
})
