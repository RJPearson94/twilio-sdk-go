package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

type TaskRouter struct {
	client     *client.Client
	Workspace  func(string) *workspace.Client
	Workspaces *workspaces.Client
}

// Used for testing purposes only
func (s TaskRouter) GetClient() *client.Client {
	return s.client
}

func New(sess *session.Session) *TaskRouter {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "taskrouter"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

func NewWithClient(client *client.Client) *TaskRouter {
	return &TaskRouter{
		client:     client,
		Workspace:  func(sid string) *workspace.Client { return workspace.New(client, sid) },
		Workspaces: workspaces.New(client),
	}
}

func NewWithCredentials(creds *credentials.Credentials) *TaskRouter {
	return New(session.New(creds))
}
