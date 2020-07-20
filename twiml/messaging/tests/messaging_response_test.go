package tests

import (
	"io/ioutil"
	"regexp"

	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging/verbs"

	"github.com/RJPearson94/twilio-sdk-go/twiml/messaging"
	"github.com/RJPearson94/twilio-sdk-go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ignoreRegex = regexp.MustCompile(`\r?\n|\s\s`)

var _ = Describe("Messaging Response TwiML", func() {
	Describe("Given I need to generate a message response with a message", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/message.golden.xml")

			response := messaging.New()
			response.Message(utils.String("Hello World"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a message response with a message body and media", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/messageBodyAndMedia.golden.xml")

			response := messaging.New()
			message := response.Message(nil)
			message.Body("Hello world")
			message.Media("https://localhost/media")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a message response with a message containing attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/messageWithAttributes.golden.xml")

			response := messaging.New()
			response.MessageWithAttributes(verbs.MessageAttributes{
				Action: utils.String("http://localhost/action"),
				Method: utils.String("POST"),
			}, utils.String("Hello World"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a message response with multiple messages", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/multipleMessages.golden.xml")

			response := messaging.New()
			response.Message(utils.String("Hello World"))
			response.Message(utils.String("Hello World Again"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a message response with redirect", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/redirect.golden.xml")

			response := messaging.New()
			response.Redirect("http://localhost/redirect")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a message response with redirect and message", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/redirectAndMessage.golden.xml")

			response := messaging.New()
			response.Redirect("http://localhost/redirect")
			response.Message(utils.String("Hello World"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

})

func CompareXML(actual string, expected string) {
	Expect(ignoreRegex.ReplaceAllString(actual, "")).To(Equal(ignoreRegex.ReplaceAllString(expected, "")))
}
