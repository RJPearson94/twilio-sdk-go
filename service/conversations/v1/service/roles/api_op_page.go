// Package roles contains auto-generated files. DO NOT MODIFY
package roles

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// RolesPageOptions defines the query options for the api operation
type RolesPageOptions struct {
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

type PageRoleResponse struct {
	AccountSid     string     `json:"account_sid"`
	ChatServiceSid string     `json:"chat_service_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	FriendlyName   string     `json:"friendly_name"`
	Permissions    []string   `json:"permissions"`
	Sid            string     `json:"sid"`
	Type           string     `json:"type"`
	URL            string     `json:"url"`
}

// RolesPageResponse defines the response fields for the roles page
type RolesPageResponse struct {
	Meta  PageMetaResponse   `json:"meta"`
	Roles []PageRoleResponse `json:"roles"`
}

// Page retrieves a page of roles
// See https://www.twilio.com/docs/conversations/api/role-resource#read-multiple-role-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *RolesPageOptions) (*RolesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of roles
// See https://www.twilio.com/docs/conversations/api/role-resource#read-multiple-role-resources for more details
func (c Client) PageWithContext(context context.Context, options *RolesPageOptions) (*RolesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Roles",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &RolesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RolesPaginator defines the fields for makings paginated api calls
// Roles is an array of roles that have been returned from all of the page calls
type RolesPaginator struct {
	options *RolesPageOptions
	Page    *RolesPage
	Roles   []PageRoleResponse
}

// NewRolesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewRolesPaginator() *RolesPaginator {
	return c.NewRolesPaginatorWithOptions(nil)
}

// NewRolesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewRolesPaginatorWithOptions(options *RolesPageOptions) *RolesPaginator {
	return &RolesPaginator{
		options: options,
		Page: &RolesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Roles: make([]PageRoleResponse, 0),
	}
}

// RolesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageRoleResponse or error that is returned from the api call(s)
type RolesPage struct {
	client *Client

	CurrentPage *RolesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *RolesPaginator) CurrentPage() *RolesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *RolesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *RolesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *RolesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &RolesPageOptions{}
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
		p.Roles = append(p.Roles, resp.Roles...)
	}

	return p.Page.Error == nil
}
