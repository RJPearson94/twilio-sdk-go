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
	Minutes     *int
	StartDate   *time.Time
	EndDate     *time.Time
	TaskChannel *string
}

type FetchActivityDuration struct {
	Avg          int    `json:"avg"`
	FriendlyName string `json:"friendly_name"`
	Max          int    `json:"max"`
	Min          int    `json:"min"`
	Sid          string `json:"sid"`
	Total        int    `json:"total"`
}

type FetchCumulativeStatistics struct {
	ActivityDurations     []FetchActivityDuration `json:"activity_durations"`
	EndTime               time.Time               `json:"end_time"`
	ReservationsAccepted  int                     `json:"reservations_accepted"`
	ReservationsCanceled  int                     `json:"reservations_canceled"`
	ReservationsCompleted int                     `json:"reservations_completed"`
	ReservationsCreated   int                     `json:"reservations_created"`
	ReservationsRejected  int                     `json:"reservations_rejected"`
	ReservationsRescinded int                     `json:"reservations_rescinded"`
	ReservationsTimedOut  int                     `json:"reservations_timed_out"`
	ReservationsWrapUp    int                     `json:"reservations_wrapup"`
	StartTime             time.Time               `json:"start_time"`
	TasksAssigned         int                     `json:"tasks_assigned"`
}

// FetchStatisticsResponse defines the response fields for the retrieved statistics
type FetchStatisticsResponse struct {
	AccountSid   string                    `json:"account_sid"`
	Cumulative   FetchCumulativeStatistics `json:"cumulative"`
	URL          string                    `json:"url"`
	WorkerSid    string                    `json:"worker_sid"`
	WorkspaceSid string                    `json:"workspace_sid"`
}

// Fetch retrieves statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics#fetch-a-specific-workers-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchStatisticsOptions) (*FetchStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics#fetch-a-specific-workers-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchStatisticsOptions) (*FetchStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers/{workerSid}/Statistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"workerSid":    c.workerSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
