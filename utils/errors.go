package utils

// TwilioError represents the error object returned by the Twilio API when an error occurs
type TwilioError struct {
	Code     *int                    `json:"code,omitempty"`
	Details  *map[string]interface{} `json:"details,omitempty"`
	Message  string                  `json:"message"`
	MoreInfo *string                 `json:"more_info,omitempty"`
	Status   int                     `json:"status"`
}

// IsNotFoundError check if the error is a not found error
func (twilioError TwilioError) IsNotFoundError() bool {
	return twilioError.Status == 404
}

// Error returns the error message
func (twilioError TwilioError) Error() string {
	return twilioError.Message
}
