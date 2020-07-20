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
			connect := response.Connect()
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
			connect := response.Connect()
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
			connect := response.Connect()
			connect.Stream()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a stream attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/connectStreamWithAttributes.golden.xml")

			response := voice.New()
			connect := response.Connect()
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
			client.Parameter()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with a dial client with parameter attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/dialClientWithCustomParameterAttributes.golden.xml")

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

	Describe("Given I need to generate a voice response with gather verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/gather.golden.xml")

			response := voice.New()
			response.Gather()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with gather verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/gather.golden.xml")

			response := voice.New()
			response.Gather()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with gather attributes and play noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/gatherWithAttributesAndSayNoun.golden.xml")

			response := voice.New()
			gather := response.GatherWithAttributes(verbs.GatherAttributes{
				Input:     utils.String("speech dtmf"),
				Timeout:   utils.Int(3),
				NumDigits: utils.Int(1),
			})
			gather.Say("Please press 1")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with gather play noun", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/gatherPlay.golden.xml")

			response := voice.New()
			gather := response.Gather()
			gather.Play(utils.String("https://localhost/test.mp3"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with gather pause nouns", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/gatherPause.golden.xml")

			response := voice.New()
			gather := response.Gather()
			gather.Pause()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with gather say and pause nouns", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/gatherSayAndPause.golden.xml")

			response := voice.New()
			gather := response.Gather()
			gather.Say("Hello")
			gather.PauseWithAttributes(verbs.PauseAttributes{
				Length: utils.Int(10),
			})
			gather.Say("World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with hangup verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/hangup.golden.xml")

			response := voice.New()
			response.Hangup()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with leave verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/leave.golden.xml")

			response := voice.New()
			response.Leave()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with pause verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/pause.golden.xml")

			response := voice.New()
			response.Pause()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with pause attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/pauseWithAttributes.golden.xml")

			response := voice.New()
			response.PauseWithAttributes(verbs.PauseAttributes{
				Length: utils.Int(10),
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

	Describe("Given I need to generate a voice response with play verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/play.golden.xml")

			response := voice.New()
			response.Play(utils.String("https://localhost/test.mp3"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with play attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/playWithAttributes.golden.xml")

			response := voice.New()
			response.PlayWithAttributes(verbs.PlayAttributes{
				Loop: utils.Int(10),
			}, utils.String("https://localhost/test.mp3"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with pay verb", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/pay.golden.xml")

			response := voice.New()
			response.Pay()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with pay attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/payWithAttributes.golden.xml")

			response := voice.New()
			response.PayWithAttributes(verbs.PayAttributes{
				ChargeAmount: utils.String("9.99"),
				Action:       utils.String("http://localhost/pay"),
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

	Describe("Given I need to generate a voice response with pay parameter", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/payParameter.golden.xml")

			response := voice.New()
			pay := response.Pay()
			pay.Parameter()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with pay parameter attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/payParameterWithAttributes.golden.xml")

			response := voice.New()
			pay := response.Pay()
			pay.ParameterWithAttributes(nouns.ParameterAttributes{
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

	Describe("Given I need to generate a voice response with pay prompt", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/payPrompt.golden.xml")

			response := voice.New()
			pay := response.Pay()
			pay.Prompt()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with pay prompt attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/payPromptWithAttributes.golden.xml")

			response := voice.New()
			pay := response.Pay()
			prompt := pay.PromptWithAttributes(verbs.PromptAttributes{
				For:     utils.String("cvv"),
				Attempt: utils.Int(1),
			})
			prompt.Say("Please enter your cvv which is on the back of your card")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with prompt", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/prompt.golden.xml")

			response := voice.New()
			response.Prompt()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with prompt attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/promptWithAttributes.golden.xml")

			response := voice.New()
			response.PromptWithAttributes(verbs.PromptAttributes{
				For:     utils.String("cvv"),
				Attempt: utils.Int(1),
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

	Describe("Given I need to generate a voice response with prompt pause, play and say verbs", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/promptWithPauseSayAndPlayVerbs.golden.xml")

			response := voice.New()
			prompt := response.Prompt()
			prompt.Say("Please enter your cvv which is on the back of your card")
			prompt.PauseWithAttributes(verbs.PauseAttributes{
				Length: utils.Int(10),
			})
			prompt.Play(utils.String("http://localhost/cvv-prompt"))
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with queue", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/queue.golden.xml")

			response := voice.New()
			response.Queue("test")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with queue attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/queueWithAttributes.golden.xml")

			response := voice.New()
			response.QueueWithAttributes(nouns.QueueAttributes{
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

	Describe("Given I need to generate a voice response with record", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/record.golden.xml")

			response := voice.New()
			response.Record()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with record attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/recordWithAttributes.golden.xml")

			response := voice.New()
			response.RecordWithAttributes(verbs.RecordAttributes{
				Transcribe:         utils.Bool(true),
				TranscribeCallback: utils.String("http://localhost/transcript"),
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

	Describe("Given I need to generate a voice response with redirect", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/redirect.golden.xml")

			response := voice.New()
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

	Describe("Given I need to generate a voice response with redirect attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/redirectWithAttributes.golden.xml")

			response := voice.New()
			response.RedirectWithAttributes(verbs.RedirectAttributes{
				Method: utils.String("POST"),
			}, "http://localhost/redirect")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with refer", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/refer.golden.xml")

			response := voice.New()
			response.Refer()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with refer attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/referWithAttributes.golden.xml")

			response := voice.New()
			response.ReferWithAttributes(verbs.ReferAttributes{
				Action: utils.String("http://localhost/refer"),
				Method: utils.String("POST"),
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

	Describe("Given I need to generate a voice response with refer sip", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/referSip.golden.xml")

			response := voice.New()
			refer := response.Refer()
			refer.Sip("test@test.com")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with refer sip attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/referSipWithAttributes.golden.xml")

			response := voice.New()
			refer := response.Refer()
			refer.SipWithAttributes(nouns.SipAttributes{
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

	Describe("Given I need to generate a voice response with refer and refer sip", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/referWithReferSip.golden.xml")

			response := voice.New()
			refer := response.Refer()
			refer.ReferSip("refer@test.com")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with refer and refer sip attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/referWithReferSipAttributes.golden.xml")

			response := voice.New()
			refer := response.Refer()
			refer.ReferSipWithAttributes(nouns.ReferSipAttributes{
				Username: utils.String("test"),
				Password: utils.String("test"),
			}, "refer@test.com")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with reject", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/reject.golden.xml")

			response := voice.New()
			response.Reject()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with reject attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/rejectWithAttributes.golden.xml")

			response := voice.New()
			response.RejectWithAttributes(verbs.RejectAttributes{
				Reason: utils.String("busy"),
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

	Describe("Given I need to generate a voice response with say", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/say.golden.xml")

			response := voice.New()
			response.Say("Hello World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with say attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/sayWithAttributes.golden.xml")

			response := voice.New()
			response.SayWithAttributes(verbs.SayAttributes{
				Voice:    utils.String("man"),
				Language: utils.String("en-US"),
			}, "Hello World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with sms", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/sms.golden.xml")

			response := voice.New()
			response.Sms("Hello World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with sms attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/smsWithAttributes.golden.xml")

			response := voice.New()
			response.SmsWithAttributes(verbs.SmsAttributes{
				To:             utils.String("+1987654321"),
				From:           utils.String("+10123456789"),
				StatusCallback: utils.String("http://localhost/sms-callback"),
			}, "Hello World")
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with start", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/start.golden.xml")

			response := voice.New()
			response.Start()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with start attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/startWithAttributes.golden.xml")

			response := voice.New()
			response.StartWithAttributes(verbs.StartAttributes{
				Action: utils.String("http://localhost/start"),
				Method: utils.String("POST"),
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

	Describe("Given I need to generate a voice response with start stream", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/startStream.golden.xml")

			response := voice.New()
			start := response.Start()
			start.Stream()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with start stream attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/startStreamWithAttributes.golden.xml")

			response := voice.New()
			start := response.Start()
			start.StreamWithAttributes(nouns.StreamAttributes{
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

	Describe("Given I need to generate a voice response with start siprec", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/startSiprec.golden.xml")

			response := voice.New()
			start := response.Start()
			start.Siprec()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with start siprec attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/startSiprecWithAttributes.golden.xml")

			response := voice.New()
			start := response.Start()
			start.SiprecWithAttributes(nouns.SiprecAttributes{
				Name: utils.String("test"),
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

	Describe("Given I need to generate a voice response with stop", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/stop.golden.xml")

			response := voice.New()
			response.Stop()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with stop attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/stopWithAttributes.golden.xml")

			response := voice.New()
			response.StopWithAttributes(verbs.StopAttributes{
				Action: utils.String("http://localhost/stop"),
				Method: utils.String("POST"),
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

	Describe("Given I need to generate a voice response with stop stream", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/stopStream.golden.xml")

			response := voice.New()
			stop := response.Stop()
			stop.Stream()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with stop stream attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/stopStreamWithAttributes.golden.xml")

			response := voice.New()
			stop := response.Stop()
			stop.StreamWithAttributes(nouns.StreamAttributes{
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

	Describe("Given I need to generate a voice response with stop siprec", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/stopSiprec.golden.xml")

			response := voice.New()
			stop := response.Stop()
			stop.Siprec()
			twiML, err := response.ToTwiML()

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the twiML should match the golden data", func() {
				CompareXML(*twiML, string(goldenData))
			})
		})
	})

	Describe("Given I need to generate a voice response with stop siprec attributes", func() {
		Describe("When the twiML is generated", func() {
			goldenData, _ := ioutil.ReadFile("testdata/stopSiprecWithAttributes.golden.xml")

			response := voice.New()
			stop := response.Stop()
			stop.SiprecWithAttributes(nouns.SiprecAttributes{
				Name: utils.String("test"),
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
