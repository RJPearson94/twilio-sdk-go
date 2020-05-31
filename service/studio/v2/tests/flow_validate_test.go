package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var _ = Describe("FlowValidate", func() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	studioSession := v2.NewWithCredentials(creds)

	Describe("Given the Flow Validation Service", func() {
		flowService := studioSession.FlowValidation

		Describe("When the Flow is successfully validated", func() {
			flowDefinition, _ := ioutil.ReadFile("testdata/flowDefinition.json")

			validateInput := &v2.ValidateFlowInput{
				FriendlyName: "Test 2",
				Status:       "draft",
				Definition:   string(flowDefinition),
			}

			httpmock.RegisterResponder("POST", "https://studio.twilio.com/v2/Flows/Validate",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/validateFlowResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := flowService.Validate(validateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the validate flow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Valid).To(Equal(true))
			})
		})
	})
})
