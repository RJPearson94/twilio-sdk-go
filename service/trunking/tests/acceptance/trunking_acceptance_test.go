package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_url"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
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

	Describe("Given the Elastic SIP Origination URL clients", func() {

		var trunkSid string

		BeforeEach(func() {
			resp, err := trunkingSession.Trunks.Create(&trunks.CreateTrunkInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create trunk. Error %s", err.Error()))
			}
			trunkSid = resp.Sid
		})

		AfterEach(func() {
			if err := trunkingSession.Trunk(trunkSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete trunk. Error %s", err.Error()))
			}
		})

		It("Then the origination url is created, fetched, updated and deleted", func() {
			originationURLsClient := trunkingSession.Trunk(trunkSid).OriginationURLs

			createResp, createErr := originationURLsClient.Create(&origination_urls.CreateOriginationURLInput{
				Weight:       0,
				Priority:     0,
				Enabled:      false,
				FriendlyName: uuid.New().String(),
				SipURL:       "sip:test@test.com",
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := originationURLsClient.Page(&origination_urls.OriginationURLsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.OriginationURLs)).Should(BeNumerically(">=", 1))

			paginator := originationURLsClient.NewOriginationURLsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.OriginationURLs)).Should(BeNumerically(">=", 1))

			originationURLClient := trunkingSession.Trunk(trunkSid).OriginationURL(createResp.Sid)

			fetchResp, fetchErr := originationURLClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := originationURLClient.Update(&origination_url.UpdateOriginationURLInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := originationURLClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})
})
