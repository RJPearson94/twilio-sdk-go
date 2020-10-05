// Package environment contains auto-generated files. DO NOT MODIFY
package environment

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/deployment"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/deployments"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/log"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/logs"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variable"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variables"
)

// Client for managing a specific environment resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Deployment  func(string) *deployment.Client
	Deployments *deployments.Client
	Log         func(string) *log.Client
	Logs        *logs.Client
	Variable    func(string) *variable.Client
	Variables   *variables.Client
}

// ClientProperties are the properties required to manage the environment resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the environment client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Deployment: func(deploymentSid string) *deployment.Client {
			return deployment.New(client, deployment.ClientProperties{
				EnvironmentSid: properties.Sid,
				ServiceSid:     properties.ServiceSid,
				Sid:            deploymentSid,
			})
		},
		Deployments: deployments.New(client, deployments.ClientProperties{
			EnvironmentSid: properties.Sid,
			ServiceSid:     properties.ServiceSid,
		}),
		Log: func(logSid string) *log.Client {
			return log.New(client, log.ClientProperties{
				EnvironmentSid: properties.Sid,
				ServiceSid:     properties.ServiceSid,
				Sid:            logSid,
			})
		},
		Logs: logs.New(client, logs.ClientProperties{
			EnvironmentSid: properties.Sid,
			ServiceSid:     properties.ServiceSid,
		}),
		Variable: func(variableSid string) *variable.Client {
			return variable.New(client, variable.ClientProperties{
				EnvironmentSid: properties.Sid,
				ServiceSid:     properties.ServiceSid,
				Sid:            variableSid,
			})
		},
		Variables: variables.New(client, variables.ClientProperties{
			EnvironmentSid: properties.Sid,
			ServiceSid:     properties.ServiceSid,
		}),
	}
}
