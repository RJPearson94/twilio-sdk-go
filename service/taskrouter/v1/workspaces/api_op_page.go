// Package workspaces contains auto-generated files. DO NOT MODIFY
package workspaces

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// WorkspacesPageOptions defines the query options for the api operation
type WorkspacesPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
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

type PageWorkspaceResponse struct {
	AccountSid           string     `json:"account_sid"`
	DateCreated          time.Time  `json:"date_created"`
	DateUpdated          *time.Time `json:"date_updated,omitempty"`
	DefaultActivityName  string     `json:"default_activity_name"`
	DefaultActivitySid   string     `json:"default_activity_sid"`
	EventCallbackURL     *string    `json:"event_callback_url,omitempty"`
	EventsFilter         *string    `json:"events_filter,omitempty"`
	FriendlyName         string     `json:"friendly_name"`
	MultiTaskEnabled     bool       `json:"multi_task_enabled"`
	PrioritizeQueueOrder string     `json:"prioritize_queue_order"`
	Sid                  string     `json:"sid"`
	TimeoutActivityName  string     `json:"timeout_activity_name"`
	TimeoutActivitySid   string     `json:"timeout_activity_sid"`
	URL                  string     `json:"url"`
}

// WorkspacesPageResponse defines the response fields for the workspaces page
type WorkspacesPageResponse struct {
	Meta       PageMetaResponse        `json:"meta"`
	Workspaces []PageWorkspaceResponse `json:"workspaces"`
}

// Page retrieves a page of workspaces
// See https://www.twilio.com/docs/taskrouter/api/workspace#list-all-workspaces for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *WorkspacesPageOptions) (*WorkspacesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of workspaces
// See https://www.twilio.com/docs/taskrouter/api/workspace#list-all-workspaces for more details
func (c Client) PageWithContext(context context.Context, options *WorkspacesPageOptions) (*WorkspacesPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Workspaces",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &WorkspacesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// WorkspacesPaginator defines the fields for makings paginated api calls
// Workspaces is an array of workspaces that have been returned from all of the page calls
type WorkspacesPaginator struct {
	options    *WorkspacesPageOptions
	Page       *WorkspacesPage
	Workspaces []PageWorkspaceResponse
}

// NewWorkspacesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewWorkspacesPaginator() *WorkspacesPaginator {
	return c.NewWorkspacesPaginatorWithOptions(nil)
}

// NewWorkspacesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewWorkspacesPaginatorWithOptions(options *WorkspacesPageOptions) *WorkspacesPaginator {
	return &WorkspacesPaginator{
		options: options,
		Page: &WorkspacesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Workspaces: make([]PageWorkspaceResponse, 0),
	}
}

// WorkspacesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageWorkspaceResponse or error that is returned from the api call(s)
type WorkspacesPage struct {
	client *Client

	CurrentPage *WorkspacesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *WorkspacesPaginator) CurrentPage() *WorkspacesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *WorkspacesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *WorkspacesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *WorkspacesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &WorkspacesPageOptions{}
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
		p.Workspaces = append(p.Workspaces, resp.Workspaces...)
	}

	return p.Page.Error == nil
}
