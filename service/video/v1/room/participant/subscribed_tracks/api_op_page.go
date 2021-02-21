// Package subscribed_tracks contains auto-generated files. DO NOT MODIFY
package subscribed_tracks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// SubscribedTracksPageOptions defines the query options for the api operation
type SubscribedTracksPageOptions struct {
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

type PageSubscribedTrackResponse struct {
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	Enabled        bool       `json:"enabled"`
	Kind           string     `json:"kind"`
	Name           string     `json:"name"`
	ParticipantSid string     `json:"participant_sid"`
	PublisherSid   string     `json:"publisher_sid"`
	RoomSid        string     `json:"room_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
}

// SubscribedTracksPageResponse defines the response fields for the subscribed track page
type SubscribedTracksPageResponse struct {
	Meta             PageMetaResponse              `json:"meta"`
	SubscribedTracks []PageSubscribedTrackResponse `json:"subscribed_tracks"`
}

// Page retrieves a page of subscribed tracks
// See https://www.twilio.com/docs/video/api/track-subscriptions#get-stl for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SubscribedTracksPageOptions) (*SubscribedTracksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of subscribed tracks
// See https://www.twilio.com/docs/video/api/track-subscriptions#get-stl for more details
func (c Client) PageWithContext(context context.Context, options *SubscribedTracksPageOptions) (*SubscribedTracksPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Rooms/{roomSid}/Participants/{participantSid}/SubscribedTracks",
		PathParams: map[string]string{
			"roomSid":        c.roomSid,
			"participantSid": c.participantSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &SubscribedTracksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SubscribedTracksPaginator defines the fields for makings paginated api calls
// SubscribedTracks is an array of subscribedtracks that have been returned from all of the page calls
type SubscribedTracksPaginator struct {
	options          *SubscribedTracksPageOptions
	Page             *SubscribedTracksPage
	SubscribedTracks []PageSubscribedTrackResponse
}

// NewSubscribedTracksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSubscribedTracksPaginator() *SubscribedTracksPaginator {
	return c.NewSubscribedTracksPaginatorWithOptions(nil)
}

// NewSubscribedTracksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSubscribedTracksPaginatorWithOptions(options *SubscribedTracksPageOptions) *SubscribedTracksPaginator {
	return &SubscribedTracksPaginator{
		options: options,
		Page: &SubscribedTracksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		SubscribedTracks: make([]PageSubscribedTrackResponse, 0),
	}
}

// SubscribedTracksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSubscribedTrackResponse or error that is returned from the api call(s)
type SubscribedTracksPage struct {
	client *Client

	CurrentPage *SubscribedTracksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SubscribedTracksPaginator) CurrentPage() *SubscribedTracksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SubscribedTracksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SubscribedTracksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SubscribedTracksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SubscribedTracksPageOptions{}
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
		p.SubscribedTracks = append(p.SubscribedTracks, resp.SubscribedTracks...)
	}

	return p.Page.Error == nil
}
