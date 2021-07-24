// Package credentials contains auto-generated files. DO NOT MODIFY
package credentials

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credential"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/public_key"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/public_keys"
)

// Client for managing a credential resources
// See https://www.twilio.com/docs/iam/credentials/api for more details
type Client struct {
	client *client.Client

	AWSCredential  func(string) *aws_credential.Client
	AWSCredentials *aws_credentials.Client
	PublicKey      func(string) *public_key.Client
	PublicKeys     *public_keys.Client
}

// New creates a new instance of the credentials client
func New(client *client.Client) *Client {
	return &Client{
		client: client,

		AWSCredential: func(awsCredentialSid string) *aws_credential.Client {
			return aws_credential.New(client, aws_credential.ClientProperties{
				Sid: awsCredentialSid,
			})
		},
		AWSCredentials: aws_credentials.New(client),
		PublicKey: func(publicKeySid string) *public_key.Client {
			return public_key.New(client, public_key.ClientProperties{
				Sid: publicKeySid,
			})
		},
		PublicKeys: public_keys.New(client),
	}
}
