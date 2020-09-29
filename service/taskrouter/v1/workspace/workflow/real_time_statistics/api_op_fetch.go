// Package real_time_statistics contains auto-generated files. DO NOT MODIFY
package real_time_statistics

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchRealTimeStatisticsOptions defines the query options for the api operation
type FetchRealTimeStatisticsOptions struct {
	TaskChannel *string
}

// FetchRealTimeStatisticsResponse defines the response fields for the retrieved real time statistics
type FetchRealTimeStatisticsResponse struct {
	AccountSid            string         `json:"account_sid"`
	LongestTaskWaitingAge int            `json:"longest_task_waiting_age"`
	LongestTaskWaitingSid *string        `json:"longest_task_waiting_sid,omitempty"`
	TasksByPriority       map[string]int `json:"tasks_by_priority"`
	TasksByStatus         map[string]int `json:"tasks_by_status"`
	TotalTasks            int            `json:"total_tasks"`
	URL                   string         `json:"url"`
	WorkflowSid           string         `json:"workflow_sid"`
	WorkspaceSid          string         `json:"workspace_sid"`
}

// Fetch retrieves real time statistics
// See https://www.twilio.com/docs/taskrouter/api/workflow-statistics#workflow-realtime-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchRealTimeStatisticsOptions) (*FetchRealTimeStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves real time statistics
// See https://www.twilio.com/docs/taskrouter/api/workflow-statistics#workflow-realtime-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchRealTimeStatisticsOptions) (*FetchRealTimeStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workflows/{workflowSid}/RealTimeStatistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"workflowSid":  c.workflowSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchRealTimeStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
