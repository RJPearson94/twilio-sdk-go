package credentials

// Credentials respresent the field necessary to authenticate against the Twilio APIs
type Credentials struct {
	Username string
	Password string
}

// New creates a new instance of credentials using the supplied twilio credentials.
// If the credentials are invalid then an error will be returned
func New(creds TwilioCredentials) (*Credentials, error) {
	if err := creds.Validate(); err != nil {
		return nil, err
	}

	return &Credentials{
		Username: creds.username(),
		Password: creds.password(),
	}, nil
}
