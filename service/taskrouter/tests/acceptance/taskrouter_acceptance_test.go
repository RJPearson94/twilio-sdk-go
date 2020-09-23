package acceptance

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task"
	taskReservation "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task/reservation"
	taskReservations "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task/reservations"
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

var _ = Describe("Taskrouter Acceptance Tests", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		Fail(fmt.Sprintf("Failed to create credentials. Error %s", err.Error()))
	}

	taskrouterSession := twilio.NewWithCredentials(creds).TaskRouter.V1

	Describe("Given the TaskRouter Workspace clients", func() {
		It("Then the workspace is create, fetched, updated and deleted", func() {
			workspacesClient := taskrouterSession.Workspaces

			createResp, createErr := workspacesClient.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := workspacesClient.Page(&workspaces.WorkspacesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Workspaces)).Should(BeNumerically(">=", 1))

			paginator := workspacesClient.NewWorkspacesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Workspaces)).Should(BeNumerically(">=", 1))

			workspaceClient := taskrouterSession.Workspace(createResp.Sid)

			fetchResp, fetchErr := workspaceClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := workspaceClient.Update(&workspace.UpdateWorkspaceInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := workspaceClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Workspace real time statistics clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			realTimeStatisticsClient := taskrouterSession.Workspace(workspaceSid).RealTimeStatistics()

			fetchResp, fetchErr := realTimeStatisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Workspace cumulative statistics clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			cumulativeStatisticsClient := taskrouterSession.Workspace(workspaceSid).CumulativeStatistics()

			fetchResp, fetchErr := cumulativeStatisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Workspace statistics clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			statisticsClient := taskrouterSession.Workspace(workspaceSid).Statistics()

			fetchResp, fetchErr := statisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Queue clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the task queue is create, fetched, updated and deleted", func() {
			taskQueuesClient := taskrouterSession.Workspace(workspaceSid).TaskQueues

			createResp, createErr := taskQueuesClient.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := taskQueuesClient.Page(&task_queues.TaskQueuesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.TaskQueues)).Should(BeNumerically(">=", 1))

			paginator := taskQueuesClient.NewTaskQueuesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.TaskQueues)).Should(BeNumerically(">=", 1))

			taskQueueClient := taskrouterSession.Workspace(workspaceSid).TaskQueue(createResp.Sid)

			fetchResp, fetchErr := taskQueueClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := taskQueueClient.Update(&task_queue.UpdateTaskQueueInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := taskQueueClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Queue real time statistics clients", func() {

		var workspaceSid string
		var taskQueueSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid
		})

		AfterEach(func() {

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			realTimeStatisticsClient := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).RealTimeStatistics()

			fetchResp, fetchErr := realTimeStatisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Queue cumulative statistics clients", func() {

		var workspaceSid string
		var taskQueueSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			cumulativeStatisticsClient := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).CumulativeStatistics()

			fetchResp, fetchErr := cumulativeStatisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Queue statistics clients", func() {

		var workspaceSid string
		var taskQueueSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			statisticsClient := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Statistics()

			fetchResp, fetchErr := statisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Workflow clients", func() {

		var workspaceSid string
		var taskQueueSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the workflow is create, fetched, updated and deleted", func() {
			workflowsClient := taskrouterSession.Workspace(workspaceSid).Workflows

			createResp, createErr := workflowsClient.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := workflowsClient.Page(&workflows.WorkflowsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Workflows)).Should(BeNumerically(">=", 1))

			paginator := workflowsClient.NewWorkflowsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Workflows)).Should(BeNumerically(">=", 1))

			workflowClient := taskrouterSession.Workspace(workspaceSid).Workflow(createResp.Sid)

			fetchResp, fetchErr := workflowClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := workflowClient.Update(&workflow.UpdateWorkflowInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := workflowClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Workflow real time statistics clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			realTimeStatisticsClient := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).RealTimeStatistics()

			fetchResp, fetchErr := realTimeStatisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Workflow cumulative statistics clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			cumulativeStatisticsClient := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).CumulativeStatistics()

			fetchResp, fetchErr := cumulativeStatisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Workflow statistics clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the real time statistics are fetched", func() {
			statisticsClient := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Statistics()

			fetchResp, fetchErr := statisticsClient.Fetch(nil)
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Worker clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the worker is created, fetched, updated and deleted", func() {
			workersClient := taskrouterSession.Workspace(workspaceSid).Workers

			createResp, createErr := workersClient.Create(&workers.CreateWorkerInput{
				FriendlyName: uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := workersClient.Page(&workers.WorkersPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Workers)).Should(BeNumerically(">=", 1))

			paginator := workersClient.NewWorkersPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Workers)).Should(BeNumerically(">=", 1))

			workerClient := taskrouterSession.Workspace(workspaceSid).Worker(createResp.Sid)

			fetchResp, fetchErr := workerClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := workerClient.Update(&worker.UpdateWorkerInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := workerClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Worker Channel clients", func() {

		var workspaceSid string
		var workerSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			workerResp, workerErr := taskrouterSession.Workspace(workspaceSid).Workers.Create(&workers.CreateWorkerInput{
				FriendlyName: uuid.New().String(),
			})
			if workerErr != nil {
				Fail(fmt.Sprintf("Failed to create worker. Error %s", workerErr.Error()))
			}
			workerSid = workerResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete worker. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the worker channel is fetched and updated", func() {
			channelsClient := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Channels

			pageResp, pageErr := channelsClient.Page(&channels.ChannelsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Channels)).Should(BeNumerically(">=", 1))

			paginator := channelsClient.NewChannelsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Channels)).Should(BeNumerically(">=", 1))

			channelClient := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Channel(paginator.Channels[0].Sid)

			fetchResp, fetchErr := channelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := channelClient.Update(&channel.UpdateChannelInput{
				Capacity: utils.Int(5),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Worker Reservation clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string
		var activitySid string
		var unavailableActivitySid string
		var workerSid string
		var taskSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid

			activityResp, activityErr := taskrouterSession.Workspace(workspaceSid).Activities.Create(&activities.CreateActivityInput{
				FriendlyName: uuid.New().String(),
				Available:    utils.Bool(true),
			})
			if activityErr != nil {
				Fail(fmt.Sprintf("Failed to create activity. Error %s", activityErr.Error()))
			}
			activitySid = activityResp.Sid

			unavailableActivityResp, unavailableActivityErr := taskrouterSession.Workspace(workspaceSid).Activities.Create(&activities.CreateActivityInput{
				FriendlyName: uuid.New().String(),
				Available:    utils.Bool(false),
			})
			if unavailableActivityErr != nil {
				Fail(fmt.Sprintf("Failed to create activity. Error %s", unavailableActivityErr.Error()))
			}
			unavailableActivitySid = unavailableActivityResp.Sid

			workerResp, workerErr := taskrouterSession.Workspace(workspaceSid).Workers.Create(&workers.CreateWorkerInput{
				FriendlyName: uuid.New().String(),
				ActivitySid:  utils.String(activitySid),
			})
			if workerErr != nil {
				Fail(fmt.Sprintf("Failed to create worker. Error %s", workerErr.Error()))
			}
			workerSid = workerResp.Sid

			taskResp, takErr := taskrouterSession.Workspace(workspaceSid).Tasks.Create(&tasks.CreateTaskInput{
				TaskChannel: utils.String("default"),
			})
			if takErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", takErr.Error()))
			}
			taskSid = taskResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			// Can't delete a worker until there are no longer able to accept tasks
			if _, err := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Update(&worker.UpdateWorkerInput{
				ActivitySid: utils.String(unavailableActivitySid),
			}); err != nil {
				Fail(fmt.Sprintf("Failed to update worker. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete worker. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Activity(activitySid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete activity. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Activity(unavailableActivitySid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete activity. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the reservation is fetched and updated", func() {
			reservationsClient := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Reservations

			pageResp, pageErr := reservationsClient.Page(&reservations.ReservationsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Reservations)).Should(BeNumerically(">=", 1))

			paginator := reservationsClient.NewReservationsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Reservations)).Should(BeNumerically(">=", 1))

			reservationClient := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Reservation(paginator.Reservations[0].Sid)

			fetchResp, fetchErr := reservationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := reservationClient.Update(&reservation.UpdateReservationInput{
				ReservationStatus: "accepted",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})

	Describe("Given the TaskRouter Activity clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the activity is create, fetched, updated and deleted", func() {
			activitiesClient := taskrouterSession.Workspace(workspaceSid).Activities

			createResp, createErr := activitiesClient.Create(&activities.CreateActivityInput{
				FriendlyName: uuid.New().String(),
				Available:    utils.Bool(true),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := activitiesClient.Page(&activities.ActivitiesPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Activities)).Should(BeNumerically(">=", 1))

			paginator := activitiesClient.NewActivitiesPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Activities)).Should(BeNumerically(">=", 1))

			activityClient := taskrouterSession.Workspace(workspaceSid).Activity(createResp.Sid)

			fetchResp, fetchErr := activityClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := activityClient.Update(&activity.UpdateActivityInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := activityClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Channel clients", func() {

		var workspaceSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the task channel is create, fetched, updated and deleted", func() {
			taskChannelsClient := taskrouterSession.Workspace(workspaceSid).TaskChannels

			createResp, createErr := taskChannelsClient.Create(&task_channels.CreateTaskChannelInput{
				FriendlyName: uuid.New().String(),
				UniqueName:   uuid.New().String(),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := taskChannelsClient.Page(&task_channels.TaskChannelsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.TaskChannels)).Should(BeNumerically(">=", 1))

			paginator := taskChannelsClient.NewTaskChannelsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.TaskChannels)).Should(BeNumerically(">=", 1))

			taskChannelClient := taskrouterSession.Workspace(workspaceSid).TaskChannel(createResp.Sid)

			fetchResp, fetchErr := taskChannelClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := taskChannelClient.Update(&task_channel.UpdateTaskChannelInput{
				ChannelOptimizedRouting: utils.Bool(true),
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := taskChannelClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the task is create, fetched, updated and deleted", func() {
			tasksClient := taskrouterSession.Workspace(workspaceSid).Tasks

			createResp, createErr := tasksClient.Create(&tasks.CreateTaskInput{
				TaskChannel: utils.String("default"),
			})
			Expect(createErr).To(BeNil())
			Expect(createResp).ToNot(BeNil())
			Expect(createResp.Sid).ToNot(BeNil())

			pageResp, pageErr := tasksClient.Page(&tasks.TasksPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Tasks)).Should(BeNumerically(">=", 1))

			paginator := tasksClient.NewTasksPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Tasks)).Should(BeNumerically(">=", 1))

			taskClient := taskrouterSession.Workspace(workspaceSid).Task(createResp.Sid)

			fetchResp, fetchErr := taskClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := taskClient.Update(&task.UpdateTaskInput{})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())

			deleteErr := taskClient.Delete()
			Expect(deleteErr).To(BeNil())
		})
	})

	Describe("Given the TaskRouter Task Reservation clients", func() {

		var workspaceSid string
		var taskQueueSid string
		var workflowSid string
		var activitySid string
		var unavailableActivitySid string
		var workerSid string
		var taskSid string

		BeforeEach(func() {
			resp, err := taskrouterSession.Workspaces.Create(&workspaces.CreateWorkspaceInput{
				FriendlyName: uuid.New().String(),
			})
			if err != nil {
				Fail(fmt.Sprintf("Failed to create workspace. Error %s", err.Error()))
			}
			workspaceSid = resp.Sid

			taskQueueResp, taskQueueErr := taskrouterSession.Workspace(workspaceSid).TaskQueues.Create(&task_queues.CreateTaskQueueInput{
				FriendlyName: uuid.New().String(),
			})
			if taskQueueErr != nil {
				Fail(fmt.Sprintf("Failed to create task queue. Error %s", taskQueueErr.Error()))
			}
			taskQueueSid = taskQueueResp.Sid

			workflowResp, workflowErr := taskrouterSession.Workspace(workspaceSid).Workflows.Create(&workflows.CreateWorkflowInput{
				FriendlyName:  uuid.New().String(),
				Configuration: fmt.Sprintf(`{ "task_routing": { "default_filter": { "queue": "%s" } } }`, taskQueueSid),
			})
			if workflowErr != nil {
				Fail(fmt.Sprintf("Failed to create workflow. Error %s", workflowErr.Error()))
			}
			workflowSid = workflowResp.Sid

			activityResp, activityErr := taskrouterSession.Workspace(workspaceSid).Activities.Create(&activities.CreateActivityInput{
				FriendlyName: uuid.New().String(),
				Available:    utils.Bool(true),
			})
			if activityErr != nil {
				Fail(fmt.Sprintf("Failed to create activity. Error %s", activityErr.Error()))
			}
			activitySid = activityResp.Sid

			unavailableActivityResp, unavailableActivityErr := taskrouterSession.Workspace(workspaceSid).Activities.Create(&activities.CreateActivityInput{
				FriendlyName: uuid.New().String(),
				Available:    utils.Bool(false),
			})
			if unavailableActivityErr != nil {
				Fail(fmt.Sprintf("Failed to create activity. Error %s", unavailableActivityErr.Error()))
			}
			unavailableActivitySid = unavailableActivityResp.Sid

			workerResp, workerErr := taskrouterSession.Workspace(workspaceSid).Workers.Create(&workers.CreateWorkerInput{
				FriendlyName: uuid.New().String(),
				ActivitySid:  utils.String(activitySid),
			})
			if workerErr != nil {
				Fail(fmt.Sprintf("Failed to create worker. Error %s", workerErr.Error()))
			}
			workerSid = workerResp.Sid

			taskResp, takErr := taskrouterSession.Workspace(workspaceSid).Tasks.Create(&tasks.CreateTaskInput{
				TaskChannel: utils.String("default"),
			})
			if takErr != nil {
				Fail(fmt.Sprintf("Failed to create task. Error %s", takErr.Error()))
			}
			taskSid = taskResp.Sid
		})

		AfterEach(func() {
			if err := taskrouterSession.Workspace(workspaceSid).Task(taskSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Workflow(workflowSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workflow. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).TaskQueue(taskQueueSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete task queue. Error %s", err.Error()))
			}

			// Can't delete a worker until there are no longer able to accept tasks
			if _, err := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Update(&worker.UpdateWorkerInput{
				ActivitySid: utils.String(unavailableActivitySid),
			}); err != nil {
				Fail(fmt.Sprintf("Failed to update worker. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Worker(workerSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete worker. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Activity(activitySid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete activity. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Activity(unavailableActivitySid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete activity. Error %s", err.Error()))
			}

			if err := taskrouterSession.Workspace(workspaceSid).Delete(); err != nil {
				Fail(fmt.Sprintf("Failed to delete workspace. Error %s", err.Error()))
			}
		})

		It("Then the reservation is fetched and updated", func() {
			reservationsClient := taskrouterSession.Workspace(workspaceSid).Task(taskSid).Reservations

			pageResp, pageErr := reservationsClient.Page(&taskReservations.ReservationsPageOptions{})
			Expect(pageErr).To(BeNil())
			Expect(pageResp).ToNot(BeNil())
			Expect(len(pageResp.Reservations)).Should(BeNumerically(">=", 1))

			paginator := reservationsClient.NewReservationsPaginator()
			for paginator.Next() {
			}

			Expect(paginator.Error()).To(BeNil())
			Expect(len(paginator.Reservations)).Should(BeNumerically(">=", 1))

			reservationClient := taskrouterSession.Workspace(workspaceSid).Task(taskSid).Reservation(paginator.Reservations[0].Sid)

			fetchResp, fetchErr := reservationClient.Fetch()
			Expect(fetchErr).To(BeNil())
			Expect(fetchResp).ToNot(BeNil())

			updateResp, updateErr := reservationClient.Update(&taskReservation.UpdateReservationInput{
				ReservationStatus: "accepted",
			})
			Expect(updateErr).To(BeNil())
			Expect(updateResp).ToNot(BeNil())
		})
	})
})
