// Package task_channel contains auto-generated files. DO NOT MODIFY
package task_channel

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a task channel resource from the account
// See https://www.twilio.com/docs/taskrouter/api/task-channel#delete-a-taskchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a task channel resource from the account
// See https://www.twilio.com/docs/taskrouter/api/task-channel#delete-a-taskchannel-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Workspaces/{workspaceSid}/TaskChannels/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
