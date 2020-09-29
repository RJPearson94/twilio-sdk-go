// Package interactions contains auto-generated files. DO NOT MODIFY
package interactions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// InteractionsPageOptions defines the query options for the api operation
type InteractionsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageInteractionResponse struct {
	AccountSid             string     `json:"account_sid"`
	Data                   *string    `json:"data,omitempty"`
	DateCreated            time.Time  `json:"date_created"`
	DateUpdated            *time.Time `json:"date_updated,omitempty"`
	InboundParticipantSid  *string    `json:"inbound_participant_sid,omitempty"`
	InboundResourceSid     *string    `json:"inbound_resource_sid,omitempty"`
	InboundResourceStatus  *string    `json:"inbound_resource_status,omitempty"`
	InboundResourceType    *string    `json:"inbound_resource_type,omitempty"`
	InboundResourceURL     *string    `json:"inbound_resource_url,omitempty"`
	OutboundParticipantSid *string    `json:"outbound_participant_sid,omitempty"`
	OutboundResourceSid    *string    `json:"outbound_resource_sid,omitempty"`
	OutboundResourceStatus *string    `json:"outbound_resource_status,omitempty"`
	OutboundResourceType   *string    `json:"outbound_resource_type,omitempty"`
	OutboundResourceURL    *string    `json:"outbound_resource_url,omitempty"`
	ServiceSid             string     `json:"service_sid"`
	SessionSid             string     `json:"session_sid"`
	Sid                    string     `json:"sid"`
	Type                   *string    `json:"type,omitempty"`
	URL                    string     `json:"url"`
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

// InteractionsPageResponse defines the response fields for the interactions page
type InteractionsPageResponse struct {
	Interactions []PageInteractionResponse `json:"interactions"`
	Meta         PageMetaResponse          `json:"meta"`
}

// Page retrieves a page of interactions
// See https://www.twilio.com/docs/proxy/api/interaction#read-multiple-interaction-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *InteractionsPageOptions) (*InteractionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of interactions
// See https://www.twilio.com/docs/proxy/api/interaction#read-multiple-interaction-resources for more details
func (c Client) PageWithContext(context context.Context, options *InteractionsPageOptions) (*InteractionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Sessions/{sessionSid}/Interactions",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sessionSid": c.sessionSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &InteractionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// InteractionsPaginator defines the fields for makings paginated api calls
// Interactions is an array of interactions that have been returned from all of the page calls
type InteractionsPaginator struct {
	options      *InteractionsPageOptions
	Page         *InteractionsPage
	Interactions []PageInteractionResponse
}

// NewInteractionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewInteractionsPaginator() *InteractionsPaginator {
	return c.NewInteractionsPaginatorWithOptions(nil)
}

// NewInteractionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewInteractionsPaginatorWithOptions(options *InteractionsPageOptions) *InteractionsPaginator {
	return &InteractionsPaginator{
		options: options,
		Page: &InteractionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Interactions: make([]PageInteractionResponse, 0),
	}
}

// InteractionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageInteractionResponse or error that is returned from the api call(s)
type InteractionsPage struct {
	client *Client

	CurrentPage *InteractionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *InteractionsPaginator) CurrentPage() *InteractionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *InteractionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *InteractionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *InteractionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &InteractionsPageOptions{}
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
		p.Interactions = append(p.Interactions, resp.Interactions...)
	}

	return p.Page.Error == nil
}
