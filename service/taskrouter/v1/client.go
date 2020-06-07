package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

type TaskRouter struct {
	Workspace  func(string) *workspace.Client
	Workspaces *workspaces.Client
}

func New(sess *session.Session) *TaskRouter {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "taskrouter"
	config.APIVersion = "v1"

	client := client.New(sess, config)

	return &TaskRouter{
		Workspace:  func(sid string) *workspace.Client { return workspace.New(client, sid) },
		Workspaces: workspaces.New(client),
	}
}

func NewWithCredentials(creds *credentials.Credentials) *TaskRouter {
	return New(session.New(creds))
}
