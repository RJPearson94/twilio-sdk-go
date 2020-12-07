package acceptance

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Trunking Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	trunkingSession := twilio.NewWithCredentials(creds).Trunking.V1

	Describe("Given the Elastic SIP Trunk", func() {
		It("Then the trunk is created, fetched, updated and deleted", func() {
			trunksClient := trunkingSession.Trunks

			createResp, createErr := trunksClient.Create(&trunks.CreateTrunkInput{})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := trunksClient.Page(&trunks.TrunksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Trunks)).Should(BeNumerically(">=", 1))

			paginator := trunksClient.NewTrunksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Trunks)).Should(BeNumerically(">=", 1))

			trunkClient := trunkingSession.Trunk(createResp.Sid)

			fetchResp, fetchErr := trunkClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := trunkClient.Update(&trunk.UpdateTrunkInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := trunkClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})
})
