package credentials

// TwilioCredentials respresents the structure of twilio credentials
type TwilioCredentials interface {
	Validate() error
	username() string
	password() string
}
