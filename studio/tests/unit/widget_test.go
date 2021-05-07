package tests

import (
	"io/ioutil"

	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/transition"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/RJPearson94/twilio-sdk-go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Studio", func() {
	Describe("Trigger widget with basic config", func() {
		triggerJSON, _ := ioutil.ReadFile("testdata/trigger.json")

		trigger := widgets.Trigger{
			Name: "Trigger",
		}

		assertJSONMatches(trigger, triggerJSON)
	})

	Describe("Trigger widget with transitions", func() {
		triggerJSON, _ := ioutil.ReadFile("testdata/triggerComplete.json")

		trigger := widgets.Trigger{
			Name: "Trigger",
			NextTransitions: widgets.TriggerNextTransitions{
				IncomingCall:    utils.String("call"),
				IncomingMessage: utils.String("message"),
				IncomingRequest: utils.String("request"),
			},
			Properties: widgets.TriggerProperties{
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(trigger, triggerJSON)
	})

	Describe("Send To Flex widget with basic config", func() {
		sendToFlexJSON, _ := ioutil.ReadFile("testdata/sendToFlex.json")

		sendToFlex := widgets.SendToFlex{
			Name: "SendToFlex",
			Properties: widgets.SendToFlexProperties{
				Workflow:   "WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Channel:    "TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Attributes: utils.String("{\"name\":\"{{trigger.message.ChannelAttributes.from}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"channelSid\":\"{{trigger.message.ChannelSid}}\"}"),
			},
		}

		assertJSONMatches(sendToFlex, sendToFlexJSON)
	})

	Describe("Send To Flex widget with all config", func() {
		sendToFlexJSON, _ := ioutil.ReadFile("testdata/sendToFlexComplete.json")

		sendToFlex := widgets.SendToFlex{
			Name: "SendToFlex",
			NextTransitions: widgets.SendToFlexNextTransitions{
				CallComplete:    utils.String("complete"),
				CallFailure:     utils.String("failure"),
				FailedToEnqueue: utils.String("enqueue"),
			},
			Properties: widgets.SendToFlexProperties{
				Workflow:      "WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Channel:       "TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				Attributes:    utils.String("{\"name\":\"{{trigger.message.ChannelAttributes.from}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"channelSid\":\"{{trigger.message.ChannelSid}}\"}"),
				Priority:      utils.String("10"),
				WaitURL:       utils.String("https://test.com/hold"),
				WaitURLMethod: utils.String("POST"),
				Timeout:       utils.String("3600"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(sendToFlex, sendToFlexJSON)
	})

	Describe("Send To Autopilot widget with basic config", func() {
		sendToAutopilotJSON, _ := ioutil.ReadFile("testdata/sendToAutopilot.json")

		sendToAutopilot := widgets.SendToAutopilot{
			Name: "SendToAutopilot",
			Properties: widgets.SendToAutopilotProperties{
				AutopilotAssistantSid: "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				From:                  "{{flow.channel.address}}",
				Body:                  "{{trigger.message.Body}}",
				Timeout:               14400,
			},
		}

		assertJSONMatches(sendToAutopilot, sendToAutopilotJSON)
	})

	Describe("Send To Autopilot widget with all config", func() {
		sendToAutopilotJSON, _ := ioutil.ReadFile("testdata/sendToAutopilotComplete.json")

		sendToAutopilot := widgets.SendToAutopilot{
			Name: "SendToAutopilot",
			NextTransitions: widgets.SendToAutopilotNextTransitions{
				Failure:      utils.String("failure"),
				SessionEnded: utils.String("sessionEnded"),
				Timeout:      utils.String("timeout"),
			},
			Properties: widgets.SendToAutopilotProperties{
				AutopilotAssistantSid: "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				From:                  "{{flow.channel.address}}",
				Body:                  "{{trigger.message.Body}}",
				Timeout:               14400,
				ChatChannel:           utils.String("{{trigger.message.ChannelSid}}"),
				ChatAttributes:        utils.String("{\"name\":\"{{trigger.message.ChannelAttributes.from}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"channelSid\":\"{{trigger.message.ChannelSid}}\"}"),
				ChatService:           utils.String("{{trigger.message.InstanceSid}}"),
				TargetTask:            utils.String("Task"),
				MemoryParameters: &[]widgets.SendToAutopilotMemoryParameter{
					{
						Key:   "key",
						Value: "value",
					},
				},
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(sendToAutopilot, sendToAutopilotJSON)
	})

	Describe("Add TwiML redirect widget with basic config", func() {
		addTwimlRedirectJSON, _ := ioutil.ReadFile("testdata/addTwimlRedirect.json")

		addTwimlRedirect := widgets.AddTwimlRedirect{
			Name: "AddTwiMLRedirect",
			Properties: widgets.AddTwimlRedirectProperties{
				Method: utils.String("POST"),
				URL:    "https://test.com/twiml",
			},
		}

		assertJSONMatches(addTwimlRedirect, addTwimlRedirectJSON)
	})

	Describe("Add TwiML redirect widget with all config", func() {
		addTwimlRedirectJSON, _ := ioutil.ReadFile("testdata/addTwimlRedirectComplete.json")

		addTwimlRedirect := widgets.AddTwimlRedirect{
			Name: "AddTwiMLRedirect",
			NextTransitions: widgets.AddTwimlRedirectNextTransitions{
				Fail:    utils.String("fail"),
				Return:  utils.String("return"),
				Timeout: utils.String("timeout"),
			},
			Properties: widgets.AddTwimlRedirectProperties{
				Method:  utils.String("POST"),
				URL:     "https://test.com/twiml",
				Timeout: utils.String("100"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(addTwimlRedirect, addTwimlRedirectJSON)
	})

	Describe("Make HTTP request widget with basic config", func() {
		makeHTTPRequestJSON, _ := ioutil.ReadFile("testdata/makeHTTPRequest.json")

		makeHTTPRequest := widgets.MakeHTTPRequest{
			Name: "MakeHTTPRequest",
			Properties: widgets.MakeHTTPRequestProperties{
				Method:      "GET",
				URL:         "https://test.com",
				ContentType: "application/x-www-form-urlencoded;charset=utf-8",
			},
		}

		assertJSONMatches(makeHTTPRequest, makeHTTPRequestJSON)
	})

	Describe("Make HTTP request widget with all config", func() {
		makeHTTPRequestJSON, _ := ioutil.ReadFile("testdata/makeHTTPRequestComplete.json")

		makeHTTPRequest := widgets.MakeHTTPRequest{
			Name: "MakeHTTPRequest",
			NextTransitions: widgets.MakeHTTPRequestNextTransitions{
				Success: utils.String("success"),
				Failed:  utils.String("failed"),
			},
			Properties: widgets.MakeHTTPRequestProperties{
				Method:      "GET",
				URL:         "https://test.com",
				ContentType: "application/x-www-form-urlencoded;charset=utf-8",
				Body:        utils.String("Hello World"),
				Parameters: &[]widgets.MakeHTTPRequestParameter{
					{
						Value: "value",
						Key:   "key",
					},
				},
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(makeHTTPRequest, makeHTTPRequestJSON)
	})

	Describe("Run function widget with basic config", func() {
		runFunctionJSON, _ := ioutil.ReadFile("testdata/runFunction.json")

		runFunction := widgets.RunFunction{
			Name: "RunFunction",
			Properties: widgets.RunFunctionProperties{
				ServiceSid: utils.String("default"),
				URL:        "https://test-function.twil.io/test-function",
			},
		}

		assertJSONMatches(runFunction, runFunctionJSON)
	})

	Describe("Run function widget with all config", func() {
		runFunctionJSON, _ := ioutil.ReadFile("testdata/runFunctionComplete.json")

		runFunction := widgets.RunFunction{
			Name: "RunFunction",
			NextTransitions: widgets.RunFunctionNextTransitions{
				Success: utils.String("success"),
				Fail:    utils.String("fail"),
			},
			Properties: widgets.RunFunctionProperties{
				ServiceSid:     utils.String("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				URL:            "https://test-function.twil.io/test-function",
				EnvironmentSid: utils.String("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				FunctionSid:    utils.String("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				Parameters: &[]widgets.RunFunctionParameter{
					{
						Value: "value",
						Key:   "key",
					},
				},
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(runFunction, runFunctionJSON)
	})

	Describe("Send and wait for reply widget with basic config", func() {
		sendAndWaitForReplyJSON, _ := ioutil.ReadFile("testdata/sendAndWaitForReply.json")

		sendAndWaitForReply := widgets.SendAndWaitForReply{
			Name: "SendAndWaitForReply",
			NextTransitions: widgets.SendAndWaitForReplyNextTransitions{
				IncomingMessage: "test",
			},
			Properties: widgets.SendAndWaitForReplyProperties{
				From:    "{{flow.channel.address}}",
				Body:    "Test",
				Timeout: "3600",
			},
		}

		assertJSONMatches(sendAndWaitForReply, sendAndWaitForReplyJSON)
	})

	Describe("Send and wait for reply widget with all config", func() {
		sendAndWaitForReplyJSON, _ := ioutil.ReadFile("testdata/sendAndWaitForReplyComplete.json")

		sendAndWaitForReply := widgets.SendAndWaitForReply{
			Name: "SendAndWaitForReply",
			NextTransitions: widgets.SendAndWaitForReplyNextTransitions{
				IncomingMessage: "incomingMessage",
				DeliveryFailure: utils.String("deliveryFailure"),
				Timeout:         utils.String("timeout"),
			},
			Properties: widgets.SendAndWaitForReplyProperties{
				From:       "{{flow.channel.address}}",
				Body:       "Test",
				Timeout:    "3600",
				Channel:    utils.String("{{trigger.message.ChannelSid}}"),
				Service:    utils.String("{{trigger.message.InstanceSid}}"),
				Attributes: utils.String("{\"name\":\"{{trigger.message.ChannelAttributes.from}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"channelSid\":\"{{trigger.message.ChannelSid}}\"}"),
				MediaURL:   utils.String("https://test.com"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(sendAndWaitForReply, sendAndWaitForReplyJSON)
	})

	Describe("Send message widget with basic config", func() {
		sendMessageJSON, _ := ioutil.ReadFile("testdata/sendMessage.json")

		sendMessage := widgets.SendMessage{
			Name: "SendMessage",
			Properties: widgets.SendMessageProperties{
				From: "{{flow.channel.address}}",
				To:   "{{contact.channel.address}}",
				Body: "Hello World",
			},
		}

		assertJSONMatches(sendMessage, sendMessageJSON)
	})

	Describe("Send message widget with all config", func() {
		sendMessageJSON, _ := ioutil.ReadFile("testdata/sendMessageComplete.json")

		sendMessage := widgets.SendMessage{
			Name: "SendMessage",
			NextTransitions: widgets.SendMessageNextTransitions{
				Failed: utils.String("failed"),
				Sent:   utils.String("sent"),
			},
			Properties: widgets.SendMessageProperties{
				From:       "{{flow.channel.address}}",
				To:         "{{contact.channel.address}}",
				Body:       "Hello World",
				Channel:    utils.String("{{trigger.message.ChannelSid}}"),
				Service:    utils.String("{{trigger.message.InstanceSid}}"),
				Attributes: utils.String("{\"name\":\"{{trigger.message.ChannelAttributes.from}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"channelSid\":\"{{trigger.message.ChannelSid}}\"}"),
				MediaURL:   utils.String("https://test.com"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(sendMessage, sendMessageJSON)
	})

	Describe("Connect virtual agent with basic config", func() {
		connectVirtualAgentJSON, _ := ioutil.ReadFile("testdata/connectVirtualAgent.json")

		connectVirtualAgent := widgets.ConnectVirtualAgent{
			Name: "ConnectVirtualAgent",
			Properties: widgets.ConnectVirtualAgentProperties{
				Connector: "test-connector",
			},
		}

		assertJSONMatches(connectVirtualAgent, connectVirtualAgentJSON)
	})

	Describe("Connect virtual agent with all config", func() {
		connectVirtualAgentJSON, _ := ioutil.ReadFile("testdata/connectVirtualAgentComplete.json")

		connectVirtualAgent := widgets.ConnectVirtualAgent{
			Name: "ConnectVirtualAgent",
			NextTransitions: widgets.ConnectVirtualAgentNextTransitions{
				Hangup: utils.String("hangup"),
				Return: utils.String("return"),
			},
			Properties: widgets.ConnectVirtualAgentProperties{
				Connector:         "test-connector",
				SentimentAnalysis: utils.String("true"),
				Language:          utils.String("en-US"),
				StatusCallbackURL: utils.String("https://test.com"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(connectVirtualAgent, connectVirtualAgentJSON)
	})

	Describe("Fork stream with basic config", func() {
		forkStreamJSON, _ := ioutil.ReadFile("testdata/forkStream.json")

		forkStream := widgets.ForkStream{
			Name: "ForkStream",
			Properties: widgets.ForkStreamProperties{
				StreamName:          utils.String("test"),
				StreamTransportType: utils.String("websocket"),
				StreamAction:        "start",
				StreamTrack:         utils.String("inbound_track"),
				StreamURL:           utils.String("wss://test.com"),
			},
		}

		assertJSONMatches(forkStream, forkStreamJSON)
	})

	Describe("Fork stream with stop config", func() {
		forkStreamJSON, _ := ioutil.ReadFile("testdata/forkStreamStop.json")

		forkStream := widgets.ForkStream{
			Name: "ForkStream",
			Properties: widgets.ForkStreamProperties{
				StreamTransportType: utils.String("websocket"),
				StreamAction:        "stop",
			},
		}

		assertJSONMatches(forkStream, forkStreamJSON)
	})

	Describe("Fork stream with all config", func() {
		forkStreamJSON, _ := ioutil.ReadFile("testdata/forkStreamComplete.json")

		forkStream := widgets.ForkStream{
			Name: "ForkStream",
			NextTransitions: widgets.ForkStreamNextTransitions{
				Next: utils.String("next"),
			},
			Properties: widgets.ForkStreamProperties{
				StreamName:          utils.String("test"),
				StreamConnector:     utils.String("connector"),
				StreamTransportType: utils.String("siprec"),
				StreamAction:        "start",
				StreamTrack:         utils.String("inbound_track"),
				StreamParameters: &[]widgets.ForkStreamStreamParameter{
					{
						Key:   "key",
						Value: "value",
					},
				},
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(forkStream, forkStreamJSON)
	})

	Describe("Say Play with digit config", func() {
		sayPlayJSON, _ := ioutil.ReadFile("testdata/sayPlayDigits.json")

		sayPlay := widgets.SayPlay{
			Name: "SayPlay",
			Properties: widgets.SayPlayProperties{
				Digits: utils.String("123"),
			},
		}

		assertJSONMatches(sayPlay, sayPlayJSON)
	})

	Describe("Say Play with say message config", func() {
		sayPlayJSON, _ := ioutil.ReadFile("testdata/sayPlaySayMessage.json")

		sayPlay := widgets.SayPlay{
			Name: "SayPlay",
			Properties: widgets.SayPlayProperties{
				Say: utils.String("Test"),
			},
		}

		assertJSONMatches(sayPlay, sayPlayJSON)
	})

	Describe("Say Play with play message config", func() {
		sayPlayJSON, _ := ioutil.ReadFile("testdata/sayPlayPlayMessage.json")

		sayPlay := widgets.SayPlay{
			Name: "SayPlay",
			Properties: widgets.SayPlayProperties{
				Play: utils.String("http://localhost.com"),
			},
		}

		assertJSONMatches(sayPlay, sayPlayJSON)
	})

	Describe("Say Play with all config", func() {
		sayPlayJSON, _ := ioutil.ReadFile("testdata/sayPlayComplete.json")

		sayPlay := widgets.SayPlay{
			Name: "SayPlay",
			NextTransitions: widgets.SayPlayNextTransitions{
				AudioComplete: utils.String("audioComplete"),
			},
			Properties: widgets.SayPlayProperties{
				Say:      utils.String("Test"),
				Language: utils.String("en-US"),
				Voice:    utils.String("alice"),
				Loop:     utils.Int(2),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(sayPlay, sayPlayJSON)
	})

	Describe("Enqueue call with TaskRouter config", func() {
		enqueueCallJSON, _ := ioutil.ReadFile("testdata/enqueueCallTaskRouter.json")

		enqueueCall := widgets.EnqueueCall{
			Name: "EnqueueCall",
			Properties: widgets.EnqueueCallProperties{
				WorkflowSid: utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			},
		}

		assertJSONMatches(enqueueCall, enqueueCallJSON)
	})

	Describe("Enqueue call with queue config", func() {
		enqueueCallJSON, _ := ioutil.ReadFile("testdata/enqueueCallQueue.json")

		enqueueCall := widgets.EnqueueCall{
			Name: "EnqueueCall",
			Properties: widgets.EnqueueCallProperties{
				QueueName: utils.String("Test"),
			},
		}

		assertJSONMatches(enqueueCall, enqueueCallJSON)
	})

	Describe("Enqueue call with all config", func() {
		enqueueCallJSON, _ := ioutil.ReadFile("testdata/enqueueCallComplete.json")

		enqueueCall := widgets.EnqueueCall{
			Name: "EnqueueCall",
			NextTransitions: widgets.EnqueueCallNextTransitions{
				CallComplete:    utils.String("callComplete"),
				CallFailure:     utils.String("callFailure"),
				FailedToEnqueue: utils.String("failedToEnqueue"),
			},
			Properties: widgets.EnqueueCallProperties{
				WorkflowSid:    utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
				WaitURL:        utils.String("http://localhost.com"),
				WaitURLMethod:  utils.String("POST"),
				Priority:       utils.Int(1),
				TaskAttributes: utils.String("{\"test\": \"test\"}"),
				Timeout:        utils.Int(10),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(enqueueCall, enqueueCallJSON)
	})

	Describe("Record call with basic config", func() {
		recordCallJSON, _ := ioutil.ReadFile("testdata/recordCall.json")

		recordCall := widgets.RecordCall{
			Name: "RecordCall",
			Properties: widgets.RecordCallProperties{
				RecordCall: false,
			},
		}

		assertJSONMatches(recordCall, recordCallJSON)
	})

	Describe("Record call with all config", func() {
		recordCallJSON, _ := ioutil.ReadFile("testdata/recordCallComplete.json")

		recordCall := widgets.RecordCall{
			Name: "RecordCall",
			NextTransitions: widgets.RecordCallNextTransitions{
				Failed:  utils.String("failed"),
				Success: utils.String("success"),
			},
			Properties: widgets.RecordCallProperties{
				RecordCall:                    true,
				Trim:                          utils.String("do-not-trim"),
				RecordingStatusCallbackURL:    utils.String("http://localhost.com"),
				RecordingStatusCallbackMethod: utils.String("GET"),
				RecordingStatusCallbackEvents: utils.String("in-progress completed"),
				RecordingChannels:             utils.String("mono"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(recordCall, recordCallJSON)
	})

	Describe("Make outgoing call with basic config", func() {
		makeOutgoingCallJSON, _ := ioutil.ReadFile("testdata/makeOutgoingCall.json")

		makeOutgoingCall := widgets.MakeOutgoingCall{
			Name: "MakeOutgoingCall",
			Properties: widgets.MakeOutgoingCallProperties{
				From: "{{flow.channel.address}}",
				To:   "{{contact.channel.address}}",
			},
		}

		assertJSONMatches(makeOutgoingCall, makeOutgoingCallJSON)
	})

	Describe("Make outgoing call with all config", func() {
		makeOutgoingCallJSON, _ := ioutil.ReadFile("testdata/makeOutgoingCallComplete.json")

		makeOutgoingCall := widgets.MakeOutgoingCall{
			Name: "MakeOutgoingCall",
			NextTransitions: widgets.MakeOutgoingCallNextTransitions{
				Answered: utils.String("answered"),
				Busy:     utils.String("busy"),
				Failed:   utils.String("failed"),
				NoAnswer: utils.String("noAnswer"),
			},
			Properties: widgets.MakeOutgoingCallProperties{
				From:                               "{{flow.channel.address}}",
				To:                                 "{{contact.channel.address}}",
				MachineDetection:                   utils.String("Enable"),
				MachineDetectionSpeechThreshold:    utils.String("100"),
				MachineDetectionSpeechEndThreshold: utils.String("150"),
				MachineDetectionTimeout:            utils.String("10"),
				MachineDetectionSilenceTimeout:     utils.String("1000"),
				DetectAnsweringMachine:             utils.Bool(true),
				SipAuthUsername:                    utils.String("test"),
				SipAuthPassword:                    utils.String("test2"),
				SendDigits:                         utils.String("1234"),
				Timeout:                            utils.Int(10),
				Trim:                               utils.String("trim-silence"),
				Record:                             utils.Bool(true),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(makeOutgoingCall, makeOutgoingCallJSON)
	})

	Describe("Record voicemail with basic config", func() {
		recordVoicemailJSON, _ := ioutil.ReadFile("testdata/recordVoicemail.json")

		recordVoicemail := widgets.RecordVoicemail{
			Name:       "RecordVoicemail",
			Properties: widgets.RecordVoicemailProperties{},
		}

		assertJSONMatches(recordVoicemail, recordVoicemailJSON)
	})

	Describe("Record voicemail with all config", func() {
		recordVoicemailJSON, _ := ioutil.ReadFile("testdata/recordVoicemailComplete.json")

		recordVoicemail := widgets.RecordVoicemail{
			Name: "RecordVoicemail",
			NextTransitions: widgets.RecordVoicemailNextTransitions{
				RecordingComplete: utils.String("recordingComplete"),
				NoAudio:           utils.String("noAudio"),
				Hangup:            utils.String("hangup"),
			},
			Properties: widgets.RecordVoicemailProperties{
				Transcribe:                 utils.Bool(true),
				Trim:                       utils.String("trim-silence"),
				TranscriptionCallbackURL:   utils.String("http://localhost.com/transcript"),
				PlayBeep:                   utils.String("true"),
				FinishOnKey:                utils.String("1"),
				RecordingStatusCallbackURL: utils.String("http://localhost.com/recording"),
				Timeout:                    utils.Int(10),
				MaxLength:                  utils.Int(1000),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(recordVoicemail, recordVoicemailJSON)
	})

	Describe("Connect call to with client config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToClient.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			Properties: widgets.ConnectCallToProperties{
				CallerID: "{{contact.channel.address}}",
				Noun:     "client",
				To:       utils.String("test"),
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Connect call to with conference config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToConference.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			Properties: widgets.ConnectCallToProperties{
				CallerID: "{{contact.channel.address}}",
				Noun:     "conference",
				To:       utils.String("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Connect call to with number config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToNumber.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			Properties: widgets.ConnectCallToProperties{
				CallerID: "{{contact.channel.address}}",
				Noun:     "number",
				To:       utils.String("+441234567890"),
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Connect call to with multi number config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToNumberMulti.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			Properties: widgets.ConnectCallToProperties{
				CallerID: "{{contact.channel.address}}",
				Noun:     "number-multi",
				To:       utils.String("+441234567890,+441234567891"),
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Connect call to with SIM config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToSIM.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			Properties: widgets.ConnectCallToProperties{
				CallerID: "{{contact.channel.address}}",
				Noun:     "sim",
				To:       utils.String("DEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Connect call to with SIP config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToSIP.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			Properties: widgets.ConnectCallToProperties{
				CallerID:    "{{contact.channel.address}}",
				Noun:        "sip",
				SipEndpoint: utils.String("sip:test@test.com"),
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Connect call to with all SIP config", func() {
		connectCallToJSON, _ := ioutil.ReadFile("testdata/connectCallToSIPComplete.json")

		connectCallTo := widgets.ConnectCallTo{
			Name: "ConnectCallTo",
			NextTransitions: widgets.ConnectCallToNextTransitions{
				CallCompleted: utils.String("callCompleted"),
				Hangup:        utils.String("hangup"),
			},
			Properties: widgets.ConnectCallToProperties{
				CallerID:    "{{contact.channel.address}}",
				Noun:        "sip",
				SipEndpoint: utils.String("sip:test@test.com"),
				SipUsername: utils.String("test"),
				SipPassword: utils.String("test2"),
				Timeout:     utils.Int(30),
				Record:      utils.Bool(true),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(connectCallTo, connectCallToJSON)
	})

	Describe("Set variables with basic config", func() {
		setVariablesJSON, _ := ioutil.ReadFile("testdata/setVariables.json")

		setVariables := widgets.SetVariables{
			Name: "SetVariables",
			Properties: widgets.SetVariablesProperties{
				Variables: &[]widgets.SetVariablesVariable{
					{
						Key:   "key",
						Value: "value",
					},
				},
			},
		}

		assertJSONMatches(setVariables, setVariablesJSON)
	})

	Describe("Set variables with complete config", func() {
		setVariablesJSON, _ := ioutil.ReadFile("testdata/setVariablesComplete.json")

		setVariables := widgets.SetVariables{
			Name: "SetVariables",
			NextTransitions: widgets.SetVariablesNextTransitions{
				Next: utils.String("next"),
			},
			Properties: widgets.SetVariablesProperties{
				Variables: &[]widgets.SetVariablesVariable{
					{
						Key:   "key",
						Value: "value",
					},
				},
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(setVariables, setVariablesJSON)
	})

	Describe("Split based on with basic config", func() {
		splitBasedOnJSON, _ := ioutil.ReadFile("testdata/splitBasedOn.json")

		splitBasedOn := widgets.SplitBasedOn{
			Name: "SplitBasedOn",
			NextTransitions: widgets.SplitBasedOnNextTransitions{
				Matches: &[]transition.Conditional{
					{
						Next: "match",
						Conditions: &[]flow.Condition{
							{
								FriendlyName: "If value equal_to test",
								Arguments:    []string{"{{contact.channel.address}}"},
								Type:         "equal_to",
								Value:        "test",
							},
						},
					},
				},
			},
			Properties: widgets.SplitBasedOnProperties{
				Input: "{{contact.channel.address}}",
			},
		}

		assertJSONMatches(splitBasedOn, splitBasedOnJSON)
	})

	Describe("Split based on with complete config", func() {
		splitBasedOnJSON, _ := ioutil.ReadFile("testdata/splitBasedOnComplete.json")

		splitBasedOn := widgets.SplitBasedOn{
			Name: "SplitBasedOn",
			NextTransitions: widgets.SplitBasedOnNextTransitions{
				NoMatch: utils.String("noMatch"),
				Matches: &[]transition.Conditional{
					{
						Next: "match",
						Conditions: &[]flow.Condition{
							{
								FriendlyName: "If value equal_to test",
								Arguments:    []string{"{{contact.channel.address}}"},
								Type:         "equal_to",
								Value:        "test",
							},
						},
					},
					{
						Next: "match2",
						Conditions: &[]flow.Condition{
							{
								FriendlyName: "If value not_equal_to test 2",
								Arguments:    []string{"{{contact.channel.address}}"},
								Type:         "not_equal_to",
								Value:        "test 2",
							},
						},
					},
				},
			},
			Properties: widgets.SplitBasedOnProperties{
				Input: "{{contact.channel.address}}",
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(splitBasedOn, splitBasedOnJSON)
	})

	Describe("Gather input on call with say message config", func() {
		gatherInputOnCallJSON, _ := ioutil.ReadFile("testdata/gatherInputOnCallSayMessage.json")

		gatherInputOnCall := widgets.GatherInputOnCall{
			Name: "GatherInputOnCall",
			Properties: widgets.GatherInputOnCallProperties{
				Say: utils.String("Hello World"),
			},
		}

		assertJSONMatches(gatherInputOnCall, gatherInputOnCallJSON)
	})

	Describe("Gather input on call with play message config", func() {
		gatherInputOnCallJSON, _ := ioutil.ReadFile("testdata/gatherInputOnCallPlayMessage.json")

		gatherInputOnCall := widgets.GatherInputOnCall{
			Name: "GatherInputOnCall",
			Properties: widgets.GatherInputOnCallProperties{
				Play: utils.String("http://localhost.com"),
			},
		}

		assertJSONMatches(gatherInputOnCall, gatherInputOnCallJSON)
	})

	Describe("Gather input on call with say message complete config", func() {
		gatherInputOnCallJSON, _ := ioutil.ReadFile("testdata/gatherInputOnCallSayMessageComplete.json")

		gatherInputOnCall := widgets.GatherInputOnCall{
			Name: "GatherInputOnCall",
			NextTransitions: widgets.GatherInputOnCallNextTransitions{
				Keypress: utils.String("keypress"),
				Speech:   utils.String("speech"),
				Timeout:  utils.String("timeout"),
			},
			Properties: widgets.GatherInputOnCallProperties{
				Say:             utils.String("Hello World"),
				Voice:           utils.String("alice"),
				Hints:           utils.String("test,test2"),
				FinishOnKey:     utils.String("1"),
				Language:        utils.String("en-US"),
				StopGather:      utils.Bool(true),
				SpeechModel:     utils.String("phone_call"),
				ProfanityFilter: utils.String("true"),
				Timeout:         utils.Int(5),
				NumberOfDigits:  utils.Int(3),
				Loop:            utils.Int(1),
				SpeechTimeout:   utils.String("auto"),
				GatherLanguage:  utils.String("en-US"),
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(gatherInputOnCall, gatherInputOnCallJSON)
	})

	Describe("Capture payments with basic config", func() {
		capturePaymentsJSON, _ := ioutil.ReadFile("testdata/capturePayments.json")

		capturePayments := widgets.CapturePayments{
			Name: "CapturePayments",
			Properties: widgets.CapturePaymentsProperties{
				PaymentTokenType: utils.String("reusable"),
			},
		}

		assertJSONMatches(capturePayments, capturePaymentsJSON)
	})

	Describe("Capture payments with complete config", func() {
		capturePaymentsJSON, _ := ioutil.ReadFile("testdata/capturePaymentsComplete.json")

		capturePayments := widgets.CapturePayments{
			Name: "CapturePayments",
			NextTransitions: widgets.CapturePaymentsNextTransitions{
				Success:           utils.String("success"),
				MaxFailedAttempts: utils.String("maxFailedAttempts"),
				ProviderError:     utils.String("providerError"),
				PayInterrupted:    utils.String("payInterrupted"),
				Hangup:            utils.String("hangup"),
				ValidationError:   utils.String("validationError"),
			},
			Properties: widgets.CapturePaymentsProperties{
				PaymentTokenType:    utils.String("reusable"),
				MinPostalCodeLength: utils.Int(3),
				PaymentConnector:    utils.String("stripe"),
				PaymentAmount:       utils.String("10.99"),
				Description:         utils.String("Pay Bill"),
				Language:            utils.String("en-GB"),
				Timeout:             utils.Int(5),
				SecurityCode:        utils.Bool(true),
				MaxAttempts:         utils.Int(2),
				Currency:            utils.String("usd"),
				PostalCode:          utils.String("false"),
				PaymentMethod:       utils.String("ACH_DEBIT"),
				ValidCardTypes:      &[]string{"visa", "amex"},
				Parameters: &[]widgets.CapturePaymentsParameter{
					{
						Key:   "key",
						Value: "value",
					},
				},
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		assertJSONMatches(capturePayments, capturePaymentsJSON)
	})

})

type Widget interface {
	Validate() error
	ToState() (*flow.State, error)
}

func assertJSONMatches(widget Widget, goldenData []byte) {
	It("Then the no validation errors should be return", func() {
		Expect(widget.Validate()).To(BeNil())
	})

	state, stateErr := widget.ToState()

	It("Then the no state errors should be return", func() {
		Expect(stateErr).To(BeNil())
	})

	json, jsonErr := state.ToString()

	It("Then the no json errors should be return", func() {
		Expect(jsonErr).To(BeNil())
	})

	It("Then the json should match", func() {
		Expect(*json).To(MatchJSON(goldenData))
	})
}
