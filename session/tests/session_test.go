package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("Session", func() {
	Describe("When I initialise a session with credentials", func() {
		credentials := &credentials.Credentials{
			Username: "Test Username",
			Password: "Test Password",
		}

		sess := session.New(credentials)

		It("The credentials should be set on the session", func() {
			Expect(sess.Credentials).To(Equal(credentials))
		})
	})
})
