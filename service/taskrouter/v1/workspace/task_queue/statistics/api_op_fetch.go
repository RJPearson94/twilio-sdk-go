// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FetchStatisticsOptions defines the query options for the api operation
type FetchStatisticsOptions struct {
	Minutes         *int
	StartDate       *time.Time
	EndDate         *time.Time
	TaskChannel     *string
	SplitByWaitTime *string
}

type FetchActivityStatistic struct {
	FriendlyName string `json:"friendly_name"`
	Sid          string `json:"sid"`
	Workers      int    `json:"workers"`
}

type FetchCumulativeStatistics struct {
	AvgTaskAcceptanceTime            int                       `json:"avg_task_acceptance_time"`
	EndTime                          time.Time                 `json:"end_time"`
	ReservationsAccepted             int                       `json:"reservations_accepted"`
	ReservationsCanceled             int                       `json:"reservations_canceled"`
	ReservationsCreated              int                       `json:"reservations_created"`
	ReservationsRejected             int                       `json:"reservations_rejected"`
	ReservationsRescinded            int                       `json:"reservations_rescinded"`
	ReservationsTimedOut             int                       `json:"reservations_timed_out"`
	SplitByWaitTime                  *map[string]FetchWaitTime `json:"split_by_wait_time,omitempty"`
	StartTime                        time.Time                 `json:"start_time"`
	TasksCanceled                    int                       `json:"tasks_canceled"`
	TasksCompleted                   int                       `json:"tasks_completed"`
	TasksDeleted                     int                       `json:"tasks_deleted"`
	TasksEntered                     int                       `json:"tasks_entered"`
	TasksMoved                       int                       `json:"tasks_moved"`
	WaitDurationInQueueUntilAccepted FetchStatisticsBreakdown  `json:"wait_duration_in_queue_until_accepted"`
	WaitDurationUntilAccepted        FetchStatisticsBreakdown  `json:"wait_duration_until_accepted"`
	WaitDurationUntilCanceled        FetchStatisticsBreakdown  `json:"wait_duration_until_canceled"`
}

type FetchRealTimeStatistics struct {
	ActivityStatistics            []FetchActivityStatistic `json:"activity_statistics"`
	LongestRelativeTaskAgeInQueue int                      `json:"longest_relative_task_age_in_queue"`
	LongestRelativeTaskSidInQueue *string                  `json:"longest_relative_task_sid_in_queue,omitempty"`
	LongestTaskWaitingAge         int                      `json:"longest_task_waiting_age"`
	LongestTaskWaitingSid         *string                  `json:"longest_task_waiting_sid,omitempty"`
	TasksByPriority               map[string]int           `json:"tasks_by_priority"`
	TasksByStatus                 map[string]int           `json:"tasks_by_status"`
	TotalAvailableWorkers         int                      `json:"total_available_workers"`
	TotalEligibleWorkers          int                      `json:"total_eligible_workers"`
	TotalTasks                    int                      `json:"total_tasks"`
}

type FetchStatisticsBreakdown struct {
	Avg   int `json:"avg"`
	Max   int `json:"max"`
	Min   int `json:"min"`
	Total int `json:"total"`
}

type FetchWaitTime struct {
	Above FetchWaitTimeTasks `json:"above"`
	Below FetchWaitTimeTasks `json:"below"`
}

type FetchWaitTimeTasks struct {
	ReservationsAccepted int `json:"reservations_accepted"`
	TasksCanceled        int `json:"tasks_canceled"`
}

// FetchStatisticsResponse defines the response fields for the retrieved statistics
type FetchStatisticsResponse struct {
	AccountSid   string                    `json:"account_sid"`
	Cumulative   FetchCumulativeStatistics `json:"cumulative"`
	RealTime     FetchRealTimeStatistics   `json:"realtime"`
	TaskQueueSid string                    `json:"task_queue_sid"`
	URL          string                    `json:"url"`
	WorkspaceSid string                    `json:"workspace_sid"`
}

// Fetch retrieves statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics#taskqueue-instance-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchStatisticsOptions) (*FetchStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves statistics
// See https://www.twilio.com/docs/taskrouter/api/taskqueue-statistics#taskqueue-instance-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchStatisticsOptions) (*FetchStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/TaskQueues/{taskQueueSid}/Statistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"taskQueueSid": c.taskQueueSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
