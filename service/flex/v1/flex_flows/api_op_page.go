// Package flex_flows contains auto-generated files. DO NOT MODIFY
package flex_flows

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FlexFlowsPageOptions defines the query options for the api operation
type FlexFlowsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFlexFlowIntegrationResponse struct {
	Channel           *string `json:"channel,omitempty"`
	CreationOnMessage *bool   `json:"creation_on_message,omitempty"`
	FlowSid           *string `json:"flow_sid,omitempty"`
	Priority          *int    `json:"priority,omitempty"`
	RetryCount        *int    `json:"retry_count,omitempty"`
	Timeout           *int    `json:"timeout,omitempty"`
	URL               *string `json:"url,omitempty"`
	WorkflowSid       *string `json:"workflow_sid,omitempty"`
	WorkspaceSid      *string `json:"workspace_sid,omitempty"`
}

type PageFlexFlowResponse struct {
	AccountSid      string                           `json:"account_sid"`
	ChannelType     string                           `json:"channel_type"`
	ChatServiceSid  string                           `json:"chat_service_sid"`
	ContactIdentity *string                          `json:"contact_identity,omitempty"`
	DateCreated     time.Time                        `json:"date_created"`
	DateUpdated     *time.Time                       `json:"date_updated,omitempty"`
	Enabled         bool                             `json:"enabled"`
	FriendlyName    string                           `json:"friendly_name"`
	Integration     *PageFlexFlowIntegrationResponse `json:"integration,omitempty"`
	IntegrationType *string                          `json:"integration_type,omitempty"`
	JanitorEnabled  *bool                            `json:"janitor_enabled,omitempty"`
	LongLived       *bool                            `json:"long_lived,omitempty"`
	Sid             string                           `json:"sid"`
	URL             string                           `json:"url"`
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

// FlexFlowsPageResponse defines the response fields for the flex flows page
type FlexFlowsPageResponse struct {
	FlexFlows []PageFlexFlowResponse `json:"flex_flows"`
	Meta      PageMetaResponse       `json:"meta"`
}

// Page retrieves a page of flex flows
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FlexFlowsPageOptions) (*FlexFlowsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of flex flows
func (c Client) PageWithContext(context context.Context, options *FlexFlowsPageOptions) (*FlexFlowsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/FlexFlows",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FlexFlowsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FlexFlowsPaginator defines the fields for makings paginated api calls
// FlexFlows is an array of flexflows that have been returned from all of the page calls
type FlexFlowsPaginator struct {
	options   *FlexFlowsPageOptions
	Page      *FlexFlowsPage
	FlexFlows []PageFlexFlowResponse
}

// NewFlexFlowsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFlexFlowsPaginator() *FlexFlowsPaginator {
	return c.NewFlexFlowsPaginatorWithOptions(nil)
}

// NewFlexFlowsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFlexFlowsPaginatorWithOptions(options *FlexFlowsPageOptions) *FlexFlowsPaginator {
	return &FlexFlowsPaginator{
		options: options,
		Page: &FlexFlowsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		FlexFlows: make([]PageFlexFlowResponse, 0),
	}
}

// FlexFlowsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFlexFlowResponse or error that is returned from the api call(s)
type FlexFlowsPage struct {
	client *Client

	CurrentPage *FlexFlowsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FlexFlowsPaginator) CurrentPage() *FlexFlowsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FlexFlowsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FlexFlowsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FlexFlowsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FlexFlowsPageOptions{}
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
		p.FlexFlows = append(p.FlexFlows, resp.FlexFlows...)
	}

	return p.Page.Error == nil
}
