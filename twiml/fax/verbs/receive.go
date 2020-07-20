package verbs

import "encoding/xml"

type ReceiveAttributes struct {
	Action     *string `xml:"action,attr,omitempty"`
	MediaType  *string `xml:"mediaType,attr,omitempty"`
	Method     *string `xml:"method,attr,omitempty"`
	PageSize   *string `xml:"pageSize,attr,omitempty"`
	StoreMedia *bool   `xml:"storeMedia,attr,omitempty"`
}

type Receive struct {
	XMLName xml.Name `xml:"Receive"`

	ReceiveAttributes
}
