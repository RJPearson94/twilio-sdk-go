package unit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channels"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/tasks"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker/channels"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker/reservation"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker/reservations"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflow"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflows"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Taskrouter V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	taskrouterSession := taskrouter.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(taskrouterSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given the workspaces client", func() {
		workspacesClient := taskrouterSession.Workspaces

		Describe("When the workspace is successfully created", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName:     "Test 2",
				MultiTaskEnabled: utils.Bool(false),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workspacesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the workspace request does not contain a friendly name", func() {
			createInput := &workspaces.CreateWorkspaceInput{}

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workspaces api returns a 500 response", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of workspaces are successfully retrieved", func() {
			pageOptions := &workspaces.WorkspacesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspacesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workspacesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the workspaces page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("workspaces"))

				workspaces := resp.Workspaces
				Expect(workspaces).ToNot(BeNil())
				Expect(len(workspaces)).To(Equal(1))

				Expect(workspaces[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workspaces[0].DefaultActivityName).To(Equal("Offline"))
				Expect(workspaces[0].DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workspaces[0].DateUpdated).To(BeNil())
				Expect(workspaces[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(workspaces[0].EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
				Expect(workspaces[0].EventsFilter).To(BeNil())
				Expect(workspaces[0].FriendlyName).To(Equal("NewWorkspace"))
				Expect(workspaces[0].MultiTaskEnabled).To(Equal(false))
				Expect(workspaces[0].PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(workspaces[0].TimeoutActivityName).To(Equal("Offline"))
				Expect(workspaces[0].TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(workspaces[0].Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workspaces[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of workspaces api returns a 500 response", func() {
			pageOptions := &workspaces.WorkspacesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workspacesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the workspaces page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated workspaces are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspacesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspacesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := workspacesClient.NewWorkspacesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated workspaces current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated workspaces results should be returned", func() {
				Expect(len(paginator.Workspaces)).To(Equal(3))
			})
		})

		Describe("When the workspaces api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspacesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := workspacesClient.NewWorkspacesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated workspaces current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a workspace sid", func() {
		workspaceClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the workspace is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workspaceClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workspace response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WS71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspace is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &workspace.UpdateWorkspaceInput{}

			resp, err := workspaceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.EventCallbackURL).To(Equal(utils.String("https://ngrok.com")))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update workspace api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &workspace.UpdateWorkspaceInput{
				FriendlyName: utils.String("Test Workspace"),
			}

			resp, err := taskrouterSession.Workspace("WS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workspace is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workspaceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workspace response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the task queues client", func() {
		taskQueuesClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueues

		Describe("When the Task Queue is successfully created", func() {
			createInput := &task_queues.CreateTaskQueueInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := taskQueuesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.AssignmentActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ReservationActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task queue request does not contain a friendly name", func() {
			createInput := &task_queues.CreateTaskQueueInput{}

			resp, err := taskQueuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task queue api returns a 500 response", func() {
			createInput := &task_queues.CreateTaskQueueInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := taskQueuesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of task queues are successfully retrieved", func() {
			pageOptions := &task_queues.TaskQueuesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueuesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskQueuesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the task queues page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("task_queues"))

				taskQueues := resp.TaskQueues
				Expect(taskQueues).ToNot(BeNil())
				Expect(len(taskQueues)).To(Equal(1))

				Expect(taskQueues[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(taskQueues[0].AssignmentActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(taskQueues[0].AssignmentActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(taskQueues[0].DateUpdated).To(BeNil())
				Expect(taskQueues[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(taskQueues[0].FriendlyName).To(Equal("English"))
				Expect(taskQueues[0].MaxReservedWorkers).To(Equal(1))
				Expect(taskQueues[0].ReservationActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(taskQueues[0].ReservationActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(taskQueues[0].Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(taskQueues[0].TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(taskQueues[0].TaskOrder).To(Equal("FIFO"))
				Expect(taskQueues[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(taskQueues[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of task queues api returns a 500 response", func() {
			pageOptions := &task_queues.TaskQueuesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := taskQueuesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the task queues page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated task queues are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueuesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueuesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := taskQueuesClient.NewTaskQueuesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated task queues current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated task queues results should be returned", func() {
				Expect(len(paginator.TaskQueues)).To(Equal(3))
			})
		})

		Describe("When the task queues api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueuesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := taskQueuesClient.NewTaskQueuesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated task queues current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a task queue sid", func() {
		taskQueueClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task queue is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskQueueClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.AssignmentActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ReservationActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get task queue api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task queue is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task_queue.UpdateTaskQueueInput{}

			resp, err := taskQueueClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.AssignmentActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal(utils.String("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.ReservationActivitySid).To(Equal(utils.String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal(utils.String("languages HAS \"english\"")))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update task queue api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task_queue.UpdateTaskQueueInput{
				FriendlyName: utils.String("Test Queue"),
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task queue is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskQueueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete task queue api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the workflows client", func() {
		workflowsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflows

		Describe("When the workflow is successfully created", func() {
			configuration := "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"},\"filters\":[{\"expression\":\"type=='sales'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"}]},{\"expression\":\"type=='marketing'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1\"}]},{\"expression\":\"type=='support'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2\"}]}]}}"
			createInput := &workflows.CreateWorkflowInput{
				FriendlyName:  "Test 2",
				Configuration: configuration,
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workflowsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workflow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))
				Expect(resp.Configuration).To(Equal(configuration))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the workflow request does not contain a friendly name", func() {
			configuration := "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"},\"filters\":[{\"expression\":\"type=='sales'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"}]},{\"expression\":\"type=='marketing'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1\"}]},{\"expression\":\"type=='support'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2\"}]}]}}"
			createInput := &workflows.CreateWorkflowInput{
				Configuration: configuration,
			}

			resp, err := workflowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow request does not contain a configuration", func() {
			createInput := &workflows.CreateWorkflowInput{
				FriendlyName: "Test 2",
			}

			resp, err := workflowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow api returns a 500 response", func() {
			configuration := "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"},\"filters\":[{\"expression\":\"type=='sales'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"}]},{\"expression\":\"type=='marketing'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1\"}]},{\"expression\":\"type=='support'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2\"}]}]}}"
			createInput := &workflows.CreateWorkflowInput{
				FriendlyName:  "Test 2",
				Configuration: configuration,
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workflowsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of workflows are successfully retrieved", func() {
			pageOptions := &workflows.WorkflowsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workflowsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the workflows page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("workflows"))

				workflows := resp.Workflows
				Expect(workflows).ToNot(BeNil())
				Expect(len(workflows)).To(Equal(1))

				Expect(workflows[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workflows[0].AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))

				configuration := "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"},\"filters\":[{\"expression\":\"type=='sales'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"}]},{\"expression\":\"type=='marketing'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1\"}]},{\"expression\":\"type=='support'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2\"}]}]}}"
				Expect(workflows[0].Configuration).To(Equal(configuration))

				Expect(workflows[0].DateUpdated).To(BeNil())
				Expect(workflows[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(workflows[0].DocumentContentType).To(Equal("application/json"))
				Expect(workflows[0].FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
				Expect(workflows[0].FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(workflows[0].Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workflows[0].TaskReservationTimeout).To(Equal(120))
				Expect(workflows[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workflows[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of workflows api returns a 500 response", func() {
			pageOptions := &workflows.WorkflowsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workflowsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the workflows page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated workflows are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := workflowsClient.NewWorkflowsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated workflows current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated workflows results should be returned", func() {
				Expect(len(paginator.Workflows)).To(Equal(3))
			})
		})

		Describe("When the workflows api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := workflowsClient.NewWorkflowsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated workflows current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a workflow sid", func() {
		workflowClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the workflow is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workflowsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workflowClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get workflow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))

				configuration := "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"},\"filters\":[{\"expression\":\"type=='sales'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"}]},{\"expression\":\"type=='marketing'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1\"}]},{\"expression\":\"type=='support'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2\"}]}]}}"
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workflow response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WF71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkflowsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &workflow.UpdateWorkflowInput{}

			resp, err := workflowClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workflow response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentCallbackURL).To(Equal(utils.String("https://example.com/")))

				configuration := "{\"task_routing\":{\"default_filter\":{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"},\"filters\":[{\"expression\":\"type=='sales'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\"}]},{\"expression\":\"type=='marketing'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1\"}]},{\"expression\":\"type=='support'\",\"targets\":[{\"queue\":\"WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX2\"}]}]}}"
				Expect(resp.Configuration).To(Equal(configuration))

				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DocumentContentType).To(Equal("application/json"))
				Expect(resp.FallbackAssignmentCallbackURL).To(Equal(utils.String("https://example2.com/")))
				Expect(resp.FriendlyName).To(Equal("Sales, Marketing, Support Workflow"))
				Expect(resp.Sid).To(Equal("WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskReservationTimeout).To(Equal(120))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update workflow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &workflow.UpdateWorkflowInput{
				FriendlyName: utils.String("Test Queue"),
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WF71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update workflow response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workflow is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workflowClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workflow API returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WF71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workflow("WF71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the workers client", func() {
		workersClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Workers

		Describe("When the Worker is successfully created", func() {
			createInput := &workers.CreateWorkerInput{
				FriendlyName: "NewWorker",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workersClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workers response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewWorker"))
				Expect(resp.ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ActivityName).To(Equal("Offline"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Available).To(Equal(false))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateStatusChanged.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the worker request does not contain a friendly name", func() {
			createInput := &workers.CreateWorkerInput{}

			resp, err := workersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker api returns a 500 response", func() {
			createInput := &workers.CreateWorkerInput{
				FriendlyName: "NewWorker",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workersClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of workers are successfully retrieved", func() {
			pageOptions := &workers.WorkersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workersClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the workers page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("workers"))

				workers := resp.Workers
				Expect(workers).ToNot(BeNil())
				Expect(len(workers)).To(Equal(1))

				Expect(workers[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workers[0].FriendlyName).To(Equal("NewWorker"))
				Expect(workers[0].ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workers[0].ActivityName).To(Equal("Offline"))
				Expect(workers[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(workers[0].Attributes).To(Equal("{}"))
				Expect(workers[0].Available).To(Equal(false))
				Expect(workers[0].DateUpdated).To(BeNil())
				Expect(workers[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(workers[0].DateStatusChanged.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(workers[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of workers api returns a 500 response", func() {
			pageOptions := &workers.WorkersPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workersClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the workers page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated workers are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := workersClient.NewWorkersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated workers current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated workers results should be returned", func() {
				Expect(len(paginator.Workers)).To(Equal(3))
			})
		})

		Describe("When the workers api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := workersClient.NewWorkersPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated workers current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a worker sid", func() {
		workerClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the worker is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workerClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get worker response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewWorker"))
				Expect(resp.ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ActivityName).To(Equal("Offline"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Available).To(Equal(false))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateStatusChanged.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get worker api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WK71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkersResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &worker.UpdateWorkerInput{}

			resp, err := workerClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update worker response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewWorker"))
				Expect(resp.ActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ActivityName).To(Equal("Offline"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Attributes).To(Equal("{}"))
				Expect(resp.Available).To(Equal(false))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateStatusChanged.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update worker response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &worker.UpdateWorkerInput{
				FriendlyName: utils.String("Test Worker"),
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WK71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update worker response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workerClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete worker response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WK71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WK71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the worker channels client", func() {
		channelsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channels

		Describe("When the page of worker channels are successfully retrieved", func() {
			pageOptions := &channels.ChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerChannelsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the worker channels page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("channels"))

				channels := resp.Channels
				Expect(channels).ToNot(BeNil())
				Expect(len(channels)).To(Equal(1))

				Expect(channels[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].AssignedTasks).To(Equal(0))
				Expect(channels[0].Available).To(Equal(true))
				Expect(channels[0].AvailableCapacityPercentage).To(Equal(100))
				Expect(channels[0].ConfiguredCapacity).To(Equal(10))
				Expect(channels[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(channels[0].DateUpdated).To(BeNil())
				Expect(channels[0].Sid).To(Equal("WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].TaskChannelSid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].TaskChannelUniqueName).To(Equal("default"))
				Expect(channels[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].WorkerSid).To(Equal("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(channels[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))

			})
		})

		Describe("When the page of worker channels api returns a 500 response", func() {
			pageOptions := &channels.ChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := channelsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the worker channels page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated worker channels are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerChannelsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := channelsClient.NewChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated worker channels current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated worker channels results should be returned", func() {
				Expect(len(paginator.Channels)).To(Equal(3))
			})
		})

		Describe("When the worker channels api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := channelsClient.NewChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated worker channels current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a worker channel sid", func() {
		channelClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the worker is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerChannelsResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := channelClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get worker channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignedTasks).To(Equal(0))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.AvailableCapacityPercentage).To(Equal(100))
				Expect(resp.ConfiguredCapacity).To(Equal(10))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.Sid).To(Equal("WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskChannelSid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskChannelUniqueName).To(Equal("default"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkerSid).To(Equal("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get worker channel api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("WC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get worker channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker channel is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkerChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &channel.UpdateChannelInput{}

			resp, err := channelClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update worker channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignedTasks).To(Equal(0))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.AvailableCapacityPercentage).To(Equal(100))
				Expect(resp.ConfiguredCapacity).To(Equal(10))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.Sid).To(Equal("WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskChannelSid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskChannelUniqueName).To(Equal("default"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkerSid).To(Equal("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update worker channel response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/WC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &channel.UpdateChannelInput{
				Capacity: utils.Int(5),
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Channel("WC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update worker channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the worker reservations client", func() {
		reservationsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Reservations

		Describe("When the page of worker reservations are successfully retrieved", func() {
			pageOptions := &reservations.ReservationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerReservationsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := reservationsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the worker reservations page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("reservations"))

				reservations := resp.Reservations
				Expect(reservations).ToNot(BeNil())
				Expect(len(reservations)).To(Equal(1))

				Expect(reservations[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(reservations[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(reservations[0].DateUpdated).To(BeNil())
				Expect(reservations[0].ReservationStatus).To(Equal("rejected"))
				Expect(reservations[0].Sid).To(Equal("WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(reservations[0].TaskSid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(reservations[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(reservations[0].WorkerName).To(Equal("Test"))
				Expect(reservations[0].WorkerSid).To(Equal("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(reservations[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of worker reservations api returns a 500 response", func() {
			pageOptions := &reservations.ReservationsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := reservationsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the worker reservations page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated worker reservations are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerReservationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerReservationsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := reservationsClient.NewReservationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated worker reservations current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated worker reservations results should be returned", func() {
				Expect(len(paginator.Reservations)).To(Equal(3))
			})
		})

		Describe("When the worker reservations api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerReservationsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := reservationsClient.NewReservationsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated worker reservations current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a worker reservation sid", func() {
		reservationClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Reservation("WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the worker is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workerReservationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := reservationClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get worker reservation response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.ReservationStatus).To(Equal("rejected"))
				Expect(resp.Sid).To(Equal("WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkerName).To(Equal("Test"))
				Expect(resp.WorkerSid).To(Equal("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get worker reservation api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Reservation("WR71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get worker reservation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the worker reservation is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkerReservationResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &reservation.UpdateReservationInput{
				ReservationStatus: "accepted",
			}

			resp, err := reservationClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update worker reservation response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.ReservationStatus).To(Equal("accepted"))
				Expect(resp.Sid).To(Equal("WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TaskSid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkerName).To(Equal("Test"))
				Expect(resp.WorkerSid).To(Equal("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update reservation request does not contain a reservation status", func() {
			updateInput := &reservation.UpdateReservationInput{}

			resp, err := reservationClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update reservation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update worker reservation response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Reservations/WR71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &reservation.UpdateReservationInput{
				ReservationStatus: "accepted",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Worker("WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Reservation("WR71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update worker reservation response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given the activities client", func() {
		activitiesClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activities

		Describe("When the activity is successfully created", func() {
			createInput := &activities.CreateActivityInput{
				FriendlyName: "NewAvailableActivity",
				Available:    utils.Bool(true),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := activitiesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create activity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewActivity"))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the activity request does not contain a friendly name", func() {
			createInput := &activities.CreateActivityInput{}

			resp, err := activitiesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the activities api returns a 500 response", func() {
			createInput := &activities.CreateActivityInput{
				FriendlyName: "NewAvailableActivity",
				Available:    utils.Bool(true),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := activitiesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of activities are successfully retrieved", func() {
			pageOptions := &activities.ActivitiesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activitiesPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := activitiesClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the activities page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("activities"))

				activities := resp.Activities
				Expect(activities).ToNot(BeNil())
				Expect(len(activities)).To(Equal(1))

				Expect(activities[0].Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(activities[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(activities[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(activities[0].FriendlyName).To(Equal("NewActivity"))
				Expect(activities[0].Available).To(Equal(true))
				Expect(activities[0].DateUpdated).To(BeNil())
				Expect(activities[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(activities[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of activities api returns a 500 response", func() {
			pageOptions := &activities.ActivitiesPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := activitiesClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the activities page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated activities are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activitiesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activitiesPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := activitiesClient.NewActivitiesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated activities current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated activities results should be returned", func() {
				Expect(len(paginator.Activities)).To(Equal(3))
			})
		})

		Describe("When the activities api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activitiesPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := activitiesClient.NewActivitiesPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated activities current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a activity sid", func() {
		activityClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the activity is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/activityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := activityClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get activity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("NewActivity"))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get activity response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WA71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the activity is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedActivityResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &activity.UpdateActivityInput{}

			resp, err := activityClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update activity response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Activity"))
				Expect(resp.Available).To(Equal(true))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update activity response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &activity.UpdateActivityInput{
				FriendlyName: utils.String("Test Activity"),
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WA71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update activity response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the activity is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := activityClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete activity api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WA71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Activity("WA71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the task channels client", func() {
		taskChannelsClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannels

		Describe("When the task channel is successfully created", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				FriendlyName: "Test 2",
				UniqueName:   "Unique Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := taskChannelsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task channel request does not contain a friendly name", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				UniqueName: "Unique Test 2",
			}

			resp, err := taskChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel request does not contain a unique name", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				FriendlyName: "Test 2",
			}

			resp, err := taskChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel api returns a 500 response", func() {
			createInput := &task_channels.CreateTaskChannelInput{
				FriendlyName: "Test 2",
				UniqueName:   "Unique Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := taskChannelsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of task channels are successfully retrieved", func() {
			pageOptions := &task_channels.TaskChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskChannelsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the task channels page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("channels"))

				taskChannels := resp.TaskChannels
				Expect(taskChannels).ToNot(BeNil())
				Expect(len(taskChannels)).To(Equal(1))

				Expect(taskChannels[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(taskChannels[0].WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(taskChannels[0].Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(taskChannels[0].FriendlyName).To(Equal("Test 2"))
				Expect(taskChannels[0].UniqueName).To(Equal("Unique Test 2"))
				Expect(taskChannels[0].ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(taskChannels[0].DateUpdated).To(BeNil())
				Expect(taskChannels[0].DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(taskChannels[0].URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of task channels api returns a 500 response", func() {
			pageOptions := &task_channels.TaskChannelsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := taskChannelsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the task channels page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated task channels are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := taskChannelsClient.NewTaskChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated task channels current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated task channels results should be returned", func() {
				Expect(len(paginator.TaskChannels)).To(Equal(3))
			})
		})

		Describe("When the task channels api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := taskChannelsClient.NewTaskChannelsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated task channels current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a task channel sid", func() {
		taskChannelClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task channel is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskChannelClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get task channel response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TC71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskChannelResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task_channel.UpdateTaskChannelInput{}

			resp, err := taskChannelClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update task channel response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FriendlyName).To(Equal("Test 2"))
				Expect(resp.UniqueName).To(Equal("Unique Test 2"))
				Expect(resp.ChannelOptimizedRouting).To(Equal(utils.Bool(true)))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update task channel response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task_channel.UpdateTaskChannelInput{
				FriendlyName: utils.String("Test Task Channel"),
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TC71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task channel response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task channel is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskChannelClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete task channel response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TC71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskChannel("TC71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the tasks client", func() {
		tasksClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Tasks

		Describe("When the task is successfully created", func() {
			createInput := &tasks.CreateTaskInput{
				TaskChannel: utils.String("default"),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := tasksClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Age).To(Equal(25200))
				Expect(resp.AssignmentStatus).To(Equal("pending"))

				attributes := make(map[string]interface{})
				attributes["type"] = "support"
				Expect(resp.Attributes).To(Equal(attributes))

				Expect(resp.Priority).To(Equal(utils.Int(1)))
				Expect(resp.Reason).To(Equal(utils.String("Test Reason")))
				Expect(resp.TaskQueueSid).To(Equal(utils.String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelSid).To(Equal(utils.String("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelUniqueName).To(Equal(utils.String("unique")))
				Expect(resp.WorkflowSid).To(Equal(utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.WorkflowFriendlyName).To(Equal(utils.String("Example Workflow")))
				Expect(resp.TaskQueueEnteredDate).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the task api returns a 500 response", func() {
			createInput := &tasks.CreateTaskInput{
				TaskChannel: utils.String("default"),
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := tasksClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of tasks are successfully retrieved", func() {
			pageOptions := &tasks.TasksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/tasksPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := tasksClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the tasks page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("tasks"))

				tasks := resp.Tasks
				Expect(tasks).ToNot(BeNil())
				Expect(len(tasks)).To(Equal(1))

			})
		})

		Describe("When the page of tasks api returns a 500 response", func() {
			pageOptions := &tasks.TasksPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := tasksClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the tasks page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated tasks are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/tasksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/tasksPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := tasksClient.NewTasksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated tasks current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated tasks results should be returned", func() {
				Expect(len(paginator.Tasks)).To(Equal(3))
			})
		})

		Describe("When the tasks api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/tasksPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := tasksClient.NewTasksPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated tasks current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a task sid", func() {
		taskClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the task is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Age).To(Equal(25200))
				Expect(resp.AssignmentStatus).To(Equal("pending"))

				attributes := make(map[string]interface{})
				attributes["type"] = "support"
				Expect(resp.Attributes).To(Equal(attributes))

				Expect(resp.Priority).To(Equal(utils.Int(1)))
				Expect(resp.Reason).To(Equal(utils.String("Test Reason")))
				Expect(resp.TaskQueueSid).To(Equal(utils.String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelSid).To(Equal(utils.String("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelUniqueName).To(Equal(utils.String("unique")))
				Expect(resp.WorkflowSid).To(Equal(utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.WorkflowFriendlyName).To(Equal(utils.String("Example Workflow")))
				Expect(resp.TaskQueueEnteredDate).To(BeNil())
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get task response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WT71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task.UpdateTaskInput{}

			resp, err := taskClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update task response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Sid).To(Equal("WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Age).To(Equal(25200))
				Expect(resp.AssignmentStatus).To(Equal("pending"))

				attributes := make(map[string]interface{})
				attributes["type"] = "support"
				Expect(resp.Attributes).To(Equal(attributes))

				Expect(resp.Priority).To(Equal(utils.Int(1)))
				Expect(resp.Reason).To(Equal(utils.String("Test Reason")))
				Expect(resp.TaskQueueSid).To(Equal(utils.String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelSid).To(Equal(utils.String("TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.TaskChannelUniqueName).To(Equal(utils.String("unique")))
				Expect(resp.WorkflowSid).To(Equal(utils.String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")))
				Expect(resp.WorkflowFriendlyName).To(Equal(utils.String("Example Workflow")))
				Expect(resp.TaskQueueEnteredDate).To(BeNil())
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-27T23:10:00Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-27T23:00:00Z"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update task response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task.UpdateTaskInput{}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WT71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete task response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/WT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Task("WT71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})
})

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectNotFoundError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(Equal(utils.Int(20404)))
	Expect(twilioErr.Message).To(Equal("The requested resource /Flows/FW71 was not found"))
	Expect(twilioErr.MoreInfo).To(Equal(utils.String("https://www.twilio.com/docs/errors/20404")))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}
