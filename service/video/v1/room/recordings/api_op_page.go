// Package recordings contains auto-generated files. DO NOT MODIFY
package recordings

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// RecordingsPageOptions defines the query options for the api operation
type RecordingsPageOptions struct {
	PageSize          *int
	Page              *int
	PageToken         *string
	Status            *string
	SourceSid         *string
	GroupingSid       *string
	MediaType         *string
	DateCreatedAfter  *time.Time
	DateCreatedBefore *time.Time
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

type PageRecordingGroupingSidsResponse struct {
	ParticipantSid string `json:"participant_sid"`
	RoomSid        string `json:"room_sid"`
}

type PageRecordingResponse struct {
	AccountSid      string                            `json:"account_sid"`
	Codec           string                            `json:"codec"`
	ContainerFormat string                            `json:"container_format"`
	DateCreated     time.Time                         `json:"date_created"`
	Duration        int                               `json:"duration"`
	GroupingSids    PageRecordingGroupingSidsResponse `json:"grouping_sids"`
	Offset          int                               `json:"offset"`
	RoomSid         string                            `json:"room_sid"`
	Sid             string                            `json:"sid"`
	Size            int                               `json:"size"`
	SourceSid       string                            `json:"source_sid"`
	Status          string                            `json:"status"`
	TrackName       string                            `json:"track_name"`
	Type            string                            `json:"type"`
	URL             string                            `json:"url"`
}

// RecordingsPageResponse defines the response fields for the recordings page
type RecordingsPageResponse struct {
	Meta       PageMetaResponse        `json:"meta"`
	Recordings []PageRecordingResponse `json:"recordings"`
}

// Page retrieves a page of recordings
// See https://www.twilio.com/docs/video/api/recordings-resource#get-list-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *RecordingsPageOptions) (*RecordingsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of recordings
// See https://www.twilio.com/docs/video/api/recordings-resource#get-list-resource for more details
func (c Client) PageWithContext(context context.Context, options *RecordingsPageOptions) (*RecordingsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Recordings",
		PathParams: map[string]string{
			"roomSid": c.roomSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &RecordingsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RecordingsPaginator defines the fields for makings paginated api calls
// Recordings is an array of recordings that have been returned from all of the page calls
type RecordingsPaginator struct {
	options    *RecordingsPageOptions
	Page       *RecordingsPage
	Recordings []PageRecordingResponse
}

// NewRecordingsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewRecordingsPaginator() *RecordingsPaginator {
	return c.NewRecordingsPaginatorWithOptions(nil)
}

// NewRecordingsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewRecordingsPaginatorWithOptions(options *RecordingsPageOptions) *RecordingsPaginator {
	return &RecordingsPaginator{
		options: options,
		Page: &RecordingsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Recordings: make([]PageRecordingResponse, 0),
	}
}

// RecordingsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageRecordingResponse or error that is returned from the api call(s)
type RecordingsPage struct {
	client *Client

	CurrentPage *RecordingsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *RecordingsPaginator) CurrentPage() *RecordingsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *RecordingsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *RecordingsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *RecordingsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &RecordingsPageOptions{}
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
		p.Recordings = append(p.Recordings, resp.Recordings...)
	}

	return p.Page.Error == nil
}
