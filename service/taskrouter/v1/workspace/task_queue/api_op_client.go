// Package task_queue contains auto-generated files. DO NOT MODIFY
package task_queue

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue/cumulative_statistics"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue/real_time_statistics"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue/statistics"
)

// Client for managing a specific task queue resource
// See https://www.twilio.com/docs/taskrouter/api/task-queue for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string

	CumulativeStatistics func() *cumulative_statistics.Client
	RealTimeStatistics   func() *real_time_statistics.Client
	Statistics           func() *statistics.Client
}

// ClientProperties are the properties required to manage the task queue resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the task queue client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,

		CumulativeStatistics: func() *cumulative_statistics.Client {
			return cumulative_statistics.New(client, cumulative_statistics.ClientProperties{
				TaskQueueSid: properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		RealTimeStatistics: func() *real_time_statistics.Client {
			return real_time_statistics.New(client, real_time_statistics.ClientProperties{
				TaskQueueSid: properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				TaskQueueSid: properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
	}
}
