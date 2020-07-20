package verbs

import "encoding/xml"

type RecordAttributes struct {
	Action                        *string `xml:"action,attr,omitempty"`
	FinishOnKey                   *string `xml:"finishOnKey,attr,omitempty"`
	MaxLength                     *int    `xml:"maxLength,attr,omitempty"`
	Method                        *string `xml:"method,attr,omitempty"`
	PlayBeep                      *bool   `xml:"playBeep,attr,omitempty"`
	RecordingStatusCallback       *string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackEvent  *string `xml:"recordingStatusCallbackEvent,attr,omitempty"`
	RecordingStatusCallbackMethod *string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	Timeout                       *int    `xml:"timeout,attr,omitempty"`
	Transcribe                    *bool   `xml:"transcribe,attr,omitempty"`
	TranscribeCallback            *string `xml:"transcribeCallback,attr,omitempty"`
	Trim                          *string `xml:"trim,attr,omitempty"`
}

type Record struct {
	XMLName xml.Name `xml:"Record"`

	*RecordAttributes
}
