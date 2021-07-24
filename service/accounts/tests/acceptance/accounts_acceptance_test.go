package acceptance

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credential"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/public_key"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/public_keys"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Accounts Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	accountsSession := twilio.NewWithCredentials(creds).Accounts.V1

	Describe("Given the accounts public key clients", func() {
		It("Then the public key is created, fetched, updated and deleted", func() {
			createResp, createErr := accountsSession.Credentials.PublicKeys.Create(&public_keys.CreatePublicKeyInput{
				PublicKey: os.Getenv("TWILIO_PUBLIC_KEY"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := accountsSession.Credentials.PublicKeys.Page(&public_keys.PublicKeysPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Credentials)).Should(BeNumerically(">=", 1))

			paginator := accountsSession.Credentials.PublicKeys.NewPublicKeysPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Credentials)).Should(BeNumerically(">=", 1))

			publicKeyClient := accountsSession.Credentials.PublicKey(createResp.Sid)

			fetchResp, fetchErr := publicKeyClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := publicKeyClient.Update(&public_key.UpdatePublicKeyInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := publicKeyClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the accounts aws credentials clients", func() {
		It("Then the aws credential is created, fetched, updated and deleted", func() {
			createResp, createErr := accountsSession.Credentials.AWSCredentials.Create(&aws_credentials.CreateAWSCredentialInput{
				Credentials: fmt.Sprintf("%s:%s", os.Getenv("TWILIO_AWS_ACCESS_KEY_ID"), os.Getenv("TWILIO_AWS_SECRET_ACCESS_KEY")),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := accountsSession.Credentials.AWSCredentials.Page(&aws_credentials.AWSCredentialsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Credentials)).Should(BeNumerically(">=", 1))

			paginator := accountsSession.Credentials.AWSCredentials.NewAWSCredentialsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Credentials)).Should(BeNumerically(">=", 1))

			awsCredentialClient := accountsSession.Credentials.AWSCredential(createResp.Sid)

			fetchResp, fetchErr := awsCredentialClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := awsCredentialClient.Update(&aws_credential.UpdateAWSCredentialInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := awsCredentialClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})
})
