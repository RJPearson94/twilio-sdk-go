// Package activities contains auto-generated files. DO NOT MODIFY
package activities

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// ActivitiesPageOptions defines the query options for the api operation
type ActivitiesPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
	Available    *bool
}

type PageActivityResponse struct {
	AccountSid   string     `json:"account_sid"`
	Available    bool       `json:"available"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	WorkspaceSid string     `json:"workspace_sid"`
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

// ActivitiesPageResponse defines the response fields for the activities page
type ActivitiesPageResponse struct {
	Activities []PageActivityResponse `json:"activities"`
	Meta       PageMetaResponse       `json:"meta"`
}

// Page retrieves a page of activities
// See https://www.twilio.com/docs/taskrouter/api/activity#read-multiple-activity-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ActivitiesPageOptions) (*ActivitiesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of activities
// See https://www.twilio.com/docs/taskrouter/api/activity#read-multiple-activity-resources for more details
func (c Client) PageWithContext(context context.Context, options *ActivitiesPageOptions) (*ActivitiesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Activities",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ActivitiesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ActivitiesPaginator defines the fields for makings paginated api calls
// Activities is an array of activities that have been returned from all of the page calls
type ActivitiesPaginator struct {
	options    *ActivitiesPageOptions
	Page       *ActivitiesPage
	Activities []PageActivityResponse
}

// NewActivitiesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewActivitiesPaginator() *ActivitiesPaginator {
	return c.NewActivitiesPaginatorWithOptions(nil)
}

// NewActivitiesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewActivitiesPaginatorWithOptions(options *ActivitiesPageOptions) *ActivitiesPaginator {
	return &ActivitiesPaginator{
		options: options,
		Page: &ActivitiesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Activities: make([]PageActivityResponse, 0),
	}
}

// ActivitiesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageActivityResponse or error that is returned from the api call(s)
type ActivitiesPage struct {
	client *Client

	CurrentPage *ActivitiesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ActivitiesPaginator) CurrentPage() *ActivitiesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ActivitiesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ActivitiesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ActivitiesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ActivitiesPageOptions{}
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
		p.Activities = append(p.Activities, resp.Activities...)
	}

	return p.Page.Error == nil
}
