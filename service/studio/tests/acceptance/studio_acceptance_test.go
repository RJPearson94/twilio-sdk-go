package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/revisions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/test_users"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var flowDefinition = fmt.Sprintf(`
{
	"description": "A New Flow",
	"states": [
	  {
		"name": "Trigger",
		"type": "trigger",
		"transitions": [
		  {
			"next": "delay",
			"event": "incomingMessage"
		  },
		  {
			"next": "delay",
			"event": "incomingCall"
		  },
		  {
			"next": "delay",
			"event": "incomingRequest"
		  }
		],
		"properties": {
		  "offset": {
			"x": 0,
			"y": 0
		  }
		}
	  },
	  {
		"name": "delay",
		"type": "add-twiml-redirect",
		"transitions": [
		  {
			"event": "return"
		  },
		  {
			"event": "timeout"
		  },
		  {
			"event": "fail"
		  }
		],
		"properties": {
		  "offset": {
			"x": 150,
			"y": 160
		  },
		  "method": "POST",
		  "url": "%s",
		  "timeout": "14400"
		}
	  }
	],
	"initial_state": "Trigger",
	"flags": {
	  "allow_concurrent_calls": true
	}
}`, os.Getenv("TWILIO_DELAY_TWIML_URL"))

var _ = Describe("Studio Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	studioSession := twilio.NewWithCredentials(creds).Studio.V2

	Describe("Given the Studio Flow", func() {
		It("Then the flow is created, fetched, updated and deleted", func() {
			flowsClient := studioSession.Flows

			createResp, createErr := flowsClient.Create(&flows.CreateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "draft",
				Definition:   flowDefinition,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := flowsClient.Page(&flows.FlowsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Flows)).Should(BeNumerically(">=", 1))

			paginator := flowsClient.NewFlowsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Flows)).Should(BeNumerically(">=", 1))

			flowClient := studioSession.Flow(createResp.Sid)

			fetchResp, fetchErr := flowClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := flowClient.Update(&flow.UpdateFlowInput{
				Status: "published",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := flowClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Studio Flow Validation", func() {
		It("Then a valid flow is validated", func() {
			flowValidationClient := studioSession.FlowValidation

			createResp, createErr := flowValidationClient.Validate(&flow_validation.ValidateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "draft",
				Definition:   flowDefinition,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Valid).To(Equal(true))
		})

		It("Then an invalid flow is validated", func() {
			flowValidationClient := studioSession.FlowValidation

			createResp, createErr := flowValidationClient.Validate(&flow_validation.ValidateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "draft",
				Definition: `{
	"description": "An invalid Flow",
	"states": [
		{
		"name": "Trigger",
		"type": "trigger",
		"transitions": [
			{
			"event": "incomingMessage"
			},
			{
			"next": "invalidTransition",
			"event": "incomingCall"
			},
			{
			"event": "incomingRequest"
			}
		],
		"properties": {
			"offset": {
			"x": 0,
			"y": 0
			}
		}
		}
	],
	"initial_state": "Trigger",
	"flags": {
		"allow_concurrent_calls": true
	}
}`,
			})
			Expect(createErr).ToNot(BeNil())
			Expect(createResp).To(BeNil())

			twilioErr := createErr.(*utils.TwilioError)
			Expect(twilioErr.Status).To(Equal(400))
			Expect(twilioErr.Details).ToNot(BeNil())
		})
	})

	Describe("Given the Flow Execution", func() {

		var flowSid string

		BeforeEach(func() {
			resp, err := studioSession.Flows.Create(&flows.CreateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "published",
				Definition:   flowDefinition,
			})
			if err != nil {
				Fail(fmt.Sprintf("Flow failed to create. Error %s", err.Error()))
			}
			flowSid = resp.Sid
		})

		AfterEach(func() {
			if err := studioSession.Flow(flowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Flow failed to delete. Error %s", err.Error()))
			}
		})

		It("Then the execution is created, fetched, updated and deleted", func() {
			executionsClient := studioSession.Flow(flowSid).Executions

			createResp, createErr := executionsClient.Create(&executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+18001234568",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := executionsClient.Page(&executions.ExecutionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Executions)).Should(BeNumerically(">=", 1))

			paginator := executionsClient.NewExecutionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Executions)).Should(BeNumerically(">=", 1))

			executionClient := studioSession.Flow(flowSid).Execution(createResp.Sid)

			fetchResp, fetchErr := executionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := executionClient.Update(&execution.UpdateExecutionInput{
				Status: "ended",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := executionClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the Flow Execution Context", func() {

		var flowSid string
		var executionSid string

		BeforeEach(func() {
			resp, err := studioSession.Flows.Create(&flows.CreateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "published",
				Definition:   flowDefinition,
			})
			if err != nil {
				Fail(fmt.Sprintf("Flow failed to create. Error %s", err.Error()))
			}
			flowSid = resp.Sid

			executionResp, executionErr := studioSession.Flow(flowSid).Executions.Create(&executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+18001234568",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			})
			if executionErr != nil {
				Fail(fmt.Sprintf("Execution failed to create. Error %s", executionErr.Error()))
			}
			executionSid = executionResp.Sid
		})

		AfterEach(func() {
			if err := studioSession.Flow(flowSid).Execution(executionSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Execution failed to delete. Error %s", err.Error()))
			}

			if err := studioSession.Flow(flowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Flow failed to delete. Error %s", err.Error()))
			}
		})

		It("Then the execution context is fetched", func() {
			contextClient := studioSession.Flow(flowSid).Execution(executionSid).Context()

			fetchResp, fetchErr := contextClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the Flow Revision", func() {

		var flowSid string
		var flowRevision int

		BeforeEach(func() {
			resp, err := studioSession.Flows.Create(&flows.CreateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "published",
				Definition:   flowDefinition,
			})
			if err != nil {
				Fail(fmt.Sprintf("Flow failed to create. Error %s", err.Error()))
			}
			flowSid = resp.Sid
			flowRevision = resp.Revision
		})

		AfterEach(func() {
			if err := studioSession.Flow(flowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Flow failed to delete. Error %s", err.Error()))
			}
		})

		It("Then the revision is fetched", func() {
			revisionsClient := studioSession.Flow(flowSid).Revisions

			pageResp, pageErr := revisionsClient.Page(&revisions.RevisionsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Revisions)).Should(BeNumerically(">=", 1))

			paginator := revisionsClient.NewRevisionsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Revisions)).Should(BeNumerically(">=", 1))

			revisionClient := studioSession.Flow(flowSid).Revision(flowRevision)

			fetchResp, fetchErr := revisionClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the Flow Test Users", func() {

		var flowSid string

		BeforeEach(func() {
			resp, err := studioSession.Flows.Create(&flows.CreateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "published",
				Definition:   flowDefinition,
			})
			if err != nil {
				Fail(fmt.Sprintf("Flow failed to create. Error %s", err.Error()))
			}
			flowSid = resp.Sid
		})

		AfterEach(func() {
			if err := studioSession.Flow(flowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Flow failed to delete. Error %s", err.Error()))
			}
		})

		It("Then the test users are fetched and updated", func() {
			testUsersClient := studioSession.Flow(flowSid).TestUsers()

			fetchResp, fetchErr := testUsersClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := testUsersClient.Update(&test_users.UpdateTestUsersInput{
				TestUsers: []string{"+14155551212"},
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	// TODO Add Step and Step Context tests - requires Get Page Support
})
