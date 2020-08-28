package credentials

// TwilioCredentials respresents the structure of twilio credentials
type TwilioCredentials interface {
	Validate() error
	AccountSid() string
	username() string
	password() string
}
