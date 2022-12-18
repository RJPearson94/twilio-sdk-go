package credentials

// Credentials represents the field necessary to authenticate against the Twilio APIs
type Credentials struct {
	AccountSid string
	Username   string
	Password   string
}

// New creates a new instance of credentials using the supplied twilio credentials.
// If the credentials are invalid then an error will be returned
func New(creds TwilioCredentials) (*Credentials, error) {
	if err := creds.Validate(); err != nil {
		return nil, err
	}

	return &Credentials{
		AccountSid: creds.AccountSid(),
		Username:   creds.username(),
		Password:   creds.password(),
	}, nil
}

// NewWithNoValidation creates a new instance of credentials using the supplied twilio credentials.
// This skips validation as some tools may initialise the client before credentials are supplied
func NewWithNoValidation(creds TwilioCredentials) *Credentials {
	return &Credentials{
		AccountSid: creds.AccountSid(),
		Username:   creds.username(),
		Password:   creds.password(),
	}
}
