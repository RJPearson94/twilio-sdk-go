package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	apiCredentialLists "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_lists"
	apiIpAccessControlLists "github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/credential_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/ip_access_control_lists"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_url"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var accountSid = os.Getenv("TWILIO_ACCOUNT_SID")

var _ = Describe("Trunking Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       accountSid,
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	twilioClient := twilio.NewWithCredentials(creds)
	trunkingClient := twilioClient.Trunking.V1
	apiClient := twilioClient.API.V2010

	Describe("Given the Elastic SIP Trunk", func() {
		It("Then the trunk is created, fetched, updated and deleted", func() {
			trunksClient := trunkingClient.Trunks

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

			trunkClient := trunkingClient.Trunk(createResp.Sid)

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
			resp, err := trunkingClient.Trunks.Create(&trunks.CreateTrunkInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create trunk. Error %s", err.Error()))
			}
			trunkSid = resp.Sid
		})

		AfterEach(func() {
			if err := trunkingClient.Trunk(trunkSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete trunk. Error %s", err.Error()))
			}
		})

		It("Then the origination url is created, fetched, updated and deleted", func() {
			originationURLsClient := trunkingClient.Trunk(trunkSid).OriginationURLs

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

			originationURLClient := trunkingClient.Trunk(trunkSid).OriginationURL(createResp.Sid)

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

	Describe("Given the Elastic SIP Phone Number clients", func() {

		var trunkSid string

		BeforeEach(func() {
			resp, err := trunkingClient.Trunks.Create(&trunks.CreateTrunkInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create trunk. Error %s", err.Error()))
			}
			trunkSid = resp.Sid
		})

		AfterEach(func() {
			if err := trunkingClient.Trunk(trunkSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete trunk. Error %s", err.Error()))
			}
		})

		It("Then the phone number is created, fetched and deleted", func() {
			phoneNumbersClient := trunkingClient.Trunk(trunkSid).PhoneNumbers

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

			phoneNumberClient := trunkingClient.Trunk(trunkSid).PhoneNumber(createResp.Sid)

			fetchResp, fetchErr := phoneNumberClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := phoneNumberClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Elastic SIP Recording client", func() {

		var trunkSid string

		BeforeEach(func() {
			resp, err := trunkingClient.Trunks.Create(&trunks.CreateTrunkInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create trunk. Error %s", err.Error()))
			}
			trunkSid = resp.Sid
		})

		AfterEach(func() {
			if err := trunkingClient.Trunk(trunkSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete trunk. Error %s", err.Error()))
			}
		})

		It("Then the recording is fetched and updated", func() {
			recordingClient := trunkingClient.Trunk(trunkSid).Recording()

			fetchResp, fetchErr := recordingClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := recordingClient.Update(&recording.UpdateRecordingInput{
				Trim: utils.String("trim-silence"),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the Elastic SIP Credential List clients", func() {

		var trunkSid string
		var credentialListSid string

		BeforeEach(func() {
			resp, err := trunkingClient.Trunks.Create(&trunks.CreateTrunkInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create trunk. Error %s", err.Error()))
			}
			trunkSid = resp.Sid

			credentialListResp, credentialListErr := apiClient.Account(accountSid).Sip.CredentialLists.Create(&apiCredentialLists.CreateCredentialListInput{
				FriendlyName: uuid.New().String(),
			})
			if credentialListErr != nil {
				Fail(fmt.Sprintf("Failed to create credential list. Error %s", credentialListErr.Error()))
			}
			credentialListSid = credentialListResp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.CredentialList(credentialListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete credential list. Error %s", err.Error()))
			}

			if err := trunkingClient.Trunk(trunkSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete trunk. Error %s", err.Error()))
			}
		})

		It("Then the credential list is created, fetched and deleted", func() {
			credentialListsClient := trunkingClient.Trunk(trunkSid).CredentialLists

			createResp, createErr := credentialListsClient.Create(&credential_lists.CreateCredentialListInput{
				CredentialListSid: credentialListSid,
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

			credentialListClient := trunkingClient.Trunk(trunkSid).CredentialList(createResp.Sid)

			fetchResp, fetchErr := credentialListClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := credentialListClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Elastic SIP IP Access Control List clients", func() {

		var trunkSid string
		var ipAccessControlListSid string

		BeforeEach(func() {
			resp, err := trunkingClient.Trunks.Create(&trunks.CreateTrunkInput{})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create trunk. Error %s", err.Error()))
			}
			trunkSid = resp.Sid

			ipAccessControlListResp, ipAccessControlListErr := apiClient.Account(accountSid).Sip.IpAccessControlLists.Create(&apiIpAccessControlLists.CreateIpAccessControlListInput{
				FriendlyName: uuid.New().String(),
			})
			if ipAccessControlListErr != nil {
				Fail(fmt.Sprintf("Failed to create IP access control list. Error %s", ipAccessControlListErr.Error()))
			}
			ipAccessControlListSid = ipAccessControlListResp.Sid
		})

		AfterEach(func() {
			if err := apiClient.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete IP access control list. Error %s", err.Error()))
			}

			if err := trunkingClient.Trunk(trunkSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete trunk. Error %s", err.Error()))
			}
		})

		It("Then the IP access control list is created, fetched and deleted", func() {
			ipAccessControlListsClient := trunkingClient.Trunk(trunkSid).IpAccessControlLists

			createResp, createErr := ipAccessControlListsClient.Create(&ip_access_control_lists.CreateIpAccessControlListInput{
				IpAccessControlListSid: ipAccessControlListSid,
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

			ipAccessControlListClient := trunkingClient.Trunk(trunkSid).IpAccessControlList(createResp.Sid)

			fetchResp, fetchErr := ipAccessControlListClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			deleteErr := ipAccessControlListClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

})
