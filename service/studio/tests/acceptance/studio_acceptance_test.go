package acceptance

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/test_users"
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
		log.Panicf("%s", err)
	}

	studioSession := twilio.NewWithCredentials(creds).Studio.V2

	Describe("Given the Studio Flow", func() {
		It("Then the flow is created, fetched, updated and deleted", func() {
			createResp, createErr := studioSession.Flows.Create(&flows.CreateFlowInput{
				FriendlyName: uuid.New().String(),
				Status:       "draft",
				Definition:   flowDefinition,
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			syncClient := studioSession.Flow(createResp.Sid)

			fetchResp, fetchErr := syncClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := syncClient.Update(&flow.UpdateFlowInput{
				Status: "published",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := syncClient.Delete()
			Expect(deleteErr).To(BeNil())
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
			createResp, createErr := studioSession.Flow(flowSid).Executions.Create(&executions.CreateExecutionInput{
				To:         "+18001234567",
				From:       "+18001234568",
				Parameters: utils.String("{\"name\": \"RJPearson94\"}"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

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

	Describe("Given the Flow Execution Content", func() {

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

		It("Then the execution content is fetched", func() {
			executionClient := studioSession.Flow(flowSid).Execution(executionSid).Context()

			fetchResp, fetchErr := executionClient.Fetch()
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
