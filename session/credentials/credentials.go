package credentials

type Credentials struct {
	Username string
	Password string
}

func New(creds TwilioCredentials) (*Credentials, error) {
	if err := creds.Validate(); err != nil {
		return nil, err
	}

	return &Credentials{
		Username: creds.username(),
		Password: creds.password(),
	}, nil
}
