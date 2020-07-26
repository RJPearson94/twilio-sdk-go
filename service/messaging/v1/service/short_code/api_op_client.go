// This is an autogenerated file. DO NOT MODIFY
package short_code

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// The properties required to manage the shortcode resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
