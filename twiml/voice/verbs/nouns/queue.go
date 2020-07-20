package nouns

import "encoding/xml"

type QueueAttributes struct {
	Method              *string `xml:"method,attr,omitempty"`
	PostWorkActivitySid *string `xml:"postWorkActivitySid,attr,omitempty"`
	ReservationSid      *string `xml:"reservationSid,attr,omitempty"`
	URL                 *string `xml:"url,attr,omitempty"`
}

type Queue struct {
	XMLName xml.Name `xml:"Queue"`
	Text    string   `xml:",chardata"`

	QueueAttributes
}
