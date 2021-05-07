package tests

import (
	"io/ioutil"

	"github.com/RJPearson94/twilio-sdk-go/studio"
	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Studio Flow", func() {
	Describe("Flow with trigger widget", func() {
		flowTriggerJSON, _ := ioutil.ReadFile("testdata/flowTrigger.json")

		trigger := widgets.Trigger{
			Name: "Trigger",
			Properties: widgets.TriggerProperties{
				Offset: &properties.Offset{
					X: 0,
					Y: 0,
				},
			},
		}

		triggerState, _ := trigger.ToState()

		flow := studio.Flow{
			Description: "A New Flow",
			Flags: &studio.FlowFlags{
				AllowConcurrentCalls: true,
			},
			InitialState: triggerState.Name,
			States: []flow.State{
				*triggerState,
			},
		}

		It("Then the no validation errors should be return", func() {
			Expect(flow.Validate()).To(BeNil())
		})

		json, jsonErr := flow.ToString()

		It("Then the no json errors should be return", func() {
			Expect(jsonErr).To(BeNil())
		})

		It("Then the json should match", func() {
			Expect(*json).To(MatchJSON(flowTriggerJSON))
		})
	})

})
