package credentials

type TwilioCredentials interface {
	Validate() error
	username() string
	password() string
}
