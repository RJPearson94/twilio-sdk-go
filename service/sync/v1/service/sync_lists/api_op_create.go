// This is an autogenerated file. DO NOT MODIFY
package sync_lists

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateSyncListInput struct {
	CollectionTtl *int    `form:"CollectionTtl,omitempty"`
	Ttl           *int    `form:"Ttl,omitempty"`
	UniqueName    *string `form:"UniqueName,omitempty"`
}

type CreateSyncListResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Revision    string     `json:"revision"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

func (c Client) Create(input *CreateSyncListInput) (*CreateSyncListResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateSyncListInput) (*CreateSyncListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	response := &CreateSyncListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}