package tests

import (
	"io/ioutil"
	"regexp"

	"github.com/RJPearson94/twilio-sdk-go/twiml/voice"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs"
	"github.com/RJPearson94/twilio-sdk-go/twiml/voice/verbs/nouns"
	"github.com/RJPearson94/twilio-sdk-go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ignoreRegex = regexp.MustCompile(`\r?\n|\s\s`)

var _ = Describe("Voice Response TwiML", func() {
	Describe("Given I need to generate a voice response with a autopilot connect", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/connectAutopilot.golden.xml")

			response := voice.New()
			connect := response.ConnectWithAttributes(verbs.ConnectAttributes{})
			connect.Autopilot("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a room connect", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/connectRoom.golden.xml")

			response := voice.New()
			connect := response.ConnectWithAttributes(verbs.ConnectAttributes{})
			connect.Room("HelloWorld")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a room connect with attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/connectRoomWithAttributes.golden.xml")

			response := voice.New()
			connect := response.ConnectWithAttributes(verbs.ConnectAttributes{})
			connect.RoomWithAttributes(nouns.RoomAttributes{
				ParticipantIdentity: utils.String("RJPearson94"),
			}, "HelloWorld")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a stream connect", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/connectStream.golden.xml")

			response := voice.New()
			connect := response.ConnectWithAttributes(verbs.ConnectAttributes{})
			connect.StreamWithAttributes(nouns.StreamAttributes{
				URL: utils.String("wss://localhost/stream"),
			})
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
