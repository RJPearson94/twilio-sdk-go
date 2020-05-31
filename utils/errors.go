package utils

type TwilioError struct {
	Code     *int    `json:"code"`
	Message  string  `json:"message"`
	MoreInfo *string `json:"more_info"`
	Status   int     `json:"status"`
}

func (twilioError TwilioError) IsNotFoundError() bool {
	return twilioError.Status == 404
}

func (twilioError TwilioError) Error() string {
	return twilioError.Message
}
