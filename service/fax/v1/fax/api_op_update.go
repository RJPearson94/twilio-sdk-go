// This is an autogenerated file. DO NOT MODIFY
package fax

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateFaxInput struct {
	Status *string `form:"Status,omitempty"`
}

type UpdateFaxResponse struct {
	APIVersion  string     `json:"api_version"`
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Direction   string     `json:"direction"`
	Duration    *int       `json:"duration,omitempty"`
	From        string     `json:"from"`
	MediaSid    *string    `json:"media_sid,omitempty"`
	MediaURL    *string    `json:"media_url,omitempty"`
	NumPages    *int       `json:"num_pages,omitempty"`
	Price       *string    `json:"price,omitempty"`
	PriceUnit   *string    `json:"price_unit,omitempty"`
	Quality     string     `json:"quality"`
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	To          string     `json:"to"`
	URL         string     `json:"url"`
}

func (c Client) Update(input *UpdateFaxInput) (*UpdateFaxResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateFaxInput) (*UpdateFaxResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Faxes/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &UpdateFaxResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}