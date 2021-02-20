// Package compositions contains auto-generated files. DO NOT MODIFY
package compositions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// CompositionsPageOptions defines the query options for the api operation
type CompositionsPageOptions struct {
	PageSize          *int
	Page              *int
	PageToken         *string
	Status            *string
	RoomSid           *string
	DateCreatedAfter  *string
	DateCreatedBefore *string
}

type PageCompositionResponse struct {
	AccountSid           string                 `json:"account_sid"`
	AudioSources         []string               `json:"audio_sources"`
	AudioSourcesExcluded []string               `json:"audio_sources_excluded"`
	Bitrate              int                    `json:"bitrate"`
	DateCompleted        *time.Time             `json:"date_completed,omitempty"`
	DateCreated          time.Time              `json:"date_created"`
	DateDeleted          *time.Time             `json:"date_deleted,omitempty"`
	Duration             int                    `json:"duration"`
	Format               string                 `json:"format"`
	Resolution           string                 `json:"resolution"`
	RoomSid              string                 `json:"room_sid"`
	Sid                  string                 `json:"sid"`
	Size                 int                    `json:"size"`
	Status               string                 `json:"status"`
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

// CompositionsPageResponse defines the response fields for the compositions page
type CompositionsPageResponse struct {
	Compositions []PageCompositionResponse `json:"compositions"`
	Meta         PageMetaResponse          `json:"meta"`
}

// Page retrieves a page of compositions
// See https://www.twilio.com/docs/video/api/compositions-resource#get-list-http-get for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *CompositionsPageOptions) (*CompositionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of compositions
// See https://www.twilio.com/docs/video/api/compositions-resource#get-list-http-get for more details
func (c Client) PageWithContext(context context.Context, options *CompositionsPageOptions) (*CompositionsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Compositions",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &CompositionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CompositionsPaginator defines the fields for makings paginated api calls
// Compositions is an array of compositions that have been returned from all of the page calls
type CompositionsPaginator struct {
	options      *CompositionsPageOptions
	Page         *CompositionsPage
	Compositions []PageCompositionResponse
}

// NewCompositionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewCompositionsPaginator() *CompositionsPaginator {
	return c.NewCompositionsPaginatorWithOptions(nil)
}

// NewCompositionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewCompositionsPaginatorWithOptions(options *CompositionsPageOptions) *CompositionsPaginator {
	return &CompositionsPaginator{
		options: options,
		Page: &CompositionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Compositions: make([]PageCompositionResponse, 0),
	}
}

// CompositionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageCompositionResponse or error that is returned from the api call(s)
type CompositionsPage struct {
	client *Client

	CurrentPage *CompositionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *CompositionsPaginator) CurrentPage() *CompositionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *CompositionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *CompositionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *CompositionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &CompositionsPageOptions{}
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
		p.Compositions = append(p.Compositions, resp.Compositions...)
	}

	return p.Page.Error == nil
}
