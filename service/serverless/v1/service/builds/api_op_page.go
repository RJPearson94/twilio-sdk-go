// Package builds contains auto-generated files. DO NOT MODIFY
package builds

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// BuildsPageOptions defines the query options for the api operation
type BuildsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageAssetVersion struct {
	AccountSid  string    `json:"account_sid"`
	AssetSid    string    `json:"asset_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
}

type PageBuildResponse struct {
	AccountSid       string                 `json:"account_sid"`
	AssetVersions    *[]PageAssetVersion    `json:"asset_versions,omitempty"`
	DateCreated      time.Time              `json:"date_created"`
	DateUpdated      *time.Time             `json:"date_updated,omitempty"`
	Dependencies     *[]PageDependency      `json:"dependencies,omitempty"`
	FunctionVersions *[]PageFunctionVersion `json:"function_versions,omitempty"`
	Runtime          string                 `json:"runtime"`
	ServiceSid       string                 `json:"service_sid"`
	Sid              string                 `json:"sid"`
	Status           string                 `json:"status"`
	URL              string                 `json:"url"`
}

type PageDependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PageFunctionVersion struct {
	AccountSid  string    `json:"account_sid"`
	DateCreated time.Time `json:"date_created"`
	FunctionSid string    `json:"function_sid"`
	Path        string    `json:"path"`
	ServiceSid  string    `json:"service_sid"`
	Sid         string    `json:"sid"`
	URL         string    `json:"url"`
	Visibility  string    `json:"visibility"`
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

// BuildsPageResponse defines the response fields for the builds page
type BuildsPageResponse struct {
	Builds []PageBuildResponse `json:"builds"`
	Meta   PageMetaResponse    `json:"meta"`
}

// Page retrieves a page of builds
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build#read-multiple-build-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *BuildsPageOptions) (*BuildsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of builds
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/build#read-multiple-build-resources for more details
func (c Client) PageWithContext(context context.Context, options *BuildsPageOptions) (*BuildsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Builds",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &BuildsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// BuildsPaginator defines the fields for makings paginated api calls
// Builds is an array of builds that have been returned from all of the page calls
type BuildsPaginator struct {
	options *BuildsPageOptions
	Page    *BuildsPage
	Builds  []PageBuildResponse
}

// NewBuildsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewBuildsPaginator() *BuildsPaginator {
	return c.NewBuildsPaginatorWithOptions(nil)
}

// NewBuildsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewBuildsPaginatorWithOptions(options *BuildsPageOptions) *BuildsPaginator {
	return &BuildsPaginator{
		options: options,
		Page: &BuildsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Builds: make([]PageBuildResponse, 0),
	}
}

// BuildsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageBuildResponse or error that is returned from the api call(s)
type BuildsPage struct {
	client *Client

	CurrentPage *BuildsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *BuildsPaginator) CurrentPage() *BuildsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *BuildsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *BuildsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *BuildsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &BuildsPageOptions{}
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
		p.Builds = append(p.Builds, resp.Builds...)
	}

	return p.Page.Error == nil
}
