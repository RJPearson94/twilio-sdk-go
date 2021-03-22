package nouns

import (
	"encoding/xml"
)

type TaskAttributes struct {
	Priority *int `xml:"priority,attr,omitempty"`
	Timeout  *int `xml:"timeout,attr,omitempty"`
}

type Task struct {
	XMLName xml.Name `xml:"Task"`
	Text    string   `xml:",chardata"`

	TaskAttributes
}
