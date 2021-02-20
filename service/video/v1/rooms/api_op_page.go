// Package rooms contains auto-generated files. DO NOT MODIFY
package rooms

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// RoomsPageOptions defines the query options for the api operation
type RoomsPageOptions struct {
	PageSize          *int
	Page              *int
	PageToken         *string
	Status            *string
	UniqueName        *string
	DateCreatedAfter  *string
	DateCreatedBefore *string
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

type PageRoomResponse struct {
	AccountSid                   string     `json:"account_sid"`
	DateCreated                  time.Time  `json:"date_created"`
	DateUpdated                  *time.Time `json:"date_updated,omitempty"`
	Duration                     *int       `json:"duration,omitempty"`
	EndTime                      *time.Time `json:"end_time,omitempty"`
	MaxConcurrentPublishedTracks *int       `json:"max_concurrent_published_tracks,omitempty"`
	MaxParticipants              int        `json:"max_participants"`
	MediaRegion                  *string    `json:"media_region,omitempty"`
	RecordParticipantsOnConnect  bool       `json:"record_participants_on_connect"`
	Sid                          string     `json:"sid"`
	Status                       string     `json:"status"`
	StatusCallback               *string    `json:"status_callback,omitempty"`
	StatusCallbackMethod         string     `json:"status_callback_method"`
	Type                         string     `json:"type"`
	URL                          string     `json:"url"`
	UniqueName                   string     `json:"unique_name"`
	VideoCodecs                  *[]string  `json:"video_codecs,omitempty"`
}

// RoomsPageResponse defines the response fields for the rooms page
type RoomsPageResponse struct {
	Meta  PageMetaResponse   `json:"meta"`
	Rooms []PageRoomResponse `json:"rooms"`
}

// Page retrieves a page of rooms
// See https://www.twilio.com/docs/video/api/rooms-resource#get-list-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *RoomsPageOptions) (*RoomsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of rooms
// See https://www.twilio.com/docs/video/api/rooms-resource#get-list-resource for more details
func (c Client) PageWithContext(context context.Context, options *RoomsPageOptions) (*RoomsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Rooms",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &RoomsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RoomsPaginator defines the fields for makings paginated api calls
// Rooms is an array of rooms that have been returned from all of the page calls
type RoomsPaginator struct {
	options *RoomsPageOptions
	Page    *RoomsPage
	Rooms   []PageRoomResponse
}

// NewRoomsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewRoomsPaginator() *RoomsPaginator {
	return c.NewRoomsPaginatorWithOptions(nil)
}

// NewRoomsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewRoomsPaginatorWithOptions(options *RoomsPageOptions) *RoomsPaginator {
	return &RoomsPaginator{
		options: options,
		Page: &RoomsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Rooms: make([]PageRoomResponse, 0),
	}
}

// RoomsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageRoomResponse or error that is returned from the api call(s)
type RoomsPage struct {
	client *Client

	CurrentPage *RoomsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *RoomsPaginator) CurrentPage() *RoomsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *RoomsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *RoomsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *RoomsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &RoomsPageOptions{}
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
		p.Rooms = append(p.Rooms, resp.Rooms...)
	}

	return p.Page.Error == nil
}
