// Package workers contains auto-generated files. DO NOT MODIFY
package workers

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers/cumulative_statistics"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers/real_time_statistics"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers/statistics"
)

// Client for managing worker resources
// See https://www.twilio.com/docs/taskrouter/api/worker for more details
type Client struct {
	client *client.Client

	workspaceSid string

	CumulativeStatistics func() *cumulative_statistics.Client
	RealTimeStatistics   func() *real_time_statistics.Client
	Statistics           func() *statistics.Client
}

// ClientProperties are the properties required to manage the workers resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the workers client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,

		CumulativeStatistics: func() *cumulative_statistics.Client {
			return cumulative_statistics.New(client, cumulative_statistics.ClientProperties{
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		RealTimeStatistics: func() *real_time_statistics.Client {
			return real_time_statistics.New(client, real_time_statistics.ClientProperties{
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
	}
}
