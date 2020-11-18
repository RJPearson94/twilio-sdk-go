// Package queries contains auto-generated files. DO NOT MODIFY
package queries

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// QueriesPageOptions defines the query options for the api operation
type QueriesPageOptions struct {
	Language    *string
	ModelBuild  *string
	Status      *string
	DialogueSid *string
	PageSize    *int
	Page        *int
	PageToken   *string
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

type PageQueryFieldResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PageQueryResponse struct {
	AccountSid    string                  `json:"account_sid"`
	AssistantSid  string                  `json:"assistant_sid"`
	DateCreated   time.Time               `json:"date_created"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	DialogueSid   *string                 `json:"dialogue_sid,omitempty"`
	Language      string                  `json:"language"`
	ModelBuildSid string                  `json:"model_build_sid"`
	Query         string                  `json:"query"`
	Results       PageQueryResultResponse `json:"results"`
	SampleSid     string                  `json:"sample_sid"`
	Sid           string                  `json:"sid"`
	SourceChannel string                  `json:"source_channel"`
	Status        string                  `json:"status"`
	URL           string                  `json:"url"`
}

type PageQueryResultResponse struct {
	Fields []PageQueryFieldResponse `json:"fields"`
	Task   string                   `json:"task"`
}

// QueriesPageResponse defines the response fields for the query page
type QueriesPageResponse struct {
	Meta    PageMetaResponse    `json:"meta"`
	Queries []PageQueryResponse `json:"queries"`
}

// Page retrieves a page of queries
// See https://www.twilio.com/docs/autopilot/api/query?code-sample=code-read-list-all-queries-for-an-assistant for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *QueriesPageOptions) (*QueriesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of queries
// See https://www.twilio.com/docs/autopilot/api/query?code-sample=code-read-list-all-queries-for-an-assistant for more details
func (c Client) PageWithContext(context context.Context, options *QueriesPageOptions) (*QueriesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Queries",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &QueriesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// QueriesPaginator defines the fields for makings paginated api calls
// Queries is an array of queries that have been returned from all of the page calls
type QueriesPaginator struct {
	options *QueriesPageOptions
	Page    *QueriesPage
	Queries []PageQueryResponse
}

// NewQueriesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewQueriesPaginator() *QueriesPaginator {
	return c.NewQueriesPaginatorWithOptions(nil)
}

// NewQueriesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewQueriesPaginatorWithOptions(options *QueriesPageOptions) *QueriesPaginator {
	return &QueriesPaginator{
		options: options,
		Page: &QueriesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Queries: make([]PageQueryResponse, 0),
	}
}

// QueriesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageQueryResponse or error that is returned from the api call(s)
type QueriesPage struct {
	client *Client

	CurrentPage *QueriesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *QueriesPaginator) CurrentPage() *QueriesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *QueriesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *QueriesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *QueriesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &QueriesPageOptions{}
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
		p.Queries = append(p.Queries, resp.Queries...)
	}

	return p.Page.Error == nil
}
