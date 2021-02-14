// Package events contains auto-generated files. DO NOT MODIFY
package events

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// EventsPageOptions defines the query options for the api operation
type EventsPageOptions struct {
	PageSize        *int
	Page            *int
	PageToken       *string
	ActorSid        *string
	EventType       *string
	ResourceSid     *string
	SourceIpAddress *string
	StartDate       *time.Time
	EndDate         *time.Time
}

type PageEventResponse struct {
	AccountSid      string                 `json:"account_sid"`
	ActorSid        string                 `json:"actor_sid"`
	ActorType       string                 `json:"actor_type"`
	Description     *string                `json:"description,omitempty"`
	EventData       map[string]interface{} `json:"event_data"`
	EventDate       time.Time              `json:"event_date"`
	EventType       string                 `json:"event_type"`
	ResourceSid     string                 `json:"resource_sid"`
	ResourceType    string                 `json:"resource_type"`
	Sid             string                 `json:"sid"`
	Source          string                 `json:"source"`
	SourceIpAddress string                 `json:"source_ip_address"`
	URL             string                 `json:"url"`
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

// EventsPageResponse defines the response fields for the events page
type EventsPageResponse struct {
	Events []PageEventResponse `json:"events"`
	Meta   PageMetaResponse    `json:"meta"`
}

// Page retrieves a page of events
// See https://www.twilio.com/docs/usage/monitor-events#read-multiple-event-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *EventsPageOptions) (*EventsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of events
// See https://www.twilio.com/docs/usage/monitor-events#read-multiple-event-resources for more details
func (c Client) PageWithContext(context context.Context, options *EventsPageOptions) (*EventsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Events",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &EventsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// EventsPaginator defines the fields for makings paginated api calls
// Events is an array of events that have been returned from all of the page calls
type EventsPaginator struct {
	options *EventsPageOptions
	Page    *EventsPage
	Events  []PageEventResponse
}

// NewEventsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewEventsPaginator() *EventsPaginator {
	return c.NewEventsPaginatorWithOptions(nil)
}

// NewEventsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewEventsPaginatorWithOptions(options *EventsPageOptions) *EventsPaginator {
	return &EventsPaginator{
		options: options,
		Page: &EventsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Events: make([]PageEventResponse, 0),
	}
}

// EventsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageEventResponse or error that is returned from the api call(s)
type EventsPage struct {
	client *Client

	CurrentPage *EventsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *EventsPaginator) CurrentPage() *EventsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *EventsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *EventsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *EventsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &EventsPageOptions{}
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
		p.Events = append(p.Events, resp.Events...)
	}

	return p.Page.Error == nil
}
