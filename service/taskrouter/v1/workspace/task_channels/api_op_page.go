// Package task_channels contains auto-generated files. DO NOT MODIFY
package task_channels

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// TaskChannelsPageOptions defines the query options for the api operation
type TaskChannelsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageMetaResponse struct {
	FirstPageURL    string  `json:"first_page_url"`
	Key             string  `json:"key"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
	URL             string  `json:"url"`
}

type PageTaskChannelResponse struct {
	AccountSid              string     `json:"account_sid"`
	ChannelOptimizedRouting *bool      `json:"channel_optimized_routing,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// TaskChannelsPageResponse defines the response fields for the task channels page
type TaskChannelsPageResponse struct {
	Meta         PageMetaResponse          `json:"meta"`
	TaskChannels []PageTaskChannelResponse `json:"channels"`
}

// Page retrieves a page of task channels
// See https://www.twilio.com/docs/taskrouter/api/task-channel#read-multiple-taskchannel-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *TaskChannelsPageOptions) (*TaskChannelsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of task channels
// See https://www.twilio.com/docs/taskrouter/api/task-channel#read-multiple-taskchannel-resources for more details
func (c Client) PageWithContext(context context.Context, options *TaskChannelsPageOptions) (*TaskChannelsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/TaskChannels",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &TaskChannelsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// TaskChannelsPaginator defines the fields for makings paginated api calls
// TaskChannels is an array of taskchannels that have been returned from all of the page calls
type TaskChannelsPaginator struct {
	options      *TaskChannelsPageOptions
	Page         *TaskChannelsPage
	TaskChannels []PageTaskChannelResponse
}

// NewTaskChannelsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewTaskChannelsPaginator() *TaskChannelsPaginator {
	return c.NewTaskChannelsPaginatorWithOptions(nil)
}

// NewTaskChannelsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewTaskChannelsPaginatorWithOptions(options *TaskChannelsPageOptions) *TaskChannelsPaginator {
	return &TaskChannelsPaginator{
		options: options,
		Page: &TaskChannelsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		TaskChannels: make([]PageTaskChannelResponse, 0),
	}
}

// TaskChannelsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageTaskChannelResponse or error that is returned from the api call(s)
type TaskChannelsPage struct {
	client *Client

	CurrentPage *TaskChannelsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *TaskChannelsPaginator) CurrentPage() *TaskChannelsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *TaskChannelsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *TaskChannelsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *TaskChannelsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &TaskChannelsPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().Meta.NextPageURL

		if nextPage == nil {
			return false
		}

		parsedURL, err := url.Parse(*nextPage)
		if err != nil {
			p.Page.Error = err
			return false
		}

		options.PageToken = utils.String(parsedURL.Query().Get("PageToken"))

		page, pageErr := strconv.Atoi(parsedURL.Query().Get("Page"))
		if pageErr != nil {
			p.Page.Error = pageErr
			return false
		}
		options.Page = utils.Int(page)

		pageSize, pageSizeErr := strconv.Atoi(parsedURL.Query().Get("PageSize"))
		if pageSizeErr != nil {
			p.Page.Error = pageSizeErr
			return false
		}
		options.PageSize = utils.Int(pageSize)
	}

	resp, err := p.Page.client.PageWithContext(context, options)
	p.Page.CurrentPage = resp
	p.Page.Error = err

	if p.Page.Error == nil {
		p.TaskChannels = append(p.TaskChannels, resp.TaskChannels...)
	}

	return p.Page.Error == nil
}
