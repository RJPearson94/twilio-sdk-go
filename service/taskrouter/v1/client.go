// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/RJPearson94/twilio-sdk-go/session"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// TaskRouter client is used to manage resources for Twilio TaskRouter
// See https://www.twilio.com/docs/taskrouter for more details
type TaskRouter struct {
	client *client.Client

	Workspace  func(string) *workspace.Client
	Workspaces *workspaces.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *TaskRouter {
	return &TaskRouter{
		client: client,

		Workspace: func(workspaceSid string) *workspace.Client {
			return workspace.New(client, workspace.ClientProperties{
				Sid: workspaceSid,
			})
		},
		Workspaces: workspaces.New(client),
	}
}

// GetClient is used for testing purposes only
func (s TaskRouter) GetClient() *client.Client {
	return s.client
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *TaskRouter {
	return New(session.New(creds))
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *TaskRouter {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "taskrouter"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}
