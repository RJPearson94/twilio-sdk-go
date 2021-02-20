// Package composition_hooks contains auto-generated files. DO NOT MODIFY
package composition_hooks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CompositionHooksPageOptions defines the query options for the api operation
type CompositionHooksPageOptions struct {
	PageSize          *int
	Page              *int
	PageToken         *string
	Enabled           *bool
	FriendlyName      *string
	DateCreatedAfter  *string
	DateCreatedBefore *string
}

type PageCompositionHookResponse struct {
	AccountSid           string                 `json:"account_sid"`
	AudioSources         []string               `json:"audio_sources"`
	AudioSourcesExcluded []string               `json:"audio_sources_excluded"`
	DateCreated          time.Time              `json:"date_created"`
	DateUpdated          *time.Time             `json:"date_updated,omitempty"`
	Enabled              bool                   `json:"enabled"`
	Format               string                 `json:"format"`
	FriendlyName         string                 `json:"friendly_name"`
	Resolution           string                 `json:"resolution"`
	Sid                  string                 `json:"sid"`
	StatusCallback       *string                `json:"status_callback,omitempty"`
	StatusCallbackMethod string                 `json:"status_callback_method"`
	Trim                 bool                   `json:"trim"`
	URL                  string                 `json:"url"`
	VideoLayout          map[string]interface{} `json:"video_layout"`
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

// CompositionHooksPageResponse defines the response fields for the composition hooks page
type CompositionHooksPageResponse struct {
	CompositionHooks []PageCompositionHookResponse `json:"composition_hooks"`
	Meta             PageMetaResponse              `json:"meta"`
}

// Page retrieves a page of composition hooks
// See https://www.twilio.com/docs/video/api/composition-hooks#hks-get-parameters for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *CompositionHooksPageOptions) (*CompositionHooksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of composition hooks
// See https://www.twilio.com/docs/video/api/composition-hooks#hks-get-parameters for more details
func (c Client) PageWithContext(context context.Context, options *CompositionHooksPageOptions) (*CompositionHooksPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/CompositionHooks",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &CompositionHooksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CompositionHooksPaginator defines the fields for makings paginated api calls
// CompositionHooks is an array of compositionhooks that have been returned from all of the page calls
type CompositionHooksPaginator struct {
	options          *CompositionHooksPageOptions
	Page             *CompositionHooksPage
	CompositionHooks []PageCompositionHookResponse
}

// NewCompositionHooksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewCompositionHooksPaginator() *CompositionHooksPaginator {
	return c.NewCompositionHooksPaginatorWithOptions(nil)
}

// NewCompositionHooksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewCompositionHooksPaginatorWithOptions(options *CompositionHooksPageOptions) *CompositionHooksPaginator {
	return &CompositionHooksPaginator{
		options: options,
		Page: &CompositionHooksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		CompositionHooks: make([]PageCompositionHookResponse, 0),
	}
}

// CompositionHooksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageCompositionHookResponse or error that is returned from the api call(s)
type CompositionHooksPage struct {
	client *Client

	CurrentPage *CompositionHooksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *CompositionHooksPaginator) CurrentPage() *CompositionHooksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *CompositionHooksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *CompositionHooksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *CompositionHooksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &CompositionHooksPageOptions{}
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
		p.CompositionHooks = append(p.CompositionHooks, resp.CompositionHooks...)
	}

	return p.Page.Error == nil
}
