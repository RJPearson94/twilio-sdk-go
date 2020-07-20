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

	Describe("Given I need to generate a voice response with dial noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dial.golden.xml")

			response := voice.New()
			response.Dial(utils.String("415-123-9876"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a dial client", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialClient.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Client(utils.String("RJPearson94"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a dial client with attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialClientWithAttributes.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.ClientWithAttributes(nouns.ClientAttributes{
				StatusCallbackEvent: utils.String("initiated ringing"),
				StatusCallback:      utils.String("http://localhost/callback"),
			}, utils.String("RJPearson94"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a dial client with parameter", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialClientWithCustomParameter.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			client := dial.Client(nil)
			client.Identity("RJPearson94")
			client.ParameterWithAttributes(nouns.ParameterAttributes{
				Name:  utils.String("VIP"),
				Value: utils.String("true"),
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

	Describe("Given I need to generate a voice response with simultaneous dialing", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialSimultaneousDialing.golden.xml")

			response := voice.New()
			dial := response.DialWithAttributes(verbs.DialAttributes{
				CallerID: utils.String("+4151234567"),
			}, nil)
			dial.Number("415-123-9876")
			dial.Client(utils.String("RJPearson94"))
			dial.Client(utils.String("Rob"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial conference noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialConference.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Conference("Room Hello World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial conference attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialConferenceWithAttributes.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.ConferenceWithAttributes(nouns.ConferenceAttributes{
				StartConferenceOnEnter: utils.Bool(false),
			}, "Room Hello World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial number noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialNumber.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Number("415-123-9876")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial number attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialNumberWithAttributes.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.NumberWithAttributes(nouns.NumberAttributes{
				StatusCallbackEvent:  utils.String("answered"),
				StatusCallback:       utils.String("https://localhost/answered"),
				StatusCallbackMethod: utils.String("POST"),
			}, "415-123-9876")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with simultaneous number dialing", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialSimultaneousNumberDialing.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Number("415-123-9876")
			dial.Number("415-123-9877")
			dial.Number("415-123-9878")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial queue noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialQueue.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Queue("test")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial queue attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialQueueWithAttributes.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.QueueWithAttributes(nouns.QueueAttributes{
				URL: utils.String("http://localhost/test.xml"),
			}, "test")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial sim noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialSim.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Sim("DEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial sip noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialSip.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.Sip("test@test.com")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with dial sip attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialSipWithAttributes.golden.xml")

			response := voice.New()
			dial := response.Dial(nil)
			dial.SipWithAttributes(nouns.SipAttributes{
				Username: utils.String("test"),
				Password: utils.String("test"),
			}, "test@test.com")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with enqueue verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/enqueue.golden.xml")

			response := voice.New()
			response.Enqueue(utils.String("test"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with enqueue attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/enqueueWithAttributes.golden.xml")

			response := voice.New()
			response.EnqueueWithAttributes(verbs.EnqueueAttributes{
				WaitURL: utils.String("http://localhost/test.xml"),
			}, utils.String("test"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with enqueue task noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/enqueueTask.golden.xml")

			response := voice.New()
			enqueue := response.Enqueue(nil)
			enqueue.Task("test")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with enqueue attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/enqueueTaskWithAttributes.golden.xml")

			response := voice.New()
			enqueue := response.Enqueue(nil)
			enqueue.TaskWithAttributes(nouns.TaskAttributes{
				Priority: utils.Int(1),
				Timeout:  utils.Int(2),
			}, "test")
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
