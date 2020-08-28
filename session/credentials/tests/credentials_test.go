package tests

import (
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Credentials", func() {
	Describe("When invalid Account credential are supplied", func() {
		Context("No Sid supplied", func() {
			creds, err := credentials.New(credentials.Account{
				AuthToken: "Test Token",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})

		Context("No Auth Token supplied", func() {
			creds, err := credentials.New(credentials.Account{
				Sid: "ACxxxxxxxxxxx",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})

		Context("A invalid sid format", func() {
			creds, err := credentials.New(credentials.Account{
				Sid:       "Test Sid",
				AuthToken: "Test Token",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})
	})

	Describe("When valid Account credential are supplied", func() {
		creds, err := credentials.New(credentials.Account{
			Sid:       "ACxxxxxxxxxxx",
			AuthToken: "Test Token",
		})

		It("Then err should be nil", func() {
			Expect(err).To(BeNil())
		})

		It("Then credentials are nil", func() {
			Expect(creds).ToNot(BeNil())
			Expect(creds.Username).To(Equal("ACxxxxxxxxxxx"))
			Expect(creds.Password).To(Equal("Test Token"))
		})
	})

	Describe("When invalid Api Key credential are supplied", func() {
		Context("No Sid supplied", func() {
			creds, err := credentials.New(credentials.APIKey{
				Account: "ACxxxxxxxxxxx",
				Value:   "Test Api Key",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})

		Context("No Value supplied", func() {
			creds, err := credentials.New(credentials.APIKey{
				Account: "ACxxxxxxxxxxx",
				Sid:     "SKxxxxxxxxxxx",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})

		Context("No Account SID supplied", func() {
			creds, err := credentials.New(credentials.APIKey{
				Sid:   "SKxxxxxxxxxxx",
				Value: "Test Api Key",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})

		Context("A invalid sid format", func() {
			creds, err := credentials.New(credentials.APIKey{
				Account: "ACxxxxxxxxxxx",
				Sid:     "Test Sid",
				Value:   "Test API Key",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})

		Context("A invalid account sid format", func() {
			creds, err := credentials.New(credentials.APIKey{
				Account: "Test account",
				Sid:     "SKxxxxxxxxxxx",
				Value:   "Test API Key",
			})

			It("Then an error is returned", func() {
				Expect(err).To(HaveOccurred())
			})

			It("Then credentials are nil", func() {
				Expect(creds).To(BeNil())
			})
		})
	})

	Describe("When valid API Key credential are supplied", func() {
		creds, err := credentials.New(credentials.APIKey{
			Account: "ACxxxxxxxxxxx",
			Sid:     "SKxxxxxxxxxxx",
			Value:   "Test Api Key",
		})

		It("Then err should be nil", func() {
			Expect(err).To(BeNil())
		})

		It("Then credentials are nil", func() {
			Expect(creds).ToNot(BeNil())
			Expect(creds.AccountSid).To(Equal("ACxxxxxxxxxxx"))
			Expect(creds.Username).To(Equal("SKxxxxxxxxxxx"))
			Expect(creds.Password).To(Equal("Test Api Key"))
		})
	})
})
